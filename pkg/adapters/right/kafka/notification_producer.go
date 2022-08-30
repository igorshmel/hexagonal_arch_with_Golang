package kafka

import (
	"fmt"

	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

func (a *Adapter) NotificationProducer(message string, topic string) error {

	value := pb.NotificationProducer{
		Name:    "file",
		Message: message,
	}

	err := a.produce(&value, topic)
	if err != nil {
		fmt.Printf("Failed NotificationProduce: %s\n", err)
		return err
	}
	return nil
}
