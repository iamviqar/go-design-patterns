# Command Pattern

## Intent
Encapsulate a request as an object, thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.

## Problem Solved
When you want to decouple the object that invokes the operation from the object that performs it, and when you need to support undo operations, logging, or queueing.

## Go Implementation

```go
// Command interface
type Command interface {
    Execute() string
    Undo() string
}

// Receiver
type Bulb struct {
    isOn bool
}

func (b *Bulb) TurnOn() string {
    b.isOn = true
    return "Bulb has been lit"
}

// Concrete command
type TurnOnCommand struct {
    bulb *Bulb
}

func (t TurnOnCommand) Execute() string {
    return t.bulb.TurnOn()
}

func (t TurnOnCommand) Undo() string {
    return t.bulb.TurnOff()
}

// Invoker
type RemoteControl struct{}

func (r RemoteControl) Submit(command Command) string {
    return command.Execute()
}
```

## Key Features

1. **Encapsulation**: Encapsulates requests as objects
2. **Decoupling**: Decouples invoker from receiver
3. **Undo Support**: Can implement undo operations
4. **Queuing**: Commands can be queued and executed later

## Benefits

- **Decoupling**: Invoker and receiver are decoupled
- **Undo/Redo**: Easy to implement undo and redo operations
- **Logging**: Commands can be logged for audit trails
- **Macro Commands**: Can combine multiple commands
- **Queuing**: Commands can be queued and executed later

## When to Use

- When you want to parameterize objects by an action to perform
- When you want to queue operations, schedule their execution, or execute them remotely
- When you want to support undo operations
- When you want to log changes so they can be reapplied after a system crash

## Real-world Examples

- **Text Editors**: Undo/redo operations
- **GUI Applications**: Button clicks, menu selections
- **Database Operations**: Transaction logs
- **Task Scheduling**: Job queues and schedulers

## Go-Specific Notes

- Go's interfaces make command pattern implementation clean
- Function types can be used for simple commands
- Go's goroutines work well with command queuing
- Context can be passed through commands for cancellation
- Error handling can be naturally integrated into command execution
