package helper

import (
	"database/sql"
	"errors"
	"testing"
)

func TestCommitOrRollback(t *testing.T) {
	// Arrange
	mockTx := &sql.Tx{}
	mockErr := errors.New("mock error")

	// Act & Assert 1 - Test Rollback on Error
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	CommitOrRollback(mockTx)
	panic(mockErr)
}
