package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/protobuf"
	"hexagonal_arch_with_Golang/internal/adapters/ports"
	"hexagonal_arch_with_Golang/pkg/config"
)

type Adapter struct {
	cfg        *config.Config
	producer   *kafka.Producer
	serializer *protobuf.Serializer
}

var _ ports.KafkaPort = &Adapter{}

func New(cfg *config.Config) (*Adapter, error) {

	bootstrapServers := cfg.Kafka.BootStrapServers
	schemaRegistryURL := cfg.Kafka.SchemaRegistryURL

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		cfg.Logger.Error("Failed to create producer: %s", err)
		return nil, err
	}
	cfg.Logger.Info("Created Producer %v", producer)

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(schemaRegistryURL))
	if err != nil {
		cfg.Logger.Error("Failed to create schema registry client: %s", err)
		return nil, err
	}

	serializer, err := protobuf.NewSerializer(client, serde.ValueSerde, protobuf.NewSerializerConfig())
	if err != nil {
		cfg.Logger.Error("Failed to create serializer: %s", err)
		return nil, err
	}

	return &Adapter{
		cfg:        cfg,
		producer:   producer,
		serializer: serializer,
	}, nil
}

func (ths *Adapter) produce(value interface{}, topic string) error {
	deliveryChan := make(chan kafka.Event)

	payload, err := ths.serializer.Serialize(topic, value)
	if err != nil {
		ths.cfg.Logger.Error("Failed to serialize payload: %s", err)
		return err
	}

	err = ths.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, deliveryChan)
	if err != nil {
		ths.cfg.Logger.Error("Produce failed: %v", err)
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		ths.cfg.Logger.Error("Delivery failed: %v", m.TopicPartition.Error)
	} else {
		ths.cfg.Logger.Info("Delivered message to topic %s [%d] at offset %v",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)

	return nil
}
