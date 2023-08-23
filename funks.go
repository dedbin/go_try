package main

import "fmt"

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

func loops() {
	// Example of a for loop
	fmt.Println("example of simple for loop")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example of a while loop
	fmt.Println("go have while loop")
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	// Example of a nested loop
	fmt.Println("go have nested loops")
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			fmt.Println(x, y)
		}
	}

	// Example of a loop with range
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("loop with range through slice or array")
	for index, value := range nums {
		fmt.Printf("the index is %v and the value is %v \n", index, value)
	}
}

func buls() {
	age := 16

	fmt.Println("age = ", age)
	fmt.Printf("age <= 18 is %v \n", age <= 18)
	fmt.Printf("age >= 18 is %v \n", age >= 18)
	fmt.Printf("age == 18 is %v \n", age == 18)
	fmt.Printf("age != 18 is %v \n", age != 18)

	age = 61
	if age < 30 {
		fmt.Println("you are young")
	} else if age < 60 {
		fmt.Println("you are middle")
	} else {
		fmt.Println("you are old")
	}

	names := []string{"gosha", "vasya", "petya", "vanya", "masha"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at index = ", index)
			continue
		}
		if value == "vanya" {
			fmt.Println("breaking at value = ", value)
			break
		} else {
			fmt.Println("index = ", index, "value = ", value)
		}
	}
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
