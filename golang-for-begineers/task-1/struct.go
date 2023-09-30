package main

import "fmt"

const (
	ADMIN int = iota + 1
	INTERN
)

type IEmployee interface {
	GetId() int
	GetType() int
	GetName() string
	GetMobile() string
	GetEmail() string
}
type Employee struct {
	id, empType         int
	name, email, mobile string
}

func (emp *Employee) GetId() int {
	return emp.id
}

func (emp *Employee) GetType() int {
	return emp.empType
}

func (emp *Employee) GetName() string {
	return emp.name
}
func (emp *Employee) GetMobile() string {
	return emp.mobile
}
func (emp *Employee) GetEmail() string {
	return emp.email
}

func (emp *Employee) isValidMobileNumber() (valid bool, err error) {
	valid, err = true, nil
	if len(emp.mobile) < 10 {
		valid, err = false, fmt.Errorf("mobile number should have atleast 10 digits")
	}
	return
}

func main() {
	var employee1 IEmployee = &Employee{
		id:      1,
		empType: ADMIN,
		name:    "John Wick",
		email:   "abc@atoz.com",
		mobile:  "123456789",
	}
	employee1.isValidMobileNumber()

	fmt.Printf("Employee name %s\n", employee1.GetName())
	fmt.Printf("Employee Type %d\n", employee1.GetType())

	employee2 := Employee{
		id:      2,
		empType: ADMIN,
		name:    "James Bond",
		email:   "bond007@atoz.com",
		mobile:  "0070070070",
	}

	fmt.Printf("Employee name %s\n", employee2.GetName())
	fmt.Printf("Employee Type %d\n", employee2.GetType())

}
