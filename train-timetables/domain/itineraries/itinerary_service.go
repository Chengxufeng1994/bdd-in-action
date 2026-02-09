package itineraries

import (
	"sort"
	"time"

	"github.com/Chegnxufeng1994/bdd-in-action/train-timetables/domain/timetables"
)

// ItineraryService provides itinerary search and management operations
type ItineraryService struct {
	timeTable timetables.TimeTable
}

// NewItineraryService creates a new itinerary service
func NewItineraryService(timeTable timetables.TimeTable) *ItineraryService {
	return &ItineraryService{
		timeTable: timeTable,
	}
}

// FindNextDepartures finds the next N departing trains from origin to destination after the given time
func (s *ItineraryService) FindNextDepartures(
	from, to string,
	after time.Time,
	limit int,
) ([]time.Time, error) {
	// Find all lines that go from 'from' to 'to'
	lines := s.timeTable.FindLinesThrough(from, to)

	// Collect departure times from each line
	var allTimes []time.Time
	for _, line := range lines {
		times := s.timeTable.GetDepartures(line, from, after)
		allTimes = append(allTimes, times...)
	}

	// Sort by departure time
	sort.Slice(allTimes, func(i, j int) bool {
		return allTimes[i].Before(allTimes[j])
	})

	// Return only the first N departures
	if len(allTimes) > limit {
		allTimes = allTimes[:limit]
	}

	return allTimes, nil
}
