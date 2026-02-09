package timetables

import "time"

// InMemoryTimeTable is an in-memory implementation of TimeTable
type InMemoryTimeTable struct {
	trains []*ScheduledTrain
}

// NewTimeTable creates a new in-memory timetable
func NewTimeTable() TimeTable {
	return &InMemoryTimeTable{
		trains: []*ScheduledTrain{},
	}
}

// AddTrain adds a scheduled train to the timetable
func (tt *InMemoryTimeTable) AddTrain(train *ScheduledTrain) {
	tt.trains = append(tt.trains, train)
}

// FindTrains returns all trains that go from departure to destination
func (tt *InMemoryTimeTable) FindTrains(from, to string) []*ScheduledTrain {
	var matchingTrains []*ScheduledTrain

	for _, train := range tt.trains {
		if train.Departure() == from && train.Destination() == to {
			matchingTrains = append(matchingTrains, train)
		}
	}

	return matchingTrains
}

// FindLinesThrough returns all unique train lines that go from one station to another
func (tt *InMemoryTimeTable) FindLinesThrough(from, to string) []string {
	linesMap := make(map[string]bool)
	var lines []string

	for _, train := range tt.trains {
		// Check if train goes from 'from' to 'to'
		if train.Departure() == from && train.Destination() == to {
			if !linesMap[train.Line()] {
				linesMap[train.Line()] = true
				lines = append(lines, train.Line())
			}
		}
	}

	return lines
}

// GetDepartures returns departure times for a specific line from a station after a given time
func (tt *InMemoryTimeTable) GetDepartures(
	lineName string,
	from string,
	after time.Time,
) []time.Time {
	var times []time.Time

	// Find all trains on the specified line that depart from the station
	for _, train := range tt.trains {
		if train.Line() == lineName && train.Departure() == from {
			// Collect departure times after the given time
			for _, departureTime := range train.DepartureTimes() {
				if departureTime.After(after) || departureTime.Equal(after) {
					times = append(times, departureTime)
				}
			}
		}
	}

	return times
}

// GetAllTrains returns all trains in the timetable
func (tt *InMemoryTimeTable) GetAllTrains() []*ScheduledTrain {
	return tt.trains
}
