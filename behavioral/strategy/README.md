# Strategy Pattern

## Intent
Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

## Problem Solved
When you have multiple ways to perform a task and want to switch between them at runtime, or when you want to avoid conditional statements for algorithm selection.

## Go Implementation

```go
// Strategy interface
type SortStrategy interface {
    Sort(data []int) []int
    GetName() string
}

// Concrete strategies
type BubbleSort struct{}

func (bs BubbleSort) Sort(data []int) []int {
    // Bubble sort implementation
    result := make([]int, len(data))
    copy(result, data)
    // ... sorting logic
    return result
}

type QuickSort struct{}

func (qs QuickSort) Sort(data []int) []int {
    // Quick sort implementation
    result := make([]int, len(data))
    copy(result, data)
    sort.Ints(result)
    return result
}

// Context
type Sorter struct {
    strategy SortStrategy
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
    s.strategy = strategy
}

func (s *Sorter) Sort(data []int) []int {
    return s.strategy.Sort(data)
}
```

## Key Features

1. **Algorithm Family**: Defines a family of algorithms
2. **Interchangeable**: Algorithms can be switched at runtime
3. **Encapsulation**: Each algorithm is encapsulated in its own class

## Benefits

- **Flexibility**: Algorithms can be switched at runtime
- **Open/Closed Principle**: New algorithms can be added without modifying existing code
- **Eliminates Conditionals**: Removes algorithm selection conditionals
- **Reusability**: Algorithms can be reused in different contexts

## When to Use

- When you have multiple ways to perform a task
- When you want to switch algorithms at runtime
- When you want to avoid large conditional statements
- When you want to hide algorithm implementation details from clients

## Real-world Examples

- **Sorting Algorithms**: Different sorting strategies based on data size
- **Payment Processing**: Different payment methods (credit card, PayPal, cryptocurrency)
- **Compression**: Different compression algorithms based on file type
- **Routing**: Different pathfinding algorithms in navigation systems

## Go-Specific Notes

- Go's interfaces make strategy pattern implementation natural
- Function types can be used for simple strategies
- Go's composition over inheritance fits perfectly with strategy pattern
- Context can be passed to strategies for cancellation
- Go's type system ensures strategy contracts are met
