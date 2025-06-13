package main

import "fmt"

type MyError struct {
	Msg string
}

func (m MyError) Error() string {
	return m.Msg
}

func doSomething() error {
	return MyError{Msg: "Something went horribly wrong"}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
	}
}
