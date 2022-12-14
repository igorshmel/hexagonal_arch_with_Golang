package ports

import (
	"hexagonal_arch_with_Golang/internal/domain/file"
	"hexagonal_arch_with_Golang/internal/domain/notification"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

type DbPort interface {
	GetRandomNotification(string) string
	NewRecordFile(model *fileDomain.File) error
	ReadFile(file *fileDomain.File) error
	IsFileExists(fileName string) bool
	ChangeStatusFile(fileName, fileStatus string) error
	NotificationRecord(model *notificationDomain.Notification) error
}

type KafkaPort interface {
	NotificationProducer(producer *pb.NotificationProducer) error
	FileProducer(producer *pb.FileProducer) error
}

type MemoryPort interface {
	GetRandomNotification() string
}
