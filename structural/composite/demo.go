package main

import (
	"fmt"
	"strings"
)

// Employee interface - component
type Employee interface {
	GetName() string
	GetSalary() int
	GetRoles() []string
	GetDetails(indent int) string
}

// Developer - leaf
type Developer struct {
	Name   string
	Salary int
	Roles  []string
}

func NewDeveloper(name string, salary int) *Developer {
	return &Developer{
		Name:   name,
		Salary: salary,
		Roles:  []string{"Development", "Testing"},
	}
}

func (d *Developer) GetName() string {
	return d.Name
}

func (d *Developer) GetSalary() int {
	return d.Salary
}

func (d *Developer) GetRoles() []string {
	return d.Roles
}

func (d *Developer) GetDetails(indent int) string {
	indentStr := strings.Repeat("  ", indent)
	return fmt.Sprintf("%s%s (Developer) - Salary: $%d", indentStr, d.Name, d.Salary)
}

// Designer - leaf
type Designer struct {
	Name   string
	Salary int
	Roles  []string
}

func NewDesigner(name string, salary int) *Designer {
	return &Designer{
		Name:   name,
		Salary: salary,
		Roles:  []string{"Design", "UI/UX"},
	}
}

func (d *Designer) GetName() string {
	return d.Name
}

func (d *Designer) GetSalary() int {
	return d.Salary
}

func (d *Designer) GetRoles() []string {
	return d.Roles
}

func (d *Designer) GetDetails(indent int) string {
	indentStr := strings.Repeat("  ", indent)
	return fmt.Sprintf("%s%s (Designer) - Salary: $%d", indentStr, d.Name, d.Salary)
}

// Organization - composite
type Organization struct {
	Name      string
	Employees []Employee
}

func NewOrganization(name string) *Organization {
	return &Organization{
		Name:      name,
		Employees: make([]Employee, 0),
	}
}

func (o *Organization) AddEmployee(employee Employee) {
	o.Employees = append(o.Employees, employee)
}

func (o *Organization) GetName() string {
	return o.Name
}

func (o *Organization) GetSalary() int {
	totalSalary := 0
	for _, employee := range o.Employees {
		totalSalary += employee.GetSalary()
	}
	return totalSalary
}

func (o *Organization) GetRoles() []string {
	roles := make([]string, 0)
	for _, employee := range o.Employees {
		roles = append(roles, employee.GetRoles()...)
	}
	return roles
}

func (o *Organization) GetDetails(indent int) string {
	indentStr := strings.Repeat("  ", indent)
	result := fmt.Sprintf("%s%s (Organization) - Total Salary: $%d\n", indentStr, o.Name, o.GetSalary())
	
	for _, employee := range o.Employees {
		result += employee.GetDetails(indent+1) + "\n"
	}
	
	return strings.TrimRight(result, "\n")
}

func main() {
	fmt.Println("=== Composite Pattern Demo ===")
	
	// Create individual employees
	john := NewDeveloper("John Doe", 12000)
	jane := NewDesigner("Jane Doe", 10000)
	
	// Create organization and add employees
	organization := NewOrganization("Tech Company")
	organization.AddEmployee(john)
	organization.AddEmployee(jane)
	
	// Create sub-organization
	subOrg := NewOrganization("Development Team")
	subOrg.AddEmployee(NewDeveloper("Alice Smith", 15000))
	subOrg.AddEmployee(NewDeveloper("Bob Johnson", 13000))
	
	// Add sub-organization to main organization
	organization.AddEmployee(subOrg)
	
	// Display the hierarchy
	fmt.Println(organization.GetDetails(0))
	
	fmt.Printf("\nTotal company salary: $%d\n", organization.GetSalary())
	
	fmt.Println("\nComposite pattern treats individual objects and compositions uniformly!")
}
