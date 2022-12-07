package main

import "sync"

//thread safe mode

var lock sync.Mutex

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		//lock has huge impact on the performance
		if instance == nil {
			instance = new(singleton)
		}
	}
	return instance
}

func (s *singleton) DoSomething() {
	println("Do something")
}

func main() {
	s := GetInstance()
	s.DoSomething()
}
