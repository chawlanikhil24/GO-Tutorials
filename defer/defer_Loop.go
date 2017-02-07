package main

import "fmt"

func main() {
	for i := 10; i != -1; i-- {
		func() {
			defer func() {
				r := recover()
				if r != nil {
					fmt.Println("Caught the Error: ", r)
				}
			}()

			if i%2 != 0 {
				fmt.Println("I am an Odd number, how can you expect me produce 0 remainder")
				panic(i)
			} else {
				fmt.Println(i, "is an even number")
			}
		}()

	}
}
