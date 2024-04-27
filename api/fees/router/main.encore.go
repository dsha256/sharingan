package router

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/dsha256/sharingan/api/fees/config"
	"github.com/dsha256/sharingan/api/fees/controller"
	"github.com/dsha256/sharingan/api/fees/service"
	"github.com/dsha256/sharingan/internal/temporal/activity"
	"github.com/dsha256/sharingan/internal/temporal/workflow"
)

//encore:service
type Fees struct {
	client            client.Client
	worker            worker.Worker
	billingController *controller.Billing
}

func initFees() (*Fees, error) {
	c, err := client.Dial(client.Options{})
	if err != nil {
		return nil, fmt.Errorf("create temporal client: %v", err)
	}

	w := worker.New(c, config.FeesTaskQueue, worker.Options{})

	w.RegisterWorkflow(workflow.Billing)
	w.RegisterActivity(activity.MoneyAddition)

	err = w.Start()
	if err != nil {
		return nil, fmt.Errorf("start temporal worker: %v", err)
	}

	newBillingService := service.NewBillingService(c)
	newBillingController := controller.NewBillingController(newBillingService)

	return &Fees{client: c, worker: w, billingController: newBillingController}, nil
}

func (s *Fees) Shutdown(force context.Context) {
	s.client.Close()
	s.worker.Stop()
}
