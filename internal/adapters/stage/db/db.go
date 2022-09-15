package pgql

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hexagonal_arch_with_Golang/internal/adapters/ports"
	"hexagonal_arch_with_Golang/internal/adapters/stage/db/models"
	"hexagonal_arch_with_Golang/pkg/config"
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
	if err != nil {
		return nil, err
	}

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
		err = db.AutoMigrate(
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

func (ths *Adapter) NewRecordFile(fileModel *models.PsqlFile) error {

	err := ths.db.Table(fileModel.TableName()).
		Create(fileModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (ths *Adapter) ChangeStatusFile(fileName, fileStatus string) error {
	err := ths.db.Table("files").
		Where("file_name=?", fileName).UpdateColumn("file_status", fileStatus).Error
	if err != nil {
		return err
	}
	return nil
}

// IsFileExists check file exists from repository
func (ths *Adapter) IsFileExists(fileName string) bool {
	if ths == nil || ths.db == nil || len(fileName) <= 0 {
		return true
	}

	var exists bool
	var psqlFile models.PsqlFile
	err := ths.db.Model(psqlFile).
		Select("count(*) > 0").
		Where("file_name = ?", fileName).
		Find(&exists).
		Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return false
	}
	return exists
}
