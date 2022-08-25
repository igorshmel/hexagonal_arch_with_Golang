package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
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
	Logger  *log.Logger

	HostName     string
	FrontEndPath string
	Debug        bool
	Server       Server
	GRPC         GRPC
	Db           Db
	NATS         NATS
	Kafka        Kafka
}

type Server struct {
	Origins []string
	Host    string
	Address string
}

type GRPC struct {
	Network string
	Address string
}

type Kafka struct {
	BootStrapServers  string
	SchemaRegistryUrl string
}

type NATS struct {
	URL      string
	Username string
	Password string
	TopicID  string
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

	return &Config{
		Context: ctx,
		Logger:  log.Default(),
	}
}

func (c *Config) Read() error {
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

	return c.parseConfig(viper.GetViper())
}

func (c *Config) Write() error {
	return fmt.Errorf("not implemented")
}

func (c *Config) parseConfig(v *viper.Viper) error {
	*c = Config{
		HostName:     v.GetString("HOSTNAME"),
		FrontEndPath: v.GetString("FRONT_END_PATH"),
		Debug:        v.GetBool("DEBUG"),
		Server: Server{
			Origins: strings.Split(v.GetString("ORIGINS"), ","),
			Host:    v.GetString("HOST"),
			Address: v.GetString("ADDRESS"),
		},
		GRPC: GRPC{
			Network: v.GetString("GRPC_NETWORK"),
			Address: v.GetString("GRPC_ADDRESS"),
		},
		Kafka: Kafka{
			BootStrapServers:  v.GetString("KAFKA_BOOTSTRAP_SERVERS"),
			SchemaRegistryUrl: v.GetString("SCHEMA_REGISTRY_URL"),
		},
		NATS: NATS{
			URL:      v.GetString("NATS_URL"),
			Username: v.GetString("NATS_USERNAME"),
			Password: v.GetString("NATS_PASSWORD"),
			TopicID:  v.GetString("NATS_TOPIC_ID"),
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

	if c.Debug {
		log.Println("WARNING: DEBUG Mode enabled!")
	}

	return nil
}
