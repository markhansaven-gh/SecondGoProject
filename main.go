package main

import "fmt"

func main() {

	// Ints
	println("Hello this is testing what I've already memorized for the Go progrmaming language!")

	var x int = 5
	var y int = 6
	var z int = 7
	fmt.Printf("Coords are: %d %d %d", x, y, z)
	println("\n\n")

	// Strings
	println("Now for some strings...")

	var firstName string = "Mark"
	var lastName string = "Haven"

	fmt.Printf("First and last name is: %s %s", firstName, lastName)

	var fullName string = firstName + " " + lastName
	fmt.Printf("\nFull name is: %s", fullName)

	// Set Arrays
	println("Now for set arrays...")
	arr := [3]string{"Mark", "Jacob", "Aaron"}
	for x := 0; x < len(arr); x++ {
		println(arr[x])
	}

	//

}
