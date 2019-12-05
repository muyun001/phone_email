package services

import (
	"errors"
	"phone_email/databases"
	"phone_email/structs/models"
	"strings"
	"time"
)

// SaveNumber 保存手机号信息
func SaveNumber(callId string, number string, emailType int) error {
	phoneNumber := &models.PhoneNumber{
		CallId:    callId,
		Number:    number,
		Email:     EmailAddress(emailType),
		Type:      emailType,
		CreatedAt: time.Time{},
	}

	err := databases.Db.Save(phoneNumber).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("Duplicate entry")
		}
		return err
	}

	return nil
}
