package activities

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/pnnguyen58/go-project-layout/internal/core/models"
	"github.com/pnnguyen58/go-project-layout/internal/core/ports/repositories"
	"github.com/pnnguyen58/go-project-layout/internal/utils/defined"
	protogen "github.com/pnnguyen58/go-project-layout/pkg/proto_generated"
)

func CreateRepayment(ctx context.Context, req *protogen.RepaymentCreateRequest,
) (*protogen.RepaymentCreateResponse, error) {
	amount := roundAmount(float64(req.Amount))
	unpaidRepayments, err := repositories.W.RepaymentRepo.GetUnpaidByLoanID(ctx, uuid.MustParse(req.LoanId))
	if err != nil {
		return nil, err
	}
	if len(unpaidRepayments) == 0 {
		return nil, fmt.Errorf("no unpaid repayment")
	}
	repayment := &models.Repayment{}
	for _, r := range unpaidRepayments {
		if amount >= r.Amount {
			repayment = r
			break
		}
	}
	if repayment.ID == uuid.Nil {
		return nil, fmt.Errorf("unacceptable repayment amount")
	}

	err = repositories.W.RepaymentRepo.UpdateState(ctx, repayment.ID, defined.PAID)
	if err != nil {
		return nil, err
	}
	log := &models.RepaymentLog{
		LoanID: uuid.MustParse(req.LoanId),
		Amount: amount,
	}
	_ = repositories.W.RepaymentRepo.CreateLog(ctx, log)

	// TODO check repayment amount exceed
	if amount > repayment.Amount {
		prepay := &models.RepaymentPrepay{
			LoanID: repayment.LoanID,
			Amount: amount - repayment.Amount,
		}
		_ = repositories.W.RepaymentRepo.CreatePrepay(ctx, prepay)
	}
	// Update loan to PAID when all repayment paid
	if len(unpaidRepayments) == 1 {
		err = repositories.W.LoanRepo.UpdateState(ctx, repayment.LoanID, defined.PAID)
		if err != nil {
			return nil, err
		}
	}
	return &protogen.RepaymentCreateResponse{
		Data: &protogen.Repayment{
			Id:        repayment.ID.String(),
			LoanId:    repayment.LoanID.String(),
			Amount:    float32(repayment.Amount),
			State:     string(defined.PAID),
			DueDate:   timestamppb.New(repayment.DueDate),
			CreatedAt: timestamppb.New(repayment.CreatedAt),
			UpdatedAt: timestamppb.New(repayment.UpdatedAt),
		},
	}, nil
}

func CreateRepaymentCompensation(ctx context.Context, req *protogen.RepaymentCreateRequest,
) (*protogen.RepaymentCreateResponse, error) {
	return nil, nil
}
