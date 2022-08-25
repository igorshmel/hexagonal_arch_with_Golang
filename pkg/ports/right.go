package ports

type DbPort interface {
	GetRandomGreeting(string) string
	GetGreetings() []string
}

type KafkaPort interface {
	TestProduce(string, string) string
}
