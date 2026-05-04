package helpers

import (
	"github.com/trinitymorphy69/distributed-systems/types"
	"slices"
)

func EventOrdering(events []types.Event) ([]types.Event, error) { 

    
    if err := CheckEvent(events); err != nil {
            return []types.Event{}, err
    }
    

    eventSequenceMap := map[string][]int{}
    processSlice := []string{}
    eventsOrder := []types.Event{}

    for _, event := range events {

        //Takes all the processes into a slice for pre-processing
        if !slices.Contains(processSlice, event.Process) {
            processSlice = append(processSlice, event.Process)
        }

    }

    // The goal of this function is to order the events into how they were executed on their process based on
    // their sequence number. The block below iterates through events to get the numbers of events that belong to
    // the current process in the iteration and add it to a slice mapped to each process. the end goal is to sort
    // each of these slices in an ascending order and use it to assign events into the main process-to-event map.
    
    for _, process := range processSlice {
        for _, event := range events {
            if event.Process == process {
                eventSequenceMap[process] = append(eventSequenceMap[process], event.Number)
            } 
        }
        slices.Sort(eventSequenceMap[process])
    }

    // The eventSequenceMap maps processes submitted to their right order of events. This right order of event
    // is gotten by getting the sequence number of events for a particular process and sort it in an ascending order.
    // What this block of code does is to go into the eventSequenceMap, range on the values (slice) of the map for each
    // of its keys, range on events, and then it appends any event whose event number matches any element of v
    // and process k into the slice of process k in the processMap. The end result is a processMap with processes
    // as keys and slices of ordered events based on eventSequenceMap as values.
    
    
    for k, v := range eventSequenceMap {
        
        for i := 0; i < len(v); i++ {

            for _, event := range events {
                if event.Process == k && event.Number == v[i] {
                    eventsOrder = append(eventsOrder, event)
                }
            }

        }


    }


    return eventsOrder, nil

}