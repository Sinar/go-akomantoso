package state_assembly

import (
	akomantoso "github.com/Sinar/go-akomantoso/internal/akomantoso"
	"github.com/davecgh/go-spew/spew"
	"github.com/ledongthuc/pdf"
)

// Session represents  the State Assembly  session  for the day
type Session struct {
	SessionName string
	TermName    string
	MeetingName string
	Meta        string // add location, date, start time? end tine?
	Attended    []akomantoso.Representative
	Missed      []akomantoso.Representative
	QAHansard   akomantoso.QAHansard
}

// ProcessHansard will  output QAHansard ..
func ProcessHansard() {
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

	spew.Dump(qaHansard)
	spew.Dump(qaSingleContent)
	//  Extract by row
	f, r, err := pdf.Open("/Users/leow/GOMOD/go-dundocs/raw/Debate/HANDSARD 8 NOVEMBER 2019 (JUMAAT).pdf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	spew.Dump(f.Name())
	spew.Dump(r.NumPage())
	rows, perr := r.Page(7).GetTextByRow()
	if perr != nil {
		panic(perr)
	}
	allLines := make([]string, len(rows))
	for _, row := range rows {
		singleLine := ""
		for _, singleChar := range row.Content {
			singleLine = singleLine + singleChar.S
		}
		allLines = append(allLines, singleLine)
	}
	qaHansard.QAContent[0].Content = allLines
	spew.Dump(qaHansard)
}
