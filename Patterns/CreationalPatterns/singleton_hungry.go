package main

type singleton struct{}

var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (s *singleton) DoSomething() {
	println("Do something")
}

func main() {
	s := GetInstance()
	s.DoSomething()
}
