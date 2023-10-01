package workflows

import (
	"go.temporal.io/sdk/workflow"
	"go.uber.org/multierr"
	"time"

	"github.com/pnnguyen58/go-project-layout/internal/core/usecases/activities"
	protogen "github.com/pnnguyen58/go-project-layout/pkg/proto_generated"
)

// CreateLoan workflows definition
func CreateLoan(ctx workflow.Context, flowInput *protogen.LoanCreateRequest) (
	*protogen.LoanCreateResponse, error) {
	// Workflow has to check input valid or not
	//inputErr := flowInput.CheckValid()
	//if inputErr != nil {
	//	return nil,
	//		temporal.NewNonRetryableApplicationError("Invalid flow input", common.ErrInvalidInput, inputErr, nil)
	//}

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// This is how you log
	// workflows.GetLogger(ctx).Info("jobInput.Inputs", flowInput.Inputs)

	result := &protogen.LoanCreateResponse{}
	err := workflow.ExecuteActivity(ctx, activities.CreateLoan, flowInput).Get(ctx, result)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, activities.CreateLoanCompensation, flowInput).
				Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()
	workflow.GetLogger(ctx).Info("Workflow completed.")

	return result, err
}
