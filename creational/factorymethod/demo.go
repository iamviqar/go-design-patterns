package main

import "fmt"

// Interviewer interface defines what an interviewer can do
type Interviewer interface {
	AskQuestions() string
}

// Developer interviewer
type Developer struct{}

func (d Developer) AskQuestions() string {
	return "Asking about design patterns!"
}

// CommunityExecutive interviewer  
type CommunityExecutive struct{}

func (c CommunityExecutive) AskQuestions() string {
	return "Asking about community building"
}

// HiringManager interface defines the factory method
type HiringManager interface {
	MakeInterviewer() Interviewer
	TakeInterview() string
}

// BaseHiringManager provides common functionality
type BaseHiringManager struct{}

func (b BaseHiringManager) TakeInterview(interviewer Interviewer) string {
	return interviewer.AskQuestions()
}

// DevelopmentManager concrete factory
type DevelopmentManager struct {
	BaseHiringManager
}

func (d DevelopmentManager) MakeInterviewer() Interviewer {
	return Developer{}
}

func (d DevelopmentManager) TakeInterview() string {
	interviewer := d.MakeInterviewer()
	return d.BaseHiringManager.TakeInterview(interviewer)
}

// MarketingManager concrete factory
type MarketingManager struct {
	BaseHiringManager
}

func (m MarketingManager) MakeInterviewer() Interviewer {
	return CommunityExecutive{}
}

func (m MarketingManager) TakeInterview() string {
	interviewer := m.MakeInterviewer()
	return m.BaseHiringManager.TakeInterview(interviewer)
}

func main() {
	fmt.Println("=== Factory Method Pattern Demo ===")
	
	// Development manager hires developers
	devManager := DevelopmentManager{}
	fmt.Printf("Development Manager Interview: %s\n", devManager.TakeInterview())
	
	// Marketing manager hires community executives
	marketingManager := MarketingManager{}
	fmt.Printf("Marketing Manager Interview: %s\n", marketingManager.TakeInterview())
	
	fmt.Println("\nFactory Method delegates object creation to subclasses!")
}
