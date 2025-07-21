# Facade Pattern

## Real World Example
📦 How do you turn on the computer? "Hit the power button" you say! That is what you believe because you are using a simple interface that computer provides on the outside, internally it has to do a lot of stuff to make it happen. This simple interface to the complex subsystem is a facade.

## In Plain Words
Facade pattern provides a simplified interface to a complex subsystem.

## Wikipedia Definition
A facade is an object that provides a simplified interface to a larger body of code, such as a class library.

## Intent
Provide a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.

## Problem Solved
When you want to hide the complexity of a subsystem and provide a simple interface to clients.

## Go Implementation

```go
// Subsystem components
type Computer struct{}

func (c Computer) GetElectricShock() string { return "Ouch!" }
func (c Computer) MakeSound() string { return "Beep beep!" }
func (c Computer) ShowLoadingScreen() string { return "Loading.." }
func (c Computer) Bam() string { return "Ready to be used!" }

// Facade
type ComputerFacade struct {
    computer Computer
}

func (cf ComputerFacade) TurnOn() string {
    result := ""
    result += cf.computer.GetElectricShock() + "\n"
    result += cf.computer.MakeSound() + "\n"
    result += cf.computer.ShowLoadingScreen() + "\n"
    result += cf.computer.Bam()
    return result
}
```

## Key Features

1. **Simplified Interface**: Provides simple interface to complex subsystem
2. **Decoupling**: Reduces dependencies between clients and subsystem
3. **Layered Architecture**: Supports layered system design

## Benefits

- **Simplicity**: Makes complex subsystems easier to use
- **Decoupling**: Reduces coupling between clients and subsystem
- **Flexibility**: Subsystem can change without affecting clients
- **Layering**: Supports structuring systems into layers

## When to Use

- When you want to provide a simple interface to a complex subsystem
- When there are many dependencies between clients and implementation classes
- When you want to layer your subsystems

## Real-world Examples

- **API Gateways**: Simplifying access to multiple microservices
- **Database Layers**: Hiding complex SQL operations
- **Payment Processing**: Simplifying multiple payment provider integrations
- **Operating System APIs**: Hiding hardware complexity

## Go-Specific Notes

- Go packages naturally provide facade functionality
- Go's composition makes facade implementation straightforward
- Interface-based design works well with facades
- Error handling can be centralized in facade methods
