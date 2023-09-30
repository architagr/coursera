package main

import (
	"fmt"
)

func main() {
	var (
		// -- employee 1 --
		employeeId1       int    = 1
		employeeName1     string = "John Wick"
		email1, mobileNo1 string = "abc@atoz.com", "123456789"
		// -- employee 2 --
		employeeId2       int    = 2
		employeeName2     string = "James Bond"
		email2, mobileNo2 string = "bond007@atoz.com", "0070070070"
	)
	printPersonalInformations(employeeName1, email1, mobileNo1, employeeId1)
	printPersonalInformations(employeeName2, email2, mobileNo2, employeeId2)
}

func printPersonalInformations(name, email, mobileNo string, employeeId int) {
	validMobile, err := isValidMobileNumber(mobileNo)
	fmt.Println("*************************")
	fmt.Printf("Details for Employee %d\n", employeeId)

	fmt.Printf("Name of the person is - %s.\n", name)
	fmt.Printf("Email is %s\n", email)
	if validMobile {
		fmt.Printf("Contact details (M) %s\n", mobileNo)
	} else {
		fmt.Printf("Contact details (M) %s, (error: %s)\n", mobileNo, err.Error())
	}
	fmt.Println("*************************")
}

func isValidMobileNumber(mobileNo string) (valid bool, err error) {
	valid, err = true, nil
	if len(mobileNo) < 10 {
		valid, err = false, fmt.Errorf("mobile number should have atleast 10 digits")
	}
	return
}
