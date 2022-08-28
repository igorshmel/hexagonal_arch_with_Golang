package ports

type DbPort interface {
	GetRandomGreeting(string) string
	GetGreetings() []string
	WriteFileToDownload(fileName, fileUrl, fileStatus string) error
}

type KafkaPort interface {
	TestProduce(fileName string, topic string) error
	FileProduce(fileName, fileUrl, fileStatus, topic string) error
}
