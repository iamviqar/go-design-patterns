package main

import "fmt"

// Implementation interface
type Theme interface {
	GetColor() string
}

// Concrete implementations
type DarkTheme struct{}

func (d DarkTheme) GetColor() string {
	return "Dark Black"
}

type LightTheme struct{}

func (l LightTheme) GetColor() string {
	return "Off white"
}

type AquaTheme struct{}

func (a AquaTheme) GetColor() string {
	return "Light blue"
}

// Abstraction
type WebPage struct {
	theme Theme
}

func NewWebPage(theme Theme) WebPage {
	return WebPage{theme: theme}
}

func (w WebPage) GetContent() string {
	return fmt.Sprintf("Webpage content in %s", w.theme.GetColor())
}

// Refined abstractions
type About struct {
	WebPage
}

func NewAbout(theme Theme) About {
	return About{WebPage: NewWebPage(theme)}
}

func (a About) GetContent() string {
	return fmt.Sprintf("About page in %s", a.theme.GetColor())
}

type Careers struct {
	WebPage
}

func NewCareers(theme Theme) Careers {
	return Careers{WebPage: NewWebPage(theme)}
}

func (c Careers) GetContent() string {
	return fmt.Sprintf("Careers page in %s", c.theme.GetColor())
}

func main() {
	fmt.Println("=== Bridge Pattern Demo ===")
	
	darkTheme := DarkTheme{}
	lightTheme := LightTheme{}
	aquaTheme := AquaTheme{}
	
	// Create pages with different themes
	about := NewAbout(darkTheme)
	fmt.Println(about.GetContent())
	
	careers := NewCareers(lightTheme)
	fmt.Println(careers.GetContent())
	
	// Change theme for existing page
	aboutWithAqua := NewAbout(aquaTheme)
	fmt.Println(aboutWithAqua.GetContent())
	
	fmt.Println("\nBridge pattern separates abstraction from implementation!")
}
