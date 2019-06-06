package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorer(t *testing.T) {
	m := Map{}

	t.Run("Should insert key value ", func(t *testing.T) {
		m.Put("existing-key", 1)
		val, err := m.Get("existing-key")
		assert.Equal(t, val, 1)
		assert.Nil(t, err)
	})

	t.Run("Should return error for non existing key", func(t *testing.T) {
		v, err := m.Get("non-existing-key")
		assert.NotEmpty(t, err)
		assert.Nil(t, v)
	})

	t.Run("Should delete existing key", func(t *testing.T) {
		m.Delete("existing-key")
		val, err := m.Get("existing-key")
		assert.Nil(t, val)
		assert.NotEmpty(t, err)
	})
}
