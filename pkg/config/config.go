package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
	"hexagonal_arch_with_Golang/pkg/logger"
)

var (
	EnvPrefix = "HEX"
	Version   = "0.1.0"
	Build     = ""
)

var (
	ErrExit = errors.New("simple exit")
)

type Config struct {
	Context context.Context
	Logger  logger.Log

	HostName     string
	FrontEndPath string
	Debug        bool
	Server       Server
	GRPC         GRPC
	Db           Db
	Kafka        Kafka
	Env          Env
}

type Server struct {
	Origins []string
	Host    string
	Address string
}

type Env struct {
	ListenAddr string
}

type GRPC struct {
	Network string
	Address string
}

type Kafka struct {
	BootStrapServers  string
	SchemaRegistryUrl string
}

type Db struct {
	Username string
	Password string
	Host     string
	Name     string
	Port     int
	SSLMode  string
}

func isDotEnvPresent() bool {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return false
	}
	return true
}

func New(ctx context.Context) *Config {
	// viper.SetEnvPrefix(EnvPrefix)

	viper.SetDefault("ORIGINS", "http://localhost:8080,http://localhost:3000")
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("ADDRESS", ":8080")
	viper.SetDefault("GRPC_NETWORK", "tcp")
	viper.SetDefault("GRPC_ADDRESS", ":9000")
	viper.SetDefault("DB_USERNAME", "postgres")
	viper.SetDefault("DB_PASSWORD", "12345")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "hexagon")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("KAFKA_BOOTSTRAP_SERVERS", "localhost")
	viper.SetDefault("SCHEMA_REGISTRY_URL", "http://localhost:8081")
	viper.SetDefault("ENV_LISTEN_ADDRESS", ":9001")

	return &Config{
		Context: ctx,
		Logger:  logger.Log{},
	}
}

func (ths *Config) Read() error {
	if isDotEnvPresent() {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
		viper.SetConfigType("dotenv")

		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}
	viper.AutomaticEnv()

	return ths.parseConfig(viper.GetViper())
}

func (ths *Config) Write() error {
	return fmt.Errorf("not implemented")
}

func (ths *Config) parseConfig(v *viper.Viper) error {
	*ths = Config{
		HostName:     v.GetString("HOSTNAME"),
		FrontEndPath: v.GetString("FRONT_END_PATH"),
		Debug:        v.GetBool("DEBUG"),
		Server: Server{
			Origins: strings.Split(v.GetString("ORIGINS"), ","),
			Host:    v.GetString("HOST"),
			Address: v.GetString("ADDRESS"),
		},
		Env: Env{
			ListenAddr: v.GetString("ENV_LISTEN_ADDRESS"),
		},
		GRPC: GRPC{
			Network: v.GetString("GRPC_NETWORK"),
			Address: v.GetString("GRPC_ADDRESS"),
		},
		Kafka: Kafka{
			BootStrapServers:  v.GetString("KAFKA_BOOTSTRAP_SERVERS"),
			SchemaRegistryUrl: v.GetString("SCHEMA_REGISTRY_URL"),
		},
		Db: Db{
			Username: v.GetString("DB_USERNAME"),
			Password: v.GetString("DB_PASSWORD"),
			Host:     v.GetString("DB_HOST"),
			Port:     v.GetInt("DB_PORT"),
			Name:     v.GetString("DB_NAME"),
			SSLMode:  v.GetString("DB_SSL_MODE"),
		},
	}

	if ths.Debug {
		log.Println("WARNING: DEBUG Mode enabled!")
	}

	return nil
}
