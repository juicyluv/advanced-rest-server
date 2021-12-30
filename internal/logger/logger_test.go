package logger_test

import (
	"testing"

	"github.com/juicyluv/advanced-rest-server/internal/logger"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	l := logger.New(logger.LevelDebug)
	assert.Equal(t, logger.LevelDebug, l.Level(), "logger level should be Debug")
}
