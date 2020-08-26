package set

import "sync"

type void struct{}

// String struct for dync map
type String struct {
	d map[string]void
	m sync.RWMutex
}

// Load for load data by key provided
func (m *String) Load(key string) bool {
	if key == "" {
		return false
	}
	if m.d == nil {
		m.d = map[string]void{}
		return false
	}
	m.m.RLock()
	defer m.m.RUnlock()
	_, ok := m.d[key]
	return ok
}

// Store for load data by key provided
func (m *String) Store(key string) {
	if key == "" {
		return
	}
	if m.d == nil {
		m.d = map[string]void{}
	}
	m.m.Lock()
	defer m.m.Unlock()
	m.d[key] = void{}
}

// Delete for delete data by key provided
func (m *String) Delete(key string) {
	if key == "" {
		return
	}
	if m.d == nil {
		m.d = map[string]void{}
		return
	}
	m.m.Lock()
	defer m.m.Unlock()
	delete(m.d, key)
}
