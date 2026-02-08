package acceptancetests

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/Chegnxufeng1994/bdd-in-action/train-timetables/tests/acceptancetests/testcontext"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "pretty",
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = []string{"../features"}

	os.Exit(m.Run())
}

func TestFeatures(t *testing.T) {
	opts.TestingT = t

	status := godog.TestSuite{
		Name:                "Train Timetable Features",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

func TestFeaturesWithCucumberReport(t *testing.T) {
	// Create output file for cucumber JSON report
	outputFile, err := os.Create("cucumber-report.json")
	if err != nil {
		t.Fatalf("failed to create cucumber report file: %v", err)
	}
	defer outputFile.Close()

	// Configure options for cucumber JSON format
	cucumberOpts := godog.Options{
		Format:   "cucumber",
		Output:   outputFile,
		Paths:    []string{"../features"},
		TestingT: t,
	}

	status := godog.TestSuite{
		Name:                "Train Timetable Features",
		ScenarioInitializer: InitializeScenario,
		Options:             &cucumberOpts,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	testCtx := testcontext.NewTestContext()

	// Register timetable step definitions

	// Reset context before each scenario
	sc.Before(func(ctx context.Context, scenario *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	// Cleanup after each scenario
	sc.After(
		func(ctx context.Context, scenario *godog.Scenario, err error) (context.Context, error) {
			testCtx.Reset()
			return ctx, nil
		},
	)
}
