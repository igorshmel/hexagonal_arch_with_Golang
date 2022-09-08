package api

import (
	"hexagonal_arch_with_Golang/internal/domain/notification"
)

func (ths *Application) Notification(name, message string) {
	nd := notificationDomain.New(ths.cfg)
	nd.SetMassage(message)
	nd.SetName(name)

	err := ths.db.NotificationRecord(nd)
	if err != nil {
		ths.cfg.Logger.Error("Error DB %s", err.Error())
	}

}
