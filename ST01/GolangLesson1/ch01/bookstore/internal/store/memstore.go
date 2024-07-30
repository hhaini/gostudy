package store

import (
	mystore "bookstore/store"
	factory "bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{books: make(map[string]*mystore.Book)})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

// Create implements store.Store.
func (m *MemStore) Create(*mystore.Book) error {
	panic("unimplemented")
}

// Delete implements store.Store.
func (m *MemStore) Delete(string) error {
	panic("unimplemented")
}

// Get implements store.Store.
func (m *MemStore) Get(string) (mystore.Book, error) {
	panic("unimplemented")
}

// GetAll implements store.Store.
func (m *MemStore) GetAll() ([]mystore.Book, error) {
	panic("unimplemented")
}

// Update implements store.Store.
func (m *MemStore) Update(*mystore.Book) error {
	panic("unimplemented")
}
