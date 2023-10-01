package clients

import (
	"context"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"time"
)

type TempoConfig struct {
	HostPort  string               `json:"hostPort"`
	Namespace *Namespace           `json:"namespace"`
	Workflows map[string]*Workflow `json:"workflows"`
}

type Namespace struct {
	Name                             string        `json:"namespace"`
	WorkflowExecutionRetentionPeriod time.Duration `json:"workflowExecutionRetentionPeriod"` // seconds
}

type Workflow struct {
	TaskQueueName          string                 `json:"taskQueueName"`
	SearchAttributes       map[string]interface{} `json:"searchAttributes"`
	ExecutionTimeout       time.Duration          `json:"executionTimeout"`       // seconds
	RunTimeout             time.Duration          `json:"runTimeout"`             // seconds
	TaskTimeout            time.Duration          `json:"taskTimeout"`            // seconds
	ScheduleToCloseTimeout time.Duration          `json:"scheduleToCloseTimeout"` // seconds
	StartToCloseTimeout    time.Duration          `json:"startToCloseTimeout"`    // seconds
	HeartbeatTimeout       time.Duration          `json:"heartbeatTimeout"`       // seconds
	MaximumInterval        time.Duration          `json:"maximumInterval"`        // seconds
	InitialInterval        time.Duration          `json:"initialInterval"`        // seconds
	BackoffCoefficient     float64                `json:"backoffCoefficient"`
	WaitForCancellation    bool                   `json:"waitForCancellation"`
}

func NewTemporal(ctx context.Context, tc *TempoConfig) (client.Client, error) {
	cl, err := client.Dial(client.Options{
		HostPort:  tc.HostPort,
		Namespace: tc.Namespace.Name,
	})
	if err != nil {
		return nil, err
	}
	namespace, err := cl.WorkflowService().DescribeNamespace(ctx, &workflowservice.DescribeNamespaceRequest{
		Namespace: tc.Namespace.Name,
	})
	if namespace != nil && err == nil {
		return cl, nil
	}
	_, err = cl.WorkflowService().RegisterNamespace(ctx, &workflowservice.RegisterNamespaceRequest{
		Namespace:                        tc.Namespace.Name,
		WorkflowExecutionRetentionPeriod: &tc.Namespace.WorkflowExecutionRetentionPeriod,
	})
	return cl, err
}
