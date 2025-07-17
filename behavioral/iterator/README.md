# Iterator Pattern

## Intent
Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.

## Problem Solved
When you need to traverse a collection without exposing its internal structure or when you want to support multiple traversals of collections.

## Go Implementation

```go
// Iterator interface
type Iterator interface {
    HasNext() bool
    Next() interface{}
}

// Collection
type StationList struct {
    stations []RadioStation
}

func (sl *StationList) GetIterator() Iterator {
    return &StationListIterator{
        stations: sl.stations,
        index:    0,
    }
}

// Concrete iterator
type StationListIterator struct {
    stations []RadioStation
    index    int
}

func (sli *StationListIterator) HasNext() bool {
    return sli.index < len(sli.stations)
}

func (sli *StationListIterator) Next() interface{} {
    if sli.HasNext() {
        station := sli.stations[sli.index]
        sli.index++
        return station
    }
    return nil
}
```

## Key Features

1. **Sequential Access**: Provides sequential access to collection elements
2. **Encapsulation**: Hides the internal structure of the collection
3. **Multiple Iterators**: Supports multiple simultaneous iterations

## Benefits

- **Uniform Interface**: Provides uniform way to traverse different collections
- **Encapsulation**: Internal structure of collection is hidden
- **Multiple Iterations**: Multiple iterators can traverse the same collection
- **Flexibility**: Different traversal algorithms can be implemented

## When to Use

- When you need to access a collection's elements without exposing its internal structure
- When you want to support multiple traversals of the same collection
- When you want to provide a uniform interface for traversing different collections
- When you need to support different traversal algorithms

## Real-world Examples

- **Database Cursors**: Iterating through query results
- **File Systems**: Traversing directory structures
- **Data Structures**: Iterating through lists, trees, graphs
- **Stream Processing**: Processing data streams element by element

## Go-Specific Notes

- Go's range keyword provides built-in iteration for slices, maps, and channels
- Custom iterators are useful for complex data structures
- Go's interfaces make iterator implementation straightforward
- Channels can be used to implement iterator-like patterns
- Go's for-range loop is the idiomatic way to iterate in most cases
