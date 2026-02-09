package stepdefinitions

import (
	"fmt"
	"strings"
	"time"

	"github.com/Chegnxufeng1994/bdd-in-action/train-timetables/domain/itineraries"
	"github.com/Chegnxufeng1994/bdd-in-action/train-timetables/domain/timetables"
	"github.com/Chegnxufeng1994/bdd-in-action/train-timetables/tests/acceptancetests/testcontext"
	"github.com/cucumber/godog"
)

// DepartingTrainsSteps implements step definitions for train departure scenarios
type DepartingTrainsSteps struct {
	ctx *testcontext.TestContext
}

// NewDepartingTrainsSteps creates a new step definitions handler
func NewDepartingTrainsSteps(ctx *testcontext.TestContext) *DepartingTrainsSteps {
	return &DepartingTrainsSteps{ctx: ctx}
}

// RegisterSteps registers all step definitions for this feature
func (dts *DepartingTrainsSteps) RegisterSteps(sc *godog.ScenarioContext) {
	sc.Step(`^the T(\d+) train to (\w+) leaves (\w+) at (.+)$`, dts.theTrainToLeavesAt)
	sc.Step(
		`^(\w+) wants? to travel from (\w+) to (\w+) at (\d+):(\d+)$`,
		dts.travelerWantsToTravel,
	)
	sc.Step(`^he should be told about the trains at: (.+)$`, dts.shouldBeToldAboutTrains)
}

// parseTimeOfDay parses a time string in "HH:MM" format to time.Time
// Uses a fixed date (2024-01-01) since we only care about time of day
func parseTimeOfDay(timeStr string) (time.Time, error) {
	// Parse time in "HH:MM" format with a fixed date
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %s (expected HH:MM)", timeStr)
	}

	// Return time with fixed date (we only care about HH:MM)
	return time.Date(2024, 1, 1, t.Hour(), t.Minute(), 0, 0, time.UTC), nil
}

// formatTimeOfDay formats a time.Time to "HH:MM" string
func formatTimeOfDay(t time.Time) string {
	return t.Format("15:04")
}

// theTrainToLeavesAt sets up a train schedule
// Example: "the T1 train to Chatswood leaves Hornsby at 08:02, 08:15, 08:21"
func (dts *DepartingTrainsSteps) theTrainToLeavesAt(
	lineNumber int,
	destination, departure, timesStr string,
) error {
	// Parse the times from comma-separated string (e.g., "08:02, 08:15, 08:21")
	timeStrings := strings.Split(timesStr, ",")
	var departureTimes []time.Time

	for _, timeStr := range timeStrings {
		timeStr = strings.TrimSpace(timeStr)
		t, err := parseTimeOfDay(timeStr)
		if err != nil {
			return fmt.Errorf("failed to parse time %s: %w", timeStr, err)
		}
		departureTimes = append(departureTimes, t)
	}

	// Create a line name (e.g., "T1")
	lineName := fmt.Sprintf("T%d", lineNumber)

	// Add this train schedule to our test context
	train := timetables.NewScheduledTrain(lineName, destination, departure, departureTimes)
	dts.ctx.AddScheduledTrain(train)

	return nil
}

// travelerWantsToTravel performs a journey search
// Example: "Travis wants to travel from Hornsby to Chatswood at 08:00"
func (dts *DepartingTrainsSteps) travelerWantsToTravel(
	traveler, from, to string,
	hour, minute int,
) error {
	// Create departure time with fixed date (we only care about time of day)
	departureTime := time.Date(2024, 1, 1, hour, minute, 0, 0, time.UTC)

	// Create an itinerary service if not already created
	if dts.ctx.ItineraryService == nil {
		dts.ctx.ItineraryService = itineraries.NewItineraryService(dts.ctx.TimeTable)
	}

	// Search for next departing trains
	times, err := dts.ctx.ItineraryService.FindNextDepartures(from, to, departureTime, 2)
	if err != nil {
		dts.ctx.LastError = err
		return nil // Don't fail the step, store error for Then step to verify
	}

	// Store the results in context
	dts.ctx.SearchResults = times
	return nil
}

// shouldBeToldAboutTrains verifies the search results
// Example: "he should be told about the trains at: 08:02, 08:15"
func (dts *DepartingTrainsSteps) shouldBeToldAboutTrains(expectedTimesStr string) error {
	// Check if there was an error during search
	if dts.ctx.LastError != nil {
		return fmt.Errorf("search failed: %w", dts.ctx.LastError)
	}

	// Parse expected times
	timeStrings := strings.Split(expectedTimesStr, ",")
	var expectedTimes []time.Time

	for _, timeStr := range timeStrings {
		timeStr = strings.TrimSpace(timeStr)
		t, err := parseTimeOfDay(timeStr)
		if err != nil {
			return fmt.Errorf("failed to parse expected time %s: %w", timeStr, err)
		}
		expectedTimes = append(expectedTimes, t)
	}

	// Verify we got the right number of results
	if len(dts.ctx.SearchResults) != len(expectedTimes) {
		return fmt.Errorf(
			"expected %d results, got %d",
			len(expectedTimes),
			len(dts.ctx.SearchResults),
		)
	}

	// Verify each result matches expected time
	for i, resultTime := range dts.ctx.SearchResults {
		if !resultTime.Equal(expectedTimes[i]) {
			return fmt.Errorf("result %d: expected departure at %s, got %s",
				i, formatTimeOfDay(expectedTimes[i]), formatTimeOfDay(resultTime))
		}
	}

	return nil
}
