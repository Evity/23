package main

import "fmt"

//abstract

//abstract observer
type Listener interface {
	OnTeatcherComming()
}

type Notifier interface {
	AddListener(l Listener)
	RemoveListener(l Listener)
	Notify()
}

//concrete observer
//observer student
type StuZhang3 struct {
	BadBehavior string
}

func (s *StuZhang3) OnTeatcherComming() {
	println("StuZhang3 Stop", s.BadBehavior)

}

func (s *StuZhang3) DoBadThing() {
	fmt.Println(
		"StuZhang3 DoBadThing",
	)
}

type StuZhao4 struct {
	BadBehavior string
}

func (s *StuZhao4) OnTeatcherComming() {
	println("StuZhang3 Stop", s.BadBehavior)

}

func (s *StuZhao4) DoBadThing() {
	fmt.Println(
		"StuZhang3 DoBadThing",
	)
}

type StuWang5 struct {
	BadBehavior string
}

func (s *StuWang5) OnTeatcherComming() {
	println("StuZhang3 Stop", s.BadBehavior)

}

func (s *StuWang5) DoBadThing() {
	fmt.Println(
		"StuZhang3 DoBadThing",
	)
}

type ClassMonitor struct {
	listenerList []Listener
}

func (c *ClassMonitor) AddListener(l Listener) {
	c.listenerList = append(c.listenerList, l)
}

func (c *ClassMonitor) RemoveListener(l Listener) {
	for i, v := range c.listenerList {
		if v == l {
			c.listenerList = append(c.listenerList[:i], c.listenerList[i+1:]...)
			break
		}
	}
}

func (c *ClassMonitor) Notify() {
	for _, v := range c.listenerList {
		v.OnTeatcherComming()
	}
}

func main() {
	zhang3 := new(StuZhang3)
	zhao4 := new(StuZhao4)
	wang5 := new(StuWang5)

	monitor := new(ClassMonitor)
	monitor.AddListener(zhang3)
	monitor.AddListener(zhao4)
	monitor.AddListener(wang5)

	zhang3.DoBadThing()
	zhao4.DoBadThing()
	wang5.DoBadThing()

	monitor.Notify()
}
