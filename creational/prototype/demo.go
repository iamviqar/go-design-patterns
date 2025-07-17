package main

import "fmt"

// Prototype interface for cloning
type Prototype interface {
	Clone() Prototype
	GetDetails() string
}

// Sheep struct represents a sheep
type Sheep struct {
	Name     string
	Category string
}

// Clone creates a copy of the sheep
func (s *Sheep) Clone() Prototype {
	return &Sheep{
		Name:     s.Name,
		Category: s.Category,
	}
}

// GetDetails returns sheep details
func (s *Sheep) GetDetails() string {
	return fmt.Sprintf("%s is a %s sheep", s.Name, s.Category)
}

// SetName sets the sheep's name
func (s *Sheep) SetName(name string) {
	s.Name = name
}

// SetCategory sets the sheep's category
func (s *Sheep) SetCategory(category string) {
	s.Category = category
}

func main() {
	fmt.Println("=== Prototype Pattern Demo ===")
	
	// Create original sheep
	original := &Sheep{
		Name:     "Dolly",
		Category: "Mountain Sheep",
	}
	
	fmt.Printf("Original: %s\n", original.GetDetails())
	
	// Clone the sheep
	cloned := original.Clone().(*Sheep)
	cloned.SetName("Jolly")
	
	fmt.Printf("Cloned: %s\n", cloned.GetDetails())
	
	// Create another clone with different properties
	anotherClone := original.Clone().(*Sheep)
	anotherClone.SetName("Molly")
	anotherClone.SetCategory("Farm Sheep")
	
	fmt.Printf("Another Clone: %s\n", anotherClone.GetDetails())
	
	// Original remains unchanged
	fmt.Printf("Original (unchanged): %s\n", original.GetDetails())
	
	fmt.Println("\nPrototype pattern creates objects by cloning existing instances!")
}
