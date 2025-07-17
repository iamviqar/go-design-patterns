# Builder Pattern

## Intent
Separate the construction of a complex object from its representation so that the same construction process can create different representations.

## Problem Solved
When you need to create complex objects step by step, or when an object requires many optional parameters (avoiding the telescoping constructor anti-pattern).

## Go Implementation

```go
// Product
type Burger struct {
    Size      int
    Cheese    bool
    Pepperoni bool
    Lettuce   bool
    Tomato    bool
}

// Builder
type BurgerBuilder struct {
    burger Burger
}

func NewBurgerBuilder(size int) *BurgerBuilder {
    return &BurgerBuilder{burger: Burger{Size: size}}
}

func (bb *BurgerBuilder) AddCheese() *BurgerBuilder {
    bb.burger.Cheese = true
    return bb
}

func (bb *BurgerBuilder) Build() Burger {
    return bb.burger
}
```

## Key Features

1. **Step-by-Step Construction**: Build objects incrementally
2. **Method Chaining**: Fluent interface for easy use
3. **Flexibility**: Different representations of the same object

## Benefits

- **Readable Code**: Method chaining creates readable, self-documenting code
- **Flexible Construction**: Can create different variations of the same object
- **Immutable Products**: Can ensure the final product is immutable
- **Avoids Telescoping Constructors**: No need for multiple constructors with different parameters

## When to Use

- When creating objects with many optional parameters
- When the construction process must allow different representations
- When you want to isolate the construction code from the representation
- When building complex objects step by step

## Real-world Examples

- **HTTP Request Builder**: Building HTTP requests with optional headers, parameters
- **SQL Query Builder**: Building complex SQL queries step by step
- **Configuration Builder**: Building configuration objects with optional settings
- **Test Data Builder**: Creating test objects with varying properties

## Go-Specific Notes

- Method chaining works naturally with pointer receivers
- Go's zero values make optional parameters easy to handle
- Struct literals can be an alternative for simple cases
- Functional options pattern is another Go idiom for similar problems
- Builder pattern works well with Go's composition over inheritance
