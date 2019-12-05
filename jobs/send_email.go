package jobs

import (
	"fmt"
	"phone_email/databases"
	"phone_email/services"
	"phone_email/settings"
	"phone_email/structs/models"
	"phone_email/structs/models/logics"
	"strings"
	"time"
)

// SendEmail 发送邮件
// 一次性发送多条手机号
func SendEmail() {
	goalEmailMapType1 := []models.PhoneNumber{}
	goalEmailMapType2 := []models.PhoneNumber{}
	goalEmailMapType3 := []models.PhoneNumber{}
	goalEmailMapType4 := []models.PhoneNumber{}

	phoneNumbers := []models.PhoneNumber{}
	databases.Db.Model(phoneNumbers).Where("status in (?)", []int{logics.STATUS_未发送, logics.STATUS_发送失败}).Find(&phoneNumbers)

	if len(phoneNumbers) == 0 {
		fmt.Println("没有号码，休眠一分钟")
		time.Sleep(time.Minute)
		return
	}

	for i := range phoneNumbers {
		switch phoneNumbers[i].Type {
		case 1:
			goalEmailMapType1 = append(goalEmailMapType1, phoneNumbers[i])
		case 2:
			goalEmailMapType2 = append(goalEmailMapType2, phoneNumbers[i])
		case 3:
			goalEmailMapType3 = append(goalEmailMapType3, phoneNumbers[i])
		case 4:
			goalEmailMapType4 = append(goalEmailMapType4, phoneNumbers[i])
		}
	}

	typeGoalEmailMap := map[string][]models.PhoneNumber{
		settings.EmailGoalEmailType_1: goalEmailMapType1,
		settings.EmailGoalEmailType_2: goalEmailMapType2,
		settings.EmailGoalEmailType_3: goalEmailMapType3,
		settings.EmailGoalEmailType_4: goalEmailMapType4,
	}

	for goalEmail, numbers := range typeGoalEmailMap {
		if len(numbers) == 0 {
			continue
		}

		numberIds := []int{}
		for i := range numbers {
			numberIds = append(numberIds, numbers[i].ID)
		}

		needSendNumbers := []string{}
		for i := range numbers {
			needSendNumbers = append(needSendNumbers, numbers[i].Number)
		}

		err := services.SendEmail(strings.Join(needSendNumbers, ", "), goalEmail)
		if err != nil {
			fmt.Println(fmt.Sprintf("邮件发送失败, 目标邮箱: %s； 错误信息:%s", goalEmail, err.Error()))
			databases.Db.Model(models.PhoneNumber{}).
				Where("id in (?)", numberIds).
				Where("status = ?", logics.STATUS_未发送).
				Update("status", logics.STATUS_发送失败)
			time.Sleep(time.Second)
			continue
		}

		databases.Db.Model(models.PhoneNumber{}).
			Where("id in (?)", numberIds).
			Update("status", logics.STATUS_已发送)
		fmt.Println("邮件发送成功")
	}
}
