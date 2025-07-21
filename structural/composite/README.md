# Composite Pattern

## Real World Example
🌿 Every organization is composed of employees. Each of the employees has the same features i.e. has a salary, has some responsibilities, may or may not report to someone, may or may not have some subordinates etc.

## In Plain Words
Composite pattern lets clients treat the individual objects in a uniform manner.

## Wikipedia Definition
In software engineering, the composite pattern is a partitioning design pattern. The composite pattern describes that a group of objects is to be treated in the same way as a single instance of an object. The intent of a composite is to "compose" objects into tree structures to represent part-whole hierarchies. Implementing the composite pattern lets clients treat individual objects and compositions uniformly.

## Intent
Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.

## Problem Solved
When you need to work with tree structures where individual objects and compositions of objects should be treated the same way.

## Go Implementation

```go
// Component interface
type Employee interface {
    GetName() string
    GetSalary() int
    GetDetails(indent int) string
}

// Leaf
type Developer struct {
    Name   string
    Salary int
}

func (d *Developer) GetDetails(indent int) string {
    return fmt.Sprintf("%s%s (Developer) - Salary: $%d", 
        strings.Repeat("  ", indent), d.Name, d.Salary)
}

// Composite
type Organization struct {
    Name      string
    Employees []Employee
}

func (o *Organization) AddEmployee(employee Employee) {
    o.Employees = append(o.Employees, employee)
}
```

## Key Features

1. **Tree Structure**: Represents part-whole hierarchies
2. **Uniform Treatment**: Treats individual and composite objects the same
3. **Recursive Composition**: Composites can contain other composites

## Benefits

- **Simplicity**: Clients can treat individual objects and compositions uniformly
- **Flexibility**: Easy to add new component types
- **Recursive Structure**: Natural support for tree-like structures
- **Open/Closed Principle**: Easy to add new component types

## When to Use

- When you want to represent part-whole hierarchies
- When you want clients to ignore the difference between individual and composite objects
- When you have tree structures with uniform operations

## Real-world Examples

- **File Systems**: Files and directories
- **GUI Components**: Windows, panels, buttons
- **Organization Charts**: Employees, departments, teams
- **Mathematical Expressions**: Numbers, operators, complex expressions

## Go-Specific Notes

- Go's interfaces make the composite pattern very natural
- Slices work well for managing child components
- Go's composition over inheritance fits perfectly
- Recursive structures are well-supported in Go
- Error handling can be naturally integrated into composite operations
