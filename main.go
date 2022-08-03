package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{}
	var num1, num2, num3, num4, num5 int

	fmt.Println("Enter the value of num1")
	fmt.Scanln(&num1)
	arr = append(arr, num1)
	fmt.Println("Enter the value of num2")
	fmt.Scanln(&num2)
	arr = append(arr, num2)
	fmt.Println("enter the value of num3")
	fmt.Scanln(&num3)
	arr = append(arr, num3)
	fmt.Println("enter the value of num4")
	fmt.Scanln(&num4)
	arr = append(arr, num4)
	fmt.Println("enter the value of num5")
	fmt.Scanln(&num5)
	arr = append(arr, num5)

	s := []int{num1, num2, num3, num4, num5}
	sort.Ints(s)
	fmt.Println(s)

	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	fmt.Println(s)
}
