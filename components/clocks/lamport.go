package clocks

import (
	"github.com/trinitymorphy69/distributed-execution-fundamentals/types"
	"github.com/trinitymorphy69/distributed-execution-fundamentals/internal/helpers"
)

func LamportClock(events []types.Event) ([]types.Event, error) {

	err := helpers.CheckEvent(events)

	if err != nil {
		return []types.Event{}, err
	}

	orderedEvents, err := helpers.EventOrdering(events)

	if err != nil {
		return []types.Event{}, err
	}


	// This map stores processes and their lamport clocks.
	// The lamport clock of a process is simply the clock of its 
	// most recent event. We will use this to compute the clocks for 
	// each event. 
	
	localClocks := map[string]int{}

	for _, event := range events {
		localClocks[event.Process] = 0
	}


	// The are three parts of the lamport clock algorithm;
	// Each process has a counter initialised to zero
	// The counter increments by 1 for every event that occurs
	// on the process. A send event carries the lamport clock to
	// the receive event. The receive event compares its local lamport 
	// clock and the received lamport clock and updates its clock to 
	// the max of both. Then it increments the clock by 1.

	for i := 0; i < len(orderedEvents); i++ {

		
		switch orderedEvents[i].Type {
		case 1, 3:

			// If the event is a send or an internal event, this increment the process 
			// lamport clock by 1 and assigns the new value as the lamport
			// clock of the event.
			localClocks[orderedEvents[i].Process] ++
			orderedEvents[i].LamportClock = localClocks[orderedEvents[i].Process]

		case 2:

			// If the event is a receive event, we get its corresponding send event and
			// compare the lamport clock the lamport clock of the send event and the
			// local clock of the receive event and set the latter to the max of both
			//  and increment by 1.

			// The for loop below seeks to get the matching send event to the receive event.
			sendEvent := types.Event{}

			for _, event := range orderedEvents {

				if helpers.EventMatch(event, orderedEvents[i]) == true {
					sendEvent = event
				}
			}

			// This updates the local clock of the receive event to the max of its counter
			// and the lamport clock of the send event + 1. It also sets the lamport clock
			// of the receive event to the updated local clock

			localClocks[orderedEvents[i].Process] = max(sendEvent.LamportClock, localClocks[orderedEvents[i].Process]) + 1
			orderedEvents[i].LamportClock = localClocks[orderedEvents[i].Process]

		}

	}

	return orderedEvents, nil

}