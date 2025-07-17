package main

import "fmt"

// Command interface
type Command interface {
	Execute() string
	Undo() string
}

// Receiver - Bulb
type Bulb struct {
	isOn bool
}

func (b *Bulb) TurnOn() string {
	b.isOn = true
	return "Bulb has been lit"
}

func (b *Bulb) TurnOff() string {
	b.isOn = false
	return "Darkness!"
}

func (b *Bulb) IsOn() bool {
	return b.isOn
}

// Concrete Commands
type TurnOnCommand struct {
	bulb *Bulb
}

func NewTurnOnCommand(bulb *Bulb) TurnOnCommand {
	return TurnOnCommand{bulb: bulb}
}

func (t TurnOnCommand) Execute() string {
	return t.bulb.TurnOn()
}

func (t TurnOnCommand) Undo() string {
	return t.bulb.TurnOff()
}

type TurnOffCommand struct {
	bulb *Bulb
}

func NewTurnOffCommand(bulb *Bulb) TurnOffCommand {
	return TurnOffCommand{bulb: bulb}
}

func (t TurnOffCommand) Execute() string {
	return t.bulb.TurnOff()
}

func (t TurnOffCommand) Undo() string {
	return t.bulb.TurnOn()
}

// Invoker - RemoteControl
type RemoteControl struct{}

func (r RemoteControl) Submit(command Command) string {
	return command.Execute()
}

func (r RemoteControl) Undo(command Command) string {
	return command.Undo()
}

func main() {
	fmt.Println("=== Command Pattern Demo ===")
	
	bulb := &Bulb{}
	remote := RemoteControl{}
	
	// Create commands
	turnOn := NewTurnOnCommand(bulb)
	turnOff := NewTurnOffCommand(bulb)
	
	// Execute commands
	fmt.Printf("Bulb is on: %t\n", bulb.IsOn())
	
	fmt.Println(remote.Submit(turnOn))
	fmt.Printf("Bulb is on: %t\n", bulb.IsOn())
	
	fmt.Println(remote.Submit(turnOff))
	fmt.Printf("Bulb is on: %t\n", bulb.IsOn())
	
	// Undo commands
	fmt.Println("\nUndo last command:")
	fmt.Println(remote.Undo(turnOff))
	fmt.Printf("Bulb is on: %t\n", bulb.IsOn())
	
	fmt.Println("\nUndo turn on:")
	fmt.Println(remote.Undo(turnOn))
	fmt.Printf("Bulb is on: %t\n", bulb.IsOn())
	
	fmt.Println("\nCommand pattern encapsulates requests as objects!")
}
