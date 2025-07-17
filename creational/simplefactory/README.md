# Simple Factory Pattern

## Intent
Simple Factory isn't actually a design pattern per se but a programming technique. It provides a way to encapsulate the instantiation logic of objects.

## Problem Solved
When you need to create objects but don't want to expose the instantiation logic to the client, and you want to reference the created object through a common interface.

## Go Implementation

The Go implementation uses interfaces and structs:

```go
// Door interface defines the contract
type Door interface {
    GetDescription() string
    GetWidth() float64
    GetHeight() float64
}

// WoodenDoor implements the Door interface
type WoodenDoor struct {
    Width  float64
    Height float64
}

// DoorFactory creates doors
type DoorFactory struct{}

func (df DoorFactory) MakeDoor(width, height float64) Door {
    return WoodenDoor{Width: width, Height: height}
}
```

## Key Features

1. **Encapsulation**: Object creation logic is encapsulated in the factory
2. **Simplicity**: Easy to understand and implement
3. **Flexibility**: Easy to modify object creation without affecting client code

## Benefits

- **Centralized Creation**: All object creation logic is in one place
- **Loose Coupling**: Client code doesn't depend on concrete implementations
- **Easy Maintenance**: Changes to object creation only affect the factory

## When to Use

- When the creation process involves some logic that shouldn't be part of the client
- When you want to generate different instances depending on the configuration
- When working with many related classes

## Real-world Examples

- **Database Connection Factory**: Creating different types of database connections
- **Logger Factory**: Creating different types of loggers (file, console, network)
- **HTTP Client Factory**: Creating HTTP clients with different configurations
- **Document Factory**: Creating different document types (PDF, Word, Excel)

## Go-Specific Notes

- Go's interface{} type provides flexibility for generic factories
- Struct embedding can be used for complex factory hierarchies
- Go's zero values make initialization straightforward
- Error handling can be easily integrated into factory methods
