# Memento Pattern

## Real World Example
💾 Take the example of calculator (i.e. originator), where whenever you perform some calculation the last calculation is saved in memory (i.e. memento) so that you can get back to it and maybe get it restored using some action buttons (i.e. caretaker).

## In Plain Words
Memento pattern is about capturing and storing the current state of an object in a manner that it can be restored later on in a smooth manner.

## Wikipedia Definition
The memento pattern is a software design pattern that provides the ability to restore an object to its previous state (undo via rollback).

Usually useful when you need to provide some sort of undo functionality.

## Intent
Without violating encapsulation, capture and externalize an object's internal state so that the object can be restored to this state later.

## Problem Solved
When you need to provide undo functionality or save/restore object states without exposing the object's internal structure.

## Go Implementation

```go
// Memento
type EditorMemento struct {
    content string
}

func NewEditorMemento(content string) EditorMemento {
    return EditorMemento{content: content}
}

func (em EditorMemento) GetContent() string {
    return em.content
}

// Originator
type Editor struct {
    content string
}

func (e *Editor) Type(words string) {
    e.content = e.content + " " + words
}

func (e *Editor) GetContent() string {
    return e.content
}

func (e *Editor) Save() EditorMemento {
    return NewEditorMemento(e.content)
}

func (e *Editor) Restore(memento EditorMemento) {
    e.content = memento.GetContent()
}

// Caretaker
type History struct {
    mementos []EditorMemento
}

func (h *History) Push(memento EditorMemento) {
    h.mementos = append(h.mementos, memento)
}

func (h *History) Pop() EditorMemento {
    if len(h.mementos) == 0 {
        return EditorMemento{}
    }
    
    lastIndex := len(h.mementos) - 1
    memento := h.mementos[lastIndex]
    h.mementos = h.mementos[:lastIndex]
    return memento
}
```

## Key Features

1. **State Capture**: Captures object state without exposing internal structure
2. **State Restoration**: Allows restoration to previous states
3. **Encapsulation**: Maintains object encapsulation principles

## Benefits

- **Undo Functionality**: Easy to implement undo/redo operations
- **State Management**: Clean way to manage object states
- **Encapsulation**: Object internals remain hidden

## When to Use

- When you need to implement undo functionality
- When you need to create snapshots of object state
- When direct access to object fields would violate encapsulation

## Example Usage

```go
func main() {
    editor := &Editor{}

    // Type some stuff
    editor.Type("This is the first sentence.")
    editor.Type("This is second.")

    // Save the state to restore to : This is the first sentence. This is second.
    saved := editor.Save()

    // Type some more
    editor.Type("And this is third.")

    fmt.Println("Content before restoring:", editor.GetContent())
    // Output: Content before restoring:  This is the first sentence. This is second. And this is third.

    // Restoring to last saved state
    editor.Restore(saved)

    fmt.Println("Content after restoring:", editor.GetContent())
    // Output: Content after restoring:  This is the first sentence. This is second.
}
```
