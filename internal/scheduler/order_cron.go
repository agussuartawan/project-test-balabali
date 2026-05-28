package scheduler

import (
	"context"
	"log"

	"github.com/agussuartawan/project-test-balabali/internal/bootstrap"
	"github.com/robfig/cron/v3"
)

func RegisterOrderCron(c *cron.Cron, deps *bootstrap.Dependencies) {
	_, err := c.AddFunc("*/5 * * * *", func() {
		log.Println("running order cron...")

		orderServices := deps.OrderService
		if err := orderServices.ProcessOrder(context.Background()); err != nil {
			log.Printf("error processing order: %v\n", err)
		}

		log.Println("order cron finished")
	}); if err != nil {
		log.Printf("error adding order cron: %v\n", err)
	}
}