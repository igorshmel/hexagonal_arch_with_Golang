package ports

import (
	"hexagonal_arch_with_Golang/internal/app/domain/file"
	"hexagonal_arch_with_Golang/internal/models"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

type DbPort interface {
	GetRandomNotification(string) string
	NewRecordFile(model *models.PsqlFile) error
	ReadFile(file *fileDomain.File) error
	IsFileExists(fileName string) bool
	ChangeStatusFile(fileName, fileStatus string) error
	NotificationRecord(model *models.PsqlNotification) error
}

type KafkaPort interface {
	NotificationProducer(producer *pb.NotificationProducer) error
	FileProducer(producer *pb.FileProducer) error
}

type MemoryPort interface {
	GetRandomNotification() string
}
