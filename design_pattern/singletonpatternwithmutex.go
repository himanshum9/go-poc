package main

import "sync"

type SingletonObj struct {
	intValue int
}

var obj *SingletonObj

var mutex sync.Mutex

func GetInstanceWithMutex(value int) *SingletonObj {
	if obj == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if obj == nil {
			obj = &SingletonObj{intValue: value}
			return obj
		}
	}
	return obj
}
