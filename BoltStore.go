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

//SaveKey puts key/val in boltdb
func (e *BoltStore) SaveKey(key string, val []byte) error {
	return e.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(e.bucket)
		return b.Put([]byte(key), val)
	})
}

// DeleteKeys is actually DeleteBucket when in BoltDB
func (e *BoltStore) DeleteKeys(prefix string) error {
	return e.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(prefix))
	})
}
