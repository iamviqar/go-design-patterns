# Proxy Pattern

## Real World Example
🎱 Have you ever used an access card to go through a door? There are multiple options to open that door i.e. it can be opened either using access card or by pressing a button that bypasses the security. The door's main functionality is to open but there is a proxy added on top of it to add some functionality. Let me better explain it using the code example below.

## In Plain Words
Using the proxy pattern, a class represents the functionality of another class.

## Wikipedia Definition
A proxy, in its most general form, is a class functioning as an interface to something else. A proxy is a wrapper or agent object that is being called by the client to access the real serving object behind the scenes. Use of the proxy can simply be forwarding to the real object, or can provide additional logic. In the proxy extra functionality can be provided, for example caching when operations on the real object are resource intensive, or checking preconditions before operations on the real object are invoked.

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
