package main

import (
	"sync"
)

// thread safe mode improved
var once sync.Once

var initialized uint32

var lock sync.Mutex

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	/* 	if atomic.LoadUint32(&initialized) == 1 {
	   		return instance
	   	}

	   	lock.Lock()
	   	defer lock.Unlock()

	   	if initialized == 0 {
	   		instance = new(singleton)
	   		atomic.StoreUint32(&initialized, 1)
	   	} */
	once.Do(func() {
		instance = new(singleton)
	})

	return instance
}

func (s *singleton) DoSomething() {
	println("Do something")
}

func main() {
	s := GetInstance()
	s.DoSomething()
}
