package parliament

import (
	"fmt"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
	"github.com/davecgh/go-spew/spew"
)

type DebateAnalyzer struct {
	pdfPath string
}

// Location, Date? 8 NOVEMBER 2019 (JUMAAT)
// ParliamentDebate for the  day!
type ParliamentDebate struct {
	ID        string
	Date      string
	Attended  []akomantoso.Representative
	Missed    []akomantoso.Representative
	QAHansard akomantoso.QAHansard
}

func extractSessionInfo(coverPageContent []string) string {
	return "Bil. 1 Isnin 18 Mei 2020 "
}
func getSessionDate() string {
	return "Selasa 28 Julai 2020"
}

func NewParliamentDebate(pdfPath string) ParliamentDebate {
	// Extract out Section Metadata for attachment
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
		// DEBUG
		//spew.Dump(pdfDocument.Pages)
		panic("Should NOT be here!!")
	}
	// Questions are usaully 2 pages or so  ..
	allLines := make([]string, 5*len(pdfDocument.Pages[0].PDFTxtSameLines))
	for _, singlePageRows := range pdfDocument.Pages {
		allLines = append(allLines, singlePageRows.PDFTxtSameLines...)
	}
	//  DEBUG
	//fmt.Println("========= Cover Pages ====================")
	//fmt.Println("NO LINES: ", len(allLines))
	// Debug allLines
	//for _, line := range allLines {
	//	fmt.Println("\"", line, "\",")
	//}
	//fmt.Println("========= END ====================")

	// Extract CoverPage Info by doing the below concurrently
	//  Detect section markers:
	sectionMarkers := extractSectionMarkers(allLines)
	// DEBUG
	//spew.Dump(sectionMarkers)
	// Extract Representatives Detected

	// Namespace: Name-Term-Meeting
	// Persist into data?

	// Will have session.yml out representing the  detected info
	return sectionMarkers.ParliamentDebate
}

func (da DebateAnalyzer) Process() (error, []akomantoso.Representative) {
	// From the Analyzer; we get the start of session; start from there
	// Extract out Section Metadata for attachment
	extractOptions := akomantoso.ExtractPDFOptions{
		StartPage:  2,
		NumPages:   15,
		MaxSampled: 10000,
	}
	pdfDocument, perr := akomantoso.NewPDFDocument(da.pdfPath, &extractOptions)
	if perr != nil {
		panic(perr)
	}
	//spew.Dump(pdfDocument.Pages)
	// Sanity  checks ..
	if len(pdfDocument.Pages) < 1 {
		// DEBUG
		//spew.Dump(pdfDocument.Pages)
		panic("Should NOT be here!!")
	}
	// Questions are usaully 2 pages or so  ..
	allLines := make([]string, 15*len(pdfDocument.Pages[0].PDFTxtSameLines))
	for _, singlePageRows := range pdfDocument.Pages {
		allLines = append(allLines, singlePageRows.PDFTxtSameStyles...)
	}
	extractDebaters(allLines)

	return nil, []akomantoso.Representative{}
}

func (pd ParliamentDebate) ExtractQAHansard() error {
	da := DebateAnalyzer{
		pdfPath: "../../raw/Parliament/Hansard/DR-28072020.pdf",
	}
	da.Process()
	return nil
}

// ProcessHansard will  output QAHansard ..
func processHansard(pdfPath string) akomantoso.QAHansard {
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
		ID:        "DEWAN RAKYAT PARLIMEN KEEMPAT BELAS",
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
