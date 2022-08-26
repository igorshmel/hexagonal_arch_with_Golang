package ports

import "hexagonal_arch_with_Golang/pkg/adapters/dto/pb"

type DbPort interface {
	GetRandomGreeting(string) string
	GetGreetings() []string
	WriteFileToDownload(fileName string, fileUrl string) error
}

type KafkaPort interface {
	TestProduce(fileName string, topic string) error
	FileProduce(pr *pb.FilePr, topic string) error
}
