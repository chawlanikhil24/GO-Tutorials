package main

import "fmt"

// SYNTAX OF basic switch
//
// switch expression {
// case condition:
// default{
//
//     }
// }

func main() {
	for i := 1; i <= 20; i++ {
		switch i {
		case 1:
			fmt.Println("I am", i, "from the basic switch")
		case 2:
			fmt.Println("I am", i, "from the basic switch")
		case 3:
			fmt.Println("I am", i, "from the basic switch")
		default:
			fmt.Println("MatcheS Nothing")
		}
	}
}


fmt.Print("---------------------------------------------------------------\n")
//IMPLEMENTATION OF SWITCH AS IF-ELSE STATEMENT
for i := 1; i <= 20; i++ {
	switch {
	case i%2 == 0:
		fmt.Print(i, " is an even number\n")
		break
	case i%3 == 0:
		fmt.Print(i, " is a multiple of 3\n")
		if i%2 == 0 {
			fmt.Print(i, " is a multiple of 6\n")
		}
		break
	case i%7 == 0, i%10 == 0: //we can execute 2 different cases in 1 CASE by separating them with comma
		fmt.Println(i, " I am either divisible by 7 or 10")
	default:
		fmt.Print(i, " neither satisfies any of the cases\n")
	}
}
