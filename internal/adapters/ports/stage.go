package ports

import (
	"hexagonal_arch_with_Golang/internal/adapters/stage/db/models"
)

type DbPort interface {
	NewRecordFile(model *models.PsqlFile) error
	IsFileExists(fileName string) bool
	ChangeStatusFile(fileName, fileStatus string) error
}
