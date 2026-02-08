package testcontext

// TestContext holds the state for acceptance tests
type TestContext struct {
	// Add shared test state here
	// Example:
	// trainService *domain.TrainService
	// searchResults []domain.Journey
	// lastError error
}

// NewTestContext creates a new test context
func NewTestContext() *TestContext {
	return &TestContext{}
}

// Reset resets the test context to initial state
func (tc *TestContext) Reset() {
	// Reset all state to initial values
	// Example:
	// tc.trainService = nil
	// tc.searchResults = nil
	// tc.lastError = nil
}
