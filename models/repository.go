package models

import (
	"ecommerce/shared"
	"sync"
	"time"
)

type Repository interface {
	SaveProduct(*ProductModel) error
	GetProductById(uint64) (*ProductModel, error)
	GetProducts() ([]*ProductModel, error)
	GetProductsByCat(string) ([]*ProductModel, error)
}

// implements a simple memory based repository,
// It could be postgres, mongo, file, redis etc.
type MemoryRepository struct {
	_lock *sync.RWMutex
	// this is counter, only going up
	// and is used for setting a new
	// product's ID
	_idncr     uint64
	products   map[uint64]*ProductModel
	categories map[string]map[uint64]bool
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		_lock:      &sync.RWMutex{},
		_idncr:     0,
		products:   make(map[uint64]*ProductModel),
		categories: make(map[string]map[uint64]bool),
	}
}

func (mr *MemoryRepository) SaveProduct(p *ProductModel) error {
	mr._lock.Lock()
	defer mr._lock.Unlock()
	if v, ok := mr.products[p.ID]; p.ID != 0 && !ok {
		mr._idncr++
		p.ID = mr._idncr
		p.CreatedAt = time.Now().Unix()
		p.UpdatedAt = p.CreatedAt
		if _, okk := mr.categories[p.Category]; !okk {
			mr.categories[p.Category] = make(map[uint64]bool)
		}
		mr.categories[p.Category][p.ID] = true
	} else {
		p.ID = v.ID
		p.UpdatedAt = time.Now().Unix()
		p.CreatedAt = v.CreatedAt
		if p.Category != v.Category {
			if _, okk := mr.categories[p.Category]; !okk {
				mr.categories[p.Category] = make(map[uint64]bool)
			}
			mr.categories[p.Category][p.ID] = true
			delete(mr.categories[v.Category], p.ID)
		}
	}
	mr.products[p.ID] = p.Clone()
	return nil
}

func (mr *MemoryRepository) GetProductById(id uint64) (*ProductModel, error) {
	mr._lock.RLock()
	defer mr._lock.RUnlock()
	if p, ok := mr.products[id]; !ok {
		return nil, shared.ErrNotFound
	} else {
		return p.Clone(), nil
	}
}

func (mr *MemoryRepository) GetProducts() ([]*ProductModel, error) {
	mr._lock.RLock()
	defer mr._lock.RUnlock()

	// we can use some sort of short term caching
	// to avoid copying memory each time, but due
	// to demonstration-type of this project, we skip
	// such a step (and it is not even a concern of low level
	// repository)
	var list = make([]*ProductModel, 0)

	for _, v := range mr.products {
		list = append(list, v.Clone())
	}

	return list, nil
}

func (mr *MemoryRepository) GetProductsByCat(cat string) ([]*ProductModel, error) {
	mr._lock.RLock()
	defer mr._lock.RUnlock()

	// we can use some sort of short term caching
	// to avoid copying memory each time, but due
	// to demonstration-type of this project, we skip
	// such a step (and it is not even a concern of low level
	// repository)
	var list = make([]*ProductModel, 0)

	for pId := range mr.categories[cat] {
		list = append(list, mr.products[pId])
	}

	return list, nil
}
