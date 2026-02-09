package testcontext

import (
	"time"

	"github.com/Chegnxufeng1994/bdd-in-action/train-timetables/domain/itineraries"
	"github.com/Chegnxufeng1994/bdd-in-action/train-timetables/domain/timetables"
)

// TestContext holds the state for acceptance tests
type TestContext struct {
	TimeTable        timetables.TimeTable
	ItineraryService *itineraries.ItineraryService
	SearchResults    []time.Time
	LastError        error
}

// NewTestContext creates a new test context
func NewTestContext() *TestContext {
	return &TestContext{
		TimeTable:     timetables.NewTimeTable(),
		SearchResults: []time.Time{},
	}
}

// Reset resets the test context to initial state
func (tc *TestContext) Reset() {
	tc.TimeTable = timetables.NewTimeTable()
	tc.ItineraryService = nil
	tc.SearchResults = []time.Time{}
	tc.LastError = nil
}

// AddScheduledTrain adds a train to the timetable
func (tc *TestContext) AddScheduledTrain(train *timetables.ScheduledTrain) {
	tc.TimeTable.AddTrain(train)
}
