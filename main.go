package main

import (
	"log"
	"phone_email/databases"
	"phone_email/jobs"
	"phone_email/routers"
)

func foreverGo(run func(), routineLimits int) {
	for i := 0; i < routineLimits; i++ {
		go func() {
			for {
				run()
			}
		}()
	}
}

func main() {
	databases.AutoMigrate()

	foreverGo(jobs.SendEmail, 1)

	router := routers.Load()
	err := router.Run(":8191")
	if err != nil {
		log.Fatalln(err)
	}
}
