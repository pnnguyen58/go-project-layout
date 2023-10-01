package usecases

import (
	"context"
	"errors"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"

	"github.com/pnnguyen58/go-project-layout/internal/core/ports/repositories"
	"github.com/pnnguyen58/go-project-layout/internal/core/usecases/workflows"
	"github.com/pnnguyen58/go-project-layout/pkg/clients"
	protogen "github.com/pnnguyen58/go-project-layout/pkg/proto_generated"
)

type Loan interface {
	Create(context.Context, *protogen.LoanCreateRequest) (*protogen.LoanCreateResponse, error)
	Approve(context.Context, *protogen.LoanApproveRequest) (*protogen.LoanApproveResponse, error)
	Get(context.Context, *protogen.LoanGetRequest) (*protogen.LoanGetResponse, error)
	AddRepayment(context.Context, *protogen.RepaymentCreateRequest) (*protogen.RepaymentCreateResponse, error)
}

type loan struct {
	logger         *zap.Logger
	temporalClient client.Client
	tempoWorkflow  *clients.Workflow
	loanRepo       repositories.Loan
	repaymentRepo  repositories.Repayment
}

func NewLoan(logger *zap.Logger, cl client.Client, tcf *clients.TempoConfig,
	loanRepo repositories.Loan, repaymentRepo repositories.Repayment) Loan {
	return &loan{
		logger:         logger,
		temporalClient: cl,
		tempoWorkflow:  tcf.Workflows["loan-workflow"],
		loanRepo:       loanRepo,
		repaymentRepo:  repaymentRepo,
	}
}

// Create new loan use case
func (l *loan) Create(ctx context.Context, req *protogen.LoanCreateRequest,
) (*protogen.LoanCreateResponse, error) {
	res, err := ExecuteWorkflow[*protogen.LoanCreateRequest, *protogen.LoanCreateResponse](
		ctx, l.logger, l.temporalClient, l.tempoWorkflow, workflows.CreateLoan, req)
	if err != nil {
		l.logger.Error("execute create loan workflow fail", zap.Error(err))
		return nil, errors.New("create loan failed")
	}
	if res == nil {
		return new(protogen.LoanCreateResponse), nil
	}
	return res, nil
}

// Approve loan use case
func (l *loan) Approve(ctx context.Context, req *protogen.LoanApproveRequest,
) (*protogen.LoanApproveResponse, error) {
	res, err := ExecuteWorkflow[*protogen.LoanApproveRequest, *protogen.LoanApproveResponse](
		ctx, l.logger, l.temporalClient, l.tempoWorkflow, workflows.ApproveLoan, req)
	if err != nil {
		l.logger.Error("execute approve loan workflow fail", zap.Error(err))
		return nil, errors.New("approve loan failed")
	}
	if res == nil {
		return new(protogen.LoanApproveResponse), nil
	}
	return res, nil
}

// Get loan use case
func (l *loan) Get(ctx context.Context, req *protogen.LoanGetRequest,
) (*protogen.LoanGetResponse, error) {
	res, err := ExecuteWorkflow[*protogen.LoanGetRequest, *protogen.LoanGetResponse](
		ctx, l.logger, l.temporalClient, l.tempoWorkflow, workflows.GetLoan, req)
	if err != nil {
		l.logger.Error("execute get loan workflow fail", zap.Error(err))
		return nil, errors.New("get loan failed")
	}
	if res == nil {
		return new(protogen.LoanGetResponse), nil
	}
	return res, nil
}

// AddRepayment add repayment use case
func (l *loan) AddRepayment(ctx context.Context, req *protogen.RepaymentCreateRequest,
) (*protogen.RepaymentCreateResponse, error) {
	res, err := ExecuteWorkflow[*protogen.RepaymentCreateRequest, *protogen.RepaymentCreateResponse](
		ctx, l.logger, l.temporalClient, l.tempoWorkflow, workflows.CreateRepayment, req)
	if err != nil {
		l.logger.Error("execute add repayment workflow fail", zap.Error(err))
		return nil, errors.New("add repayment failed")
	}
	if res == nil {
		return new(protogen.RepaymentCreateResponse), nil
	}
	return res, nil
}
