# Flyweight Pattern

## Real World Example
🍃 Did you ever have fresh tea from some stall? They often make more than one cup that you demanded and save the rest for any other customer so to save the resources e.g. gas etc. Flyweight pattern is all about that i.e. sharing.

## In Plain Words
It is used to minimize memory usage or computational expenses by sharing as much as possible with similar objects.

## Wikipedia Definition
In computer programming, flyweight is a software design pattern. A flyweight is an object that minimizes memory use by sharing as much data as possible with other similar objects; it is a way to use objects in large numbers when a simple repeated representation would use an unacceptable amount of memory.

## Intent
Use sharing to support large numbers of fine-grained objects efficiently.

## Problem Solved
When you need to create a large number of similar objects and want to minimize memory usage by sharing common data.

## Go Implementation

```go
// Flyweight interface
type TeaType interface {
    Serve(table int) string
}

// Concrete flyweight
type KarakTea struct{}

func (k KarakTea) Serve(table int) string {
    return fmt.Sprintf("Serving Karak tea to table #%d", table)
}

// Flyweight factory
type TeaFactory struct {
    teas map[string]TeaType
}

func (tf *TeaFactory) GetTea(teaType string) TeaType {
    if tea, exists := tf.teas[teaType]; exists {
        return tea
    }
    // Create and store new flyweight
    tea := KarakTea{}
    tf.teas[teaType] = tea
    return tea
}
```

## Key Features

1. **Sharing**: Objects share common intrinsic state
2. **Memory Efficiency**: Reduces memory usage significantly
3. **Factory Management**: Factory manages flyweight instances

## Benefits

- **Memory Savings**: Dramatically reduces memory usage
- **Performance**: Fewer objects mean less memory allocation
- **Scalability**: Can handle large numbers of objects efficiently
- **Centralized Management**: Factory controls object creation

## When to Use

- When you need to create a large number of similar objects
- When object storage costs are high
- When object state can be made extrinsic
- When groups of objects can be replaced by few shared objects

## Real-world Examples

- **Text Editors**: Character objects sharing font and formatting
- **Game Development**: Particle systems, bullet objects
- **Web Browsers**: DOM nodes sharing styling information
- **Graphics Systems**: Graphical primitives sharing rendering data

## Go-Specific Notes

- Go's map type works well for flyweight factories
- Go's garbage collector helps with memory management
- Struct composition can be used for context data
- Go's zero values can represent default flyweight states
- sync.Map can be used for concurrent flyweight access
