package tests

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestMain(m *testing.M) {
	status := godog.TestSuite{
		Name:                "integration",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"features"},
			Output: colors.Colored(os.Stdout),
		},
	}.Run()

	if runStatus := m.Run(); runStatus > status {
		status = runStatus
	}

	os.Exit(status)
}
