package state_assembly

import (
	"reflect"
	"testing"
)

func Test_extractSectionMarkers(t *testing.T) {
	type args struct {
		allLines []string
	}
	tests := []struct {
		name string
		args args
		want SectionMarkers
	}{
		{"happy #1", args{
			[]string{
				"  ",
				" 8 NOVEMBER 2019 (JUMAAT) ",
				" 1  ",
				" DEWAN NEGERI SELANGOR YANG KEEMPAT BELAS TAHUN 2019  ",
				" YANG HADIR  ",
				" Y.B. Tuan Ng Suee Lim (Sekinchan)  ",
				" (Tuan Speaker)  ",
				" 8 NOVEMBER 2019 (JUMAAT) ",
				" 3  ",
				" Y.B. Tuan Mohd Shaid bin Rosli (Jeram)   ",
				" TURUT HADIR  ",
				" (Mengikut Fasal LII (3) Undang-undang Tubuh Kerajaan  ",
				" Selangor, 1959)  ",
				" Y.B. Dato’ Mo�hd Amin bin Ahmad Ahya, D.P.M.S., B.C.M., B.K.T., P.J.K.  ",
				" 8 NOVEMBER 2019 (JUMAAT) ",
				" 7  ",
				"  (TUAN SPEAKER MEMPENGERUSIKAN MESYUARAT)  ",
			}},
			SectionMarkers{
				DatePageMarker:         "8 NOVEMBER 2019 (JUMAAT)",
				PresentMarker:          4,
				AbsentMarker:           0,
				ParticipatedMarker:     10,
				OfficersPresentMarker:  0,
				SessionStartMarkerLine: 16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractSectionMarkers(tt.args.allLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractSectionMarkers() = %v, want %v", got, tt.want)
			}
		})
	}
}
