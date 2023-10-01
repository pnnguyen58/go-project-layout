package activities

import (
	"context"
	"github.com/google/uuid"

	"github.com/pnnguyen58/go-project-layout/internal/core/ports/repositories"
	"github.com/pnnguyen58/go-project-layout/internal/utils/defined"
	protogen "github.com/pnnguyen58/go-project-layout/pkg/proto_generated"
)

func ApproveLoan(ctx context.Context, req *protogen.LoanApproveRequest) (*protogen.LoanApproveResponse, error) {
	err := repositories.W.LoanRepo.UpdateState(ctx, uuid.MustParse(req.Id), defined.APPROVED)
	if err != nil {
		return nil, err
	}
	err = repositories.W.RepaymentRepo.UpdateStateByLoanID(ctx, uuid.MustParse(req.Id), defined.APPROVED)
	if err != nil {
		return nil, err
	}
	return &protogen.LoanApproveResponse{
		Data: req.Id,
	}, nil
}

func ApproveLoanCompensation(ctx context.Context, req *protogen.LoanApproveRequest) (*protogen.LoanApproveResponse, error) {
	// TODO: implement approve loan compensate
	return nil, nil
}
