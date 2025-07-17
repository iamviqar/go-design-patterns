# Decorator Pattern

## Intent
Attach additional responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing for extending functionality.

## Problem Solved
When you want to add behavior to objects without altering their structure, and when extension by subclassing would be impractical.

## Go Implementation

```go
// Component interface
type Coffee interface {
    GetCost() float64
    GetDescription() string
}

// Concrete component
type SimpleCoffee struct{}

func (s SimpleCoffee) GetCost() float64 {
    return 10.0
}

// Decorator
type MilkDecorator struct {
    coffee Coffee
}

func (m MilkDecorator) GetCost() float64 {
    return m.coffee.GetCost() + 2.0
}

func (m MilkDecorator) GetDescription() string {
    return m.coffee.GetDescription() + ", milk"
}
```

## Key Features

1. **Dynamic Behavior**: Add behavior at runtime
2. **Composition**: Uses composition instead of inheritance
3. **Flexibility**: Can combine multiple decorators

## Benefits

- **Runtime Enhancement**: Add responsibilities at runtime
- **Flexible Alternative**: More flexible than static inheritance
- **Single Responsibility**: Each decorator has one responsibility
- **Combinable**: Decorators can be combined in different ways

## When to Use

- When you want to add responsibilities to objects dynamically
- When extension by subclassing is impractical
- When you want to withdraw responsibilities from objects
- When a large number of independent extensions are possible

## Real-world Examples

- **Stream Processing**: BufferedReader, GzipReader wrapping base readers
- **HTTP Middleware**: Authentication, logging, compression middleware
- **UI Components**: Bordered, scrollable, resizable windows
- **Text Formatting**: Bold, italic, underline text decorators

## Go-Specific Notes

- Go's interfaces make decorators very natural
- Composition is preferred over inheritance in Go
- Go's io.Reader/Writer interfaces are great examples of decorator pattern
- Functional decorators can be implemented using higher-order functions
- Go's embedding can be used for simpler decorators
