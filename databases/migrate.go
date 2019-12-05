package databases

import "phone_email/structs/models"

func AutoMigrate() {
	Db.AutoMigrate(&models.PhoneNumber{})
}
