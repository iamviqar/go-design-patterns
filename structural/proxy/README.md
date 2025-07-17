# Proxy Pattern

## Intent
Provide a placeholder or surrogate for another object to control access to it.

## Problem Solved
When you need to control access to an object, add extra functionality, or defer expensive operations until they're needed.

## Go Implementation

```go
// Subject interface
type Door interface {
    Open() string
    Close() string
}

// Real subject
type LabDoor struct{}

func (l LabDoor) Open() string {
    return "Opening lab door"
}

// Proxy
type SecuredDoor struct {
    door     LabDoor
    password string
}

func (sd SecuredDoor) Open(password string) string {
    if sd.authenticate(password) {
        return sd.door.Open()
    }
    return "Access denied!"
}

func (sd SecuredDoor) authenticate(password string) bool {
    return sd.password == password
}
```

## Key Features

1. **Access Control**: Controls access to the real object
2. **Lazy Loading**: Can defer expensive operations
3. **Additional Behavior**: Can add functionality without changing the real object

## Benefits

- **Security**: Can add authentication and authorization
- **Performance**: Can implement lazy loading and caching
- **Monitoring**: Can log and monitor access
- **Remote Access**: Can provide local representative for remote objects

## Types of Proxies

- **Virtual Proxy**: Controls access to expensive objects
- **Protection Proxy**: Controls access based on permissions
- **Remote Proxy**: Provides local representative for remote objects
- **Smart Proxy**: Adds additional behavior when accessing objects

## When to Use

- When you need to control access to an object
- When you want to add functionality without changing the object
- When you need lazy initialization of expensive objects
- When you need to add security, logging, or caching

## Real-world Examples

- **Web Proxies**: Caching and filtering web requests
- **Database Proxies**: Connection pooling and query optimization
- **Security Proxies**: Authentication and authorization
- **Virtual Proxies**: Lazy loading of images or large objects

## Go-Specific Notes

- Go's interfaces make proxy implementation straightforward
- Composition is used instead of inheritance
- Go's HTTP reverse proxy is a real-world example
- Context can be used for request-scoped proxy behavior
- Go's concurrency features work well with proxy patterns
