package main

import "fmt"

func printInterface(x interface{}) { //Function, that is only accepting interface type parameters
	fmt.Println(x)
}

func main() {
	interfaceEx := []interface{}{"Nikhil", "XYZZ", "ABCD"} //Created an array of interface{} type
	fmt.Println(interfaceEx)
	printInterface(interfaceEx)
}
