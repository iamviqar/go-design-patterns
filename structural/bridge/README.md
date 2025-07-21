# Bridge Pattern

## Real World Example
🚡 Consider you have a website with different pages and you are supposed to allow the user to change the theme. What would you do? Create multiple copies of each of the pages for each of the themes or would you just create separate theme and load them based on the user's preferences? Bridge pattern allows you to do the second i.e.

## In Plain Words
Bridge pattern is about preferring composition over inheritance. Implementation details are pushed from a hierarchy to another object with a separate hierarchy.

## Wikipedia Definition
The bridge pattern is a design pattern used in software engineering that is meant to "decouple an abstraction from its implementation so that the two can vary independently".

## Intent
Decouple an abstraction from its implementation so that the two can vary independently.

## Problem Solved
When you want to avoid permanent binding between an abstraction and its implementation, or when you want to share an implementation among multiple objects.

## Go Implementation

```go
// Implementation interface
type Theme interface {
    GetColor() string
}

// Concrete implementations
type DarkTheme struct{}
func (d DarkTheme) GetColor() string {
    return "Dark Black"
}

// Abstraction
type About struct {
    theme Theme
}

func NewAbout(theme Theme) About {
    return About{theme: theme}
}

func (a About) GetContent() string {
    return fmt.Sprintf("About page in %s", a.theme.GetColor())
}
```

## Key Features

1. **Separation**: Separates abstraction from implementation
2. **Independence**: Both can vary independently
3. **Runtime Binding**: Implementation can be selected at runtime

## Benefits

- **Platform Independence**: Abstractions can work across different platforms
- **Implementation Flexibility**: Can change implementation without affecting clients
- **Extensibility**: Both abstractions and implementations can be extended independently
- **Hiding Details**: Implementation details are hidden from clients

## When to Use

- When you want to avoid permanent binding between abstraction and implementation
- When both abstractions and implementations should be extensible through subclassing
- When you want to share an implementation among multiple objects
- When you need to switch implementations at runtime

## Real-world Examples

- **GUI Frameworks**: Different rendering engines for different platforms
- **Database Drivers**: Different databases with same query interface
- **Graphics APIs**: Different graphics backends (OpenGL, DirectX, Vulkan)
- **Messaging Systems**: Different message transports (HTTP, TCP, WebSocket)

## Go-Specific Notes

- Go's interfaces naturally support the bridge pattern
- Composition is used instead of inheritance
- Go's implicit interface satisfaction provides flexibility
- Dependency injection works well with bridge pattern
- Go's struct embedding can be used for shared behavior
