package types

type (
	CreateSequenceRequest struct {
		Name                 string
		OpenTrackingEnabled  bool
		ClickTrackingEnabled bool
		StepsInterval        int32
	}

	UpdateSequenceRequest struct {
		ID                   string
		Name                 *string
		OpenTrackingEnabled  *bool
		ClickTrackingEnabled *bool
		StepsInterval        *int
	}
)
