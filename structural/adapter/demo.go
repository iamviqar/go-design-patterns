package main

import "fmt"

// Lion interface - what we want to use
type Lion interface {
	Roar() string
}

// AfricanLion implements Lion
type AfricanLion struct{}

func (a AfricanLion) Roar() string {
	return "Roooaaar!"
}

// Hunter can hunt lions
type Hunter struct{}

func (h Hunter) Hunt(lion Lion) string {
	return "Hunter hunts: " + lion.Roar()
}

// WildDog - incompatible interface we need to adapt
type WildDog struct{}

func (w WildDog) Bark() string {
	return "Woof woof!"
}

// WildDogAdapter adapts WildDog to Lion interface
type WildDogAdapter struct {
	wildDog WildDog
}

// NewWildDogAdapter creates a new adapter
func NewWildDogAdapter(wildDog WildDog) WildDogAdapter {
	return WildDogAdapter{wildDog: wildDog}
}

// Roar adapts bark to roar
func (w WildDogAdapter) Roar() string {
	return w.wildDog.Bark()
}

func main() {
	fmt.Println("=== Adapter Pattern Demo ===")
	
	hunter := Hunter{}
	
	// Hunter can hunt African lions
	africanLion := AfricanLion{}
	fmt.Println(hunter.Hunt(africanLion))
	
	// Hunter cannot hunt wild dogs directly
	wildDog := WildDog{}
	fmt.Printf("Wild dog says: %s\n", wildDog.Bark())
	
	// But with adapter, hunter can hunt wild dogs too
	wildDogAdapter := NewWildDogAdapter(wildDog)
	fmt.Println(hunter.Hunt(wildDogAdapter))
	
	fmt.Println("\nAdapter allows incompatible interfaces to work together!")
}
