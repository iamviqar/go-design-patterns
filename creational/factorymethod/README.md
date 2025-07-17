# Factory Method Pattern

## Intent
Define an interface for creating an object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.

## Problem Solved
When a class needs to instantiate one of several possible classes, or when a class wants its subclasses to specify the objects it creates.

## Go Implementation

```go
// Product interface
type Interviewer interface {
    AskQuestions() string
}

// Creator interface with factory method
type HiringManager interface {
    MakeInterviewer() Interviewer
    TakeInterview() string
}

// Concrete creators implement the factory method
type DevelopmentManager struct{}

func (d DevelopmentManager) MakeInterviewer() Interviewer {
    return Developer{}
}
```

## Key Features

1. **Delegation**: Object creation is delegated to subclasses
2. **Flexibility**: Easy to add new product types
3. **Encapsulation**: Creation logic is encapsulated in specific managers

## Benefits

- **Open/Closed Principle**: Open for extension, closed for modification
- **Single Responsibility**: Each creator has one reason to change
- **Loose Coupling**: Client code works with abstractions

## When to Use

- When a class can't anticipate the class of objects it needs to create
- When a class wants its subclasses to specify the objects it creates
- When you want to localize the knowledge of which class gets created

## Real-world Examples

- **UI Framework**: Different platforms create different UI controls
- **Database Drivers**: Different databases create different connection objects
- **Payment Processing**: Different providers create different payment processors
- **File Parsers**: Different parsers for different file formats

## Go-Specific Notes

- Go's composition over inheritance works well with factory methods
- Interface satisfaction is implicit, making implementations flexible
- Go's struct embedding can simulate inheritance for shared behavior
- Error handling can be incorporated into factory method returns
