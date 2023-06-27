package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func NewEmployee(id int, name string, org *Organization) *Employee {
	return &Employee{
		Id:   id,
		Name: name,
		Org:  org,
	}
}

func main() {
	cisco := &Organization{Name: "Cisco", City: "Bangalore"}
	e1 := &Employee{Id: 100, Name: "Magesh"}
	e1.Org = cisco
	// e2 := Employee{Id: 101, Name: "Suresh", Org: cisco}
	e2 := NewEmployee(101, "Suresh", cisco)

	fmt.Println(e1)
	fmt.Println(e2)

	fmt.Println("Change cisco city to Mysuru")
	cisco.City = "Mysuru"

	fmt.Println(e1.Org.City)
	fmt.Println(e2.Org.City)
	fmt.Println(cisco)
}
