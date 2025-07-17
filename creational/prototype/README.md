# Prototype Pattern

## Intent
Specify the kinds of objects to create using a prototypical instance, and create new objects by copying this prototype.

## Problem Solved
When object creation is expensive or complex, and you want to create new objects by copying existing ones rather than creating from scratch.

## Go Implementation

```go
// Prototype interface
type Prototype interface {
    Clone() Prototype
    GetDetails() string
}

// Concrete prototype
type Sheep struct {
    Name     string
    Category string
}

func (s *Sheep) Clone() Prototype {
    return &Sheep{
        Name:     s.Name,
        Category: s.Category,
    }
}
```

## Key Features

1. **Cloning**: Creates new objects by copying existing ones
2. **Performance**: Avoids expensive initialization
3. **Flexibility**: Can create variations of existing objects

## Benefits

- **Performance**: Cloning can be faster than creating from scratch
- **Simplicity**: No need to know the exact class of object being created
- **Dynamic Configuration**: Can add/remove prototypes at runtime
- **Reduced Subclassing**: Alternative to Factory Method pattern

## When to Use

- When object creation is more expensive than copying
- When you need to create objects whose class is determined at runtime
- When you want to avoid building a hierarchy of factory classes
- When instances of a class can have one of only a few different combinations of state

## Real-world Examples

- **Game Development**: Cloning game objects with predefined configurations
- **Document Templates**: Creating documents from templates
- **Database Records**: Copying records with slight modifications
- **Configuration Objects**: Creating similar configurations for different environments

## Go-Specific Notes

- Go's struct copying works naturally for simple prototypes
- For deep copying, you need to implement custom Clone methods
- Go's interfaces make the prototype pattern very flexible
- JSON marshaling/unmarshaling can be used for deep cloning
- sync.Pool can be used for object pooling with prototype pattern
