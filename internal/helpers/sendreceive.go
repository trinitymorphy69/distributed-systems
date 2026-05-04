package helpers

import (
	"slices"
	"reflect"
    "github.com/trinitymorphy69/distributed-execution-fundamentals/types"
)

func MatchSendReceive(events []types.Event) ([]types.Event, error) {

    if err := CheckEvent(events); err != nil {
        return []types.Event{}, err
    }

    // Map an overall lamport clock counter initialised to zero to all the processes in the program
    processMap := map[string]int{}

    eventCh := make(chan types.Event, len(events))

    for _, event := range events {
         // Map an overall lamport clock counter initialised to zero to all the processes in the program
        processMap[event.Process] = 0

        // Sends all events into a channel to compare each value so as to map send to receives.
        eventCh <- event
    }

    close(eventCh)

     
    // Sorting the slice into smaller slices of [send, receive]
    sortedEvent := []types.Event{}

    for eventCh != nil {
        value, ok := <-eventCh
        if !ok {
            eventCh = nil
            continue
        }


       // This checks whether the event has already been added to the slice by virtue of being the corresponding receive event of
       // an already processed send event. 

       if found := slices.ContainsFunc(sortedEvent, func(e types.Event) bool {
        	return e.Process == value.Process && e.Number == value.Number && 
            e.LamportClock == value.LamportClock && reflect.DeepEqual(e.VectorClock, value.VectorClock) &&
            e.Type == value.Type && e.Message == value.Message && e.To == value.To && e.From == value.From
   			}); found == true {continue}
        
        //if len(sortedEvent) == len(events)-1 {
            // These checks whether the event to be processed is the last event passed by the program executor and appends it to the array of 
            // sorted events. There's no need for the last event to be checked against anything because it will be the last event regardless 
            // sortedEvent = append(sortedEvent, value)
			//continue
            
		WorkingEvent := value
		for _, event := range events {

			
		//This if conditional checks whether the message of both events are the same. In the case where there are the same;
		//It checks whether their process names match either their corresponding From field or To field.
		//This satisfies the check that if you have two events A and B where either A or B is the send or receive, their messages 
		//should match each other. If A is the send, its To field should match the process of B. If A is receive, its From field should match
		//the process of B.
		
		if WorkingEvent.Message == event.Message {

			if WorkingEvent.Type == 1 && event.Type == 2 {
				// This checks to see whether the To field of the send event matches the From field of the receive event 
				// and the process name of the receive event.
				if EventMatch(WorkingEvent, event) {
                    sortedEvent = append(sortedEvent, WorkingEvent)
                    sortedEvent = append(sortedEvent, event)
                }

			} else if event.Type == 1 && WorkingEvent.Type == 2 {
				if EventMatch(event, WorkingEvent) {
                    sortedEvent = append(sortedEvent, event)
                    sortedEvent = append(sortedEvent, WorkingEvent)
                }

			}
		}

               
        }

    }

    return sortedEvent, nil
    
}



