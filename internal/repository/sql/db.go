package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/adotomov/sf-sequence-api/internal/models"
	"github.com/adotomov/sf-sequence-api/internal/types"
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	_ "github.com/lib/pq"
)

type DB struct {
	dbrConn *dbr.Connection
}

func NewDB(db *sql.DB) (*DB, error) {
	dbrConn := &dbr.Connection{
		DB:      db,
		Dialect: dialect.PostgreSQL,
	}

	d := &DB{
		dbrConn: dbrConn,
	}

	return d, nil
}

func (d *DB) GetSequences(ctx context.Context) ([]*models.Sequence, error) {
	var result []*models.Sequence

	sess := d.dbrConn.NewSession(nil)
	rows, err := sess.Query(GetSequencesQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id                   string
			name                 string
			openTrackingEnabled  bool
			clickTrackingEnabled bool
			stepsInterval        int32
		)

		var sq models.Sequence

		err = rows.Scan(
			&id,
			&name,
			&openTrackingEnabled,
			&clickTrackingEnabled,
			&stepsInterval,
		)

		if err != nil {
			log.Print("Error reading row")
		}

		sq.ID = id
		sq.Name = name
		sq.OpenTrackingEnabled = openTrackingEnabled
		sq.ClickTrackingEnabled = clickTrackingEnabled
		sq.StepsInterval = int(stepsInterval)
		sq.Mailboxes = d.getSequenceMailboxes(id)
		sq.Steps = d.getSequenceSteps(id)

		result = append(result, &sq)
	}

	return result, nil
}

func (d *DB) CreateSequence(ctx context.Context, rb types.CreateSequenceRequest) (*models.Sequence, error) {
	var id string

	sess := d.dbrConn.NewSession(nil)
	err := sess.QueryRow(`INSERT INTO sequence (SEQUENCE_NAME, OPEN_TRACKING_ENABLED, CLICK_TRACKING_ENABLED, STEPS_INTERVAL) VALUES ($1, $2, $3, $4) RETURNING ID`, rb.Name, rb.OpenTrackingEnabled, rb.ClickTrackingEnabled, rb.StepsInterval).Scan(&id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &models.Sequence{
		ID: id,
	}, nil
}

func (d *DB) GetSequenceByID(ctx context.Context, sid string) (*models.Sequence, error) {
	var (
		id                   string
		name                 string
		openTrackingEnabled  bool
		clickTrackingEnabled bool
		stepsInterval        int32
	)

	var sq models.Sequence

	sess := d.dbrConn.NewSession(nil)
	err := sess.QueryRowContext(ctx, GetSequenceByIdQuery, sid).Scan(
		&id,
		&name,
		&openTrackingEnabled,
		&clickTrackingEnabled,
		&stepsInterval,
	)

	if err != nil {
		return nil, err
	}

	sq.ID = id
	sq.Name = name
	sq.OpenTrackingEnabled = openTrackingEnabled
	sq.ClickTrackingEnabled = clickTrackingEnabled
	sq.StepsInterval = int(stepsInterval)
	sq.Mailboxes = d.getSequenceMailboxes(id)
	sq.Steps = d.getSequenceSteps(sid)

	return &sq, nil
}

func (d *DB) UpdateSequence(ctx context.Context, rb types.UpdateSequenceRequest) (*models.Sequence, error) {
	sess := d.dbrConn.NewSession(nil)
	_, err := sess.ExecContext(ctx,
		`UPDATE sequence SET 
		SEQUENCE_NAME = COALESCE($2, SEQUENCE_NAME), OPEN_TRACKING_ENABLED = COALESCE($3, OPEN_TRACKING_ENABLED), CLICK_TRACKING_ENABLED = COALESCE($4, CLICK_TRACKING_ENABLED), STEPS_INTERVAL = COALESCE($5, STEPS_INTERVAL) WHERE sequence.ID = $1`,
		rb.ID, rb.Name, rb.OpenTrackingEnabled, rb.ClickTrackingEnabled, rb.StepsInterval)

	if err != nil {
		return nil, err
	}

	return &models.Sequence{
		ID: rb.ID,
	}, nil
}

func (d *DB) CreateStep(ctx context.Context, rb types.CreateStepRequest) (*models.Step, error) {
	var id string

	sess := d.dbrConn.NewSession(nil)
	err := sess.QueryRowContext(ctx, `INSERT INTO step (SEQUENCE_ID, EMAIL_SUBJECT, EMAIL_CONTENT) VALUES ($1, $2, $3) RETURNING ID`, rb.SequenceID, rb.EmailSubject, rb.EmailContent).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &models.Step{
		ID: id,
	}, nil
}

func (d *DB) UpdateStep(ctx context.Context, rb types.UpdateStepRequest) (*models.Step, error) {
	sess := d.dbrConn.NewSession(nil)
	_, err := sess.ExecContext(ctx, `UPDATE step SET EMAIL_SUBJECT = COALESCE($2, EMAIL_SUBJECT), EMAIL_CONTENT = COALESCE($3, EMAIL_CONTENT) WHERE step.ID = $1`, rb.ID, rb.EmailSubject, rb.EmailContent)

	if err != nil {
		return nil, err
	}

	return &models.Step{
		ID: rb.ID,
	}, nil
}

func (d *DB) DeleteStep(ctx context.Context, sid string) error {
	sess := d.dbrConn.NewSession(nil)
	_, err := sess.ExecContext(ctx, `DELETE FROM step WHERE ID = $1`, sid)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) getSequenceMailboxes(sid string) []*models.Mailbox {
	var result []*models.Mailbox

	sess := d.dbrConn.NewSession(nil)
	rows, err := sess.Query(GetSequenceMailboxesQuery, sid)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id       string
			email    string
			maxLimit int32
		)

		var sm models.Mailbox

		err := rows.Scan(
			&id,
			&email,
			&maxLimit,
		)
		if err != nil {
			return nil
		}

		sm.ID = id
		sm.Email = email
		sm.MaxLimit = int(maxLimit)

		result = append(result, &sm)
	}

	return result
}

func (d *DB) getSequenceSteps(sid string) []*models.Step {
	var result []*models.Step

	sess := d.dbrConn.NewSession(nil)
	rows, err := sess.Query(GetSequenceStepsQuery, sid)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id           string
			emailSubject string
			emailContent string
		)

		var st models.Step

		err := rows.Scan(
			&id,
			&emailSubject,
			&emailContent,
		)

		if err != nil {
			return nil
		}

		st.ID = id
		st.EmailContent = emailContent
		st.EmailSubject = emailSubject

		result = append(result, &st)
	}

	return result
}
