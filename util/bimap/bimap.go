package bimap

import (
	"sync"
)

//example:
//	a := NewBiMap()
//	a.Set(1, "hello")
//	a.Set(2, "world")
//	fmt.Println(a.Get(1))
//	fmt.Println(a.Get(2))

type BiMap struct {
	lock *sync.RWMutex
	kmap map[interface{}]interface{}
	vmap map[interface{}]interface{}
}

func NewBiMap() *BiMap {
	return &BiMap{
		lock: &sync.RWMutex{},
		kmap: make(map[interface{}]interface{}),
		vmap: make(map[interface{}]interface{}),
	}
}

func (b *BiMap) Set(k interface{}, v interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.kmap[k] = v
	b.vmap[v] = k
}

func (b *BiMap) ValueExists(v interface{}) bool {
	b.lock.RLock()
	defer b.lock.RUnlock()
	_, ok := b.vmap[v]
	return ok
}

func (b *BiMap) Get(k interface{}) (interface{}, bool) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	value, ok := b.kmap[k]
	return value, ok
}

func (b *BiMap) GetKey(v interface{}) (interface{}, bool) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	key, ok := b.vmap[v]
	return key, ok
}

func (b *BiMap) Pop(k interface{}) (interface{}, bool) {
	v, ok := b.Get(k)
	b.lock.Lock()
	defer b.lock.Unlock()
	if ok {
		delete(b.kmap, k)
		delete(b.vmap, v)
		return v, true
	}
	return nil, false
}

func (b *BiMap) PopValue(v interface{}) (interface{}, bool) {
	k, ok := b.GetKey(v)
	b.lock.Lock()
	defer b.lock.Unlock()
	if ok {
		delete(b.kmap, k)
		delete(b.vmap, v)
		return k, true
	}
	return nil, false
}

func (b *BiMap) Len() int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return len(b.kmap)
}

func (b *BiMap) Keys() []interface{} {
	b.lock.RLock()
	defer b.lock.RUnlock()
	keys := make([]interface{}, 0)
	for key, _ := range b.kmap {
		keys = append(keys, key)
	}
	return keys
}

func (b *BiMap) Values() []interface{} {
	b.lock.RLock()
	defer b.lock.RUnlock()
	values := make([]interface{}, 0)
	for value, _ := range b.vmap {
		values = append(values, value)
	}
	return values
}
