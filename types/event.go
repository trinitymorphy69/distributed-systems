package types


/* 
Process - The process the event is occurring on.
Number - The ID/sequence number of the event on its Process.
Clock - This is the lamport clock of the event.
Type - There are three major types of events in a distributed execution; send (1), receive (2), and internal (3)
Message - This represents the event's payload. 
To - This is the Process a send event will go to. 
From - This is the Process a receive event is coming from.
*/


type Event struct {
    Process string
    Number int
    Type int
    Message string
    To string
    From string
    LamportClock int
    VectorClock map[string]int

}

