package main

import "fmt"

func main() {
	var ptr *int //INITIALISED A INTEGER TYPE POINTER
	i := 4       //DEINED A INTEGER WITH VALUE = 4
	ptr = &i     //ASSIGNED THE ADDRESS OF "i", to POINTER "ptr"
	fmt.Println("The address of 'i' in memory: ", ptr)
	fmt.Println("The value of 'i' stored at that address: ", *ptr) //Printing the address of "i", and value stored at that address
	*ptr = *ptr << 1
	fmt.Println("The value of 'i' after operation: ", i)
}
