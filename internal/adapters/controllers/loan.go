package controllers

import (
	"context"
	"go.uber.org/zap"

	"github.com/pnnguyen58/go-project-layout/internal/core/usecases"
	protogen "github.com/pnnguyen58/go-project-layout/pkg/proto_generated"
)

func NewLoan(logger *zap.Logger, loanUseCase usecases.Loan) *Loan {
	return &Loan{logger: logger, loanUseCase: loanUseCase}
}

type Loan struct {
	protogen.UnimplementedLoanServiceServer
	logger      *zap.Logger
	loanUseCase usecases.Loan
}

func (l *Loan) CreateLoan(ctx context.Context, req *protogen.LoanCreateRequest) (*protogen.LoanCreateResponse, error) {
	return l.loanUseCase.Create(ctx, req)
}

func (l *Loan) ApproveLoan(ctx context.Context, req *protogen.LoanApproveRequest) (*protogen.LoanApproveResponse, error) {
	return l.loanUseCase.Approve(ctx, req)
}

func (l *Loan) GetLoan(ctx context.Context, req *protogen.LoanGetRequest) (*protogen.LoanGetResponse, error) {
	return l.loanUseCase.Get(ctx, req)
}

func (l *Loan) AddRepayment(ctx context.Context, req *protogen.RepaymentCreateRequest) (*protogen.RepaymentCreateResponse, error) {
	return l.loanUseCase.AddRepayment(ctx, req)
}
