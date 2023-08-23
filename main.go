package main

import "fmt"

func main() {
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
	fmt.Printf("smth = %q", str2)
}
