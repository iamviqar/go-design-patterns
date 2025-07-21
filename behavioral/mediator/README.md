# Mediator Pattern

## Real World Example
👽 A general example would be when you talk to someone on your mobile phone, there is a network provider sitting between you and them and your conversation goes through it instead of being directly sent. In this case network provider is mediator.

## In Plain Words
Mediator pattern adds a third party object (called mediator) to control the interaction between two objects (called colleagues). It helps reduce the coupling between the classes communicating with each other. Because now they don't need to have the knowledge of each other's implementation.

## Wikipedia Definition
In software engineering, the mediator pattern defines an object that encapsulates how a set of objects interact. This pattern is considered to be a behavioral pattern due to the way it can alter the program's running behavior.

## Intent
Define an object that encapsulates how a set of objects interact. Mediator promotes loose coupling by keeping objects from referring to each other explicitly, and lets you vary their interaction independently.

## Problem Solved
When you want to reduce dependencies between communicating objects and centralize complex communications and control logic.

## Go Implementation

```go
// Mediator interface
type ChatRoomMediator interface {
    ShowMessage(user User, message string) string
}

// Concrete mediator
type ChatRoom struct{}

func (cr ChatRoom) ShowMessage(user User, message string) string {
    time := "Jan 1, 2022 12:00"
    sender := user.GetName()
    return fmt.Sprintf("%s [%s]: %s", time, sender, message)
}

// Colleague
type User struct {
    name      string
    mediator  ChatRoomMediator
}

func NewUser(name string, mediator ChatRoomMediator) User {
    return User{name: name, mediator: mediator}
}

func (u User) GetName() string {
    return u.name
}

func (u User) Send(message string) string {
    return u.mediator.ShowMessage(u, message)
}
```

## Key Features

1. **Decoupling**: Colleagues don't reference each other directly
2. **Centralized Control**: All communication logic is in the mediator
3. **Reusability**: Mediator can be reused with different colleague configurations

## Benefits

- **Loose Coupling**: Objects don't need to know about each other
- **Centralized Logic**: Communication logic is in one place
- **Easy Maintenance**: Changes to interaction patterns only affect the mediator

## When to Use

- When a set of objects communicate in well-defined but complex ways
- When reusing an object is difficult because it refers to many other objects
- When behavior distributed between several classes should be customizable without subclassing

## Example Usage

```go
func main() {
    mediator := ChatRoom{}

    john := NewUser("John Doe", mediator)
    jane := NewUser("Jane Doe", mediator)

    fmt.Println(john.Send("Hi there!"))
    fmt.Println(jane.Send("Hey!"))
    
    // Output:
    // Jan 1, 2022 12:00 [John Doe]: Hi there!
    // Jan 1, 2022 12:00 [Jane Doe]: Hey!
}
```
