package repositories

type wire struct {
	LoanRepo      Loan
	RepaymentRepo Repayment
}

var W wire

func Wire(loanRepo Loan, repaymentRepo Repayment) {
	w := &W
	w.LoanRepo = loanRepo
	w.RepaymentRepo = repaymentRepo
}
