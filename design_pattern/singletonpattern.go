package main

import "sync"

type Singleton struct {
	value int
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance(initialValue int) *Singleton {
	once.Do(func() {
		instance = &Singleton{value: initialValue}
	})
	return instance
}
