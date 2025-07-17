package main

import (
	"fmt"
	"sync"
)

// President represents a singleton object
type President struct {
	name string
}

// GetName returns the president's name
func (p *President) GetName() string {
	return p.name
}

// SetName sets the president's name
func (p *President) SetName(name string) {
	p.name = name
}

var (
	presidentInstance *President
	once              sync.Once
)

// GetPresident returns the singleton instance of President
func GetPresident() *President {
	once.Do(func() {
		presidentInstance = &President{name: "Default President"}
	})
	return presidentInstance
}

// Alternative implementation using sync.Mutex for demonstration
type Database struct {
	name string
}

var (
	dbInstance *Database
	dbMutex    sync.Mutex
)

// GetDatabase returns singleton database instance using mutex
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

func (d *Database) GetName() string {
	return d.name
}

func (d *Database) SetName(name string) {
	d.name = name
}

func main() {
	fmt.Println("=== Singleton Pattern Demo ===")
	
	// Get president instances
	president1 := GetPresident()
	president2 := GetPresident()
	
	fmt.Printf("President 1: %s\n", president1.GetName())
	fmt.Printf("President 2: %s\n", president2.GetName())
	fmt.Printf("Are they the same instance? %t\n", president1 == president2)
	
	// Modify through one instance
	president1.SetName("John Doe")
	fmt.Printf("After setting name via president1: %s\n", president2.GetName())
	
	fmt.Println()
	
	// Demonstrate database singleton
	db1 := GetDatabase()
	db2 := GetDatabase()
	
	fmt.Printf("Database 1: %s\n", db1.GetName())
	fmt.Printf("Database 2: %s\n", db2.GetName())
	fmt.Printf("Are they the same instance? %t\n", db1 == db2)
	
	db1.SetName("Production DB")
	fmt.Printf("After setting name via db1: %s\n", db2.GetName())
	
	fmt.Println("\nSingleton ensures only one instance exists!")
}
