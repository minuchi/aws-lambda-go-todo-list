package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckError(t *testing.T) {
	t.Run("should panic if error is not nil", func(t *testing.T) {
		assert.PanicsWithError(t, "some error", func() {
			CheckError(errors.New("some error"))
		})
	})

	t.Run("should not panic if error is nil", func(t *testing.T) {
		assert.NotPanics(t, func() {
			CheckError(nil)
		})
	})
}
