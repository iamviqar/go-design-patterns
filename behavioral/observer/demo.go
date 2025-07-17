package main

import "fmt"

// Observer interface
type Observer interface {
	Update(jobTitle string)
	GetName() string
}

// Subject interface
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(jobTitle string)
}

// JobSeeker - concrete observer
type JobSeeker struct {
	name string
}

func NewJobSeeker(name string) *JobSeeker {
	return &JobSeeker{name: name}
}

func (js *JobSeeker) Update(jobTitle string) {
	fmt.Printf("Hi %s! New job posted: %s\n", js.name, jobTitle)
}

func (js *JobSeeker) GetName() string {
	return js.name
}

// JobPostings - concrete subject
type JobPostings struct {
	observers []Observer
}

func NewJobPostings() *JobPostings {
	return &JobPostings{observers: make([]Observer, 0)}
}

func (jp *JobPostings) Attach(observer Observer) {
	jp.observers = append(jp.observers, observer)
}

func (jp *JobPostings) Detach(observer Observer) {
	for i, obs := range jp.observers {
		if obs.GetName() == observer.GetName() {
			jp.observers = append(jp.observers[:i], jp.observers[i+1:]...)
			break
		}
	}
}

func (jp *JobPostings) Notify(jobTitle string) {
	for _, observer := range jp.observers {
		observer.Update(jobTitle)
	}
}

func (jp *JobPostings) AddJob(jobTitle string) {
	jp.Notify(jobTitle)
}

func main() {
	fmt.Println("=== Observer Pattern Demo ===")
	
	// Create job postings (subject)
	jobPostings := NewJobPostings()
	
	// Create job seekers (observers)
	johnDoe := NewJobSeeker("John Doe")
	janeDoe := NewJobSeeker("Jane Doe")
	
	// Subscribe job seekers
	jobPostings.Attach(johnDoe)
	jobPostings.Attach(janeDoe)
	
	// Post new job
	fmt.Println("Posting new job:")
	jobPostings.AddJob("Software Engineer")
	
	fmt.Println("\nPosting another job:")
	jobPostings.AddJob("Data Scientist")
	
	// Unsubscribe one observer
	fmt.Printf("\n%s unsubscribes...\n", johnDoe.GetName())
	jobPostings.Detach(johnDoe)
	
	fmt.Println("Posting job after John unsubscribed:")
	jobPostings.AddJob("Product Manager")
	
	fmt.Println("\nObserver pattern enables loose coupling between subjects and observers!")
}
