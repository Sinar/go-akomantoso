package state_assembly

import (
	"reflect"
	"testing"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
)

func TestProcessHansard(t *testing.T) {
	type args struct {
		pdfPath string
	}
	tests := []struct {
		name string
		args args
		want akomantoso.QAHansard
	}{
		// TODO: Add test cases.
		{"Happy path", args{"/Users/leow/GOMOD/go-dundocs/raw/Debate/HANDSARD 8 NOVEMBER 2019 (JUMAAT).pdf"}, akomantoso.QAHansard{
			ID:        "DEWAN NEGERI SELANGOR YANG KEEMPAT BELAS TAHUN 2019",
			QAContent: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessHansard(tt.args.pdfPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessHansard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStateAssemblyDebate(t *testing.T) {
	type args struct {
		pdfPath string
	}
	tests := []struct {
		name string
		args args
		want StateAssemblyDebate
	}{
		{"happy  #1",
			args{"/Users/leow/GOMOD/go-dundocs/raw/Debate/HANDSARD 8 NOVEMBER 2019 (JUMAAT).pdf"},
			StateAssemblyDebate{
				StateAssemblySession: StateAssemblySession{
					ID:      "dunsel14-2019-p1-m1",
					Name:    "DEWAN NEGERI SELANGOR YANG KEEMPAT BELAS TAHUN 2019",
					Term:    "PENGGAL 1",
					Meeting: "MESYUARAT 1",
				},
				Location:  "",
				Date:      "",
				Attended:  nil,
				Missed:    nil,
				QAHansard: akomantoso.QAHansard{},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStateAssemblyDebate(tt.args.pdfPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStateAssemblyDebate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDebateAnalyzer_Process(t *testing.T) {
	type fields struct {
		pdfPath string
	}
	tests := []struct {
		name   string
		fields fields
		want   error
		want1  []akomantoso.Representative
	}{
		{"case #1a", fields{"../../raw/StateAssembly/Hansard/HANSARD-16-JULAI-2020.pdf"}, nil,
			[]akomantoso.Representative{},
		},
		{"case #1b", fields{"../../raw/StateAssembly/Hansard/HANSARD-15-JULAI-2020.pdf"}, nil,
			[]akomantoso.Representative{},
		},
		{"case #1c", fields{"../../raw/StateAssembly/Hansard/HANSARD-14-JULAI-2020.pdf"}, nil,
			[]akomantoso.Representative{},
		},
		{"case #1d", fields{"../../raw/StateAssembly/Hansard/HANSARD-13-JULAI-2020-1.pdf"}, nil,
			[]akomantoso.Representative{},
		},
		//{"case #2", fields{"../../raw/StateAssembly/Hansard/HANSARD-16-MAC-2020.pdf"}, nil,
		//	[]akomantoso.Representative{},
		//},
		//{"case #3", fields{"../../raw/StateAssembly/Hansard/HANSARD-17-MAC-2020.pdf"}, nil,
		//	[]akomantoso.Representative{},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := DebateAnalyzer{
				pdfPath: tt.fields.pdfPath,
			}
			got, got1 := da.Process()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Process() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

//func TestDebateProcessPages(t *testing.T) {
//	type args struct {
//		pdfDocument *akomantoso.PDFDocument
//		dps         DebateProcessorState
//	}
//	// Setup different PDFs and DPSs
//	dps := DebateProcessorState{
//		SectionMarkers: SectionMarkers{
//			DatePageMarker:         "",
//			SessionStartMarkerLine: 7,
//		},
//		CurrentPage:        0,
//		CurrentContent:     "",
//		LastMatchedRep:     "",
//		RepresentativesMap: nil,
//		RolesMap:           nil,
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//		{"case #1", args{nil, dps}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// Makes a copy of the base scenario
//			currentDPS := tt.args.dps
//			spew.Dump(currentDPS)
//
//			DebateProcessPages(tt.args.pdfDocument, currentDPS)
//		})
//	}
//}

func TestDebateProcessPages(t *testing.T) {
	type args struct {
		pdfDocument *akomantoso.PDFDocument
		dps         DebateProcessorState
	}
	// Setup different PDFs and DPSs
	dps := DebateProcessorState{
		SectionMarkers: SectionMarkers{
			DatePageMarker:         "",
			SessionStartMarkerLine: 7,
		},
		CurrentPage:        0,
		CurrentContent:     "",
		LastMatchedRep:     "",
		RepresentativesMap: nil,
		RolesMap:           nil,
	}
	tests := []struct {
		name string
		args args
		want StateAssemblyDebateContent
	}{
		// TODO: Add test cases.
		{"case #1", args{nil, dps}, StateAssemblyDebateContent{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Makes a copy of the base scenario
			currentDPS := tt.args.dps
			if got := DebateProcessPages(tt.args.pdfDocument, currentDPS); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DebateProcessPages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDebateProcessSinglePage(t *testing.T) {
	type args struct {
		allLines []string
		dps      DebateProcessorState
	}
	// Setup different PDFs and DPSs
	dps := DebateProcessorState{
		SectionMarkers: SectionMarkers{
			DatePageMarker:         "",
			SessionStartMarkerLine: 7,
		},
		CurrentPage:        0,
		CurrentContent:     "",
		LastMatchedRep:     "",
		RepresentativesMap: nil,
		RolesMap:           nil,
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case #", args{
			allLines: nil,
			dps:      dps,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentDPS := tt.args.dps
			if err := DebateProcessSinglePage(tt.args.allLines, &currentDPS); (err != nil) != tt.wantErr {
				t.Errorf("DebateProcessSinglePage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
