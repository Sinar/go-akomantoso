package state_assembly

import (
	"fmt"
	"regexp"
	"strings"

	akomantoso "github.com/Sinar/go-akomantoso/internal/akomantoso"
	"github.com/davecgh/go-spew/spew"
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
func ExtractSessionInfo(pdfPath string) (string, int) {
	// Sample first 10 pages to get the needed info
	extractOptions := akomantoso.ExtractPDFOptions{
		StartPage:  1,
		NumPages:   10,
		MaxSampled: 1000,
	}
	pdfDocument, perr := akomantoso.NewPDFDocument(pdfPath, &extractOptions)
	if perr != nil {
		panic(perr)
	}
	//spew.Dump(pdfDocument.Pages)
	// Sanity  checks ..
	if len(pdfDocument.Pages) < 1 {
		spew.Dump(pdfDocument.Pages)
		panic("Should NOT be here!!")
	}
	// Questions are usaully 2 pages or so  ..
	allLines := make([]string, 5*len(pdfDocument.Pages[0].PDFTxtSameLines))
	for _, singlePageRows := range pdfDocument.Pages {
		allLines = append(allLines, singlePageRows.PDFTxtSameLines...)
	}
	//  DEBUG
	//fmt.Println("========= Cover Pages ====================")
	fmt.Println("NO LINES: ", len(allLines))
	//for _, line := range allLines {
	//	fmt.Println("\"", line, "\",")
	//}
	//fmt.Println("========= END ====================")

	// Extract CoverPage Info by doing the below concurrently
	//  Detect section markers:
	sectionMarkers := extractSectionMarkers(allLines)
	// DEBUG
	spew.Dump(sectionMarkers)
	return sectionMarkers.DatePageMarker, sectionMarkers.SessionStartMarkerLine

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
