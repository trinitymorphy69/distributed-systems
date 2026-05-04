package clocks

import (
	//"cmp"
	"fmt"
	//"slices"

	"slices"
	"strconv"
	"maps"


	"github.com/trinitymorphy69/distributed-systems/internal/helpers"
	"github.com/trinitymorphy69/distributed-systems/internal/types"
)

func VectorClock(events []types.Event) ([]types.Event, error) {

	// Checks that all the events are valid.
	if err := helpers.CheckEvent(events); err != nil {
		return []types.Event{}, err
	}

	// Orders all passed events.
	orderedEvents, err := helpers.EventOrdering(events)

	if err != nil {
		return []types.Event{}, err
	}

	// This stores the vector clock for all the processes in
	// the distributed execution. A vector clock is a vector
	// of length N for N processes in the system where each 
	// position represents a process. We store this vector as
	// a map of string to int where the string value represent the
	// particular process holding that position, and the integer
	// value represents its clock.

	localClocks := map[string]map[string]int{}

	for _, event := range orderedEvents {
		localClocks[event.Process] =  map[string]int{}
	}

	// Although processes can hold any position in a vector clock,
	// the most efficient way is to follow any ascending order from
	// the first process to the last process. This is indicated by 
	// the number attached to P in the process naming convention.
	
	processSliceI := []int{}
	processSliceS := []string{}

	// The task here is to order the processes in an ascending order.
	// To do this, we first strip of the P which is the first element
	// in the process name and convert the remaining values to integers. 
	// This is stored in the processSliceI. Afterwhich we proceed to sort
	// the slice.

	for k := range localClocks {
		if b, err := strconv.Atoi(k[1:]); err != nil {
			return []types.Event{}, err
		} else {
			processSliceI = append(processSliceI, b)
		}
	}

	slices.Sort(processSliceI)

	// Here, we convert the values in ProcessSliceI back to process names
	// while still maintaining their sorted order and store in processSliceS.

	for i := 0; i < len(processSliceI); i++ {
		processSliceS = append(processSliceS, "P"+ strconv.Itoa(processSliceI[i]))
	}

	// Here, we proceed to set the value and order of the vector clocks using
	// the ordered values from processSliceS.

	for k := range localClocks {

		for i := 0; i < len(processSliceS); i++ {
			localClocks[k][processSliceS[i]] = 0
		}

	}

	// The are three parts of the vector clock algorithm;
	// Each process has a vector clock initialised to zero where each position
	// of the vector clock respresents a process. For its position in the clock, 
	// the counter increments by 1 for every event that occurs on the process. 
	// A send event carries the vector clock to the receive event. The receive event 
	// compares its local vector clock and the received vector clock in a pointwise 
	// maximum manner and updates its clock to the max of each position. 
	// Then it increments its position by 1.

	for i := 0; i < len(orderedEvents); i++ {

		switch orderedEvents[i].Type {
		case 1, 3:

			// If the event is a send or an internal event, this increment the process 
			// position in the vector clock by 1 and assigns the new vector as the vector
			// clock of the event.

			localClocks[orderedEvents[i].Process][orderedEvents[i].Process] += 1 
			orderedEvents[i].VectorClock = maps.Clone(localClocks[orderedEvents[i].Process])
			

		case 2:

			// If the event is a receive event, we get its corresponding send event and
			// compare the vector clock of the send event and the local vector clock of
			// the receive event and set the latter to the max of both and increment the 
			// process position by 1.

			// The for loop below seeks to get the matching send event to the receive event.
			sendEvent := types.Event{}

			for _, event := range orderedEvents {

				if helpers.EventMatch(event, orderedEvents[i]) == true {
					sendEvent = event
				}
			}

			// This updates the local clock of the receive event to the max of its counter
			// and the vector clock of the send event and adds 1 to the position of the 
			// process of the receive event. It also sets the vector clock of the receive 
			// event to the updated local clock.

			// Here we range in the maps that store the vector clocks of each process.  in the position of the 
			// process it was sent from is 
			for k, v := range localClocks[orderedEvents[i].Process] {

				// We check whether the local vector clock is greater than the vector clock
				// of the send event by comparing all its positions. We then update the positions
				// where the vector clock of the send event is higher than the local vector clock
				// of the process.

				if sendEvent.VectorClock[orderedEvents[i].Process] > v {


					localClocks[orderedEvents[i].Process][k] = sendEvent.VectorClock[k]

				}
				
				// After updating the vector clock, we then increment the position of the process
				// the event is on by 1 
				localClocks[orderedEvents[i].Process][k]++

				// Here we set the vector clock of the event to the vector clock of the process.
				orderedEvents[i].VectorClock = maps.Clone(localClocks[orderedEvents[i].Process])

			}

		}
		 

	}

	for i := 0; i < len(orderedEvents); i++ {
		fmt.Printf("%v(%v) ⟶ %v\n", orderedEvents[i].Process, orderedEvents[i].Number, orderedEvents[i].VectorClock)
	}
	
	return orderedEvents, nil
	
}