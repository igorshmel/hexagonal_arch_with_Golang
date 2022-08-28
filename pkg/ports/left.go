package ports

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type APIPort interface {
	SayHello(name string) string
	FileDownload(fileUrl string) error
	Download(fileUrl, fileName string)
}

type KafkaConsumerPort interface {
	Consumer([]string, ...protoreflect.MessageType)
}
