package ports

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type KafkaConsumerPort interface {
	Consumer(topics []string, mt ...protoreflect.MessageType)
}

type RestPort interface {
	Run()
	FileHandlers()
}
