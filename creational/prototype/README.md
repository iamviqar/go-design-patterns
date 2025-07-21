# Prototype Pattern

## Real World Example
🐑 Remember dolly? The sheep that was cloned! Lets not get into the details but the key point here is that it is all about cloning.

## In Plain Words
Create object based on an existing object through cloning.

## Wikipedia Definition
The prototype pattern is a creational design pattern in software development. It is used when the type of objects to create is determined by a prototypical instance, which is cloned to produce new objects.

In short, it allows you to create a copy of an existing object and modify it to your needs, instead of going through the trouble of creating an object from scratch and setting it up.

## Intent
Specify the kinds of objects to create using a prototypical instance, and create new objects by copying this prototype.

## Problem Solved
When an object is required that is similar to existing object or when the creation would be expensive as compared to cloning.

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
