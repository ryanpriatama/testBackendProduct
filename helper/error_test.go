package helper

import (
	"errors"
	"testing"
)

func TestPanicIfError(t *testing.T) {
	t.Run("Panic if error is not nil", func(t *testing.T) {
		err := errors.New("some error")
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic as expected")
			}
		}()
		PanicIfError(err)
	})

	t.Run("Does not panic if error is nil", func(t *testing.T) {
		PanicIfError(nil)
	})
}
