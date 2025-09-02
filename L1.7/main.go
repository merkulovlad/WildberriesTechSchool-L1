package main
import (
	"sync"
)

type syncronizedMap struct {
	sync.RWMutex
	data map[string]interface{}
}

func NewSyncronizedMap() *syncronizedMap {
	return &syncronizedMap{
		data: make(map[string]interface{}),
	}
}

func (m *syncronizedMap) Get(key string) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()
	value, ok := m.data[key]
	return value, ok
}

func (m *syncronizedMap) Set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}

func main() {
	m := NewSyncronizedMap()
	m.Set("key", "value")
	value, ok := m.Get("key")
	if ok {
		println(value)
	}
}