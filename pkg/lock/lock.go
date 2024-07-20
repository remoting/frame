package lock

import "sync"

var _keyLock = NewKeyLock()

type KeyLock struct {
	mutexes sync.Map
}

func NewKeyLock() *KeyLock {
	return &KeyLock{}
}
func Lock(key string) func() {
	return _keyLock.Lock(key)
}
func (m *KeyLock) Lock(key string) func() {
	value, _ := m.mutexes.LoadOrStore(key, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()

	return func() {
		mtx.Unlock()
		m.mutexes.Delete(key)
	}
}
