# Flyweight Pattern

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
