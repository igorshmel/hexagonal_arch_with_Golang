package ports

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"hexagonal_arch_with_Golang/pkg/adapters/dto/pb"
)

type APIPort interface {
	SayHello(name string) string
	FileDownload(pr *pb.FilePr) error
}

type KafkaConsumerPort interface {
	TestConsumer([]string, protoreflect.MessageType) error
}
