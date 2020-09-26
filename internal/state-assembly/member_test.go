package state_assembly

import (
	"reflect"
	"testing"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
)

func Test_extractDebaters(t *testing.T) {
	type args struct {
		allLines  []string
		pdfPath   string
		startPage int
		numPages  int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"case: missing bkt melawati", args{
			allLines:  nil,
			pdfPath:   "../../raw/StateAssembly/Hansard/HANSARD-16-JULAI-2020.pdf",
			startPage: 11,
			numPages:  2,
		}, []string{"bob"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allLines := tt.args.allLines
			if tt.args.pdfPath != "" {
				allLines = setupExtractDebaters(tt.args.pdfPath, tt.args.startPage, tt.args.numPages)
			}
			if got := extractDebaters(allLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractDebaters() = %v, want %v", got, tt.want)
			}
		})
	}
}

// setupExtractDebaters is helper function to replicate tircky issues
func setupExtractDebaters(pdfPath string, startPage int, numPages int) []string {
	extractOptions := akomantoso.ExtractPDFOptions{
		StartPage:  startPage,
		NumPages:   numPages,
		MaxSampled: 10000,
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
	allLines := make([]string, 30*len(pdfDocument.Pages[0].PDFTxtSameLines))
	for _, singlePageRows := range pdfDocument.Pages {
		allLines = append(allLines, singlePageRows.PDFTxtSameStyles...)
	}
	return allLines
}

func Test_looksLikeRep(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := looksLikeRep(tt.args.line)
			if got != tt.want {
				t.Errorf("looksLikeRep() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("looksLikeRep() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
