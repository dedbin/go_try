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

	names := []string{"gosha", "vasya", "petya"}

	fmt.Println("Arrays and Slices")
	fmt.Printf("Arrays in go are static arr = %v\n", arr)
	fmt.Printf("Slices in go are dynamic arr = %v\n", slice)
	fmt.Println()
	fmt.Printf("array of names = %v\n", names)
	names[1] = "vanya"
	fmt.Printf("u can change arrays in go %v\n", names)
	fmt.Println()
	fmt.Printf("go have slices names[1:3] = %v\n", names[1:3])
	fmt.Printf("names[1:] = %v\n", names[1:])
	fmt.Printf("names[:3] = %v\n", names[:3])
	fmt.Println()
	fmt.Printf("u can add to slice. \nslise before = %v\n", slice)
	slice = append(slice, 6)
	fmt.Printf("slise after = %v\n", slice)
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
