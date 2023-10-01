package main

import (
	"github.com/pnnguyen58/go-project-layout/internal/core/usecases/activities"
	"github.com/pnnguyen58/go-project-layout/internal/core/usecases/workflows"
	"github.com/pnnguyen58/go-project-layout/pkg/clients"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func registerLoan(cl client.Client, conf *clients.TempoConfig) {
	for _, taskQueue := range conf.Workflows {
		switch taskQueue.TaskQueueName {
		case "loan-workflow":
			w := worker.New(cl, taskQueue.TaskQueueName, worker.Options{})
			w.RegisterWorkflow(workflows.CreateLoan)
			w.RegisterWorkflow(workflows.ApproveLoan)
			w.RegisterWorkflow(workflows.GetLoan)
			w.RegisterWorkflow(workflows.CreateRepayment)

			w.RegisterActivity(activities.CreateLoan)
			w.RegisterActivity(activities.CreateLoanCompensation)
			w.RegisterActivity(activities.ApproveLoan)
			w.RegisterActivity(activities.ApproveLoanCompensation)
			w.RegisterActivity(activities.GetLoan)
			w.RegisterActivity(activities.GetLoanCompensation)
			w.RegisterActivity(activities.CreateRepayment)
			w.RegisterActivity(activities.CreateRepaymentCompensation)
			// TODO: add more workflows and activities
			go func() {
				err := w.Run(worker.InterruptCh())
				if err != nil {
					log.Fatalln("Unable to start worker", err)
				}
			}()
		default:
			// TODO: add more workers
			log.Println("app not defined")
		}
	}
}
