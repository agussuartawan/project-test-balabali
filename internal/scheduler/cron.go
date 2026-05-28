package scheduler

import (
	"log"

	"github.com/robfig/cron/v3"
)

func NewCron() *cron.Cron {
	c := cron.New()

	log.Println("Cron scheduler initialized")

	return c
}