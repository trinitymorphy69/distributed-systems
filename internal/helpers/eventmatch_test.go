package helpers

import (
	"testing"
	"pgregory.net/rapid"
	"github.com/trinitymorphy69/distributed-systems/internal/helpers/test-helpers"
)

func TestEventMatch(t *testing.T) {
	
	rapid.Check(t, func(t *rapid.T) {

		send := testhelpers.EventGenerator(t)
		receive := testhelpers.EventGenerator(t)

		outputBool := EventMatch(send, receive)

		if outputBool {

			if !testhelpers.CheckTrue(send, receive) {
				t.Fatalf("Main validation passed, but %v and %v do not match.", send, receive)
			}

		}

	})
}

