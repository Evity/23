package main

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
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
