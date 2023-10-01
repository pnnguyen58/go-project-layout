package usecases

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.uber.org/zap"

	"github.com/pnnguyen58/go-project-layout/pkg/clients"
)

func ExecuteWorkflow[Req any, Res any](ctx context.Context, logger *zap.Logger, cl client.Client, conf *clients.Workflow,
	workflow interface{}, req Req) (Res, error) {
	var res Res
	if cl == nil {
		return res, errors.New("not found client")
	}
	if conf == nil {
		return res, errors.New("empty config")
	}
	if workflow == nil {
		return res, errors.New("invalid workflow")
	}
	// Get task config
	taskQueueName := conf.TaskQueueName
	taskQueueID := uuid.New().String()
	taskTimeout := conf.TaskTimeout

	// Get workflow config
	attributes := conf.SearchAttributes
	executionTimeout := conf.ExecutionTimeout
	runTimeout := conf.RunTimeout

	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    conf.InitialInterval,
		BackoffCoefficient: conf.BackoffCoefficient,
		MaximumInterval:    conf.MaximumInterval,
	}
	workflowOptions := client.StartWorkflowOptions{
		ID:                       taskQueueName + "_" + taskQueueID,
		TaskQueue:                taskQueueName,
		SearchAttributes:         attributes,
		WorkflowExecutionTimeout: executionTimeout,
		WorkflowRunTimeout:       runTimeout,
		RetryPolicy:              retryPolicy,
	}

	we, err := cl.ExecuteWorkflow(ctx, workflowOptions, workflow, req)
	if err != nil {
		logger.Error("execute workflow failed")
		return res, err
	}

	ctxWithTimeout, cancelHandler := context.WithTimeout(context.Background(), taskTimeout)
	defer cancelHandler()
	err = we.Get(ctxWithTimeout, &res)
	if err != nil {
		return res, err
	}
	logger.Info(fmt.Sprintf("execute workflow ID: %v successfully", we.GetID()))
	return res, nil
}
