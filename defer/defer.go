package main

import "fmt"

func cleanser() {
	fmt.Println("I am the Cleanser Function")
}

func main() {
	defer cleanser()
	fmt.Println("I am the statement which will execute before defer")
}
