package main

import (
	"fmt"
)

// const (
// 	ADMIN int = iota + 1
// 	INTERN
// )

var lastId int = 0

// type Employee struct {
// 	id, empType         int
// 	name, email, mobile string
// }

// func (emp *Employee) GetId() int {
// 	return emp.id
// }

// func (emp *Employee) GetType() int {
// 	return emp.empType
// }

// func (emp *Employee) GetName() string {
// 	return emp.name
// }

// func (emp *Employee) GetMobile() string {
// 	return emp.mobile
// }
// func (emp *Employee) GetEmail() string {
// 	return emp.email
// }

func InitEmployee(name, email, mobile string, empType int) (Employee, error) {
	lastId++
	if empType != ADMIN && empType != INTERN {
		return Employee{}, fmt.Errorf("%d is not a valid Employee Type", empType)
	}
	return Employee{
		name:    name,
		email:   email,
		mobile:  mobile,
		empType: empType,
		id:      lastId,
	}, nil
}

type EmployeeRepository struct {
	listOfEmployees       []Employee
	employeeTypeEmployees map[int][]Employee
}

func InitEmployeeRepository() EmployeeRepository {
	return EmployeeRepository{
		listOfEmployees:       make([]Employee, 0, 10),
		employeeTypeEmployees: make(map[int][]Employee),
	}
}
func (repo *EmployeeRepository) AddEmployeeInMap(emp Employee) {
	employees, ok := repo.employeeTypeEmployees[emp.empType]
	if !ok {
		employees = make([]Employee, 0)
	}
	employees = append(employees, emp)
	repo.employeeTypeEmployees[emp.empType] = employees
}
func (repo *EmployeeRepository) Add(emp Employee) {
	repo.listOfEmployees = append(repo.listOfEmployees, emp)
}
func (repo *EmployeeRepository) AddMany(emps []Employee) {
	repo.listOfEmployees = append(repo.listOfEmployees, emps...)

}

func (repo *EmployeeRepository) GetAll() []Employee {
	return repo.listOfEmployees
}

func (repo *EmployeeRepository) GetEmployeeById(id int) (Employee, error) {
	for idx, emp := range repo.listOfEmployees {
		if emp.id == id {
			return repo.listOfEmployees[idx], nil
		}
	}
	return Employee{}, fmt.Errorf("employee with id: %d not found", id)
}

func (repo *EmployeeRepository) Search(name string, skip, pageSize int) []Employee {
	count := 0
	result := make([]Employee, 0, pageSize)
	for _, emp := range repo.listOfEmployees {
		if len(result) >= pageSize {
			break
		}
		if emp.name == name {
			count++
			if count <= skip {
				continue
			}
			result = append(result, emp)
		}
	}
	return result
}
