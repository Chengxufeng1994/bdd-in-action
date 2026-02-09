package timetables

import "time"

// TimeTable defines the contract for managing train schedules
type TimeTable interface {
	// AddTrain adds a scheduled train to the timetable
	AddTrain(train *ScheduledTrain)

	// FindTrains returns all trains that go from departure to destination
	FindTrains(from, to string) []*ScheduledTrain

	// FindLinesThrough returns all unique train lines that go from one station to another
	FindLinesThrough(from, to string) []string

	// GetDepartures returns departure times for a specific line from a station after a given time
	GetDepartures(lineName string, from string, after time.Time) []time.Time

	// GetAllTrains returns all trains in the timetable
	GetAllTrains() []*ScheduledTrain
}
