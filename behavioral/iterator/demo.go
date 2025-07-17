package main

import "fmt"

// Iterator interface
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// RadioStation represents a radio station
type RadioStation struct {
	frequency float64
}

func NewRadioStation(frequency float64) RadioStation {
	return RadioStation{frequency: frequency}
}

func (r RadioStation) GetFrequency() float64 {
	return r.frequency
}

func (r RadioStation) String() string {
	return fmt.Sprintf("%.1f", r.frequency)
}

// StationList collection
type StationList struct {
	stations []RadioStation
}

func NewStationList() *StationList {
	return &StationList{stations: make([]RadioStation, 0)}
}

func (sl *StationList) AddStation(station RadioStation) {
	sl.stations = append(sl.stations, station)
}

func (sl *StationList) RemoveStation(frequency float64) {
	for i, station := range sl.stations {
		if station.GetFrequency() == frequency {
			sl.stations = append(sl.stations[:i], sl.stations[i+1:]...)
			break
		}
	}
}

func (sl *StationList) Count() int {
	return len(sl.stations)
}

func (sl *StationList) GetIterator() Iterator {
	return &StationListIterator{
		stations: sl.stations,
		index:    0,
	}
}

// StationListIterator concrete iterator
type StationListIterator struct {
	stations []RadioStation
	index    int
}

func (sli *StationListIterator) HasNext() bool {
	return sli.index < len(sli.stations)
}

func (sli *StationListIterator) Next() interface{} {
	if sli.HasNext() {
		station := sli.stations[sli.index]
		sli.index++
		return station
	}
	return nil
}

func main() {
	fmt.Println("=== Iterator Pattern Demo ===")
	
	stationList := NewStationList()
	
	// Add stations
	stationList.AddStation(NewRadioStation(89.1))
	stationList.AddStation(NewRadioStation(101.5))
	stationList.AddStation(NewRadioStation(104.3))
	stationList.AddStation(NewRadioStation(98.7))
	
	fmt.Printf("Total stations: %d\n", stationList.Count())
	
	// Iterate using iterator
	fmt.Println("\nIterating through stations:")
	iterator := stationList.GetIterator()
	for iterator.HasNext() {
		station := iterator.Next().(RadioStation)
		fmt.Printf("Radio Station: %s FM\n", station)
	}
	
	// Remove a station
	stationList.RemoveStation(98.7)
	fmt.Printf("\nAfter removing 98.7 FM, total stations: %d\n", stationList.Count())
	
	// Iterate again
	fmt.Println("\nIterating after removal:")
	iterator2 := stationList.GetIterator()
	for iterator2.HasNext() {
		station := iterator2.Next().(RadioStation)
		fmt.Printf("Radio Station: %s FM\n", station)
	}
	
	fmt.Println("\nIterator provides sequential access to elements!")
}
