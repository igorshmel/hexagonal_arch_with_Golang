package kafka

import (
	"fmt"

	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

func (a *Adapter) NotificationProducer(producer *pb.NotificationProducer) error {

	err := a.produce(producer, producer.Topic)
	if err != nil {
		fmt.Printf("failed NotificationProducer: %s\n", err)
		return err
	}
	return nil
}
