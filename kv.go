package storage

import "errors"

type Map map[string]interface{}

func (v Map) Put(key string, val interface{}) {
	v[key] = val
}

func (v Map) Get(key string) (interface{}, error) {
	val := v[key]
	val, ok := v[key]
	if !ok {
		return nil, errors.New("Key not found")
	}
	return val, nil
}

func (v Map) Delete(key string) {
	// checks if key exists. and deletes content only if key exists. else ignores.
	if _, ok := v[key]; ok {
		delete(v, key)
	}
}
