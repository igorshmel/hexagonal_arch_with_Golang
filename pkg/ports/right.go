package ports

import (
	"hexagonal_arch_with_Golang/pkg/dto/pb"
	"hexagonal_arch_with_Golang/pkg/models"
)

type DbPort interface {
	GetRandomNotification(string) string
	NewRecordFile(model *models.PsqlFile) error
	NotificationRecord(model *models.PsqlNotification) error
	NewPsqlFile(fileName, filePath, fileUrl, fileHash, fileStatus string) *models.PsqlFile
	NewPsqlNotification(name, message string) *models.PsqlNotification
}

type KafkaPort interface {
	NotificationProducer(message string, topic string) error
	FileProducer(producer *pb.FileProducer) error
}
