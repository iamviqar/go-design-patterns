# Builder Pattern

## Real World Example
👷 Imagine you are at Hardee's and you order a specific deal, lets say, "Big Hardee" and they hand it over to you without *any questions*; this is the example of simple factory. But there are cases when the creation logic might involve more steps. For example you want a customized Subway deal, you have several options in how your burger is made e.g what bread do you want? what types of sauces would you like? What cheese would you want? etc. In such cases builder pattern comes to the rescue.

## In Plain Words
Allows you to create different flavors of an object while avoiding constructor pollution. Useful when there could be several flavors of an object. Or when there are a lot of steps involved in creation of an object.

## Wikipedia Definition
The builder pattern is an object creation software design pattern with the intentions of finding a solution to the telescoping constructor anti-pattern.

## The Telescoping Constructor Anti-Pattern
At one point or the other we have all seen a constructor like below:
```go
func NewBurger(size int, cheese bool, pepperoni bool, tomato bool, lettuce bool) *Burger
```
As you can see; the number of constructor parameters can quickly get out of hand and it might become difficult to understand the arrangement of parameters. Plus this parameter list could keep on growing if you would want to add more options in future. This is called telescoping constructor anti-pattern.

## Intent
Separate the construction of a complex object from its representation so that the same construction process can create different representations.

## Problem Solved
When there could be several flavors of an object and to avoid the constructor telescoping. The key difference from the factory pattern is that; factory pattern is to be used when the creation is a one step process while builder pattern is to be used when the creation is a multi step process.

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
