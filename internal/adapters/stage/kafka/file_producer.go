package kafka

import (
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

func (ths *Adapter) FileProducer(producer *pb.FileProducer) error {
	err := ths.produce(producer, producer.Topic)
	if err != nil {
		ths.cfg.Logger.Error("Failed fileProduce: %s", err)
		return err
	}
	return nil
}
