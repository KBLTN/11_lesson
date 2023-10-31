package main

import (
	"errors"
	"fmt"
	"log"
)

type AppError struct {
	Message string
	Err     error
}

func (ae *AppError) Error() string {
	return ae.Message
}

func main() {
	divide(4, 0)
	fmt.Println("after panic")
}

func divide(a, b int) {
	defer func() {
		var appErr *AppError
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appErr) {
					fmt.Println("app err panic!", err)
				} else {
					fmt.Println("custom panic!")
				}
			default:
				panic("some panic")
			}
			log.Println("panic happened:", err)
		}
	}()
	fmt.Println(div(a, b))
}

func div(a, b int) int {
	if b == 0 {
		panic(&AppError{
			Message: "this is divide by zero custom error",
			Err:     nil,
		})
	}
	return a / b
}

//type name struct {
//	A, B int
//}
//
//func (n *name) method() {
//	fmt.Println("ok")
//	fmt.Println(n.A)
//}

//func main() {
//	n := &name{1, 2}
//	n = nil
//	n.method()
//}

//___________
//func main() {
//	go test()
//	time.Sleep(300 * time.Millisecond)
//}
//
//func test() int {
//	a := []int{1, 2, 3}
//	return a[3]

//__________
//defer func() {
//	fmt.Println("OK")
//}()
//
//panic("something goes wrong")
