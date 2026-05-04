package testhelpers

import (
	"github.com/trinitymorphy69/distributed-execution-fundamentals/types"
)

// For an event to match, the boolean expressions in the function below must be satisfied.
func CheckTrue(send, receive types.Event) bool {

	return send.To != "" && send.Process != "" && receive.Process != "" && receive.From != "" && 
	send.To == receive.Process && send.Process == receive.From

}