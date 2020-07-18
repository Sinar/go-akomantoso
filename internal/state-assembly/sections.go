package state_assembly

// Detect the Section Markers in the Cover Pages
// * Metadata (first date + page number); ignore line with  just number in it?
//  Output: [<Line FOUND> for each category; it may be missing]
type SectionMarkers struct {
	DatePageMarker         string
	PresentMarker          int // * HADIR
	AbsentMarker           int // * TIDAK HADIR
	ParticipatedMarker     int // * TURUT HADIR
	OfficersPresentMarker  int // * PEGAWAI BERTUGAS
	SessionStartMarkerLine int // * Session Start (END)
}

func extractSectionMarkers(allLines []string) SectionMarkers {
	// DEBUG
	//spew.Dump(allLines)
	return SectionMarkers{
		DatePageMarker:         "",
		PresentMarker:          0,
		AbsentMarker:           0,
		ParticipatedMarker:     0,
		OfficersPresentMarker:  0,
		SessionStartMarkerLine: 0,
	}
}

// Detect the Sections within Debate  Session (?)
