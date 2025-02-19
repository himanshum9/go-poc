package main

import "fmt"

// Entry represents a key-value pair
type Entry struct {
	key   string
	value any
}

// HashMap is a simple custom hashmap
type HashMap struct {
	buckets map[int][]Entry
	size    int
}

// NewHashMap initializes the HashMap
func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make(map[int][]Entry, size),
		size:    size,
	}
}

// hashFunction generates an index based on key
func (hm *HashMap) hashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % hm.size
}

// Put inserts a key-value pair
func (hm *HashMap) Put(key string, value any) {
	index := hm.hashFunction(key)
	bucket := hm.buckets[index]

	// Update if key exists
	for i, entry := range bucket {
		if entry.key == key {
			hm.buckets[index][i].value = value
			return
		}
	}

	// Insert new entry
	hm.buckets[index] = append(bucket, Entry{key, value})
}

// Get retrieves a value by key
func (hm *HashMap) Get(key string) (any, bool) {
	index := hm.hashFunction(key)
	bucket := hm.buckets[index]

	for _, entry := range bucket {
		if entry.key == key {
			return entry.value, true
		}
	}
	return nil, false
}

// Remove deletes a key-value pair
func (hm *HashMap) Remove(key string) {
	index := hm.hashFunction(key)
	bucket := hm.buckets[index]

	for i, entry := range bucket {
		if entry.key == key {
			hm.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

func main() {
	hm := NewHashMap(30)

	hm.Put("name", "John")
	hm.Put("age", "zczxcz")
	hm.Put("sds", "sds")
	hm.Put("asdsge", 30)
	hm.Put("sds", 12)

	value, found := hm.Get("name")
	if found {
		fmt.Println("Name:", value)
	}

	hm.Remove("name")

	_, found = hm.Get("name")
	fmt.Println("Found 'name' after removal:", found)
	fmt.Println(hm.buckets)
}
