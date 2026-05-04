package helpers

import (
	"testing"
	"pgregory.net/rapid"
	"github.com/trinitymorphy69/distributed-systems/types"
	"github.com/trinitymorphy69/distributed-systems/internal/helpers/test-helpers"
)

func TestTypeMatchField(t *testing.T)  {
	
	rapid.Check(t, func(t *rapid.T) {

		event := testhelpers.EventGenerator(t)

		err := TypeMatchField(event)

		if err == nil {

			if !checkMatching(event) {
				t.Fatalf("Main function passed, but event type does not match its To/From fields: %v", event)
			}

		}

	})


}


func checkMatching(event types.Event) bool {

	var validMatch bool

	switch event.Type {
	case 1:

		// 1. If it is a Send event, To must not be empty, From must be empty, and To must not be the same as process.
		validMatch = event.To != "" && event.From == "" && event.To != event.Process

	case 2:
		// 2. If it is a Receive event, From must not be empty, To must be empty and From must not be the same as process.
		validMatch = event.From != "" && event.To == "" && event.From != event.Process

	case 3:
		// 3. If it is a Internal event, From and To must be empty
		validMatch = event.To == "" && event.From == ""

	}

	return validMatch
}

