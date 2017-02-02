package main

import "fmt"

func main() {
	var array_EX2 = [5]int{2, 3, 4, 5, 6}
	//DECLARED AN ARRAY "array_EX2" of length 10,as well as initialised its values
	slice_EX2 := make([]int, 5, 20)
	slice_EX2 = []int{12, 234, 43, 34, 34}
	//DECLARED AN ARRAY using "make",of length=10 and capacity=10
	slice_EX3 := []int{12, 34, 4, 34, 35, 5}
	fmt.Println("LENGTH", "CAPACITY")
	fmt.Println(len(array_EX2), "\t", cap(array_EX2))
	fmt.Println(len(slice_EX2), "\t", cap(slice_EX2))
	fmt.Println(len(slice_EX3), "\t", cap(slice_EX3))
	slice_EX2 = append(slice_EX2, [5]int{3, 45, 65, 234, 4})
	fmt.Println(len(slice_EX2), "\t", cap(slice_EX2))

}
