package types

type (
	CreateStepRequest struct {
		SequenceID   string
		EmailSubject string
		EmailContent string
	}
	UpdateStepRequest struct {
		ID           string
		EmailSubject *string
		EmailContent *string
	}
)
