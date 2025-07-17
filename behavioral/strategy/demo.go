package main

import (
	"fmt"
	"sort"
)

// SortStrategy interface
type SortStrategy interface {
	Sort(data []int) []int
	GetName() string
}

// BubbleSort strategy
type BubbleSort struct{}

func (bs BubbleSort) GetName() string {
	return "bubble sort"
}

func (bs BubbleSort) Sort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	
	n := len(result)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// QuickSort strategy (using Go's built-in sort)
type QuickSort struct{}

func (qs QuickSort) GetName() string {
	return "quick sort"
}

func (qs QuickSort) Sort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	sort.Ints(result)
	return result
}

// Sorter context
type Sorter struct {
	strategy SortStrategy
}

func NewSorter(strategy SortStrategy) *Sorter {
	return &Sorter{strategy: strategy}
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(data []int) []int {
	fmt.Printf("Sorting using %s\n", s.strategy.GetName())
	return s.strategy.Sort(data)
}

func main() {
	fmt.Println("=== Strategy Pattern Demo ===")
	
	smallDataset := []int{1, 3, 4, 2}
	largeDataset := []int{1, 4, 3, 2, 8, 10, 5, 6, 9, 7}
	
	// Create sorter with bubble sort strategy
	sorter := NewSorter(BubbleSort{})
	
	// Sort small dataset with bubble sort
	fmt.Printf("Small dataset: %v\n", smallDataset)
	sorted := sorter.Sort(smallDataset)
	fmt.Printf("Sorted: %v\n", sorted)
	
	fmt.Println()
	
	// Switch to quick sort for large dataset
	fmt.Printf("Large dataset: %v\n", largeDataset)
	sorter.SetStrategy(QuickSort{})
	sorted = sorter.Sort(largeDataset)
	fmt.Printf("Sorted: %v\n", sorted)
	
	fmt.Println("\nStrategy pattern allows switching algorithms at runtime!")
}
