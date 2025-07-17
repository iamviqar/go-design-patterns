# Observer Pattern

## Intent
Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

## Problem Solved
When you need to notify multiple objects about state changes in another object without making them tightly coupled.

## Go Implementation

```go
// Observer interface
type Observer interface {
    Update(jobTitle string)
    GetName() string
}

// Subject interface
type Subject interface {
    Attach(observer Observer)
    Detach(observer Observer)
    Notify(jobTitle string)
}

// Concrete observer
type JobSeeker struct {
    name string
}

func (js *JobSeeker) Update(jobTitle string) {
    fmt.Printf("Hi %s! New job posted: %s\n", js.name, jobTitle)
}

// Concrete subject
type JobPostings struct {
    observers []Observer
}

func (jp *JobPostings) Attach(observer Observer) {
    jp.observers = append(jp.observers, observer)
}

func (jp *JobPostings) Notify(jobTitle string) {
    for _, observer := range jp.observers {
        observer.Update(jobTitle)
    }
}
```

## Key Features

1. **One-to-Many**: One subject can have multiple observers
2. **Loose Coupling**: Subject and observers are loosely coupled
3. **Dynamic Relationships**: Observers can be added/removed at runtime

## Benefits

- **Loose Coupling**: Subject and observers are independent
- **Dynamic Subscription**: Observers can subscribe/unsubscribe at runtime
- **Broadcast Communication**: Subject can notify all observers simultaneously
- **Open/Closed Principle**: New observers can be added without modifying subject

## When to Use

- When changes to one object require changing many others
- When an object should notify others without knowing who they are
- When you want to decouple the sender and receiver of notifications
- When you need to implement event handling systems

## Real-world Examples

- **Event Systems**: GUI event handling, game events
- **Model-View-Controller**: Model notifies views of changes
- **Pub/Sub Systems**: Message broadcasting systems
- **Stock Market**: Price changes notify multiple displays

## Go-Specific Notes

- Go's channels provide a natural way to implement observer pattern
- Goroutines can be used for asynchronous notifications
- Go's interfaces make observer implementation flexible
- Context can be used for cancellation and timeouts
- sync package can be used for thread-safe observer management
