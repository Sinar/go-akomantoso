package state_assembly

import "github.com/Sinar/go-akomantoso/internal/akomantoso"

// Session represents  the State Assembly  session  for the day
type Session struct {
	SessionName string
	TermName    string
	MeetingName string
	Meta        string // add location, date, start time? end tine?
	Attended    []akomantoso.Representative
	Missed      []akomantoso.Representative
}
