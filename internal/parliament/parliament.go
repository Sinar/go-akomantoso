package parliament

// Parliament Session
// Information extracted out by Questions
// Has helper files for Parliament in general?
// DEWAN RAKYAT, PARLIMEN KEEMPAT BELAS, PENGGAL KETIGA, MESYUARAT KEDUA
// Term (Penggal),
// Meeting (Mesyuarat), Session?
// Start /  end ..

// Session represents the Parliament Session for the Term  Range
type ParliamentSession struct {
	ID        string
	Name      string
	Term      string
	Meeting   string
	StartDate string
	EndDate   string
}

func NewParliamentSession() ParliamentSession {
	return ParliamentSession{
		ID:        "short-code-2020",
		Name:      "",
		Term:      "",
		Meeting:   "",
		StartDate: "",
		EndDate:   "",
	}
}

// Determine the namespace for session in the data folder; cache in folder
func (sas ParliamentSession) loadParliamentSession() error {
	// If have  processed this session before, whether debate or  QA
	// extract and  skip the common info  and load from YAML
	return nil
}

// Loop through the individual split PDFs
// Process PDF by Type
// TODO: Determine via use of page 1 of Parliament Hansard
