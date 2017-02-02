package main

import "fmt"

// SYNTAX OF "FOR" USED
// for initialsation;condition;operation{
//     BODY
// }
func main() {
	count := 10
	for index := 0; index < count; index++ {
		fmt.Println(index)
	}
}
