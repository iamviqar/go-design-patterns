package main

import "fmt"

// Door interface defines what a door can do
type Door interface {
	GetDescription() string
	GetWidth() float64
	GetHeight() float64
}

// WoodenDoor is a concrete implementation of Door
type WoodenDoor struct {
	Width  float64
	Height float64
}

// GetDescription returns a description of the wooden door
func (w WoodenDoor) GetDescription() string {
	return fmt.Sprintf("WoodenDoor{width=%.1f, height=%.1f}", w.Width, w.Height)
}

// GetWidth returns the width of the door
func (w WoodenDoor) GetWidth() float64 {
	return w.Width
}

// GetHeight returns the height of the door
func (w WoodenDoor) GetHeight() float64 {
	return w.Height
}

// DoorFactory is the simple factory that creates doors
type DoorFactory struct{}

// MakeDoor creates a new wooden door with given dimensions
func (df DoorFactory) MakeDoor(width, height float64) Door {
	return WoodenDoor{Width: width, Height: height}
}

// NewDoorFactory creates a new door factory
func NewDoorFactory() DoorFactory {
	return DoorFactory{}
}

func main() {
	fmt.Println("=== Simple Factory Pattern Demo ===")
	
	factory := NewDoorFactory()
	
	// Create doors using the factory
	door1 := factory.MakeDoor(100, 200)
	door2 := factory.MakeDoor(50, 100)
	
	fmt.Printf("Door 1 - Width: %.1f, Height: %.1f\n", door1.GetWidth(), door1.GetHeight())
	fmt.Printf("Door 1: %s\n", door1.GetDescription())
	
	fmt.Printf("Door 2 - Width: %.1f, Height: %.1f\n", door2.GetWidth(), door2.GetHeight())
	fmt.Printf("Door 2: %s\n", door2.GetDescription())
	
	fmt.Println("\nSimple Factory centralizes object creation logic!")
}
