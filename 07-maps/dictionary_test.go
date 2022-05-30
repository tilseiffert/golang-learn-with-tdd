package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	key := "test"
	value := "this is just a test"
	dicitonary := Dictionary{key: value}

	t.Run("kown word", func(t *testing.T) {
		got, err := dicitonary.Search(key)
		assert.NoError(t, err)
		assert.Equal(t, got, value)
	})

	t.Run("unkown word", func(t *testing.T) {
		got, err := dicitonary.Search("unkown")
		assert.ErrorIs(t, err, ErrNotFound)
		assert.Equal(t, got, "")
	})

}

func TestAdd(t *testing.T) {
	dicitonary := Dictionary{}
	key := "test"
	value := "this is just a test"

	dicitonary.Add(key, value)

	got, err := dicitonary.Search(key)
	assert.NoError(t, err)
	assert.Equal(t, got, value)
}
