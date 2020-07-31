package state_assembly

import (
	"fmt"

	akomantoso "github.com/Sinar/go-akomantoso/internal/akomantoso"
	"github.com/davecgh/go-spew/spew"
)

// Location, Date? 8 NOVEMBER 2019 (JUMAAT)
// StateAssemblyDebate for the  day!
type StateAssemblyDebate struct {
	StateAssemblySession StateAssemblySession
	Location             string
	Date                 string
	Attended             []akomantoso.Representative
	Missed               []akomantoso.Representative
	QAHansard            akomantoso.QAHansard
}

func extractSessionInfo(coverPageContent []string) string {
	return "SHAH ALAM, 8 NOVEMBER 2019 (JUMAAT) "
}
func getSessionDate() string {
	return "8 NOVEMBER 2019 (JUMAAT)"
}

func NewStateAssemblyDebate(pdfPath string) StateAssemblyDebate {
	extractOptions := akomantoso.ExtractPDFOptions{
		StartPage:  1,
		NumPages:   7,
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
	// Extract Representatives Detected

	// Namespace: Name-Term-Meeting
	// Persist into data?

	// Will have session.yml out representing the  detected info
	return StateAssemblyDebate{
		StateAssemblySession: StateAssemblySession{},
		Location:             "",
		Date:                 "",
		Attended:             nil,
		Missed:               nil,
		QAHansard:            akomantoso.QAHansard{},
	}
}

// ProcessHansard will  output QAHansard ..
func ProcessHansard(pdfPath string) akomantoso.QAHansard {
	// Test single content
	qaSingleContent := akomantoso.QAContent{
		ID:       "Q.1",
		Content:  nil,
		Title:    "",
		QContent: nil,
		QBy:      akomantoso.Representative{},
		AContent: nil,
		ABy:      akomantoso.Representative{},
	}
	qaContent := make([]akomantoso.QAContent, 0) //  Whole session for  Debate
	qaContent = append(qaContent, qaSingleContent)
	qaHansard := akomantoso.QAHansard{
		ID:        "DEWAN NEGERI SELANGOR YANG KEEMPAT BELAS TAHUN 2019",
		QAContent: qaContent,
	}

	//spew.Dump(qaHansard)
	//spew.Dump(qaSingleContent)

	//pdfPath := "/Users/leow/GOMOD/go-dundocs/raw/Debate/HANDSARD 8 NOVEMBER 2019 (JUMAAT).pdf"
	extractOptions := akomantoso.ExtractPDFOptions{
		StartPage:  1,
		NumPages:   3,
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
		panic("SHould NOT be here!!")
	}
	// Questions are usaully 2 pages or so  ..
	allLines := make([]string, 2*len(pdfDocument.Pages[0].PDFTxtSameLines))
	for _, singlePageRows := range pdfDocument.Pages {
		allLines = append(allLines, singlePageRows.PDFTxtSameLines...)
	}
	qaHansard.QAContent[0].Content = allLines
	//  DEBUG
	fmt.Println("========= QAHansard Output ====================")
	spew.Dump(qaHansard)
	fmt.Println("========= END ====================")

	return qaHansard
}

func (sad StateAssemblyDebate) ExtractQAHansard() error {
	return nil
}
