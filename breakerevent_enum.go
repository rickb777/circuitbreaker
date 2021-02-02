// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.0.2

package circuit

import (
	"fmt"
)

const breakereventEnumStrings = "BreakerTrippedBreakerResetBreakerFailBreakerReady"

var breakereventEnumIndex = [...]uint16{0, 14, 26, 37, 49}

// AllBreakerEvents lists all 4 values in order.
var AllBreakerEvents = []BreakerEvent{
	BreakerTripped, BreakerReset, BreakerFail, BreakerReady,
}

// String returns the literal string representation of a BreakerEvent, which is
// the same as the const identifier.
func (i BreakerEvent) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllBreakerEvents) {
		return fmt.Sprintf("BreakerEvent(%d)", i)
	}
	return breakereventEnumStrings[breakereventEnumIndex[o]:breakereventEnumIndex[o+1]]
}

// Tag returns the string representation of a BreakerEvent. This is an alias for String.
func (i BreakerEvent) Tag() string {
	return i.String()
}

// Ordinal returns the ordinal number of a BreakerEvent.
func (i BreakerEvent) Ordinal() int {
	switch i {
	case BreakerTripped:
		return 0
	case BreakerReset:
		return 1
	case BreakerFail:
		return 2
	case BreakerReady:
		return 3
	}
	return -1
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i BreakerEvent) Int() int {
	return int(i)
}

// BreakerEventOf returns a BreakerEvent based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid BreakerEvent is returned.
func BreakerEventOf(i int) BreakerEvent {
	if 0 <= i && i < len(AllBreakerEvents) {
		return AllBreakerEvents[i]
	}
	// an invalid result
	return BreakerTripped + BreakerReset + BreakerFail + BreakerReady + 1
}

// IsValid determines whether a BreakerEvent is one of the defined constants.
func (i BreakerEvent) IsValid() bool {
	switch i {
	case BreakerTripped, BreakerReset, BreakerFail, BreakerReady:
		return true
	}
	return false
}
