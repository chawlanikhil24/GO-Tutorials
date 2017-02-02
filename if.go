package main

import "fmt"

// syntax of IF-ELSE statment:
//
// if (condition 1){
//
//     CODE HERE
//
// }else if(condition 2){
//
//     CODE HERE
// }else{
//
//     CODE HERE
// }

func main() {
	//PROGRAM TO PRINT MULTIPLES OF 2 AND 3 IN BETWEEN 1-20

	for i := 1; i <= 20; i++ {

		//IMPLEMENTATION OF IF-ELSE

		if i%2 == 0 {
			fmt.Println(i, " is a multiple of 2\n")
		} else if i%3 == 0 {
			fmt.Println(i, " is a multiple of 3\n")
		} else {
			fmt.Println(i, " neither a multiple of 2 nor 3\n")
		}
	}
}
