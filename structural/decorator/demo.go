package main

import "fmt"

// Coffee interface
type Coffee interface {
	GetCost() float64
	GetDescription() string
}

// SimpleCoffee - base component
type SimpleCoffee struct{}

func (s SimpleCoffee) GetCost() float64 {
	return 10.0
}

func (s SimpleCoffee) GetDescription() string {
	return "Simple coffee"
}

// MilkDecorator
type MilkDecorator struct {
	coffee Coffee
}

func NewMilkDecorator(coffee Coffee) MilkDecorator {
	return MilkDecorator{coffee: coffee}
}

func (m MilkDecorator) GetCost() float64 {
	return m.coffee.GetCost() + 2.0
}

func (m MilkDecorator) GetDescription() string {
	return m.coffee.GetDescription() + ", milk"
}

// WhipDecorator
type WhipDecorator struct {
	coffee Coffee
}

func NewWhipDecorator(coffee Coffee) WhipDecorator {
	return WhipDecorator{coffee: coffee}
}

func (w WhipDecorator) GetCost() float64 {
	return w.coffee.GetCost() + 5.0
}

func (w WhipDecorator) GetDescription() string {
	return w.coffee.GetDescription() + ", whip"
}

// VanillaDecorator
type VanillaDecorator struct {
	coffee Coffee
}

func NewVanillaDecorator(coffee Coffee) VanillaDecorator {
	return VanillaDecorator{coffee: coffee}
}

func (v VanillaDecorator) GetCost() float64 {
	return v.coffee.GetCost() + 3.0
}

func (v VanillaDecorator) GetDescription() string {
	return v.coffee.GetDescription() + ", vanilla"
}

func main() {
	fmt.Println("=== Decorator Pattern Demo ===")
	
	// Simple coffee
	coffee := SimpleCoffee{}
	fmt.Printf("Cost: $%.1f, Description: %s\n", coffee.GetCost(), coffee.GetDescription())
	
	// Add milk
	coffeeWithMilk := NewMilkDecorator(coffee)
	fmt.Printf("Cost: $%.1f, Description: %s\n", coffeeWithMilk.GetCost(), coffeeWithMilk.GetDescription())
	
	// Add whip
	coffeeWithMilkAndWhip := NewWhipDecorator(coffeeWithMilk)
	fmt.Printf("Cost: $%.1f, Description: %s\n", coffeeWithMilkAndWhip.GetCost(), coffeeWithMilkAndWhip.GetDescription())
	
	// Add vanilla
	fancyCoffee := NewVanillaDecorator(coffeeWithMilkAndWhip)
	fmt.Printf("Cost: $%.1f, Description: %s\n", fancyCoffee.GetCost(), fancyCoffee.GetDescription())
	
	// Create different combination
	specialCoffee := NewVanillaDecorator(NewMilkDecorator(SimpleCoffee{}))
	fmt.Printf("\nSpecial Coffee - Cost: $%.1f, Description: %s\n", specialCoffee.GetCost(), specialCoffee.GetDescription())
	
	fmt.Println("\nDecorator pattern allows adding behavior dynamically!")
}
