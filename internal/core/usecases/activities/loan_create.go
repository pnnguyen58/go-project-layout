package activities

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/pnnguyen58/go-project-layout/internal/core/models"
	"github.com/pnnguyen58/go-project-layout/internal/core/ports/repositories"
	"github.com/pnnguyen58/go-project-layout/internal/utils/defined"
	protogen "github.com/pnnguyen58/go-project-layout/pkg/proto_generated"
)

func CreateLoan(ctx context.Context, req *protogen.LoanCreateRequest) (*protogen.LoanCreateResponse, error) {
	amount := float64(req.Amount)
	loan := &models.Loan{
		CustomerID:    uuid.MustParse(req.CustomerId),
		RepaymentType: string(defined.WEEKLY),
		Amount:        amount,
		Term:          req.Term,
		State:         string(defined.PENDING),
	}
	if err := loan.ValidateCreate(); err != nil {
		return nil, err
	}
	err := repositories.W.LoanRepo.Create(ctx, loan)
	if err != nil {
		return nil, err
	}
	repayments := makeRepayment(*loan)
	err = repositories.W.RepaymentRepo.CreateBatch(ctx, repayments, defined.BACTH_SIZE)
	if err != nil {
		return nil, err
	}

	data := &protogen.Loan{
		Id:            loan.ID.String(),
		CustomerId:    loan.CustomerID.String(),
		RepaymentType: loan.RepaymentType,
		Amount:        float32(loan.Amount),
		Term:          loan.Term,
		State:         loan.State,
		CreatedAt:     timestamppb.New(loan.CreatedAt),
		UpdatedAt:     timestamppb.New(loan.UpdatedAt),
	}
	for _, r := range repayments {
		data.Repayments = append(data.Repayments, &protogen.Repayment{
			Id:        r.ID.String(),
			LoanId:    r.LoanID.String(),
			Amount:    float32(r.Amount),
			State:     r.State,
			DueDate:   timestamppb.New(r.DueDate),
			CreatedAt: timestamppb.New(r.CreatedAt),
		})
	}
	return &protogen.LoanCreateResponse{
		Data: data,
	}, nil
}

func CreateLoanCompensation(ctx context.Context, req *protogen.LoanCreateRequest) (*protogen.LoanCreateResponse, error) {
	// TODO: implement create loan compensate
	return nil, nil
}
