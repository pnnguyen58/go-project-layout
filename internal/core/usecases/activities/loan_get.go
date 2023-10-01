package activities

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/pnnguyen58/go-project-layout/internal/core/ports/repositories"
	protogen "github.com/pnnguyen58/go-project-layout/pkg/proto_generated"
)

func GetLoan(ctx context.Context, req *protogen.LoanGetRequest) (*protogen.LoanGetResponse, error) {
	loan, err := repositories.W.LoanRepo.GetByID(ctx, uuid.MustParse(req.Id))
	if err != nil {
		return nil, err
	}
	repayments, err := repositories.W.RepaymentRepo.GetByLoanID(ctx, uuid.MustParse(req.Id))
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
			UpdatedAt: timestamppb.New(r.UpdatedAt),
		})
	}
	return &protogen.LoanGetResponse{
		Data: data,
	}, nil
}

func GetLoanCompensation(ctx context.Context, req *protogen.LoanGetRequest) (*protogen.LoanGetResponse, error) {
	// TODO: implement approve loan compensate
	return nil, nil
}
