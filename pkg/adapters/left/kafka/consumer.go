package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/protobuf"
	"google.golang.org/protobuf/reflect/protoreflect"
	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
	"hexagonal_arch_with_Golang/pkg/ports"
)

type Adapter struct {
	cfg          *config.Config
	consumer     *kafka.Consumer
	deserializer *protobuf.Deserializer
	app          ports.AppPort
}

var _ ports.KafkaConsumerPort = &Adapter{}

func New(cfg *config.Config, group string, app ports.AppPort) (*Adapter, error) {

	bootstrapServers := cfg.Kafka.BootStrapServers
	schemaRegistryUrl := cfg.Kafka.SchemaRegistryUrl

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  bootstrapServers,
		"group.id":           group,
		"session.timeout.ms": 6000,
		"auto.offset.reset":  "earliest"})

	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		if err != nil {
			return nil, err
		}
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(schemaRegistryUrl))

	if err != nil {
		fmt.Printf("Failed to create schema registry client: %s\n", err)
		return nil, err
	}

	deser, err := protobuf.NewDeserializer(client, serde.ValueSerde, protobuf.NewDeserializerConfig())
	if err != nil {
		fmt.Printf("Failed to create deserializer: %s\n", err)
		return nil, err
	}
	return &Adapter{
		cfg:          cfg,
		consumer:     c,
		deserializer: deser,
		app:          app,
	}, nil
}

func (a *Adapter) Consumer(topics []string, messageTypes ...protoreflect.MessageType) {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	err := a.consumer.SubscribeTopics(topics, nil)
	// Register the Protobuf type so that Deserialize can be called.
	// An alternative is to pass a pointer to an instance of the Protobuf type
	// to the DeserializeInto method.
	for _, mt := range messageTypes {
		err = a.deserializer.ProtoRegistry.RegisterMessage(mt)
		if err != nil {
			fmt.Printf("filed deserializer for %v \n", mt)
		}
	}

	run := true

	for run {
		select {
		case sig := <-sigChan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := a.consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				value, err := a.deserializer.Deserialize(*e.TopicPartition.Topic, e.Value)
				if err != nil {
					fmt.Printf("Failed to deserialize payload: %s\n", err)
				} else {
					switch value.(type) {
					case *pb.NotificationProducer:
						a.app.Notification(value.(*pb.NotificationProducer).Name, value.(*pb.NotificationProducer).Message)
						fmt.Printf("NotificationProducer: %v\n", value)
					case *pb.FileProducer:
						if value.(*pb.FileProducer).FileStatus == "parsing" {
							go a.app.Download(
								value.(*pb.FileProducer).FileUrl,
								value.(*pb.FileProducer).FilePath,
								value.(*pb.FileProducer).FileName,
							)
						}
					}
					fmt.Printf("%% Message on %s:\n%+v\n", e.TopicPartition, value)
				}
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				_, err = fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if err != nil {
					fmt.Printf("filed Fprint for kafkaError %s \n", err.Error())
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	err = a.consumer.Close()
	if err != nil {
		fmt.Printf("filed closing consumer %s \n", err.Error())
	}
}
