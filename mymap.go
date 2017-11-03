package main

import "sync"

type SafeMap struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeMap) Set(key string, value int) {
	c.mux.Lock()
	c.v[key] = value
	c.mux.Unlock()
}

func (c *SafeMap) Value(key string) (int, bool) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	value, ok := c.v[key]
	return value, ok
}

func (c *SafeMap) Length() int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return len(c.v)
}

func NewSafeMap() *SafeMap {
	sm := &SafeMap{v: make(map[string]int)}
	return sm
}