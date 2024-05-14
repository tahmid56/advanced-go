package main

import "fmt"

type MyError struct {
	Msg string
	Code int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Msg)
}

func myFunction() error {
	return &MyError{Msg: "Something went wrong", Code: 404}
}

func main(){
	err := myFunction()

	if err != nil {
		switch e:= err.(type){
		case *MyError:
			fmt.Println("Custom error occured:", e)
		default:
			fmt.Println("Generic error: ", err)
		}
	}
}