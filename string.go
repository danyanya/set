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

// Keys for return all keys in set
func (m *String) Keys() []string {
	if m.d == nil {
		return []string{}
	}

	m.m.Lock()
	defer m.m.Unlock()

	keys := make([]string, 0, len(m.d))
	for k := range m.d {
		keys = append(keys, k)
	}
	return keys
}
