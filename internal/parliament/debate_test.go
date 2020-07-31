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
				ID:        "bil-11-selasa-28-julai-2020",
				Date:      "Selasa, 28 Julai 2020",
				Attended:  nil,
				Missed:    nil,
				QAHansard: akomantoso.QAHansard{},
			}},
		{"happy  #2",
			args{"../../raw/Parliament/Hansard/DR-18052020.pdf"},
			ParliamentDebate{
				ID:        "bil-1-isnin-18-mei-2020",
				Date:      "Isnin, 18 Mei 2020",
				Attended:  nil,
				Missed:    nil,
				QAHansard: akomantoso.QAHansard{},
			}},
		{"happy  #3",
			args{"../../raw/Parliament/Hansard/DR-13072020 New 1.pdf"},
			ParliamentDebate{
				ID:        "bil-2-isnin-13-julai-2020",
				Date:      "Isnin, 13 Julai 2020",
				Attended:  nil,
				Missed:    nil,
				QAHansard: akomantoso.QAHansard{},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParliamentDebate(tt.args.pdfPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParliamentDebate() = %v, want %v", got, tt.want)
			}
		})
	}
}
