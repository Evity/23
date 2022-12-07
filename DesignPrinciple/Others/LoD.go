package main

import "fmt"

//Law of Demeter
/*
Each unit should have only limited knowledge about other units:
 only units "closely" related to the current unit. (Black box)
*/

type Task struct{}

var tasks []*Task

func NewTask(t *Task) {
	tasks = append(tasks, t)
}

func LoadTask() *Task {
	return &Task{}
}

type PMOnWork interface {
	FinishTheJob()
}

type Dev struct {
}

func NewDev() *Dev {
	return &Dev{}
}

type PM struct {
}

/* PM not care about TASK, but let DEV to handle the task
Dev working on task, meanwhile the task is transparent to others*/

func (p *PM) FinishTheJob() {
	dev := NewDev()
	dev.finishTheJob()
}

func (d *Dev) finishTheJob() {
	task := LoadTask()
	fmt.Printf("Task solving:%v", task)
}

func main() {
	Work(&PM{})
}

func Work(p PMOnWork) {
	p.FinishTheJob()
}
