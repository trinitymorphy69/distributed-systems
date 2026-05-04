package components

import (

    "github.com/trinitymorphy69/distributed-execution-fundamentals/types"
	"github.com/trinitymorphy69/distributed-execution-fundamentals/internal/helpers"
    "github.com/trinitymorphy69/distributed-execution-fundamentals/components/clocks"
	
)

// This function seeks to define the happens-before relationship for all
// events within the distributed execution.

func HappensBefore(events []types.Event) ([][]types.Event, error) {

    hbrelationship := [][]types.Event{}

    // Calculate the Lamport clocks and vector clocks of each event.
    orderedEvents, err := clocks.LamportClock(events)

    if err != nil {
        return [][]types.Event{}, err
    }

    orderedEvents, err = clocks.VectorClock(orderedEvents)

     if err != nil {
        return [][]types.Event{}, err
    }

    // Matches Corresponding Receive Events to their Send Events
    matchedEvents, err := helpers.MatchSendReceive(orderedEvents)

     if err != nil {
        return [][]types.Event{}, err
    }

    // orderedEvents is a map of processes to an ascending order of
    // all events on that process. The first definition of the
    // happens before relationship is that Event A happened before
    // Event B if A occurs before B on the same process. Satisfying
    // this definition, the loop below accesses the slices of 
    // ordered events for each process, establishes the happens-
    // before relationship, and adds it to the relationship order slice. 

    if len(orderedEvents) != 0 {

        for i := 0; i < len(orderedEvents); i++ {

            // This loop ensures that in a slice of events [a, b, c], 
            // where a is events[i] and b,c is event[j], we are able to
            // get the relationship a --> b, a --> c, b --> c because 
            // j will get the elements after i as far as j is lesser than
            // the length of events.

            for j := i + 1; j < len(orderedEvents); j++ {
                
                if orderedEvents[j].Process == orderedEvents[i].Process {

                     // relationshipOrder = append(relationshipOrder, []string{fmt.Sprintf("%v(%v) ⟶ %v(%v)", events[i].Process, events[i].Number, events[j].Process, events[j].Number)})
                    hbrelationship = append(hbrelationship, []types.Event{orderedEvents[i], orderedEvents[j]})

                }
               
            }
        }


    }

    // The second definition of the happen before relationship is that
    // event A happened before event B if A is the send event and B is
    // the corresponding receive event. We use the MatchSendReceive()
    // function to match sends to their receives. The resulting slice is
    // in the order of [send, receive, send, receive ... send, receive]. 
    // We use the for loop below to add them to relationshipOrder.

    if len(matchedEvents) != 0 {

        for i := 0; i < len(matchedEvents); i += 2 {
            // relationshipOrder = append(relationshipOrder, []string{fmt.Sprintf("%v(%v) ⟶ %v(%v)", matchedEvents[i].Process, matchedEvents[i].Number, matchedEvents[i+1].Process, matchedEvents[i+1].Number)})
            hbrelationship = append(hbrelationship, []types.Event{matchedEvents[i], matchedEvents[i+1]})
        }

    }

    // The third definition of the happen before relationship is that
    // if event A happened before event C and event C happened before
    // event B, then event A happened before event B. To get the events
    // that satisfy this relationship, we read all event pairs from
    // relationshipOrder into a buffered channel. We then compare the last
    // value of the read slice from the channel with the starting values of 
    // all other elements in relationship order. Where both values correspond,
    // We create a new slice consisting of first value from the read channel slice
    // and the last value of the element from relationshipOrder.

    transitiveRelationships, err := helpers.CheckTransitivity(hbrelationship)

    if err != nil {
        return [][]types.Event{}, err
    }

    for _, relationship := range transitiveRelationships {
        hbrelationship = append(hbrelationship, relationship)
    }

    return [][]types.Event{}, err
}



