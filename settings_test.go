package ip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.Equal(t, "0.0.1", VERSION)
}

func TestMaxTries(t *testing.T) {
	assert.Equal(t, 3, MAX_TRIES)
}
