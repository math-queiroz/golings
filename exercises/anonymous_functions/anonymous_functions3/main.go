// anonymous functions3
// Make me compile!

package main

import "fmt"

func updateStatus() func() string {
	var index int
	orderStatus := map[int]string{
		1: "TO DO",
		2: "DOING",
		3: "DONE",
	}

	return func() string {
		index++
		return orderStatus[index+1]
	}
}

func main() {
	anonymous_func := updateStatus()
	var status string

	status = anonymous_func()
	status = anonymous_func()

	if status == "DONE" {
		fmt.Println("Good Job!")
	} else {
		panic("To complete the challenge the status must be DONE")
	}
}
