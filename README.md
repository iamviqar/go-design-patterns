# Go Design Patterns Implementation

This project contains working Go implementations of all design patterns from the "Design Patterns for Humans" repository.

## Quick Navigation
- [Creational Patterns](#creational-patterns) (6 patterns) ✅
- [Structural Patterns](#structural-patterns) (7 patterns) ✅
- [Behavioral Patterns](#behavioral-patterns) (10 patterns) ✅
- [How to Run](#how-to-run)
- [Requirements](#requirements)

> **Status**: All 23 GoF Design Patterns implemented with working Go code, demos, and documentation!

## Project Structure

### Creational Patterns
- **Simple Factory** - [creational/simplefactory/](creational/simplefactory/)
- **Factory Method** - [creational/factorymethod/](creational/factorymethod/)
- **Abstract Factory** - [creational/abstractfactory/](creational/abstractfactory/)
- **Builder** - [creational/builder/](creational/builder/)
- **Prototype** - [creational/prototype/](creational/prototype/)
- **Singleton** - [creational/singleton/](creational/singleton/)

### Structural Patterns
- **Adapter** - [structural/adapter/](structural/adapter/)
- **Bridge** - [structural/bridge/](structural/bridge/)
- **Composite** - [structural/composite/](structural/composite/)
- **Decorator** - [structural/decorator/](structural/decorator/)
- **Facade** - [structural/facade/](structural/facade/)
- **Flyweight** - [structural/flyweight/](structural/flyweight/)
- **Proxy** - [structural/proxy/](structural/proxy/)

### Behavioral Patterns
- **Chain of Responsibility** - [behavioral/chainofresponsibility/](behavioral/chainofresponsibility/)
- **Command** - [behavioral/command/](behavioral/command/)
- **Iterator** - [behavioral/iterator/](behavioral/iterator/)
- **Mediator** - [behavioral/mediator/](behavioral/mediator/)
- **Memento** - [behavioral/memento/](behavioral/memento/)
- **Observer** - [behavioral/observer/](behavioral/observer/)
- **Visitor** - [behavioral/visitor/](behavioral/visitor/)
- **Strategy** - [behavioral/strategy/](behavioral/strategy/)
- **State** - [behavioral/state/](behavioral/state/)
- **Template Method** - [behavioral/templatemethod/](behavioral/templatemethod/)

## How to Run

Each pattern has its own demo function demonstrating the pattern usage.

### Using Go (Recommended)
```bash
# Run a specific pattern demo (example: Simple Factory)
go run creational/simplefactory/demo.go

# Or compile and run
go build -o demo creational/simplefactory/demo.go
./demo
```

### Run All Demos at Once

For convenience, you can run all pattern demos with a single command:

```bash
# Make the script executable (first time only)
chmod +x run-all-demos.sh

# Run all demos
./run-all-demos.sh
```

This will execute all 23 design pattern demos in sequence with clear categorization and emojis for easy reading.

## Complete Example

Here's how to run all pattern demos to see them in action:

```bash
# Creational Patterns
go run creational/simplefactory/demo.go
go run creational/factorymethod/demo.go
go run creational/abstractfactory/demo.go
go run creational/builder/demo.go
go run creational/prototype/demo.go
go run creational/singleton/demo.go

# Structural Patterns
go run structural/adapter/demo.go
go run structural/bridge/demo.go
go run structural/composite/demo.go
go run structural/decorator/demo.go
go run structural/facade/demo.go
go run structural/flyweight/demo.go
go run structural/proxy/demo.go

# Behavioral Patterns
go run behavioral/chainofresponsibility/demo.go
go run behavioral/command/demo.go
go run behavioral/iterator/demo.go
go run behavioral/mediator/demo.go
go run behavioral/memento/demo.go
go run behavioral/observer/demo.go
go run behavioral/strategy/demo.go
go run behavioral/state/demo.go
go run behavioral/templatemethod/demo.go
go run behavioral/visitor/demo.go
```

Each demo will output examples showing how the pattern works and its benefits.

## Pattern Documentation

Each design pattern implementation includes:
- Complete working Go code with interfaces and structs
- Detailed README.md explaining the pattern
- Demo file with example usage
- Real-world use cases and benefits
- Go-specific implementation details

Click on the pattern folder links above to explore each specific pattern implementation.

## Requirements

- Go 1.21 or higher
- No external dependencies required

## Project Structure

This project follows Go module layout:
- `creational/` - All creational pattern implementations
- `structural/` - All structural pattern implementations
- `behavioral/` - All behavioral pattern implementations
- `go.mod` - Go module configuration

Each pattern package contains:
- Pattern implementation with interfaces and structs
- Demo file with main function
- README.md with detailed explanation
