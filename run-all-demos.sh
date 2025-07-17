#!/bin/bash

# Script to run all Go Design Pattern demos
echo "🎯 Running all Go Design Pattern demos..."
echo "======================================="

echo
echo "📦 CREATIONAL PATTERNS"
echo "======================"
echo "🏭 Simple Factory:"
go run creational/simplefactory/demo.go
echo
echo "🏭 Factory Method:"
go run creational/factorymethod/demo.go
echo
echo "🏭 Abstract Factory:"
go run creational/abstractfactory/demo.go
echo
echo "🏗️ Builder:"
go run creational/builder/demo.go
echo
echo "📋 Prototype:"
go run creational/prototype/demo.go
echo
echo "👤 Singleton:"
go run creational/singleton/demo.go

echo
echo "🏗️ STRUCTURAL PATTERNS"
echo "======================"
echo "🔌 Adapter:"
go run structural/adapter/demo.go
echo
echo "🌉 Bridge:"
go run structural/bridge/demo.go
echo
echo "🌳 Composite:"
go run structural/composite/demo.go
echo
echo "🎨 Decorator:"
go run structural/decorator/demo.go
echo
echo "🎭 Facade:"
go run structural/facade/demo.go
echo
echo "🪶 Flyweight:"
go run structural/flyweight/demo.go
echo
echo "🛡️ Proxy:"
go run structural/proxy/demo.go

echo
echo "🎭 BEHAVIORAL PATTERNS"
echo "======================"
echo "⛓️ Chain of Responsibility:"
go run behavioral/chainofresponsibility/demo.go
echo
echo "📢 Command:"
go run behavioral/command/demo.go
echo
echo "🔄 Iterator:"
go run behavioral/iterator/demo.go
echo
echo "👁️ Observer:"
go run behavioral/observer/demo.go
echo
echo "🎯 Strategy:"
go run behavioral/strategy/demo.go
echo
echo "🔄 State:"
go run behavioral/state/demo.go

echo
echo "✅ Completed demos!"
echo "🔗 Check individual pattern READMEs for detailed explanations."
