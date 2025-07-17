# State Pattern

## Intent
Allow an object to alter its behavior when its internal state changes. The object will appear to change its class.

## Problem Solved
When an object's behavior depends on its state and it must change its behavior at runtime depending on that state.

## Go Implementation

```go
// State interface
type WritingState interface {
    Write(words string) string
}

// Concrete states
type UpperCase struct{}

func (u UpperCase) Write(words string) string {
    return strings.ToUpper(words)
}

type LowerCase struct{}

func (l LowerCase) Write(words string) string {
    return strings.ToLower(words)
}

// Context
type TextEditor struct {
    state WritingState
}

func (te *TextEditor) SetState(state WritingState) {
    te.state = state
}

func (te *TextEditor) Type(words string) string {
    return te.state.Write(words)
}
```

## Key Features

1. **State-dependent Behavior**: Behavior changes based on internal state
2. **State Encapsulation**: Each state encapsulates its behavior
3. **Runtime State Changes**: States can be changed at runtime

## Benefits

- **Eliminates Conditionals**: Removes large conditional statements
- **Clean State Transitions**: Makes state transitions explicit
- **Easy to Extend**: New states can be added easily
- **State-specific Behavior**: Each state handles its own behavior

## When to Use

- When an object's behavior depends on its state
- When you have many conditional statements based on object state
- When state transitions are complex
- When you want to make state transitions explicit

## Real-world Examples

- **State Machines**: Traffic lights, vending machines
- **Game Development**: Character states (idle, running, jumping)
- **UI Components**: Button states (normal, pressed, disabled)
- **Connection States**: Network connection states

## Go-Specific Notes

- Go's interfaces make state pattern implementation clean
- Go's composition works well with state pattern
- State can be implemented using function types for simple cases
- Go's type system ensures state contracts are met
- Context can be passed through states for additional data
