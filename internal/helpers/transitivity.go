package helpers

import (
	"fmt"
    "reflect"
	"github.com/trinitymorphy69/distributed-systems/internal/types"
)

// The third definition of the happen before relationship is that
// if event A happened before event C and event C happened before
// event B, then event A happened before event B. To get the events
// that satisfy this relationship, we read all event pairs from
// relationshipOrder into a buffered channel. We then compare the last
// value of the read slice from the channel with the starting values of
// all other elements in relationship order. Where both values correspond,
// We create a new slice consisting of first value from the read channel slice
// and the last value of the element from relationshipOrder.

func CheckTransitivity(hbrelationship [][]types.Event) ([][]types.Event, error) {

    if len(hbrelationship) < 2 {
        return [][]types.Event{}, fmt.Errorf("%v cannot be empty", hbrelationship)
    }

	eventCh := make(chan []types.Event, len(hbrelationship))
    workingOn := hbrelationship
    transitiveRelationship := [][]types.Event{}

    for _, event := range workingOn {

        // These checks that the minimal events within the slice are also valid ones.
        if err := CheckEvent(event); err != nil {
            return [][]types.Event{}, err
        }


        eventCh <- event
    }

	close(eventCh)

    for eventCh != nil {

        value, ok := <-eventCh

        if !ok {
            eventCh = nil
            continue
        }

		// Here, we have two event pairs; value and pairedEvent.
		// The transitivity rule of the happens-before relationship 
		// states that A --> B if A --> C and C --> B. Here we take
		// two event pairs [A, B] and [C, D]. Each pair represents
		// an already established relationship. We then check if B == C.
		// If true, we then establish that A --> D and add it to the 
		// eventPairs which is the slice containing all happens-before
		// relationship.

        for _, pairedEvent := range workingOn {

            if reflect.DeepEqual(value[1], pairedEvent[0]) {
                transitiveRelationship = append(transitiveRelationship, []types.Event{value[0], pairedEvent[1]})
            } 

        }


    }

    if len(transitiveRelationship) == 0 {
        return transitiveRelationship, fmt.Errorf("No transitive relationship")
    }

	return transitiveRelationship, nil

}