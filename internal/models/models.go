package models

import "time"

type Status int

const (
	INVALID Status = iota
	IN_PROGRESS
	PENDING
	SCHEDULED
	CANCELLED
	COMPLETED
)

type Sequence struct {
	ID                   string     `json:"id"`
	Name                 string     `json:"name,omitempty"`
	OpenTrackingEnabled  bool       `json:"openTrackingEnabled,omitempty"`
	ClickTrackingEnabled bool       `json:"clickTrackingEnabled,omitempty"`
	StepsInterval        int        `json:"stepsInterval,omitempty"`
	Mailboxes            []*Mailbox `json:"mailboxes,omitempty"`
	Steps                []*Step    `json:"steps,omitempty"`
}

type Step struct {
	ID           string `json:"id,omitempty"`
	SequenceID   string `json:"sequenceId,omitempty"`
	EmailSubject string `json:"emailSubject,omitempty"`
	EmailContent string `json:"emailContent,omitempty"`
}

type Mailbox struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	MaxLimit int    `json:"maxLimit"`
}

type SequenceMailbox struct {
	SequenceID string `json:"sequence_id"`
	MailboxID  string `json:"mailbox_id"`
}

// To-Do: Implement the logic for creating a schedule for step execution
type Schedule struct {
	ID         string
	SequenceID string
	StartDate  time.Time
	Status     Status
}

// To-Do: Implement the logic. Each task is a step to be executed
type Task struct {
	ID         string
	ScheduleID string
	StepID     string
	ExecuteAt  time.Time
	Status     Status
}
