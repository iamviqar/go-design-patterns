# Adapter Pattern

## Real World Example
🔌 Consider that you have some pictures in your memory card and you need to transfer them to your computer. In order to transfer them you need some kind of adapter that is compatible with your computer ports so that you can attach memory card to your computer. In this case card reader is an adapter.

Another example would be the famous power adapter; a three legged plug can't be connected to a two pronged outlet, it needs to use a power adapter that makes it compatible with the two pronged outlet.

Yet another example would be a translator translating words spoken by one person to another.

## In Plain Words
Adapter pattern lets you wrap an otherwise incompatible object in an adapter to make it compatible with another class.

## Wikipedia Definition
In software engineering, the adapter pattern is a software design pattern that allows the interface of an existing class to be used as another interface. It is often used to make existing classes work with others without modifying their source code.

## Intent
Convert the interface of a class into another interface that clients expect. Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.

## Problem Solved
When you have existing code that you can't modify, but you need it to work with code that expects a different interface.

## Go Implementation

```go
// Target interface
type Lion interface {
    Roar() string
}

// Existing incompatible interface
type WildDog struct{}

func (w WildDog) Bark() string {
    return "Woof woof!"
}

// Adapter
type WildDogAdapter struct {
    wildDog WildDog
}

func (w WildDogAdapter) Roar() string {
    return w.wildDog.Bark()
}
```

## Key Features

1. **Interface Conversion**: Converts one interface to another
2. **Code Reuse**: Allows reuse of existing incompatible code
3. **Decoupling**: Keeps client code separate from adapted code

## Benefits

- **Legacy Integration**: Integrate legacy code with new systems
- **Third-party Integration**: Work with third-party libraries with different interfaces
- **Code Reuse**: Reuse existing code without modification
- **Clean Separation**: Keep adaptation logic separate

## When to Use

- When you want to use an existing class with an incompatible interface
- When you want to create a reusable class that cooperates with unrelated classes
- When you need to use several existing subclasses, but it's impractical to adapt their interface by subclassing

## Real-world Examples

- **Database Drivers**: Adapting different database APIs to a common interface
- **Payment Gateways**: Adapting different payment providers to unified interface
- **File System**: Adapting different file systems (local, cloud, FTP) to common interface
- **Media Players**: Adapting different audio/video formats to common player interface

## Go-Specific Notes

- Go's implicit interface satisfaction makes adapters very natural
- Composition is preferred over inheritance for adapters
- Go's struct embedding can be used for simple adaptations
- Interface{} can be used for generic adapters when appropriate
- Go's type system helps ensure adapter contracts are met
