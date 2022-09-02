package kafka

import (
	"fmt"

	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

func (a *Adapter) FileProducer(producer *pb.FileProducer) error {
	err := a.produce(producer, producer.Topic)
	if err != nil {
		fmt.Printf("Failed fileProduce: %s\n", err)
		return err
	}
	return nil
}
