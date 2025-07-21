# Chain of Responsibility Pattern

## Real World Example
🔗 For example, you have three payment methods (`A`, `B` and `C`) setup in your account; each having a different amount in it. `A` has 100 USD, `B` has 300 USD and `C` having 1000 USD and the preference for payments is chosen as `A` then `B` then `C`. You try to purchase something that is worth 210 USD. Using Chain of Responsibility, first of all account `A` will be checked if it can make the purchase, if yes purchase will be made and the chain will be broken. If not, request will move forward to account `B` checking for amount if yes chain will be broken otherwise the request will keep forwarding till it finds the suitable handler. Here `A`, `B` and `C` are links of the chain and the whole phenomenon is Chain of Responsibility.

## In Plain Words
It helps building a chain of objects. Request enters from one end and keeps going from object to object till it finds the suitable handler.

## Wikipedia Definition
In object-oriented design, the chain-of-responsibility pattern is a design pattern consisting of a source of command objects and a series of processing objects. Each processing object contains logic that defines the types of command objects that it can handle; the rest are passed to the next processing object in the chain.

## Intent
Avoid coupling the sender of a request to its receiver by giving more than one object a chance to handle the request. Chain the receiving objects and pass the request along the chain until an object handles it.

## Problem Solved
When you want to decouple request senders from receivers and allow multiple objects to handle a request without the sender knowing which object will handle it.

## Go Implementation

```go
// Handler interface
type Account interface {
    SetNext(account Account)
    Pay(amount int) string
    CanPay(amount int) bool
}

// Base handler
type BaseAccount struct {
    successor Account
    balance   int
}

func (b *BaseAccount) SetNext(account Account) {
    b.successor = account
}

func (b *BaseAccount) Pay(amount int) string {
    if b.CanPay(amount) {
        // Handle the request
        return "Payment processed"
    } else if b.successor != nil {
        // Pass to next handler
        return b.successor.Pay(amount)
    }
    return "Cannot process payment"
}

// Concrete handler
type Bank struct {
    BaseAccount
}

func (b *Bank) CanPay(amount int) bool {
    return b.balance >= amount
}
```

## Key Features

1. **Decoupling**: Sender doesn't know which handler will process the request
2. **Dynamic Chain**: Chain can be modified at runtime
3. **Flexible Handling**: Multiple handlers can process the same request type

## Benefits

- **Reduced Coupling**: Sender and receiver are decoupled
- **Flexibility**: Can add/remove handlers dynamically
- **Responsibility Distribution**: Each handler has a specific responsibility
- **Open/Closed Principle**: Easy to add new handlers

## When to Use

- When you want to decouple request senders from receivers
- When multiple objects may handle a request and the handler isn't known beforehand
- When you want to issue a request to one of several objects without specifying the receiver explicitly
- When the set of objects that can handle a request should be specified dynamically

## Real-world Examples

- **Middleware**: HTTP middleware chains in web frameworks
- **Event Handling**: GUI event propagation
- **Logging**: Log level filtering and routing
- **Authentication**: Multi-factor authentication chains

## Go-Specific Notes

- Go's interfaces make chain implementation natural
- Composition is used instead of inheritance
- Go's HTTP middleware is a real-world example of this pattern
- Context can be passed through the chain for request-scoped data
- Error handling can be naturally integrated into the chain
