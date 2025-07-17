package main

import (
	"fmt"
	"strings"
)

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

type DefaultText struct{}

func (d DefaultText) Write(words string) string {
	return words
}

// Context
type TextEditor struct {
	state WritingState
}

func NewTextEditor(state WritingState) *TextEditor {
	return &TextEditor{state: state}
}

func (te *TextEditor) SetState(state WritingState) {
	te.state = state
}

func (te *TextEditor) Type(words string) string {
	return te.state.Write(words)
}

func main() {
	fmt.Println("=== State Pattern Demo ===")
	
	editor := NewTextEditor(DefaultText{})
	
	fmt.Println("Default state:")
	fmt.Println(editor.Type("First line"))
	
	editor.SetState(UpperCase{})
	fmt.Println("\nUpper case state:")
	fmt.Println(editor.Type("Second line"))
	
	editor.SetState(LowerCase{})
	fmt.Println("\nLower case state:")
	fmt.Println(editor.Type("Third line"))
	
	fmt.Println("\nState pattern allows object behavior to change based on internal state!")
}
