package helpers

import (
	"strconv"
	"testing"
	"pgregory.net/rapid"
	"github.com/trinitymorphy69/distributed-execution-fundamentals/internal/helpers/test-helpers"
)

func TestProcessNamingConvention(t *testing.T) {
	
	rapid.Check(t, func(t *rapid.T) {
		
		event := testhelpers.EventGenerator(t)

		err := ProcessNamingConvention(event)

		if err == nil {

			if !correctNamingConvention(event.Process) {
				t.Fatalf("Main validation passed, but Process Field is invalid: %v", event)
			}

			if !correctNamingConvention(event.From) {
				t.Fatalf("Main validation passed, but From Field is invalid: %v", event)
			}

			if !correctNamingConvention(event.To) {
				t.Fatalf("Main validation passed, but To Field is invalid: %v", event)
			}

		}

	})
}



func correctNamingConvention(eventField string) bool {

	if eventField != "" {

		// 1. The first value in the naming field must be P
		correctStartingValue := eventField[0] == 'P'

		// 2. Every value after the first in the string must
		// result to a natural number.
		_, a := strconv.Atoi(string(eventField[1:]))
		resultingEndingValues := a == nil

		return correctStartingValue && resultingEndingValues

	}

	return true

} 
