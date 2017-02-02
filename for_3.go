package main

import "fmt"

func main() {
	array1 := [10]int{12, 32, 43, 24, 55, 76, 27, 48, 69, 10}
	fmt.Println("index", "\t", "value")
	for index, value := range array1 { //index,value are IDENTIFIERS, range is KEYWORD
		fmt.Println(index, "\t", value)
	}
	//if we need only values, replace index with "_"
	fmt.Println("value")
	for _, value := range array1 { //index,value are IDENTIFIERS, range is KEYWORD
		fmt.Println(value)
	}
}
