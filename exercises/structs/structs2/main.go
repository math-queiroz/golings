// structs2
// Make me compile!

package main

import "fmt"

type Person struct {
	// don't just create the phone field here. embed a new struct
	name string
	age int
	ContactDetails
}

type ContactDetails struct {
	phone string
}

func main() {
	contactDetails := ContactDetails{phone: "+11 99 98 97"}
	person := Person{name: "John", age: 32, ContactDetails: contactDetails}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
