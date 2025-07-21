# Visitor Pattern

## Real World Example
🏃 Consider someone visiting Dubai. They just need a way (i.e. visa) to enter Dubai. After arrival, they can come and visit any place in Dubai on their own without having to ask for permission or to do some leg work in order to visit any place here; just let them know of a place and they can visit it. Visitor pattern lets you do just that, it helps you add places to visit so that they can visit as much as they can without having to do any legwork.

## In Plain Words
Visitor pattern lets you add further operations to objects without having to modify them.

## Wikipedia Definition
In object-oriented programming and software engineering, the visitor design pattern is a way of separating an algorithm from an object structure on which it operates. A practical result of this separation is the ability to add new operations to existing object structures without modifying those structures. It is one way to follow the open/closed principle.

## Intent
Represent an operation to be performed on the elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operates.

## Problem Solved
When you need to perform operations on a group of similar kind of objects and want to avoid polluting their classes with these operations.

## Go Implementation

```go
// Visitor interface
type AnimalOperation interface {
    VisitMonkey(monkey Monkey) string
    VisitLion(lion Lion) string  
    VisitDolphin(dolphin Dolphin) string
}

// Element interface  
type Animal interface {
    Accept(operation AnimalOperation) string
}

// Concrete elements
type Monkey struct{}

func (m Monkey) Shout() string {
    return "Ooh oo aa aa!"
}

func (m Monkey) Accept(operation AnimalOperation) string {
    return operation.VisitMonkey(m)
}

type Lion struct{}

func (l Lion) Roar() string {
    return "Roaaar!"
}

func (l Lion) Accept(operation AnimalOperation) string {
    return operation.VisitLion(l)
}

type Dolphin struct{}

func (d Dolphin) Speak() string {
    return "Tuut tuttu tuutt!"
}

func (d Dolphin) Accept(operation AnimalOperation) string {
    return operation.VisitDolphin(d)
}

// Concrete visitors
type Speak struct{}

func (s Speak) VisitMonkey(monkey Monkey) string {
    return monkey.Shout()
}

func (s Speak) VisitLion(lion Lion) string {
    return lion.Roar()
}

func (s Speak) VisitDolphin(dolphin Dolphin) string {
    return dolphin.Speak()
}

type Jump struct{}

func (j Jump) VisitMonkey(monkey Monkey) string {
    return "Jumped 20 feet high! on to the tree!"
}

func (j Jump) VisitLion(lion Lion) string {
    return "Jumped 7 feet! Back on the ground!"
}

func (j Jump) VisitDolphin(dolphin Dolphin) string {
    return "Walked on water a little and disappeared"
}
```

## Key Features

1. **Operation Separation**: Operations are separated from object structure
2. **Easy Extension**: New operations can be added without modifying existing classes
3. **Type Safety**: Compile-time type checking for operations

## Benefits

- **Open/Closed Principle**: Open for extension, closed for modification
- **Single Responsibility**: Each visitor has a single responsibility
- **Easy Maintenance**: Adding new operations doesn't affect existing code

## When to Use

- When you need to perform operations on objects of different types
- When you want to add new operations without modifying existing classes
- When the object structure is stable but operations change frequently

## Example Usage

```go
func main() {
    monkey := Monkey{}
    lion := Lion{}
    dolphin := Dolphin{}

    speak := Speak{}
    jump := Jump{}

    fmt.Println("Making animals speak:")
    fmt.Println("Monkey:", monkey.Accept(speak))
    fmt.Println("Lion:", lion.Accept(speak))
    fmt.Println("Dolphin:", dolphin.Accept(speak))

    fmt.Println("\nMaking animals jump:")
    fmt.Println("Monkey:", monkey.Accept(jump))
    fmt.Println("Lion:", lion.Accept(jump))
    fmt.Println("Dolphin:", dolphin.Accept(jump))
}
```

## Output
```
Making animals speak:
Monkey: Ooh oo aa aa!
Lion: Roaaar!
Dolphin: Tuut tuttu tuutt!

Making animals jump:
Monkey: Jumped 20 feet high! on to the tree!
Lion: Jumped 7 feet! Back on the ground!
Dolphin: Walked on water a little and disappeared
```
