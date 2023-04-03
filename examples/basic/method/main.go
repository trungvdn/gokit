package main

import "fmt"

type Employee struct {
	LastName  string
	FirstName string
}

func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) changeName() {
	(*e).FirstName = "A"
}

func main() {
	e := Employee{
		FirstName: "Yasuo",
		LastName:  "Zoe",
	}
	fmt.Println(e.fullName())
	/*This proves that the method changeName received just a copy of the actual struct e.
	Hence any changes made to the copy inside the method did not affect the original struct.
	When a method belongs to the pointer of a type, its receiver will receive the pointer to the object instead of a copy of the object
	don’t necessarily need to call a method on a pointer if the method has a pointer receiver.
	You are allowed to call this method on the value instead and Go will pass the pointer of the value as the receiver.*/
	e.changeName()
	fmt.Println(e.fullName())

	// Methods with the same name
	// Pointer receivers
	// Calling methods with pointer receiver on values
	// Methods on nested struct
	// Anonymously nested struct
	// Promoted methods
	// Methods can accept both pointer and value
	// Methods on non-struct type
	/* a method can receive any type as long as the type definition and method definition is in the same package.*/
}
