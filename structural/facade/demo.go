package main

import "fmt"

// Complex subsystem components
type Computer struct{}

func (c Computer) GetElectricShock() string {
	return "Ouch!"
}

func (c Computer) MakeSound() string {
	return "Beep beep!"
}

func (c Computer) ShowLoadingScreen() string {
	return "Loading.."
}

func (c Computer) Bam() string {
	return "Ready to be used!"
}

func (c Computer) CloseEverything() string {
	return "Bup bup bup buzzzz!"
}

func (c Computer) Sooth() string {
	return "Haaah!"
}

func (c Computer) PullCurrent() string {
	return "Zzzzz"
}

// Facade provides simple interface
type ComputerFacade struct {
	computer Computer
}

func NewComputerFacade() ComputerFacade {
	return ComputerFacade{computer: Computer{}}
}

func (cf ComputerFacade) TurnOn() string {
	result := ""
	result += cf.computer.GetElectricShock() + "\n"
	result += cf.computer.MakeSound() + "\n"
	result += cf.computer.ShowLoadingScreen() + "\n"
	result += cf.computer.Bam()
	return result
}

func (cf ComputerFacade) TurnOff() string {
	result := ""
	result += cf.computer.CloseEverything() + "\n"
	result += cf.computer.Sooth() + "\n"
	result += cf.computer.PullCurrent()
	return result
}

func main() {
	fmt.Println("=== Facade Pattern Demo ===")
	
	computer := NewComputerFacade()
	
	fmt.Println("Turning on computer using facade:")
	fmt.Println(computer.TurnOn())
	
	fmt.Println("\nTurning off computer using facade:")
	fmt.Println(computer.TurnOff())
	
	fmt.Println("\nFacade provides a simplified interface to complex subsystems!")
}
