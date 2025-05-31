package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Mathias", Age: 25}
	fmt.Printf("This is out person %+v\n", person)

	employee := struct {
		name string
		id   int
	}{
		name: "pizza",
		id:   123,
	}

	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	contact := Contact{
		Name: "Nate",
		Address: Address{
			Street: "123 Main street",
			City:   "Santa Monica",
		},
	}

	fmt.Println("This is an contact", contact)
	fmt.Println("This is an employee", employee)

	fmt.Println("name before: ", person.Name)
	// modifyPersonName(&person)
	person.modifyPersonName("Hope")
	fmt.Println("name after: ", person.Name)

	x := 20
	ptr := &x
	fmt.Printf("value of x: %d and address of x %p\n", x, ptr)
	*ptr = 30
	fmt.Printf("value of new x: %d and address of x %p\n", x, ptr)
}

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("inside scope: new name", p.Name)
}

// func modifyPersonName(person *Person) {
// 	person.Name = "Hope"
// 	fmt.Println("inside scope: new name", person.Name)
// }
