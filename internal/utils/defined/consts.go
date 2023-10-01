package defined

const (
	SCHEDULED State = "SCHEDULED"
	APPROVED  State = "APPROVED"
	PENDING   State = "PENDING"
	PAID      State = "PAID"
)

const (
	DAILY    RepaymentType = "daily"
	WEEKLY   RepaymentType = "weekly"
	MONTHLY  RepaymentType = "monthly"
	ANNUALLY RepaymentType = "annually"
)

const (
	PRECISION  = 2
	BACTH_SIZE = 10
)

const (
	ERROR_DATABASE_EXCEPTION = "DatabaseException"
)
