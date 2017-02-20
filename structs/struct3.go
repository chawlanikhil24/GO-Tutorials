package main

import "fmt"

type list struct {
	arr []int
}

// func showchange()

func (ob list) app1(item int) {
	fmt.Println("Current array is: ", ob.arr)
	fmt.Println("Item to append is: ", item)
	ob.arr = append(ob.arr, item)
	fmt.Println("Array after change is: ", ob.arr)
}

func (ob *list) app2(item int) {
	// ob := &list{}
	fmt.Println("Current array is: ", ob.arr)
	fmt.Println("Item to append is: ", item)
	ob.arr = append(ob.arr, item)
	fmt.Println("Array after change is: ", ob.arr)
}

func main() {
	fmt.Println("=========Usual OBJECTs==========")
	ob1 := list{}
	ob1.app1(10)
	ob1.app1(11)
	ob1.app1(12)
	fmt.Println("=========OBJECT POINTERS==========")
	//POINTER INVOKATION
	ob2 := &list{}
	ob2.app2(10)
	ob2.app2(11)
	ob2.app2(12)
}
