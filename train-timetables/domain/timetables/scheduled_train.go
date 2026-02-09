package timetables

import "time"

// ScheduledTrain represents a train with its scheduled departure times
type ScheduledTrain struct {
	line        string
	departure   string
	destination string
	times       []time.Time
}

// NewScheduledTrain creates a new scheduled train
func NewScheduledTrain(line, destination, departure string, times []time.Time) *ScheduledTrain {
	return &ScheduledTrain{
		line:        line,
		departure:   departure,
		destination: destination,
		times:       times,
	}
}

// Line returns the train line (e.g., "T1")
func (st *ScheduledTrain) Line() string {
	return st.line
}

// Destination returns the destination station
func (st *ScheduledTrain) Destination() string {
	return st.destination
}

// Departure returns the departure station
func (st *ScheduledTrain) Departure() string {
	return st.departure
}

// DepartureTimes returns all scheduled departure times
func (st *ScheduledTrain) DepartureTimes() []time.Time {
	return st.times
}

// GetNextDepartures returns the next N departures after the given time
func (st *ScheduledTrain) GetNextDepartures(after time.Time, limit int) []time.Time {
	var nextDepartures []time.Time

	for _, departureTime := range st.times {
		if departureTime.After(after) || departureTime.Equal(after) {
			nextDepartures = append(nextDepartures, departureTime)
			if len(nextDepartures) >= limit {
				break
			}
		}
	}

	return nextDepartures
}
