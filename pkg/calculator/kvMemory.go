package calculator

import (
	"errors"
)

// ErrNotFound is returned when an nonexisting key is requested
var ErrNotFound = errors.New("Key not found")

var counter = 1

// KVStorage describes basic functions required for KV storage
type KVStorage interface {
	Add(float64) error
	Get(int) (float64, error)
	Remove(int) error
	Lenght() int
}

// KV hold all the data in memory
type KV struct {
	data map[int]float64
}

// NewMemoryStorage returns an initialized instance of KV
func NewMemoryStorage() *KV {
	return &KV{map[int]float64{}}
}

// Add writes a given value under the given key
func (kv *KV) Add(v float64) error {
	kv.data[counter] = v
	counter++
	return nil
}

// Get returns value linked to given key. Returns ErrNotFound when key does not
// exist.
func (kv *KV) Get(k int) (float64, error) {
	if v, ok := kv.data[k]; ok {
		return v, nil
	}

	return 0, ErrNotFound
}

// Remove deletes the given key from the map. Returns ErrNotFound if key does
// not exist.
func (kv *KV) Remove(k int) error {
	if _, ok := kv.data[k]; !ok {
		return ErrNotFound
	}

	delete(kv.data, k)
	counter--
	return nil
}

// Lenght function returns the counter
func (kv *KV) Lenght() int {
	return len(kv.data)
}
