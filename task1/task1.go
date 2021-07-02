package task1

import "fmt"

type Action struct {
	name string
}

func NewAction(name string) *Action {
	return &Action{name: name}
}

func (a Action) DoWork() {
	fmt.Println("Do", a.name)
}

type Human struct {
	action Action
}

func (h Human) DoAction() {
	fmt.Print("Human:")
	h.action.DoWork()
}

// Task1 is task
func Task1() {
	man := Human{action: *NewAction("varit borsh")}
	man.DoAction()
}
