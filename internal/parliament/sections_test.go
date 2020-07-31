package parliament

import (
	"reflect"
	"testing"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
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
				" PENYATA RASMI PARLIMEN  ",
				" DEWAN RAKYAT  ",
				" PARLIMEN KEEMPAT BELAS  ",
				" PENGGAL KETIGA  ",
				" MESYUARAT KEDUA  ",
				" Bil. 11 Selasa        28 Julai 2020  ",
				" K A N D U N G A N  ",
				" JAWAPAN-JAWAPAN MENTERI BAGI PERTANYAAN-PERTANYAAN (Halaman      1)  ",
				"   Peraturan Mesyuarat (Halaman  108)  ",
				" DR. 28.7.2020 1  ",
				" Selasa, 28 Julai 2020  ",
				" Mesyuarat dimulakan pada pukul 10.00 pagi  ",
				" DOA  ",
				" [Tuan Yang di-Pertua mempengerusikan Mesyuarat]  ",
				" JAWAPAN-JAWAPAN MENTERI BAGI PERTANYAAN-PERTANYAAN  ",
				" 1.  Dato Sri Hajah Rohani binti Abdul Karim [Batang Lupar] minta Menteri Perumahan  ",
			}},
			SectionMarkers{
				DatePageMarker: "DR.28.7.2020",
				ParliamentDebate: ParliamentDebate{
					ID:        "bil-11-selasa-28-julai-2020",
					Date:      "Selasa, 28 Julai 2020",
					Attended:  nil,
					Missed:    nil,
					QAHansard: akomantoso.QAHansard{},
				},
				SessionStartMarkerLine: 14,
			},
		},
		{"happy #2", args{
			[]string{
				"  ",
				"  ",
				" Naskhah belum disemak (DR)  ",
				" PENYATA RASMI PARLIMEN  ",
				" PARLIMEN KEEMPAT BELAS  ",
				" PENGGAL KETIGA  ",
				" MESYUARAT PERTAMA  ",
				" Bil. 1 Isnin      18 Mei 2020  ",
				" K A N D U N G A N  ",
				" ISTIADAT PEMBUKAAN PENGGAL KETIGA  ",
				"  MAJLIS PARLIMEN YANG KEEMPAT BELAS   (Halaman       1)  ",
				" TITAH SERI PADUKA BAGINDA YANG DI-PERTUAN AGONG (Halaman       2)  ",
				" DR.18.5.2020 1  ",
				" MALAYSIA  ",
				" Isnin, 18 Mei 2020  ",
				" Mesyuarat dimulakan pada pukul 10.12 pagi  ",
				" DOA  ",
				" [Tuan Yang di-Pertua  mempengerusikan Mesyuarat]  ",
				" DR.18.5.2020 2  ",
				" TITAH KEBAWAH DULI YANG MAHA MULIA  ",
				" SERI PADUKA BAGINDA YANG DI-PERTUAN AGONG XVI  ",
				" AL-SULTAN ABDULLAH RI’AYATUDDIN�  ",
			}},
			SectionMarkers{
				DatePageMarker: "DR.18.5.2020",
				ParliamentDebate: ParliamentDebate{
					ID:        "bil-1-isnin-18-mei-2020",
					Date:      "Isnin, 18 Mei 2020",
					Attended:  nil,
					Missed:    nil,
					QAHansard: akomantoso.QAHansard{},
				},
				SessionStartMarkerLine: 17,
			},
		},
		{"happy #3", args{
			[]string{
				"  ",
				" PARLIMEN KEEMPAT BELAS  ",
				" PENGGAL KETIGA  ",
				" MESYUARAT KEDUA  ",
				" Bil. 2 Isnin        13 Julai 2020  ",
				"  - Waktu Mesyuarat dan Urusan Dibebaskan Daripada   ",
				"    Peraturan Mesyuarat (Halaman  130)  ",
				" DR. 13.7.2020 1  ",
				" Isnin, 13 Julai 2020  ",
				" Mesyuarat dimulakan pada pukul 10.00 pagi  ",
				" [Tuan Yang di-Pertua mempengerusikan Mesyuarat]  ",
				" ____________________________________________________________________  ",
				" PEMASYHURAN DARIPADA TUAN YANG DI-PERTUA  ",
				" [Setiausaha membacakan Perutusan]  ",
				" “19 Disember 2019�  ",
				" 8 NOVEMBER 2019 (JUMAAT) ",
				" YANG DI-PERTUA DEWAN NEGARA” �  ",
				" PELANTIKAN KETUA PEMBANGKANG  ",
				" Tuan Yang di-Pertua: Ahli-ahli Yang Berhormat, mengikut Peraturan Mesyuarat  ",
				" 4A(3), saya suka hendak memaklumkan iaitu mengikut maklum balas yang saya terima",
			}},
			SectionMarkers{
				DatePageMarker: "DR.13.7.2020",
				ParliamentDebate: ParliamentDebate{
					ID:        "bil-2-isnin-13-julai-2020",
					Date:      "Isnin, 13 Julai 2020",
					Attended:  nil,
					Missed:    nil,
					QAHansard: akomantoso.QAHansard{},
				},
				SessionStartMarkerLine: 10,
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

func Test_hasDatePageMarker(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasDatePageMarker(tt.args.line); got != tt.want {
				t.Errorf("hasDatePageMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasSessionStartMarkerLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSessionStartMarkerLine(tt.args.line); got != tt.want {
				t.Errorf("hasSessionStartMarkerLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
