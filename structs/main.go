package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo // sets the key to be "contactInfo" of type contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Gordon",
		contactInfo: contactInfo{
			email:   "detective@gcpd.com",
			zipCode: 94000,
		},
	}

	// updateName here actually ends up using pointer to jim under the hood.
	jim.updateName("James")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func notes() {
	// Can create a struct like the following
	alex := person{"Alex", "Anderson", contactInfo{}}      // relies on positional arguments
	bob := person{firstName: "Bob", lastName: "DeBuilder"} // slightly more explicit
	var cam person                                         // Initialised with Zero value.
	cam.firstName = "Cam"
	cam.lastName = "O'Mile"
	// Zero values: string = "", int = 0, float = 0, bool = false.

	fmt.Printf("%+v", alex) // %+v prints out struct fields along with values
	fmt.Println(alex, bob, cam)

	// Pointers:
	// updateName takes in a pointer to a person so we can update the value of cam
	camPointer := &cam
	camPointer.updateName("Cameron")
	/*
		GO has a shortcut with pointers. If the receiver has definition of receiving a type pointer to struct
		we can call it without the &variable and GO will automatically interpret it as a pointer.
		GOTCHA: slices are of "Reference types". This is because the slice structure points to
		the actual array in memory. So the slice actually gets passed by value, but it still points to the same
		underlying array. Other Reference types: maps, channels, pointers, functions.
	*/
}
