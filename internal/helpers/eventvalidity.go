package helpers

import (
	"fmt"
	"github.com/trinitymorphy69/distributed-systems/internal/types"
)

func CheckEvent(events []types.Event) error {

	var err error

	for _, event := range events {

		if event.Process == "" {
			// Checks whether the Process the event occurs on is specified.
			return fmt.Errorf("Empty Process for Event %v", event)

		} else if event.Number <= 0 {
			// Checks whether the event number is specified. This event number
			// communicates the order of events on a particular process. It must not be
			// zero or a negative number.
			return fmt.Errorf("Invalid Number for Event %v on Process %v", event.Number, event.Process)

		} else if event.LamportClock != 0 {
			// The Lamport clock of an event should be set to zero. It will be calculated based on the 
			// ascertained order of events within the distributed execution.
			return fmt.Errorf("Do not assign a Lamport clock Value for Event %v on Process %v", event.Number, event.Process)

		} else if len(event.VectorClock) != 0 {
			// The vector clock of an event should be set to zero. It will be calculated based on the 
			// ascertained order of events within the distributed execution.
			return fmt.Errorf("Do not assign a Vector clock value for Event %v on Process %v", event.Number, event.Process)

		} else if event.Type < 1 || event.Type > 3 {
			// Checks whether the Number representative of the specific event type is valid.
			return fmt.Errorf("Invalid Type for Event %v on Process %v", event.Number, event.Process)
		} 
			
		if err := ProcessNamingConvention(event); err != nil {
			// Checks that the content of the Process, From, and To field satisfies the naming
			// convention for a process
			return err
		}

		if event.Type >= 1 && event.Type <= 3 {
			// Checks that the filled fields matches the event type.
			// To must Send, From for Receive, none for Internal.
    		if err := TypeMatchField(event); err != nil {
        	return err
   			}
		}

	}
	return err

}




func CheckEventMinimal(events []types.EventMinimal) error {

	var err error 

	for _, event := range events {

		if event.Process == "" {
			// Checks whether the Process the event occurs on is specified.
			return fmt.Errorf("Empty Process for Event %v", event)

		} else if event.Number <= 0 {
			// Checks whether the event number is specified. This event number
			// communicates the order of events on a particular process. It must not be
			// zero or a negative number.
			return fmt.Errorf("Invalid Number for Event %v on Process %v", event.Number, event.Process)

		}

		if err := ProcessNamingConventionforMinimalEvent(event); err != nil {
			// Checks that the content of the Process, From, and To field satisfies the naming
			// convention for a process
			return err
		}
	}

	return err
}