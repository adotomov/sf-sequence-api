package db

import (
	_ "embed"
)

//go:embed queries/sequences.sql
var GetSequencesQuery string

//go:embed queries/get_sequence.sql
var GetSequenceByIdQuery string

//go:embed queries/get_sequence_mailboxes.sql
var GetSequenceMailboxesQuery string

//go:embed queries/get_sequence_steps.sql
var GetSequenceStepsQuery string

// go:embed queries/create_step.sql
var CreateStepQuery string
