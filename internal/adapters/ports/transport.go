package ports

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type AppPort interface {
	NewFile(fileUrl string) error
	Download(fileUrl, filePath, fileName string)
	Notification(name, message string)
}

type KafkaConsumerPort interface {
	Consumer(topics []string, mt ...protoreflect.MessageType)
}
