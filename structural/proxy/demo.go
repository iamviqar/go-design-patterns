package main

import "fmt"

// Door interface
type Door interface {
	Open() string
	Close() string
}

// LabDoor - real subject
type LabDoor struct{}

func (l LabDoor) Open() string {
	return "Opening lab door"
}

func (l LabDoor) Close() string {
	return "Closing lab door"
}

// Security - proxy that controls access
type Security struct {
	door     LabDoor
	password string
}

func NewSecurity(door LabDoor, password string) Security {
	return Security{door: door, password: password}
}

func (s Security) Open() string {
	return "Security: " + s.door.Open()
}

func (s Security) Close() string {
	return "Security: " + s.door.Close()
}

func (s Security) Authenticate(password string) bool {
	return s.password == password
}

// SecuredDoor - protected proxy
type SecuredDoor struct {
	security Security
}

func NewSecuredDoor(security Security) SecuredDoor {
	return SecuredDoor{security: security}
}

func (sd SecuredDoor) Open(password string) string {
	if sd.security.Authenticate(password) {
		return sd.security.Open()
	}
	return "Access denied! Incorrect password."
}

func (sd SecuredDoor) Close() string {
	return sd.security.Close()
}

func main() {
	fmt.Println("=== Proxy Pattern Demo ===")
	
	// Create real door
	labDoor := LabDoor{}
	
	// Create security proxy
	security := NewSecurity(labDoor, "secret")
	
	// Create secured door (proxy)
	securedDoor := NewSecuredDoor(security)
	
	// Try to open with wrong password
	fmt.Println(securedDoor.Open("invalid"))
	
	// Try to open with correct password
	fmt.Println(securedDoor.Open("secret"))
	
	// Close the door
	fmt.Println(securedDoor.Close())
	
	fmt.Println("\nProxy controls access to the real object!")
}
