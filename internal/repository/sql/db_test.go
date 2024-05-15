package db

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetSequences(t *testing.T) {
	var testSequenceId string = "899ef688-b7dc-4833-91e6-9a8d74d28277"
	testRows := sqlmock.NewRows([]string{"SEQUENCE"})
	tests := map[string]struct {
		mockDB     func() (*sql.DB, sqlmock.Sqlmock)
		entityType string
		err        error
	}{
		"OK Sequences": {
			mockDB: func() (*sql.DB, sqlmock.Sqlmock) {
				mockDB, mock, _ := sqlmock.New()
				testRows1 := testRows.AddRow(testSequenceId)
				mock.ExpectQuery(regexp.QuoteMeta(GetSequencesQuery)).WillReturnRows(testRows1)
				return mockDB, mock
			},
		},
	}
	ctx := context.Background()

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			mockDB, mock := tt.mockDB()
			db, _ := NewDB(mockDB)
			got, err := db.GetSequences(ctx)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
			if tt.err != nil {
				require.Error(t, err)
				require.Equal(t, testSequenceId, got)
			}
		})
	}
}

// TO-DO: Complete the test suite
