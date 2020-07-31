package parliament

import (
	"reflect"
	"testing"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
)

func TestNewParliamentDebate(t *testing.T) {
	type args struct {
		pdfPath string
	}
	tests := []struct {
		name string
		args args
		want ParliamentDebate
	}{
		{"happy  #1",
			args{"../../raw/Parliament/Hansard/DR-28072020.pdf"},
			ParliamentDebate{
				Date:      "",
				Attended:  nil,
				Missed:    nil,
				QAHansard: akomantoso.QAHansard{},
			}},
		//{"happy  #2",
		//	args{"../../raw/Parliament/Hansard/DR-18052020.pdf"},
		//	ParliamentDebate{
		//		Date:      "",
		//		Attended:  nil,
		//		Missed:    nil,
		//		QAHansard: akomantoso.QAHansard{},
		//	}},
		//{"happy  #3",
		//	args{"../../raw/Parliament/Hansard/DR-13072020 New 1.pdf"},
		//	ParliamentDebate{
		//		Date:      "",
		//		Attended:  nil,
		//		Missed:    nil,
		//		QAHansard: akomantoso.QAHansard{},
		//	}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParliamentDebate(tt.args.pdfPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParliamentDebate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParliamentDebate_ExtractQAHansard(t *testing.T) {
	type fields struct {
		ParliamentSession ParliamentSession
		Location          string
		Date              string
		Attended          []akomantoso.Representative
		Missed            []akomantoso.Representative
		QAHansard         akomantoso.QAHansard
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sad := ParliamentDebate{
				Date:      tt.fields.Date,
				Attended:  tt.fields.Attended,
				Missed:    tt.fields.Missed,
				QAHansard: tt.fields.QAHansard,
			}
			if err := sad.ExtractQAHansard(); (err != nil) != tt.wantErr {
				t.Errorf("ExtractQAHansard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessHansard(tt.args.pdfPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessHansard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractSessionInfo(t *testing.T) {
	type args struct {
		coverPageContent []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractSessionInfo(tt.args.coverPageContent); got != tt.want {
				t.Errorf("extractSessionInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSessionDate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSessionDate(); got != tt.want {
				t.Errorf("getSessionDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
