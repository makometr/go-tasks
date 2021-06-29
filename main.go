package main

import "go-tasks/task17"

func main() {
	var v interface{} = make(chan float32, 0)
	task17.Task17(v)
}
