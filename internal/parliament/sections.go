package parliament

import (
	"regexp"
	"strings"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
)

// Detect the Section Markers in the Cover Pages
// * Metadata (first date + page number); ignore line with  just number in it?
//  Output: [<Line FOUND> for each category; it may be missing]
type SectionMarkers struct {
	DatePageMarker         string
	ParliamentDebate       ParliamentDebate
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

func hasSessionStartMarkerLine(line string) bool {
	// Tuan Yang di-Pertua mempengerusikan Mesyuarat
	matched, err := regexp.MatchString(`tuan yang di-pertua mempengerusikan mesyuarat`, strings.ToLower(line))
	if err != nil {
		panic(err)
	}
	return matched
}

func extractSectionMarkers(allLines []string) SectionMarkers {
	var datePageMarker string
	var sessionStartMarkerLinePage int
	for i, line := range allLines {
		// DEBUG
		//fmt.Println("LINE: ", i)
		normalizedLine := strings.TrimSpace(line)
		if hasDatePageMarker(normalizedLine) {
			datePageMarker = normalizedLine
			continue
		}
		if hasSessionStartMarkerLine(normalizedLine) {
			sessionStartMarkerLinePage = i
			break
		}
	}
	// DEBUG
	//spew.Dump(allLines)
	return SectionMarkers{
		DatePageMarker: datePageMarker,
		ParliamentDebate: ParliamentDebate{
			ID:        "",
			Date:      "",
			Attended:  nil,
			Missed:    nil,
			QAHansard: akomantoso.QAHansard{},
		},
		SessionStartMarkerLine: sessionStartMarkerLinePage,
	}
}

// Detect the Sections within Debate  Session (?)
