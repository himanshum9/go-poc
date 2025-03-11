package main

import "fmt"

type HashMap struct {
	data map[string]int
}

func NewHashMap() *HashMap {
	return &HashMap{data: make(map[string]int)}
}

func (hmap *HashMap) Insert(key string, val int) {
	hmap.data[key] = val
}

func (hmap *HashMap) Display(key string) int {
	return hmap.data[key]
}

func (hmap *HashMap) DisplayAll() {
	for key, val := range hmap.data {
		fmt.Println(key, val)
	}

}

func main() {
	hmap := NewHashMap()
	hmap.Insert("a", 1)
	hmap.Insert("b", 2)
	hmap.Insert("c", 3)
	hmap.Insert("a", 5)
	val := hmap.Display("a")
	fmt.Println(val)
	hmap.DisplayAll()
}
