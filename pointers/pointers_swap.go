package main

import "fmt"

func swap1(arg1, arg2 *int) {
	*arg1, *arg2 = *arg2, *arg1
}

func swap2(arg1, arg2 *int) {
	*arg1 = *arg1 + *arg2
	*arg2 = *arg1 - *arg2
	*arg1 = *arg1 - *arg2
}

func swap3(arg1, arg2 *int) {
	temp := *arg1
	*arg1 = *arg2
	*arg2 = temp
}

func swap4(arg1, arg2 *int) {
	temp1 := *arg1
	temp2 := *arg2
	*arg2 = temp1
	*arg1 = temp2
}

func showresults(a, b, i int) {
	fmt.Println(" ")
	fmt.Println("Values after swapping, by method ", i)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println(" ")
}

func main() {
	var a, b int
	fmt.Println("Enter a:")
	fmt.Scanf("%d", &a)
	fmt.Println("Enter b:")
	fmt.Scanf("%d", &b)
	fmt.Println(" ")
	fmt.Println("Values before swapping")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	i := 1
	swap1(&a, &b)
	showresults(a, b, i)
	swap2(&a, &b)
	i++
	showresults(a, b, i)
	swap3(&a, &b)
	i++
	showresults(a, b, i)
	swap4(&a, &b)
	i++
	showresults(a, b, i)
	fmt.Println("Enough of Swapping !.. Phewww!")
}
