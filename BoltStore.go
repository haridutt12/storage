package storage

import (
	bolt "github.com/coreos/bbolt"
)

//BoltStore is an in-memory implementation of the Storage driver
type BoltStore struct {
	db     *bolt.DB
	bucket []byte
}

// Get returns the value of a Key
func (e *BoltStore) Get(key string) ([]byte, error) {
	v := []byte{}
	err := e.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(e.bucket)
		vb := b.Get([]byte(key))
		if len(vb) != 0 {
			v = vb
		}
		return nil
	})

	return v, err
}

//Put  puts key/val in boltdb
func (e *BoltStore) Put(key string, val []byte) {
	_ := e.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(e.bucket)
		_ = b.Put([]byte(key), val)
	})
}

// Delete is actually DeleteBucket when in BoltDB
func (e *BoltStore) Delete(prefix string) {
	_ := e.db.Update(func(tx *bolt.Tx) error {
		_ = tx.DeleteBucket([]byte(prefix))
	})
}
