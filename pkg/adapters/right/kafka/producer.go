package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/protobuf"
	"hexagonal_arch_with_Golang/pkg/adapters/dto/pb"
	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/ports"
)

type Adapter struct {
	cfg        *config.Config
	producer   *kafka.Producer
	serializer *protobuf.Serializer
}

var _ ports.KafkaPort = &Adapter{}

func New(cfg *config.Config) (*Adapter, error) {

	bootstrapServers := cfg.Kafka.BootStrapServers
	schemaRegistryUrl := cfg.Kafka.SchemaRegistryUrl

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return nil, err
	}

	fmt.Printf("Created Producer %v\n", producer)

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(schemaRegistryUrl))
	if err != nil {
		fmt.Printf("Failed to create schema registry client: %s\n", err)
		return nil, err
	}

	serializer, err := protobuf.NewSerializer(client, serde.ValueSerde, protobuf.NewSerializerConfig())
	if err != nil {
		fmt.Printf("Failed to create serializer: %s\n", err)
		return nil, err
	}

	return &Adapter{
		cfg:        cfg,
		producer:   producer,
		serializer: serializer,
	}, nil
}

func (a *Adapter) FileProduce(filePr *pb.FilePr, topic string) error {
	err := a.produce(filePr, topic)
	if err != nil {
		fmt.Printf("Failed fileProduce: %s\n", err)
	}
	return nil
}

func (a *Adapter) TestProduce(fileName string, topic string) error {

	value := pb.FilePr{
		Name: fileName,
		Url:  "green",
	}
	err := a.produce(&value, topic)
	if err != nil {
		fmt.Printf("Failed fileProduce: %s\n", err)
	}
	return nil
}

func (a *Adapter) produce(value interface{}, topic string) error {
	deliveryChan := make(chan kafka.Event)

	payload, err := a.serializer.Serialize(topic, &value)
	if err != nil {
		fmt.Printf("Failed to serialize payload: %s\n", err)
		return err
	}

	err = a.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, deliveryChan)
	if err != nil {
		fmt.Printf("Produce failed: %v\n", err)
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)

	return nil
}
