package module

import (
	// "fmt"
	"fmt"
	"sync"
)

type ThreadSafeStore interface {
	Add(key string, obj interface{})
	Get(key string) (item interface{}, exists bool)
	List() []interface{}
}

type threadSafeMap struct {
	lock  sync.RWMutex
	items map[string]interface{}
}

func NewThreadSafeStore() ThreadSafeStore {
	return &threadSafeMap{
		items:    map[string]interface{}{},
	}
}

func (c *threadSafeMap) Add(key string, obj interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items[key] = obj
}

func (c *threadSafeMap) Get(key string) (item interface{}, exists bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	item, exists = c.items[key]
	return item, exists
}

func (c *threadSafeMap) List() []interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	list := make([]interface{}, 0, len(c.items))
	for _, item := range c.items {
		list = append(list, item)
	}
	return list
}

func Store() {
	// 初期化
	store := NewThreadSafeStore()
	// 追加
	a := []int{1, 23445}
	store.Add("key", a)
	g, ok := store.Get("s")
	fmt.Println(g, ok)
}
