package settings

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	DbConnection string
	DbUsername   string
	DbPassword   string
	DbHost       string
	DbPort       string
	DbDatabase   string
)

var (
	EmailUsername      string
	EmailHost          string
	EmailPort          int
	EmailPassword      string
	EmailHeaderSubject string
)

var (
	EmailGoalEmailType_1 string
	EmailGoalEmailType_2 string
	EmailGoalEmailType_3 string
	EmailGoalEmailType_4 string
)

var Debug bool

func init() {
	checkEnv()
	LoadSetting()
}

func checkEnv() {
	_ = godotenv.Load()
	needChecks := []string{
		"DB_CONNECTION", "DB_HOST", "DB_PORT", "DB_DATABASE", "DB_USERNAME", "DB_PASSWORD",
		"EMAIL_USERNAME", "EMAIL_HOST", "EMAIL_PORT", "EMAIL_PASSWORD", "EMAIL_HEADER_SUBJECT",
		"EMAIL_GOAL_EMAIL_TYPE_1", "EMAIL_GOAL_EMAIL_TYPE_2", "EMAIL_GOAL_EMAIL_TYPE_3", "EMAIL_GOAL_EMAIL_TYPE_4",
	}

	for _, envKey := range needChecks {
		if os.Getenv(envKey) == "" {
			log.Fatalf("env %s missed", envKey)
		}
	}
}

func LoadSetting() {
	DbConnection = os.Getenv("DB_CONNECTION")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbDatabase = os.Getenv("DB_DATABASE")

	EmailUsername = os.Getenv("EMAIL_USERNAME")
	EmailHost = os.Getenv("EMAIL_HOST")
	EmailPort = loadIntFatal("EMAIL_PORT")
	EmailPassword = os.Getenv("EMAIL_PASSWORD")
	EmailHeaderSubject = os.Getenv("EMAIL_HEADER_SUBJECT")

	EmailGoalEmailType_1 = os.Getenv("EMAIL_GOAL_EMAIL_TYPE_1")
	EmailGoalEmailType_2 = os.Getenv("EMAIL_GOAL_EMAIL_TYPE_2")
	EmailGoalEmailType_3 = os.Getenv("EMAIL_GOAL_EMAIL_TYPE_3")
	EmailGoalEmailType_4 = os.Getenv("EMAIL_GOAL_EMAIL_TYPE_4")

	debug := os.Getenv("DEBUG")
	if debug != "" && debug != "false" && debug != "0" {
		Debug = true
	}
}

func loadIntFatal(e string) int {
	intVar, err := strconv.Atoi(os.Getenv(e))
	if err != nil {
		log.Fatalf("env %s invalid\n", e)
	}

	return intVar
}
