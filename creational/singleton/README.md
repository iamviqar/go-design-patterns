# Singleton Pattern

## Real World Example
💍 There can only be one president of a country at a time. The same president has to be brought to action, whenever duty calls. President here is singleton.

## In Plain Words
Ensures that only one object of a particular class is ever created.

## Wikipedia Definition
In software engineering, the singleton pattern is a software design pattern that restricts the instantiation of a class to one object. This is useful when exactly one object is needed to coordinate actions across the system.

## Important Note
Singleton pattern is actually considered an anti-pattern and overuse of it should be avoided. It is not necessarily bad and could have some valid use-cases but should be used with caution because it introduces a global state in your application and change to it in one place could affect in the other areas and it could become pretty difficult to debug. The other bad thing about them is it makes your code tightly coupled plus mocking the singleton could be difficult.

## Intent
Ensure a class has only one instance and provide a global point of access to that instance.

## Problem Solved
When you need exactly one instance of a class and want to control access to shared resources like databases, file systems, or configuration objects.

## Go Implementation

### Using sync.Once (Recommended)
```go
type President struct {
    name string
}

var (
    presidentInstance *President
    once              sync.Once
)

func GetPresident() *President {
    once.Do(func() {
        presidentInstance = &President{name: "Default President"}
    })
    return presidentInstance
}
```

### Using sync.Mutex
```go
var (
    dbInstance *Database
    dbMutex    sync.Mutex
)

func GetDatabase() *Database {
    if dbInstance == nil {
        dbMutex.Lock()
        defer dbMutex.Unlock()
        if dbInstance == nil {
            dbInstance = &Database{name: "Default Database"}
        }
    }
    return dbInstance
}
```

## Key Features

1. **Single Instance**: Only one instance can exist
2. **Global Access**: Provides a global point of access
3. **Lazy Initialization**: Instance created when first needed

## Benefits

- **Controlled Access**: Controls access to sole instance
- **Reduced Memory**: Only one instance in memory
- **Global State**: Provides global access to shared state
- **Lazy Loading**: Created only when needed

## When to Use

- When exactly one instance of a class is needed
- When you need strict control over global variables
- For objects that coordinate actions across the system
- When the sole instance should be extensible by subclassing

## Real-world Examples

- **Database Connection Pool**: Managing database connections
- **Logger**: Application-wide logging service
- **Configuration Manager**: Application settings and configuration
- **Cache Manager**: Application-wide caching service
- **Thread Pool**: Managing worker threads

## Go-Specific Notes

- `sync.Once` is the idiomatic way to implement singleton in Go
- Package-level variables can serve as singletons in Go
- Go's init() functions can be used for singleton initialization
- Avoid global state when possible - dependency injection is often better
- Consider using channels for coordination instead of singletons

## Cautions

- Can make unit testing difficult
- Can introduce hidden dependencies
- May violate Single Responsibility Principle
- Can create bottlenecks in concurrent applications
