package testhelpers

import (

	"pgregory.net/rapid"
	"github.com/trinitymorphy69/distributed-systems/internal/types"

)

func EventGenerator(t *rapid.T) types.Event {

	return types.Event{
		Process: rapid.SampledFrom([]string{"", "p1", "P1", "45764896", "0", "1999", "P!@#$", "P2", "P5", ""}).Draw(t, "Process"),
		Number: rapid.IntRange(0, 1000).Draw(t, "Number"),
		Type: rapid.IntRange(0, 4).Draw(t, "Type"),
		Message: rapid.String().Draw(t, "Message"),
		To: rapid.SampledFrom([]string{"", "p1", "P1", "45764896", "0", "1999", "P!@#$", "P2", "P5", ""}).Draw(t, "To"),
		From: rapid.SampledFrom([]string{"", "p1", "P1", "45764896", "0", "1999", "P!@#$", "P2", "P5", ""}).Draw(t, "From"),
		LamportClock: rapid.IntRange(0, 2).Draw(t, "LamportClock"),
		VectorClock: rapid.MapOf(

			rapid.SampledFrom([]string{"P1", "P2", "P3"}),
			rapid.IntRange(-1, 1000000),

		).Draw(t, "VectorClock"),

	}

}

func MinimalEventGenerator(t *rapid.T) types.EventMinimal {

	return types.EventMinimal{

		Process: rapid.SampledFrom([]string{"", "p1", "P1", "45764896", "0", "1999", "P!@#$", "P2", "P5", ""}).Draw(t, "Process"),
		Number: rapid.IntRange(0, 1000).Draw(t, "Number"),
	}
}