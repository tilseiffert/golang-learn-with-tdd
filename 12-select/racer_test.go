package selectlearn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.quii.dev"
	fastURL := "http://tilseiffert.de/linktree"

	assert.Equal(t, fastURL, Racer(slowURL, fastURL))
}
