package goenums

import "fmt"

// an enum is a type that has a fixed number of possible values each with a distinct name.
// Go doesn't have an enum tyoe as a distinct  language feature, but enums are simple to implement using an existing language idioms

// declare an enum
type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

type AttendanceStatus string

const (
	StatusPresent AttendanceStatus = "present"
	StatusAbsent  AttendanceStatus = "absent"
	StatusUnknown AttendanceStatus = "unknown"
)

type Bases uint

const (
	BaseA Bases = 'A'
	BaseC Bases = 'C'
	BaseT Bases = 'T'
	BaseG Bases = 'G'
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// iota keyword generates successive constant values automatically

// use the enum
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func GoEnums() {
	fmt.Println("This is a go enum!!")
}
