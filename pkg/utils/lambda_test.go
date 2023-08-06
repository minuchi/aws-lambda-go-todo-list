package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLambda(t *testing.T) {
	t.Run("should return true if LAMBDA_TASK_ROOT is set", func(t *testing.T) {
		os.Setenv("LAMBDA_TASK_ROOT", "some value")
		assert.True(t, IsLambda())
	})
	t.Run("should return false if LAMBDA_TASK_ROOT is not set", func(t *testing.T) {
		os.Unsetenv("LAMBDA_TASK_ROOT")
		assert.False(t, IsLambda())
	})
}
