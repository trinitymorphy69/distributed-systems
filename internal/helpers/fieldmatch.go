package helpers

import (
	
	"fmt"
	"github.com/trinitymorphy69/distributed-execution-fundamentals/types"

)

// Checks that the filled fields matches the event type.
// To must Send, From for Receive, none for Internal.
func TypeMatchField(event types.Event) error {
	
	var err error

	switch event.Type {
	case 1:
		if event.To == "" {
			// Checks that a Send Event has a process it is going to.
			return fmt.Errorf("Empty To Field for Send Event %v on Process %v", event.Number, event.Process)
		} 
		
		if event.From != "" {
			//Checks that a Send Event cannot have a filled From field.
			return fmt.Errorf("Filled From Field for Send Event %v on Process %v", event.Number, event.Process)
		} 
		
		if event.To == event.Process {
			return fmt.Errorf("Process field and To field cannot be equal.")
		}

	case 2:
		if event.From == "" {
			// Checks that a Receive Event has a process it is receiving the payload from.
			return fmt.Errorf("Empty From Field for Receive Event %v on Process %v", event.Number, event.Process)
		}
		
		if event.To != "" {
			// FILLED TO FIELD FOR RECEIVE EVENT - Checks that a Receive Event cannot have a filled To field.
			return fmt.Errorf("Filled To Field for Receive Event %v on Process %v", event.Number, event.Process)
		} 
		
		if event.From == event.Process {
			return fmt.Errorf("Process field and From field cannot be equal.")
		}

	case 3:
		if event.To != "" || event.From != "" {
			// Checks that an Event of Type Internal (3)
			// does not have a filled To or From field.
			return fmt.Errorf("Filled To or From Field for Internal Event %v on Process %v", event.Number, event.Process)
		}
	default:
		// Skipping all previous cases to default is an indicator that we have an invalid event.Type so we 
		// return an error for that.
		return fmt.Errorf("Invalid Type for Event %v on Process %v", event.Number, event.Process)
	}
	return err
}