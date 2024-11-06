package cmd

import (
	"fmt"
	"testing"

	"github.com/ShohamBit/traceectl/pkg/models"
)

var eventTests = []models.TestCase{
	//event list
	{
		TestName:        "No events list",
		OutputSlice:     []string{"event", "list"},
		ExpectedPrinter: "",
		ExpectedError:   fmt.Errorf("accepts 0 arg(s), received 1"),
	},

	//event describe
	{
		TestName:        "No events describe",
		OutputSlice:     []string{"event", "describe"},
		ExpectedPrinter: "",
		ExpectedError:   fmt.Errorf("accepts 1 arg(s), received 0"), // Update expected output
	},
	//event enable
	{
		TestName:        "No  events enable",
		OutputSlice:     []string{"event", "enable"},
		ExpectedPrinter: "",
		ExpectedError:   fmt.Errorf("accepts 1 arg(s), received 0"), // Update expected output

	},
	{
		TestName:        "Single enable event",
		OutputSlice:     []string{"event", "enable", "event"},
		ExpectedPrinter: "Enabled event: event",
		ExpectedError:   nil,
	},
	//event disable
	{
		TestName:        "No disable events",
		OutputSlice:     []string{"event", "disable"},
		ExpectedPrinter: "",
		ExpectedError:   fmt.Errorf("accepts 1 arg(s), received 0"), // Update expected output
	},
	{
		TestName:        "Single disable event",
		OutputSlice:     []string{"event", "disable", "event"},
		ExpectedPrinter: "Disabled event: event",
		ExpectedError:   nil,
	},
	//event run
	//TODO: add test when support run is added
}

func TestEvent(t *testing.T) {

	for _, testCase := range eventTests {
		t.Run(testCase.TestName, func(t *testing.T) { TestCommand(t, testCase) })
	}
}
