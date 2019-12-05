package services

import "phone_email/settings"

// EmailAddress 根据邮箱类型获取邮箱地址
func EmailAddress(emailType int) string {
	switch emailType {
	case 1:
		return settings.EmailGoalEmailType_1
	case 2:
		return settings.EmailGoalEmailType_2
	case 3:
		return settings.EmailGoalEmailType_3
	case 4:
		return settings.EmailGoalEmailType_4
	}
	return settings.EmailGoalEmailType_1
}
