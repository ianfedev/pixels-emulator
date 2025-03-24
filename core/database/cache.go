package database

import (
	"context"
	"github.com/go-gorm/caches/v4"
	"sync"
)

type MemoryCache struct {
	store *sync.Map
}

func (c *MemoryCache) Init() {
	if c.store == nil {
		c.store = &sync.Map{}
	}
}

func (c *MemoryCache) Get(_ context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
	c.Init()
	val, ok := c.store.Load(key)
	if !ok {
		return nil, nil
	}

	if err := q.Unmarshal(val.([]byte)); err != nil {
		return nil, err
	}

	return q, nil
}

func (c *MemoryCache) Store(_ context.Context, key string, val *caches.Query[any]) error {
	c.Init()
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	c.store.Store(key, res)
	return nil
}

func (c *MemoryCache) Invalidate(_ context.Context) error {
	c.store = &sync.Map{}
	return nil
}
