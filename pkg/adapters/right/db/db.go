package pgql

import (
	"fmt"
	"time"

	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/models"
	"hexagonal_arch_with_Golang/pkg/ports"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Adapter struct {
	cfg *config.Config
	db  *gorm.DB
}

var _ ports.DbPort = &Adapter{}

func New(cfg *config.Config, migrate bool) (*Adapter, error) {
	var driver gorm.Dialector
	if cfg.Db.Host == "SQLite" {
		driver = sqlite.Open(cfg.Db.Name)
	} else {
		conn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Db.Host,
			cfg.Db.Port,
			cfg.Db.Username,
			cfg.Db.Password,
			cfg.Db.Name,
			cfg.Db.SSLMode,
		)
		driver = postgres.Open(conn)
	}

	db, err := gorm.Open(driver, &gorm.Config{})

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	/*err = db.Migrator().DropTable("file")
	if err != nil {
		fmt.Println("Error Drop Table file")
	}*/

	if migrate {
		err := db.AutoMigrate(
			&models.PsqlNotification{},
			&models.PsqlFile{},
		)
		if err != nil {
			return nil, err
		}
	}

	return &Adapter{
		cfg: cfg,
		db:  db,
	}, nil
}

func (ths *Adapter) GetRandomNotification(name string) string {
	g := models.PsqlNotification{Name: name}
	err := ths.db.Table("notifications").Create(&g).Error
	if err != nil {
		return err.Error()
	}
	return "OK"
}

func (ths *Adapter) NewRecordFile(model *models.PsqlFile) error {
	err := ths.db.Table("files").Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

// NewPsqlFile	takes in an aggregate
func (ths *Adapter) NewPsqlFile(name, path, url, hash, status string) *models.PsqlFile {
	return &models.PsqlFile{
		FileName:   name,
		FilePath:   path,
		FileUrl:    url,
		FileHash:   hash,
		FileStatus: status,
	}
}

// NewPsqlNotification	takes in an aggregate
func (ths *Adapter) NewPsqlNotification(name, message string) *models.PsqlNotification {
	return &models.PsqlNotification{
		Name:    "file",
		Message: message,
	}
}
func (ths *Adapter) NotificationRecord(model *models.PsqlNotification) error {
	err := ths.db.Table("notification").Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

//// PsqlToDomain converts into a domain.File
//func (ths *Adapter) PsqlToDomain() *fileDomain.File {
//	return fileDomain.New(
//		ths.cfg,
//		domain.UserUUID(ths.UserUUID),
//		domain.ConfirmToken(ths.ConfirmToken),
//		domain.ExpiredToken(ths.ExpiredToken),
//		domain.TypeToken(ths.TypeToken),
//		domain.UsedFlag(ths.Used),
//		domain.Phone(ths.Phone),
//	)
//}
