package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	dicitonary := Dictionary{"test": "this is just a test"}

	t.Run("kown word", func(t *testing.T) {
		got, err := dicitonary.Search("test")
		assert.NoError(t, err)
		assert.Equal(t, got, "this is just a test")
	})

	t.Run("unkown word", func(t *testing.T) {
		got, err := dicitonary.Search("unkown")
		assert.ErrorIs(t, err, ErrNotFound)
		assert.Equal(t, got, "")
	})

}
