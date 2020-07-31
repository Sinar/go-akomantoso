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
	// As per: https://stackoverflow.com/a/42251527
	normalizedLine := strings.Join(strings.Fields(strings.ToLower(line)), " ")
	// Page Boundary: DR.18.5.2020 1
	// could  be DR. 28.7.2020 1 as well
	// can be as simple as `^dr\.\s*\d+`
	matched, err := regexp.MatchString(`^dr\.\s*\d+.*\d+$`, normalizedLine)
	if err != nil {
		panic(err)
	}
	return matched
}

func hasParliamentDebateID(line string, parliamentDebateID *string) bool {
	//  Bil. 1 Isnin      18 Mei 2020
	// As per: https://stackoverflow.com/a/42251527
	normalizedLine := strings.Join(strings.Fields(strings.ToLower(line)), " ")
	matched, err := regexp.MatchString(`isnin|selasa|rabu|khamis|jumaat`, normalizedLine)
	if err != nil {
		panic(err)
	}
	if matched {
		// Check it starts with bil
		bmatched, berr := regexp.MatchString(`bil`, normalizedLine)
		if berr != nil {
			panic(berr)
		}
		if bmatched {
			// Just naive matching + replace
			normalizedLine = strings.ReplaceAll(normalizedLine, ".", "")
			normalizedLine = strings.ReplaceAll(normalizedLine, " ", "-")
			// Bring out the final transformation back
			*parliamentDebateID = normalizedLine
		}
	}
	// Replace non alphanum with ''
	// Replace :space: with  '-'
	return false
}

func hasParliamentDebateDate(line string) bool {
	// Isnin, 18 Mei 2020
	// As per: https://stackoverflow.com/a/42251527
	normalizedLine := strings.Join(strings.Fields(strings.ToLower(line)), " ")
	matched, err := regexp.MatchString(`isnin|selasa|rabu|khamis|jumaat`, normalizedLine)
	if err != nil {
		panic(err)
	}
	return matched
}

func hasSessionStartMarkerLine(line string) bool {
	// As per: https://stackoverflow.com/a/42251527
	normalizedLine := strings.Join(strings.Fields(strings.ToLower(line)), " ")
	// Tuan Yang di-Pertua mempengerusikan Mesyuarat
	matched, err := regexp.MatchString(`tuan yang di-pertua mempengerusikan mesyuarat`, normalizedLine)
	if err != nil {
		panic(err)
	}
	return matched
}

func extractSectionMarkers(allLines []string) SectionMarkers {
	var datePageMarker, parliamentDebateID, parliamentDebateDate string
	var sessionStartMarkerLinePage int
	for i, line := range allLines {
		// DEBUG
		//fmt.Println("LINE: ", i)
		normalizedLine := strings.TrimSpace(line)
		if hasDatePageMarker(normalizedLine) {
			// DEBUG
			//fmt.Println("MATCHED!!!! ==> ", normalizedLine)
			// Join all but the last segment as that will be the page number
			// can be a function later  to extract out pages  ,.
			datePageFields := strings.Fields(normalizedLine)
			for i, singleField := range datePageFields {
				if i < len(datePageFields)-1 {
					datePageMarker = datePageMarker + singleField
				}
			}
			continue
		}
		if hasParliamentDebateID(normalizedLine, &parliamentDebateID) {
			// Found it, store it normalized and move on ..
			continue
		}
		if hasParliamentDebateDate(normalizedLine) {
			parliamentDebateDate = normalizedLine
			continue
		}
		if hasSessionStartMarkerLine(normalizedLine) {
			sessionStartMarkerLinePage = i
			break
		}
	}
	// Extracted out metadata for ParliamentDebate Day
	parliamentDebate := ParliamentDebate{
		ID:        parliamentDebateID,
		Date:      parliamentDebateDate,
		Attended:  nil,
		Missed:    nil,
		QAHansard: akomantoso.QAHansard{},
	}
	// DEBUG
	//spew.Dump("ID ==> ", parliamentDebate.ID)
	return SectionMarkers{
		DatePageMarker:         datePageMarker,
		ParliamentDebate:       parliamentDebate,
		SessionStartMarkerLine: sessionStartMarkerLinePage,
	}
}

// Detect the Sections within Debate  Session (?)
