package main

import "fmt"

// Abstract product interfaces
type Door interface {
	GetDescription() string
}

type DoorFittingExpert interface {
	GetDescription() string
}

// Concrete wooden products
type WoodenDoor struct{}

func (w WoodenDoor) GetDescription() string {
	return "I am a wooden door"
}

type Welder struct{}

func (w Welder) GetDescription() string {
	return "I can only fit iron doors"
}

// Concrete iron products
type IronDoor struct{}

func (i IronDoor) GetDescription() string {
	return "I am an iron door"
}

type Carpenter struct{}

func (c Carpenter) GetDescription() string {
	return "I can only fit wooden doors"
}

// Abstract factory interface
type DoorFactory interface {
	MakeDoor() Door
	MakeFittingExpert() DoorFittingExpert
}

// Concrete wooden factory
type WoodenDoorFactory struct{}

func (w WoodenDoorFactory) MakeDoor() Door {
	return WoodenDoor{}
}

func (w WoodenDoorFactory) MakeFittingExpert() DoorFittingExpert {
	return Carpenter{}
}

// Concrete iron factory
type IronDoorFactory struct{}

func (i IronDoorFactory) MakeDoor() Door {
	return IronDoor{}
}

func (i IronDoorFactory) MakeFittingExpert() DoorFittingExpert {
	return Welder{}
}

func main() {
	fmt.Println("=== Abstract Factory Pattern Demo ===")
	
	// Create wooden door family
	woodenFactory := WoodenDoorFactory{}
	woodenDoor := woodenFactory.MakeDoor()
	woodenExpert := woodenFactory.MakeFittingExpert()
	
	fmt.Printf("Wooden Door: %s\n", woodenDoor.GetDescription())
	fmt.Printf("Wooden Expert: %s\n", woodenExpert.GetDescription())
	
	fmt.Println()
	
	// Create iron door family
	ironFactory := IronDoorFactory{}
	ironDoor := ironFactory.MakeDoor()
	ironExpert := ironFactory.MakeFittingExpert()
	
	fmt.Printf("Iron Door: %s\n", ironDoor.GetDescription())
	fmt.Printf("Iron Expert: %s\n", ironExpert.GetDescription())
	
	fmt.Println("\nAbstract Factory creates families of related objects!")
}
