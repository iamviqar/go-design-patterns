# Simple Factory Pattern

## Real World Example
🏠 Consider, you are building a house and you need doors. You can either put on your carpenter clothes, bring some wood, glue, nails and all the tools required to build the door and start building it in your house or you can simply call the factory and get the built door delivered to you so that you don't need to learn anything about the door making or to deal with the mess that comes with making it.

## In Plain Words
Simple factory simply generates an instance for client without exposing any instantiation logic to the client.

## Wikipedia Definition
In object-oriented programming (OOP), a factory is an object for creating other objects – formally a factory is a function or method that returns objects of a varying prototype or class from some method call, which is assumed to be "new".

## Intent
Simple Factory isn't actually a design pattern per se but a programming technique. It provides a way to encapsulate the instantiation logic of objects.

## Problem Solved
When creating an object is not just a few assignments and involves some logic, it makes sense to put it in a dedicated factory instead of repeating the same code everywhere.

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
