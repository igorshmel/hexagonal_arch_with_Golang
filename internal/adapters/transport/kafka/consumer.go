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
	"hexagonal_arch_with_Golang/internal/adapters/ports"
	"hexagonal_arch_with_Golang/internal/app"
	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

type Adapter struct {
	cfg          *config.Config
	consumer     *kafka.Consumer
	deserializer *protobuf.Deserializer
	app          app.ApiPort
}

var _ ports.KafkaConsumerPort = &Adapter{}

func New(cfg *config.Config, group string, app app.ApiPort) (*Adapter, error) {

	bootstrapServers := cfg.Kafka.BootStrapServers
	schemaRegistryURL := cfg.Kafka.SchemaRegistryURL

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  bootstrapServers,
		"group.id":           group,
		"session.timeout.ms": 6000,
		"auto.offset.reset":  "earliest"})

	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "Failed to create consumer: %s", err)
		if err != nil {
			return nil, err
		}
		os.Exit(1)
	}

	cfg.Logger.Info("Created Consumer %v", c)

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(schemaRegistryURL))

	if err != nil {
		cfg.Logger.Error("Failed to create schema registry client: %s", err)
		return nil, err
	}

	deserializer, err := protobuf.NewDeserializer(client, serde.ValueSerde, protobuf.NewDeserializerConfig())
	if err != nil {
		cfg.Logger.Error("Failed to create deserializer: %s", err)
		return nil, err
	}
	return &Adapter{
		cfg:          cfg,
		consumer:     c,
		deserializer: deserializer,
		app:          app,
	}, nil
}

func (ths *Adapter) Consumer(topics []string, messageTypes ...protoreflect.MessageType) {
	l := ths.cfg.Logger
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	err := ths.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		l.Error("filed consumer subscribe topics %s", err.Error())
	}

	// Register the Protobuf type so that Deserialize can be called.
	// An alternative is to pass a pointer to an instance of the Protobuf type
	// to the DeserializeInto method.
	for _, mt := range messageTypes {
		err = ths.deserializer.ProtoRegistry.RegisterMessage(mt)
		if err != nil {
			l.Error("filed deserializer for %v", mt)
		}
	}

	run := true

	for run {
		select {
		case sig := <-sigChan:
			l.Debug("Caught signal %v: terminating", sig)
			run = false
		default:
			ev := ths.consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				value, errDes := ths.deserializer.Deserialize(*e.TopicPartition.Topic, e.Value)
				if errDes != nil {
					l.Error("Failed to deserialize payload: %s", errDes)
				} else {
					switch value.(type) {
					case *pb.NotificationProducer:
						v := value.(*pb.NotificationProducer)
						if v != nil {
							ths.app.Notification(v.Name, v.Message)
						}
						l.Info("NotificationProducer: %v", value)
					case *pb.FileProducer:
						v := value.(*pb.FileProducer)
						if v != nil {
							if v.FileStatus == "parsing" {
								go ths.app.Download(
									v.FileUrl,
									v.FilePath,
									v.FileName,
								)
							}
						}
					}
					l.Info("%% Message on %s: %+v", e.TopicPartition, value)
				}
				if e.Headers != nil {
					l.Debug("%% Headers: %v", e.Headers)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				_, err = fmt.Fprintf(os.Stderr, "%% Error: %v: %v", e.Code(), e)
				if err != nil {
					l.Error("failed Fprint for kafkaError %s", err.Error())
				}
			default:
				l.Debug("Ignored %v", e)
			}
		}
	}

	fmt.Printf("Closing consumer")
	err = ths.consumer.Close()
	if err != nil {
		l.Error("filed closing consumer %s", err.Error())
	}
}
