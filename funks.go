package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// printing is a Go function that demonstrates various ways to print output.
func printing() {
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

// loops is a Go function that demonstrates different types of loops.
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

// buls prints the age variable and performs various comparisons.
func bools() {
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

// arrays_and_slices is a Go function that demonstrates the usage of arrays and slices.
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

// maps_expl is a Go function that demonstrates the usage of maps in Go.
func maps_expl() {
	// Create a map to store employee information
	employees := make(map[string]int)

	// Add some employee data to the map
	employees["John"] = 35
	employees["Jane"] = 28
	employees["Michael"] = 42

	// Print all employee information
	fmt.Println("print all employee information")
	for name, age := range employees {
		fmt.Println(name, "is", age, "years old")
	}
	// Access and print the age of an employee
	fmt.Println("print only one employee information")
	fmt.Println("John's age:", employees["John"])

	// Update an employee's age
	employees["Jane"] = 29
	fmt.Println("update Jane's age \nJohn's age:", employees["John"])

	// Delete an employee from the map
	delete(employees, "Michael")
	fmt.Println("delete Michael \n")
	for name, age := range employees {
		fmt.Println(name, "is", age, "years old")
	}
}

// pointer_expl explains how pointers work in Go.
func pointer_expl() {
	var a int = 10
	var p *int = &a
	fmt.Println("value of a = ", a)
	fmt.Println("memory address of a = ", p)
	fmt.Printf("type of p = %T\n", p)
	fmt.Println("value at ", p, " = ", *p)
}

// example of a struct
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format_bill generates a formatted bill for a given bill object.
func (b *bill) format_bill() string {
	fs := "Bill breakdown\n"
	var total float64 = 0

	for name, price := range b.items {
		total += price
		fs += fmt.Sprintf("%-25v ... %v$\n", name, price)
	}
	fs += fmt.Sprintf("%-25v ... %0.2f$\n", "tip:", b.tip)
	total += b.tip
	fs += fmt.Sprintf("%-25v ... %0.2f$\n", "Total:", total)
	return fs
}

// demonstrate_struct demonstrates the usage of a struct in Go.
func demonstrate_struct() {
	b := create_bill()
	fmt.Println(b.format_bill())
	b.update_tip(10.00)
	b.add_item("milk", 1.99)
	fmt.Println(b.format_bill())
}

// update tip
func (b *bill) update_tip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) add_item(name string, price float64) {
	b.items[name] = price
}

// create bill using information for user
func create_bill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := get_input("Enter bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created bill - ", b.name)

	return b
}

// get_input is a function that takes a prompt string and a *bufio.Reader as parameters.
func get_input(promt string, r *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

// promt_options prompts the user for options and performs corresponding actions.
func promt_options(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := get_input("Choose an option (a - add, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
	switch opt {
	case "a":
		name, _ := get_input("Enter item name: ", reader)
		price, _ := get_input("Enter item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promt_options(b)
		}
		b.add_item(name, p)
		fmt.Println("item added - ", name, price)
		promt_options(b)
	case "s":
		b.save_bill()
		fmt.Println("u saved the bill", b.name)

	case "t":
		tip, _ := get_input("Enter tip: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promt_options(b)
		}
		b.update_tip(t)
		fmt.Println("tip added - ", tip)
		promt_options(b)

	default:
		fmt.Println("u chose invalid option...")
		promt_options(b)
	}
}

// save_bill saves the bill to a file.
func (b *bill) save_bill() {
	data := []byte(b.format_bill())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill saved!")
}

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// area calculates the area of a square.
func (d square) area() float64 {
	return d.length * d.length
}

// perimeter calculates the perimeter of a square.
func (d square) perimeter() float64 {
	return 4 * d.length
}

// area calculates the area of a circle.
func (d circle) area() float64 {
	return math.Pi * d.radius * d.radius
}

// perimeter calculates the perimeter of a circle.
func (d circle) perimeter() float64 {
	return 2 * math.Pi * d.radius
}

// print_shape prints the area and perimeter of a given shape.
func print_shape(s shape) {
	fmt.Printf("area of %T is %0.2f \n", s, s.area())
	fmt.Printf("perimeter of %T is %0.2f \n", s, s.perimeter())
}

// demo_interface is a function that demonstrates the use of interfaces in Go.
func demo_interface() {
	shapes := []shape{
		square{length: 10},
		circle{radius: 5},
		circle{radius: 10},
		square{length: 20},
	}
	for _, v := range shapes {
		print_shape(v)
		fmt.Println("---")
	}
}
