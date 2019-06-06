package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBolt(t *testing.T) {
	b := BoltStore{}

	t.Run("Should insert key value ", func(t *testing.T) {
		b.Put("existing-key", []byte("1"))
		val, err := b.Get("existing-key")
		assert.Equal(t, val, 1)
		assert.Nil(t, err)
	})

	t.Run("Should return error for non existing key", func(t *testing.T) {
		v, err := b.Get("non-existing-key")
		assert.NotEmpty(t, err)
		assert.Nil(t, v)
	})

	t.Run("Should delete existing key", func(t *testing.T) {
		b.Delete("existing-key")
		val, err := b.Get("existing-key")
		assert.Nil(t, val)
		assert.NotEmpty(t, err)
	})
}
