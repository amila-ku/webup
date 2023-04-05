package bootstrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsCommandAvailable(t *testing.T) {
	assert.Equal(t, commandExists("ls"), true)
	assert.Equal(t, commandExists("lsl"), false)
}
