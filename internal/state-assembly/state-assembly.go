package state_assembly

// Has helper files for State Assembly in general?
// DEWAN NEGERI SELANGOR YANG KEEMPAT BELAS TAHUN 2019
// Term (Penggal),
// Meeting (Mesyuarat), Session?
// Start /  end ..

// Session represents  the State Assembly Session for the Term  Range
type StateAssemblySession struct {
	ID        string
	Name      string
	Term      string
	Meeting   string
	StartDate string
	EndDate   string
}

func NewStateAssemblySession() StateAssemblySession {
	return StateAssemblySession{
		ID:        "short-code-2019",
		Name:      "",
		Term:      "",
		Meeting:   "",
		StartDate: "",
		EndDate:   "",
	}
}

// Determine the namespace for session in the data folder; cache in folder
func (sas StateAssemblySession) loadStateAssemblySession() error {
	// If have  processed this session before, whether debate or  QA
	// extract and  skip the common info  and load from YAML
	return nil
}
