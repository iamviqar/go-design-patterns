package main

import "fmt"

// Burger is the product being built
type Burger struct {
	Size       int
	Cheese     bool
	Pepperoni  bool
	Lettuce    bool
	Tomato     bool
}

func (b Burger) String() string {
	return fmt.Sprintf("Burger{size=%d, cheese=%t, pepperoni=%t, lettuce=%t, tomato=%t}",
		b.Size, b.Cheese, b.Pepperoni, b.Lettuce, b.Tomato)
}

// BurgerBuilder builds burgers step by step
type BurgerBuilder struct {
	burger Burger
}

// NewBurgerBuilder creates a new burger builder
func NewBurgerBuilder(size int) *BurgerBuilder {
	return &BurgerBuilder{
		burger: Burger{Size: size},
	}
}

// AddCheese adds cheese to the burger
func (bb *BurgerBuilder) AddCheese() *BurgerBuilder {
	bb.burger.Cheese = true
	return bb
}

// AddPepperoni adds pepperoni to the burger
func (bb *BurgerBuilder) AddPepperoni() *BurgerBuilder {
	bb.burger.Pepperoni = true
	return bb
}

// AddLettuce adds lettuce to the burger
func (bb *BurgerBuilder) AddLettuce() *BurgerBuilder {
	bb.burger.Lettuce = true
	return bb
}

// AddTomato adds tomato to the burger
func (bb *BurgerBuilder) AddTomato() *BurgerBuilder {
	bb.burger.Tomato = true
	return bb
}

// Build returns the final burger
func (bb *BurgerBuilder) Build() Burger {
	return bb.burger
}

func main() {
	fmt.Println("=== Builder Pattern Demo ===")
	
	// Build a custom burger using method chaining
	customBurger := NewBurgerBuilder(14).
		AddPepperoni().
		AddLettuce().
		AddTomato().
		Build()
	
	fmt.Printf("Custom Burger: %s\n", customBurger)
	
	// Build a simple cheese burger
	cheeseBurger := NewBurgerBuilder(10).
		AddCheese().
		AddTomato().
		Build()
	
	fmt.Printf("Cheese Burger: %s\n", cheeseBurger)
	
	// Build a simple burger with no toppings
	simpleBurger := NewBurgerBuilder(8).Build()
	
	fmt.Printf("Simple Burger: %s\n", simpleBurger)
	
	fmt.Println("\nBuilder pattern avoids telescoping constructor anti-pattern!")
}
