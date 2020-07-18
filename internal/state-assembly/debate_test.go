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
