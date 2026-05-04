package helpers

import (
	"github.com/trinitymorphy69/distributed-systems/types"
)

// This function takes two Events of type Send and Receive and checks to see whether 
// they correspond to each other i.e the receive event matches the send event.  
// It achieves this by comparing the To field of the Send event to the From field and 
// the Process field of the Receive event. Their matching implies that the correspond 
// to each other.

func EventMatch(send types.Event, receive types.Event) bool {

    // We first check that the events passed are valid.
    err := CheckEvent([]types.Event{send, receive})

    // If they are not, we return a false because the 
    // event matching is vacuously false.
    if err != nil {return false}

    // We first check that the fields we will compare to determine if the
    // receive event is the right one for the send event, are not empty.
    if send.Process != "" && send.To != "" && receive.Process != "" && receive.From != "" {

        // For a receive event to match the send event, the send To field must
        // match the receive Process field. And the send Process field must match
        // must match the receive From field.
        if send.To != receive.Process && send.Process != receive.From {
            return false
        }
        return true
    }
     
    return false
}