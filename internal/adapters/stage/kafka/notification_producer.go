package kafka

import (
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

func (ths *Adapter) NotificationProducer(producer *pb.NotificationProducer) error {

	err := ths.produce(producer, producer.Topic)
	if err != nil {
		ths.cfg.Logger.Error("failed NotificationProducer: %s", err)
		return err
	}
	return nil
}
