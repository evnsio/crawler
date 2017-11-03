package main

import "sync"

type SafeMap struct {
	v   map[string]int
	mux sync.Mutex
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mux.Lock()
	sm.v[key] = value
	sm.mux.Unlock()
}

func (sm *SafeMap) Value(key string) (int, bool) {
	sm.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer sm.mux.Unlock()
	value, ok := sm.v[key]
	return value, ok
}

func NewSafeMap() *SafeMap {
	sm := &SafeMap{v: make(map[string]int)}
	return sm
}
