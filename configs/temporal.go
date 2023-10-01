package configs

import (
	"time"

	"github.com/pnnguyen58/go-project-layout/pkg/clients"
)

func LoadTempoConfig() *clients.TempoConfig {
	return mockTempoConfig()
}

func mockTempoConfig() (tc *clients.TempoConfig) {
	tc = &clients.TempoConfig{
		HostPort: C.Server.TempoHost,
		Namespace: &clients.Namespace{
			Name:                             "aspire-code-challenge",
			WorkflowExecutionRetentionPeriod: 1720 * time.Hour,
		},
		Workflows: map[string]*clients.Workflow{},
	}
	tc.Workflows["loan-workflow"] = &clients.Workflow{
		TaskQueueName:      "loan-workflow",
		ExecutionTimeout:   10 * time.Second,
		RunTimeout:         10 * time.Second,
		TaskTimeout:        10 * time.Second,
		MaximumInterval:    10 * time.Second,
		InitialInterval:    10 * time.Second,
		BackoffCoefficient: 1.0,
	}
	tc.Workflows["repayment-workflow"] = &clients.Workflow{
		TaskQueueName:      "repayment-workflow",
		ExecutionTimeout:   10 * time.Second,
		RunTimeout:         10 * time.Second,
		TaskTimeout:        10 * time.Second,
		MaximumInterval:    10 * time.Second,
		InitialInterval:    10 * time.Second,
		BackoffCoefficient: 1.0,
	}
	return
}
