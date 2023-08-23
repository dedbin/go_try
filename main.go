package main

import (
	"fmt"
)

func main() {
	Printing()
	arrays_and_slices()
}

func arrays_and_slices() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var slice = []int{1, 2, 3, 4, 5}

	fmt.Println("Arrays and Slices")
	fmt.Printf("Arrays in go are static arr = %v\n", arr)
	fmt.Printf("Slices in go are dynamic arr = %v\n", slice)
}

func Printing() {
	age := 16
	name := "gosha"
	str2 := "smth"
	//Print (no \n)
	fmt.Print("Hello, ")
	fmt.Print("world \n")
	fmt.Print("I am working")

	//Println (have \n)
	fmt.Println("how r u?")
	fmt.Println("bye-bye")
	fmt.Println("My name is", name, " i am ", age, "years old")

	//Printf (formated output, no \n)
	fmt.Printf("My name is %v i am %v  years old", name, age)
	fmt.Printf("smth = %q \n", str2)
}
