package parliament

import (
	"regexp"
	"strings"
)

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

func hasDatePageMarker(line string) bool {
	// Look out for day Parliaments sit; isnin|selasa|rabu|khamis|jumaat
	matched, err := regexp.MatchString(`isnin|selasa|rabu|khamis|jumaat`, strings.ToLower(line))
	if err != nil {
		panic(err)
	}
	return matched
}

func hasPresentMarker(line string) bool {
	// YANG HADIR
	matched, err := regexp.MatchString(`yang hadir`, strings.ToLower(line))
	if err != nil {
		panic(err)
	}
	return matched
}

func hasAbsentMarker(line string) bool {
	// TIDAK HADIR
	matched, err := regexp.MatchString(`tidak hadir`, strings.ToLower(line))
	if err != nil {
		panic(err)
	}
	return matched
}

func hasParticipatedMarker(line string) bool {
	// TURUT HADIR
	matched, err := regexp.MatchString(`turut hadir`, strings.ToLower(line))
	if err != nil {
		panic(err)
	}
	return matched
}

func hasOfficersPresentMarker(line string) bool {
	// PEGAWAI BERTUGAS
	matched, err := regexp.MatchString(`pegawai bertugas`, strings.ToLower(line))
	if err != nil {
		panic(err)
	}
	return matched
}

func hasSessionStartMarkerLine(line string) bool {
	// SPEAKER MEMPENGERUSIKAN
	matched, err := regexp.MatchString(`speaker mempengerusikan`, strings.ToLower(line))
	if err != nil {
		panic(err)
	}
	return matched
}

func extractSectionMarkers(allLines []string) SectionMarkers {
	var datePageMarker string
	var presentMakerPage, absentMarkerPage, participatedMarkerPage, officersPresentMarkerPage, sessionStartMarkerLinePage int
	for i, line := range allLines {
		// DEBUG
		//fmt.Println("LINE: ", i)
		normalizedLine := strings.TrimSpace(line)
		if hasDatePageMarker(normalizedLine) {
			datePageMarker = normalizedLine
		}
		if hasPresentMarker(normalizedLine) {
			presentMakerPage = i
		}
		if hasAbsentMarker(normalizedLine) {
			absentMarkerPage = i
		}
		if hasParticipatedMarker(normalizedLine) {
			participatedMarkerPage = i
		}
		if hasOfficersPresentMarker(normalizedLine) {
			officersPresentMarkerPage = i
		}
		if hasSessionStartMarkerLine(normalizedLine) {
			sessionStartMarkerLinePage = i
		}
	}
	// DEBUG
	//spew.Dump(allLines)
	return SectionMarkers{
		DatePageMarker:         datePageMarker,
		PresentMarker:          presentMakerPage,
		AbsentMarker:           absentMarkerPage,
		ParticipatedMarker:     participatedMarkerPage,
		OfficersPresentMarker:  officersPresentMarkerPage,
		SessionStartMarkerLine: sessionStartMarkerLinePage,
	}
}

// Detect the Sections within Debate  Session (?)
