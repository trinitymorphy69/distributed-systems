package clocks

import (
	"testing"
	"pgregory.net/rapid"
	"github.com/trinitymorphy69/distributed-execution-fundamentals/internal/helpers"
	"github.com/trinitymorphy69/distributed-execution-fundamentals/internal/helpers/test-helpers"
)

func TestVectorClock(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {

		events := rapid.SliceOfN(rapid.Custom(testhelpers.EventGenerator), 10, -1).Draw(t, "Incomplete")

		output, err := VectorClock(events)

		if err == nil {

			if orderedEvents, err := helpers.EventOrdering(output); err == nil {

				for i := 0; i < len(orderedEvents)-1; i++ {

					if orderedEvents[i].Process == orderedEvents[i+1].Process {

						/*

						If event A happened before event B, then the vector clock of event A will be 
						lesser than the vector clock of event B. So what does it mean for one vector clock
						to be lesser than the other?

						Well the vector clock of A is said to be greater than the vector clock of B if:
						1. All the positions in VC(A) are lesser than or equal to the corresponding positionsin VC(B).
						2. Atleast one position in VC(B) is strictly greater than the corresponding in VC(A)
						3. For these conditions to hold, but events must be on the same process.

						*/

						a := orderedEvents[i].VectorClock
						b := orderedEvents[i+1].VectorClock
						var count int

						for k, v := range a {

							if a[k] > b[k] {
								t.Fatalf("Vector Clock function passed, but there's an error. Position {%v:%v} of event %v on process %v cannot be greater than position {%v:%v} of event %v on process %v.", k, v, orderedEvents[i].Number, orderedEvents[i].Process, k, b[k], orderedEvents[i+1].Number, orderedEvents[i+1].Process)
							}

							if b[k] > a[k] {
								// We will be using the count variable to check that the clock satisfies the second property
								// of vector clocks.
								count++
							}

						}

						// If count is lesser than one, it means that there's no position in the vector clock of the second event
						// relationship that is strictly greater than the corresponding position in the vector clock of the first event.
						if count < 1 {
							
							t.Fatalf("Vector Clock function passed, but there's an error. Event %v has no position in its vector clock %v that is greater than the positions in the vector clock %v of event %v. ", orderedEvents[i+1].Number, orderedEvents[i+1].VectorClock, orderedEvents[i].VectorClock, orderedEvents[i].Number)

						}
					}

					
				}


			}

		}

	})
}