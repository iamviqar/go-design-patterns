package main

import "fmt"

// Account interface
type Account interface {
	SetNext(account Account)
	Pay(amount int) string
	CanPay(amount int) bool
	GetBalance() int
}

// BaseAccount provides common functionality
type BaseAccount struct {
	successor Account
	balance   int
}

func (b *BaseAccount) SetNext(account Account) {
	b.successor = account
}

func (b *BaseAccount) Pay(amount int) string {
	if b.CanPay(amount) {
		b.balance -= amount
		return fmt.Sprintf("Paid $%d using this account. Remaining balance: $%d", amount, b.balance)
	} else if b.successor != nil {
		return fmt.Sprintf("Cannot pay using this account. ") + b.successor.Pay(amount)
	} else {
		return "Insufficient funds in all accounts!"
	}
}

func (b *BaseAccount) GetBalance() int {
	return b.balance
}

// Bank account
type Bank struct {
	BaseAccount
}

func NewBank(balance int) *Bank {
	return &Bank{BaseAccount: BaseAccount{balance: balance}}
}

func (b *Bank) CanPay(amount int) bool {
	return b.balance >= amount
}

func (b *Bank) Pay(amount int) string {
	if b.CanPay(amount) {
		b.balance -= amount
		return fmt.Sprintf("Paid $%d using Bank. Remaining balance: $%d", amount, b.balance)
	} else if b.successor != nil {
		return fmt.Sprintf("Cannot pay $%d using Bank. ") + b.successor.Pay(amount)
	} else {
		return "Insufficient funds in all accounts!"
	}
}

// Paypal account
type Paypal struct {
	BaseAccount
}

func NewPaypal(balance int) *Paypal {
	return &Paypal{BaseAccount: BaseAccount{balance: balance}}
}

func (p *Paypal) CanPay(amount int) bool {
	return p.balance >= amount
}

func (p *Paypal) Pay(amount int) string {
	if p.CanPay(amount) {
		p.balance -= amount
		return fmt.Sprintf("Paid $%d using Paypal. Remaining balance: $%d", amount, p.balance)
	} else if p.successor != nil {
		return fmt.Sprintf("Cannot pay $%d using Paypal. ") + p.successor.Pay(amount)
	} else {
		return "Insufficient funds in all accounts!"
	}
}

// Bitcoin account
type Bitcoin struct {
	BaseAccount
}

func NewBitcoin(balance int) *Bitcoin {
	return &Bitcoin{BaseAccount: BaseAccount{balance: balance}}
}

func (b *Bitcoin) CanPay(amount int) bool {
	return b.balance >= amount
}

func (b *Bitcoin) Pay(amount int) string {
	if b.CanPay(amount) {
		b.balance -= amount
		return fmt.Sprintf("Paid $%d using Bitcoin. Remaining balance: $%d", amount, b.balance)
	} else if b.successor != nil {
		return fmt.Sprintf("Cannot pay $%d using Bitcoin. ") + b.successor.Pay(amount)
	} else {
		return "Insufficient funds in all accounts!"
	}
}

func main() {
	fmt.Println("=== Chain of Responsibility Pattern Demo ===")
	
	// Create accounts
	bank := NewBank(100)
	paypal := NewPaypal(200)
	bitcoin := NewBitcoin(300)
	
	// Set up the chain: bank -> paypal -> bitcoin
	bank.SetNext(paypal)
	paypal.SetNext(bitcoin)
	
	// Try different payment amounts
	fmt.Println("Payment of $50:")
	fmt.Println(bank.Pay(50))
	
	fmt.Println("\nPayment of $120:")
	fmt.Println(bank.Pay(120))
	
	fmt.Println("\nPayment of $350:")
	fmt.Println(bank.Pay(350))
	
	fmt.Println("\nPayment of $500:")
	fmt.Println(bank.Pay(500))
	
	fmt.Println("\nChain of Responsibility passes requests along a chain of handlers!")
}
