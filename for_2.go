package main

import "fmt"

// SYNTAX OF "FOR" USED
// initialisation
//
// for condition {
//     BODY
//     operation
// }

func main() {
	count := 10
	index := 1          //INITIALISATION
	for index < count { //CONDITION
		fmt.Println(index)
		index++ //OPERATION
	}
}
