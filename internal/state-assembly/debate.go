package state_assembly

import (
	"fmt"
	"regexp"
	"strings"

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

type DebateContent struct {
	RepresentativeID akomantoso.RepresentativeID
	RawContent       string
	FinalContent     string
}

type StateAssemblyDebateContent struct {
	DebateContents []DebateContent
}

type DebateProcessorState struct {
	SectionMarkers     SectionMarkers
	CurrentPage        int
	CurrentContents    []DebateContent
	LastPedingContent  DebateContent
	RepresentativesMap map[string]akomantoso.RepresentativeID
	RolesMap           map[string]akomantoso.RepresentativeID
}

type DebateAnalyzer struct {
	pdfPath string
}

func (da DebateAnalyzer) Process() (error, []akomantoso.Representative) {
	// From the Analyzer; we get the start of session; start from there
	// Extract out Section Metadata for attachment
	extractOptions := akomantoso.ExtractPDFOptions{
		//StartPage: 7,
		//NumPages:   10,
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
	allLines := make([]string, 30*len(pdfDocument.Pages[0].PDFTxtSameLines))
	for _, singlePageRows := range pdfDocument.Pages {
		allLines = append(allLines, singlePageRows.PDFTxtSameStyles...)
	}
	extractDebaters(allLines)

	return nil, []akomantoso.Representative{}
}

func DebateProcessSinglePage(allLines []string, dps *DebateProcessorState) error {
	// Extract out each block and find next block of texts
	// DEBUG
	//spew.Dump(allLines)
	// Skip page headers and page number (first 2 lines)
	var pendingDebateContent DebateContent
	// If came from previous round; LastPedingContent not empty
	pendingDebateContent = dps.LastPedingContent
	for i, singleRow := range allLines {
		// Test case generation
		fmt.Println(fmt.Sprintf("\"%s\",", singleRow))
		if i > 1 {
			// Split by colon
			splitRow := strings.Split(singleRow, ":")
			// If cannot split by ':', no reps
			if len(splitRow) > 1 {
				// Remove special chars, extra spaces
				// Extra space removal
				whitespaces := regexp.MustCompile(`\s+`)
				singleRow = whitespaces.ReplaceAllString(splitRow[0], " ")
				// Remove special chars
				singleRow = removeNonASCII(singleRow)
				// Remove '.'
				singleRow = strings.ReplaceAll(singleRow, ".", "")
				singleRow = strings.Trim(singleRow, " ")
				repID := dps.RepresentativesMap[singleRow]
				if repID != "" {
					// DEBUG
					//fmt.Println("REP: ", singleRow, " ID: ", repID)
					if pendingDebateContent.RepresentativeID != "" {
						// Recognize Rep in here .. finalize previous and attach to last Rep
						// ONLY if there were RepID already; which is NOT there in first round
						dps.CurrentContents = append(dps.CurrentContents, pendingDebateContent)
					}
					// publish the DebateContent and start a new one
					pendingDebateContent = DebateContent{
						RepresentativeID: repID,
					}
					// Content start with the other half ..
					pendingDebateContent.RawContent = splitRow[1]
				} else {
					// Append the line content
					pendingDebateContent.RawContent += singleRow
				}
			} else {
				// Append the line content
				pendingDebateContent.RawContent += singleRow
			}
		}
	}
	// Last left over .. should become LeftoverContent
	dps.LastPedingContent = pendingDebateContent
	// DEBUG
	//spew.Dump(dps.CurrentContents)
	//fmt.Println("LEFT OVER: ", dps.LastPedingContent)

	return nil
}

func DebateProcessPages(pdfDocument *akomantoso.PDFDocument, dps DebateProcessorState) StateAssemblyDebateContent {
	saStateAssemblyDebateContent := StateAssemblyDebateContent{}
	for _, singlePageRow := range pdfDocument.Pages {
		DebateProcessSinglePage(singlePageRow.PDFTxtSameLines, &dps)
		// Should signsl end of debate here?
	}
	// Edge case, hit completion; append the last content to the last Representative
	return saStateAssemblyDebateContent
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
