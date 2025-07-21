# Template Method Pattern

## Real World Example
🏗️ Suppose we are building houses. The steps for building might look like:
- Prepare the base of house
- Build the walls  
- Add roof
- Add other floors

The order of these steps can never be changed i.e. you can't build the roof before building the walls etc but each of the steps can be modified for example walls can be made of wood or polyester or stone.

## In Plain Words
Template method defines the skeleton of how a certain algorithm could be performed, but defers the implementation of those steps to the children classes.

## Wikipedia Definition
In software engineering, the template method pattern is a behavioral design pattern that defines the program skeleton of an algorithm in an operation, deferring some steps to subclasses. It lets one redefine certain steps of an algorithm without changing the algorithm's structure.

## Intent
Define the skeleton of an algorithm in an operation, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.

## Problem Solved
When you have an algorithm with a fixed structure but want to allow subclasses to override specific steps of the algorithm.

## Go Implementation

```go
// Abstract builder interface
type Builder interface {
    Test() string
    Lint() string
    Assemble() string
    Deploy() string
    // Template method
    Build() string
}

// Base builder with template method
type BaseBuilder struct{}

func (bb BaseBuilder) Build() string {
    var result strings.Builder
    
    result.WriteString("Running build process:\n")
    result.WriteString("1. " + bb.Test() + "\n")
    result.WriteString("2. " + bb.Lint() + "\n") 
    result.WriteString("3. " + bb.Assemble() + "\n")
    result.WriteString("4. " + bb.Deploy() + "\n")
    result.WriteString("Build completed!")
    
    return result.String()
}

// Default implementations
func (bb BaseBuilder) Test() string {
    return "Running tests"
}

func (bb BaseBuilder) Lint() string {
    return "Running linter"
}

func (bb BaseBuilder) Assemble() string {
    return "Assembling the build"
}

func (bb BaseBuilder) Deploy() string {
    return "Deploying to server"
}

// Concrete builders
type AndroidBuilder struct {
    BaseBuilder
}

func (ab AndroidBuilder) Test() string {
    return "Running android tests"
}

func (ab AndroidBuilder) Lint() string {
    return "Running android linter"
}

func (ab AndroidBuilder) Assemble() string {
    return "Assembling the android build"
}

func (ab AndroidBuilder) Deploy() string {
    return "Deploying android build to server"
}

type IosBuilder struct {
    BaseBuilder
}

func (ib IosBuilder) Test() string {
    return "Running ios tests"
}

func (ib IosBuilder) Lint() string {
    return "Running ios linter"
}

func (ib IosBuilder) Assemble() string {
    return "Assembling the ios build"
}

func (ib IosBuilder) Deploy() string {
    return "Deploying ios build to server"
}
```

## Key Features

1. **Algorithm Structure**: Defines the skeleton of an algorithm
2. **Step Customization**: Allows subclasses to override specific steps
3. **Code Reuse**: Common algorithm structure is reused

## Benefits

- **Code Reuse**: Common parts of the algorithm are implemented once
- **Control**: Parent class controls the algorithm structure
- **Flexibility**: Subclasses can customize specific steps

## When to Use

- When you have an algorithm with a fixed structure but varying implementations
- When you want to avoid code duplication in similar algorithms
- When you want to control the points at which subclasses can extend behavior

## Example Usage

```go
func main() {
    androidBuilder := AndroidBuilder{}
    fmt.Println(androidBuilder.Build())
    
    fmt.Println("\n" + strings.Repeat("-", 40) + "\n")
    
    iosBuilder := IosBuilder{}
    fmt.Println(iosBuilder.Build())
}
```
