package main

import "fmt"

// syntax of NESTED IF-ELSE statment:
//
// if (condition 1){
//         if (condition 2){
//             if(condition 3){
//                 CONDITIONS SATISFIES
//                 DO SOMETHING HERE
//             }
//         }
// }

func main() {
	// PROGRAM TO PRINT MULTIPLES OF 6 IN BETWEEN 1-60
	for i := 1; i <= 60; i++ {

		//IMPLEMENTATION OF NESTED IF-ELSE

		if i%2 == 0 {
			if i%3 == 0 {
				fmt.Println(i, " is a multiple of 6\n")
			}
		}
	}
}
