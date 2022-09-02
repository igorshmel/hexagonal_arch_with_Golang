package api

import (
	"fmt"

	"hexagonal_arch_with_Golang/internal/models"
)

func (ths *Application) Notification(name, message string) {

	psqlNotificationModel := models.NewPsqlNotification(name, message)
	err := ths.db.NotificationRecord(psqlNotificationModel)
	if err != nil {
		fmt.Printf("Error DB %s\n", err.Error())
	}

}
