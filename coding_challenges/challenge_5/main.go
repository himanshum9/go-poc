// Challenge 5 · Custom Map API
//Prompt: Implement a map-like data structure in Go that provides get() and put() semantics. Decide how to address collisions, concurrency, and zero-value handling.

package main

import (
	"fmt"
	"sync"
)

type HashMap struct {
	lock sync.RWMutex
	hmap map[string]int
}

func (h *HashMap) Get(v string) (int, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	val, ok := h.hmap[v]
	if !ok {
		return 0, false
	}
	return val, true
}

func (h *HashMap) Set(k string, v int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.hmap[k] = v
}

func main() {
	hmap := &HashMap{hmap: make(map[string]int)}
	hmap.Set("a", 100)
	hmap.Set("b", 0)

	v, ok := hmap.Get("b") // returns 0, true
	fmt.Println(v, ok)
	v, ok2 := hmap.Get("x")
	fmt.Println(v, ok2)

}
