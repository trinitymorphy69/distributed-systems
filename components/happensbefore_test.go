package components

import (
	"testing"
	"pgregory.net/rapid"
	"github.com/trinitymorphy69/distributed-systems/internal/helpers/test-helpers"
)

func TestHappensBefore(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {

		events := rapid.SliceOfN(rapid.Custom(testhelpers.EventGenerator), 1, -1).Draw(t, "Input")

		output, err := HappensBefore(events)

		if err == nil {

			/*
			
			To confirm that indeed the provided relationships are valid happen-before relationships,
			we compare their vector clocks. If A happened before B, then VC(A) < VC(B). We use vector
			clocks to compare instead of lamports because vector clocks are consistent with causality
			and can characterise it.

			*/

			for _, relationship := range output {

				/*

				If event A happened before event B, then the vector clock of event A will be 
				lesser than the vector clock of event B. So what does it mean for one vector clock
				to be lesser than the other?

				Well the vector clock of A is said to be greater than the vector clock of B if:
				1. All the positions in VC(A) are lesser than or equal to the corresponding positionsin VC(B).
				2. Atleast one position in VC(B) is strictly greater than the corresponding in VC(A)
				3. For these conditions to hold, but events must be on the same process.

				*/

				a := relationship[0].VectorClock
				b := relationship[1].VectorClock
				var count int

				for k, v := range a {

					if a[k] > b[k] {
						t.Fatalf("Happens before relationship established between %v and %v is wrong. Position {%v:%v} of event %v on process %v cannot be greater than position {%v:%v} of event %v on process %v if there exists a valid happens-before relationship between them.", relationship[0], relationship[1], k, v, relationship[0].Number, relationship[0].Process, k, b[k], relationship[1].Number, relationship[1].Process)
					}

					if b[k] > a[k] {
						// We will be using the count variable to check that the clock satisfies the second property
						// of vector clocks.
						count++
					}

				}

				// If count is lesser than one, it means that there's no position in the second event of the happens-before
				// relationship that is strictly greater than the corresponding position in the first event of the relationship.
				if count < 1 {
					t.Fatalf("Happens before relationship established between %v and %v is wrong. There's no position in %v for Event %v that is greater than the corresponding position in %v for %v.", relationship[0], relationship[1], b, relationship[1].Number, a,  relationship[0].Number)
				}

			}
		}

	})
}