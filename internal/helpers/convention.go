package helpers

import (
	"github.com/trinitymorphy69/distributed-systems/types"
	"fmt"
	"strconv"

)

// Checks that the content of the Process field satisfies the naming
// convention for a process. The convention is p123..n where n is a
// natural number.
func ProcessNamingConvention(event types.Event) error {

	var err error

	if event.Process != "" {

		if event.Process[0] != 'P' {
		// Checks that P prefixes the process field.
		return fmt.Errorf("Prefix P to the Process Field for Event %v", event)

		} else if _, err := strconv.Atoi(string(event.Process[1:])); err != nil {
		// Checks that everything after P in the process field is a natural number.
		return fmt.Errorf("Only integers should suffix P in the Process Field for Event %v", event)

		} 
		
	}
	
	if event.To != "" {

		if event.To[0] != 'P' {
		// Checks that P prefixes the process field.
		return fmt.Errorf("Prefix P to the To field for Event %v", event)

		} else if _, err := strconv.Atoi(string(event.To[1:])); err != nil {
		// Checks that everything after P in the process field is a natural number.
		return fmt.Errorf("Only integers should suffix P to the To field for Event %v", event)

		} 

	}
	
	if event.From != "" {

		if event.From[0] != 'P' {
			// Checks that P prefixes the process field.
			return fmt.Errorf("Prefix P to From field for Event %v", event)
		} 
		
		if _, err := strconv.Atoi(string(event.From[1:])); err != nil {
			// Checks that everything after P in the process field is a natural number.
			return fmt.Errorf("Only integers should suffix P to the From field for Event %v", event)
		} 

	}

	return err
}


func ProcessNamingConventionforMinimalEvent(event types.EventMinimal) error {

	var err error

	if event.Process != "" {

		if event.Process[0] != 'P' {
		// Checks that P prefixes the process field.
		return fmt.Errorf("Prefix P to the Process Field for Event %v", event)

		} else if _, err := strconv.Atoi(string(event.Process[1:])); err != nil {
		// Checks that everything after P in the process field is a natural number.
		return fmt.Errorf("Only integers should suffix P in the Process Field for Event %v", event)

		} 
		
	}

	return err

}
