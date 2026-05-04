package clocks

import (
	"testing"
	"pgregory.net/rapid"
	"github.com/trinitymorphy69/distributed-systems/internal/helpers"
	"github.com/trinitymorphy69/distributed-systems/internal/helpers/test-helpers"
)

func TestLamportClock(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {

		events := rapid.SliceOfN(rapid.Custom(testhelpers.EventGenerator), 1, -1).Draw(t, "Inputs")

		output, err := LamportClock(events)

		if err == nil {

			// The following code orders the events produced by the generator based on the process
			// they occur on. Then it checks that two events a, b occuring on the same process must
			// satisfy the Lamport clock property where the Lamport clock of is lesser than the lamport
			// clock of b.
			
			if orderedOutput, err := helpers.EventOrdering(output); err == nil {

				for i := 0; i < len(orderedOutput)-1; i++ {

					if orderedOutput[i].Process == orderedOutput[i+1].Process {

						if orderedOutput[i].Number >= orderedOutput[i+1].Number {
							t.Fatalf("Main function passed. But Lamport clock of %v on Process %v is >= the Lamport clock of %v on Process %v even though the latter occurs after the former.", orderedOutput[i].Number, orderedOutput[i].Process, orderedOutput[i+1].Number, orderedOutput[i+1].Process)
						}
					}
				}

			}

			
		}

	})
	
}