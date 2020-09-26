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
		//{"case #1a", fields{"../../raw/StateAssembly/Hansard/HANSARD-16-JULAI-2020.pdf"}, nil,
		//	[]akomantoso.Representative{},
		//},
		{"case #1b", fields{"../../raw/StateAssembly/Hansard/HANSARD-15-JULAI-2020.pdf"}, nil,
			[]akomantoso.Representative{},
		},
		//{"case #1c", fields{"../../raw/StateAssembly/Hansard/HANSARD-14-JULAI-2020.pdf"}, nil,
		//	[]akomantoso.Representative{},
		//},
		//{"case #1d", fields{"../../raw/StateAssembly/Hansard/HANSARD-13-JULAI-2020-1.pdf"}, nil,
		//	[]akomantoso.Representative{},
		//},
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

func TestDebateProcessPages(t *testing.T) {
	type args struct {
		pdfDocument *akomantoso.PDFDocument
		dps         DebateProcessorState
	}
	// Samples of RepMap for 16, 15 respectively
	repMap16 := map[string]akomantoso.RepresentativeID{
		"TUAN SPEAKER":                                "tuan-speaker",
		"YB TUAN HEE LOY SIAN":                        "yb-tuan-hee-loy-sian",
		"YB DATO DR AHMAD YUNUS BIN HAIRI":            "yb-dato-dr-ahmad-yunus-bin-hairi",
		"YB TUAN HARUMAINI BIN HAJI OMAR":             "yb-tuan-harumaini-bin-haji-omar",
		"YB DATO MOHD IMRAN BIN TAMRIN":               "yb-dato-mohd-imran-bin-tamrin",
		"YB PUAN HANIZA BINTI MOHAMED TALHA":          "yb-puan-haniza-binti-mohamed-talha",
		"YB TUAN IR IZHAM BIN HASHIM":                 "yb-tuan-ir-izham-bin-hashim",
		"YB PUAN MICHELLE NG MEI SZE":                 "yb-puan-michelle-ng-mei-sze",
		"YB TUAN HASNUL BIN BAHARUDDIN":               "yb-tuan-hasnul-bin-baharuddin",
		"YB PUAN ROZANA BINTI ZAINAL ABIDIN":          "yb-puan-rozana-binti-zainal-abidin",
		"YB TUAN EDRY FAIZAL BIN EDDY YUSOF":          "yb-tuan-edry-faizal-bin-eddy-yusof",
		"YB TUAN LAI WAI CHONG":                       "yb-tuan-lai-wai-chong",
		"YB TUAN MAZWAN BIN JOHARI":                   "yb-tuan-mazwan-bin-johari",
		"YB TUAN MOHD SHAID BIN ROSLI":                "yb-tuan-mohd-shaid-bin-rosli",
		"YB TUAN BORHAN BIN AMAN SHAH":                "yb-tuan-borhan-bin-aman-shah",
		"YB PUAN DR DAROYAH BINTI ALWI":               "yb-puan-dr-daroyah-binti-alwi",
		"TUAN TIMBALAN SPEAKER":                       "tuan-timbalan-speaker",
		"YB TUAN KHAIRUDDIN BIN OTHMAN":               "yb-tuan-khairuddin-bin-othman",
		"YB DATUK ABDUL RASHID BIN ASARI":             "yb-datuk-abdul-rashid-bin-asari",
		"YB PUAN WONG SIEW KI":                        "yb-puan-wong-siew-ki",
		"YB TUAN MOHD FAKHRULRAZI BIN MOHD MOKHTAR":   "yb-tuan-mohd-fakhrulrazi-bin-mohd-mokhtar",
		"YB TUAN CHUA WEI KIAT":                       "yb-tuan-chua-wei-kiat",
		"YB TUAN LAU WENG SAN":                        "yb-tuan-lau-weng-san",
		"YB TUAN NG SZE HAN":                          "yb-tuan-ng-sze-han",
		"YB TUAN DR IDRIS BIN AHMAD":                  "yb-tuan-dr-idris-bin-ahmad",
		"YB TUAN LEONG TUCK CHEE":                     "yb-tuan-leong-tuck-chee",
		"YB DATO TENG CHANG KHIM":                     "yb-dato-teng-chang-khim",
		"YB TUAN AZMIZAN BIN ZAMAN HURI":              "yb-tuan-azmizan-bin-zaman-huri",
		"YB TUAN GANARAJAH A/L R GEORGE":              "yb-tuan-gunarajah-a/l-r-george",
		"YB TUAN MOHD NAJWAN BIN HALIMI":              "yb-tuan-mohd-najwan-bin-halimi",
		"YB TUAN ZAKARIA BIN HAJI HANAFI":             "yb-tuan-zakaria-bin-haji-hanafi",
		"YB TUAN GUNARAJAH A/L GEORGE":                "yb-tuan-gunarajah-a/l-r-george",
		"YB TUAN RIZAM BIN ISMAIL":                    "yb-tuan-rizam-bin-ismail",
		"YB TUAN RAJIV A/L RISHYAKARAN":               "yb-tuan-rajiv-a/l-rishyakaran",
		"YB TUAN SAARI BIN SUNGIB":                    "yb-tuan-saari-bin-sungib",
		"(TUAN SPEAKER MEMPENGERUSIKAN) TUAN SPEAKER": "(tuan-speaker-mempengerusikan)-tuan-speaker",
		"YB TUAN HAJI SAARI BIN SUNGIB":               "yb-tuan-haji-saari-bin-sungib",
		"YB TUAN MOHD SANY BIN HAMZAN":                "yb-tuan-mohd-sany-bin-hamzan",
		"YB MOHD NAJWAN BIN HALIMI":                   "yb-mohd-najwan-bin-halimi",
		"YB TUAN MOHD NAJWAN BIN HALIM":               "yb-tuan-mohd-najwan-bin-halim",
		"YB TUAN HALIMEY BIN ABU BAKAR":               "yb-tuan-halimey-bin-abu-bakar",
		"YB TUAN AZMIZAM BIN ZAMAN HURI":              "yb-tuan-azmizam-bin-zaman-huri",
		"YB TUAN PUAN JAMALIAH BINTI JAMALUDDIN":      "yb-puan-jamaliah-binti-jamaluddin",
		"YB PUAN LIM YI WEI":                          "yb-puan-lim-yi-wei",
		"YB PUAN JAMALIAH BINTI JAMALUDDIN":           "yb-puan-jamaliah-binti-jamaluddin",
		"YAB DATO MENTERI BESAR":                      "yab-dato-menteri-besar",
		"YB PUAN ELIZABETH WONG KEAT PING":            "yb-puan-elizabeth-wong-keat-ping",
		"YB TUAN SYAMSUL FIRDAUS BIN MOHAMED SUPRI":   "yb-tuan-syamsul-firdaus-bin-mohamed-supri",
		"YB TUAN DATO MOHD IMRAN BIN TAMRIN":          "yb-tuan-dato-mohd-imran-bin-tamrin",
		"YAB DATO MENTERI BESAR SELANGOR":             "yab-dato-menteri-besar-selangor",
		"YB TUAN MOHD FAKRULRAZI BIN MOHD MOKHTAR":    "yb-tuan-mohd-fakrulrazi-bin-mohd-mokhtar",
	}
	repMap15 := map[string]akomantoso.RepresentativeID{
		"TUAN SPEAKER":                              "tuan-speaker",
		"YB TUAN HAJI SAARI BIN SUNGIB":             "yb-tuan-haji-saari-bin-sungib",
		"YB PUAN DR SITI MARIAH BT MAHMUD":          "yb-puan-dr-siti-mariah-bt-mahmud",
		"YB DATO SERI MOHAMED AZMIN BIN ALI":        "yb-dato-seri-mohamed-azmin-bin-ali",
		"YB TUAN ZAKARIA BIN HANAFI":                "yb-tuan-zakaria-bin-hanafi",
		"YB TUAN GUNARAJAH A/L R GEORGE":            "yb-tuan-gunarajah-a/l-r-george",
		"YB PUAN DR SITI MARIAH BINTI MAHMUD":       "yb-puan-dr-siti-mariah-binti-mahmud",
		"YB DATO DR AHMAD YUNUS BIN HAIRI":          "yb-dato-dr-ahmad-yunus-bin-hairi",
		"YB TUAN LEONG TUCK CHEE":                   "yb-tuan-leong-tuck-chee",
		"YB TUAN LAU WENG SAN":                      "yb-tuan-lau-weng-san",
		"YB TUAN RIZAM BIN ISMAIL":                  "yb-tuan-rizam-bin-ismail",
		"YB PUAN RODZIAH BINTI ISMAIL":              "yb-puan-rodziah-binti-ismail",
		"YB TUAN MOHD FAKHRULRAZI BIN MOHD MOKHTAR": "yb-tuan-mohd-fakhrulrazi-bin-mohd-mokhtar",
		"YB TUAN LAI WAI CHONG":                     "yb-tuan-lai-wai-chong",
		"YAB DATO MENTERI BESAR":                    "yab-dato-menteri-besar",
		"YB DATUK ROSNI BT SOHAR":                   "yb-datuk-rosni-bt-sohar",
		"YB TUAN HASNUL BIN BAHARUDDIN":             "yb-tuan-hasnul-bin-baharuddin",
		"YB TUAN IR IZHAM BIN HASHIM":               "yb-tuan-ir-izham-bin-hashim",
		"YB TUAN RAJIV A/L RISHYAKARAN":             "yb-tuan-rajiv-a/l-rishyakaran",
		"YB HARUMAINI BIN HAJI OMAR":                "yb-harumaini-bin-haji-omar",
		"YB TUAN MOHD ZAWAWI BIN AHMAD MUGHNI":      "yb-tuan-mohd-zawawi-bin-ahmad-mughni",
		"YB TUAN SHATIRI BIN MANSOR":                "yb-tuan-shatiri-bin-mansor",
		"YB DATO TENG CHANG KHIM":                   "yb-dato-teng-chang-khim",
		"YB DATO; DR AHMAD YUNUS BIN HAIRI":         "yb-dato;-dr-ahmad-yunus-bin-hairi",
		"YB TUAN SYAMSUL FIRDAUS BIN MOHAMED SUPRI": "yb-tuan-syamsul-firdaus-bin-mohamed-supri",
		"YB TUAN SAARI BIN SUNGIB":                  "yb-tuan-saari-bin-sungib",
		"YB TUAN HALIMEY BIN ABU BAKAR":             "yb-tuan-halimey-bin-abu-bakar",
		"TUAN TIMBALAN SPEAKER":                     "tuan-timbalan-speaker",
		"YB PUAN MICHELLE NG MEI SZE":               "yb-puan-michelle-ng-mei-sze",
		"YB DATUK ROSNI BIN SOHAR":                  "yb-datuk-rosni-bin-sohar",
		"YB PUAN ELIZABETH WONG KEAT PING":          "yb-puan-elizabeth-wong-keat-ping",
		"YB PUAN LIM YI WEI":                        "yb-puan-lim-yi-wei",
		"YB TUAN NG SZE HAN":                        "yb-tuan-ng-sze-han",
		"YB TUAN MOHD SHAID BIN ROSLI":              "yb-tuan-mohd-shaid-bin-rosli",
		"YB TUAN MUHAMMAD HILMAN BIN IDHAM":         "yb-tuan-muhammad-hilman-bin-idham",
		"YB TUAN HILMAN BIN IDHAM":                  "yb-tuan-hilman-bin-idham",
		"YB TUAN MOHD KHAIRUDDIN BIN OTHMAN":        "yb-tuan-mohd-khairuddin-bin-othman",
		"YB TUAN MOHD NAJWAN BIN HALIMI":            "yb-tuan-mohd-najwan-bin-halimi",
		"YB TUAN SALLEHUDIN BIN AMIRUDDIN":          "yb-tuan-sallehudin-bin-amiruddin",
		"YB TUAN HEE LOY SIAN":                      "yb-tuan-hee-loy-sian",
		"YB DATO MOHD IMRAN BIN TAMRIN":             "yb-dato-mohd-imran-bin-tamrin",
		"YB PUAN DR DAROYAH BT ALWI":                "yb-puan-dr-daroyah-bt-alwi",
		"YB TUAN MOHD SANY BIN HAMZAN":              "yb-tuan-mohd-sany-bin-hamzan",
		"YB DATO DR AHAMD YUNUS BIN HAIRI":          "yb-dato-dr-ahamd-yunus-bin-hairi",
		"YB TUAN EDRY FAIZAL BIN EDDY YUSOF":        "yb-tuan-edry-faizal-bin-eddy-yusof",
		"YB PUAN ROZANA BINTI ZAINAL ABIDIN":        "yb-puan-rozana-binti-zainal-abidin",
	}
	// Setup different PDFs and DPSs
	// Extract out Section Metadata for attachment
	extractOptions := akomantoso.ExtractPDFOptions{
		StartPage:  7,
		NumPages:   5,
		MaxSampled: 10000,
	}
	pdfDocument, perr := akomantoso.NewPDFDocument("../../raw/StateAssembly/Hansard/HANSARD-16-JULAI-2020.pdf", &extractOptions)
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

	dps := DebateProcessorState{
		SectionMarkers: SectionMarkers{
			DatePageMarker:         "16 JULAI 2020 (KHAMIS)",
			SessionStartMarkerLine: 7,
		},
		RepresentativesMap: repMap16,
	}

	// For 15th test case
	pdfDocument15, perr := akomantoso.NewPDFDocument("../../raw/StateAssembly/Hansard/HANSARD-15-JULAI-2020.pdf", &extractOptions)
	if perr != nil {
		panic(perr)
	}
	//spew.Dump(pdfDocument.Pages)
	// Sanity  checks ..
	if len(pdfDocument15.Pages) < 1 {
		// DEBUG
		//spew.Dump(pdfDocument.Pages)
		panic("Should NOT be here!!")
	}

	dps15 := DebateProcessorState{
		SectionMarkers: SectionMarkers{
			DatePageMarker:         "15 JULAI 2020 (RABU)",
			SessionStartMarkerLine: 7,
		},
		RepresentativesMap: repMap15,
	}

	tests := []struct {
		name string
		args args
		want StateAssemblyDebateContent
	}{
		{"case #1", args{pdfDocument, dps}, StateAssemblyDebateContent{}},
		{"case #2", args{pdfDocument15, dps15}, StateAssemblyDebateContent{}},
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
		allLines       []string
		pdfPath        string
		dps            DebateProcessorState
		currentPage    int
		currentContent []string
	}
	// Samples of RepMap for 14, 13 respectively
	repMap14 := map[string]akomantoso.RepresentativeID{
		"TUAN SPEAKER":                              "tuan-speaker",
		"YB TUAN IR IZHAM BIN HASHIM":               "yb-tuan-ir-izham-bin-hashim",
		"YB DATO DR AHMAD YUNUS BIN HAIRI":          "yb-dato-dr-ahmad-yunus-bin-hairi",
		"YB TUAN LAU WENG SAN":                      "yb-tuan-lau-weng-san",
		"YB PUAN MICHELLE NG MEI SZE":               "yb-puan-michelle-ng-mei-sze",
		"YB DATO TENG CHANG KHIM":                   "yb-dato-teng-chang-khim",
		"YB TUAN MOHD NAJWAN BIN HALMI":             "yb-tuan-mohd-najwan-bin-halmi",
		"YB TUAN KHAIRUDDIN BIN OTHMAN":             "yb-tuan-khairuddin-bin-othman",
		"YB TUAN HALIMEY BIN ABU BAKAR":             "yb-tuan-halimey-bin-abu-bakar",
		"YB TUAN FAKHRULRAZI BIN MOHD MOKHTAR":      "yb-tuan-fakhrulrazi-bin-mohd-mokhtar",
		"YB PUAN JAMALIAH BINTI JAMALUDDIN":         "yb-puan-jamaliah-binti-jamaluddin",
		"YB TUAN MOHD FAKHRULRAZI BIN MOHD MOKHTAR": "yb-tuan-mohd-fakhrulrazi-bin-mohd-mokhtar",
		"YB TUAN HARUMAINI BIN HAJI OMAR":           "yb-tuan-harumaini-bin-haji-omar",
		"YB TUAN MOHD ZAWAWI BIN AHMAD MUGHNI":      "yb-tuan-mohd-zawawi-bin-ahmad-mughni",
		"YB PUAN RODZIAH BINTI ISMAIL":              "yb-puan-rodziah-binti-ismail",
		"YB TUAN HARUMAINI BIN HJ OMAR":             "yb-tuan-harumaini-bin-hj-omar",
		"YB TUAN HAJI BORHAN BIN AMAN SHAH":         "yb-tuan-haji-borhan-bin-aman-shah",
		"YB TUAN HEE LOY SIAN":                      "yb-tuan-hee-loy-sian",
		"YB TUAN EAN YONG HIAN WAH":                 "yb-tuan-ean-yong-hian-wah",
		"YB TUAN LAI WAI CHONG":                     "yb-tuan-lai-wai-chong",
		"YB TUAN MOHD SANY BIN HAMZAN":              "yb-tuan-mohd-sany-bin-hamzan",
		"YAB DATO MENTERI BESAR":                    "yab-dato-menteri-besar",
		"YB TUAN RIZAM BIN ISMAIL":                  "yb-tuan-rizam-bin-ismail",
		"YB TUAN CHUA WEI KIAT":                     "yb-tuan-chua-wei-kiat",
		"YB TUAN EDRY FAIZAL BIN EDDY YUSOF":        "yb-tuan-edry-faizal-bin-eddy-yusof",
		"YB TUAN DR IDRIS BIN AHMAD":                "yb-tuan-dr-idris-bin-ahmad",
		"YB TUAN LEONG TUCK CHEE":                   "yb-tuan-leong-tuck-chee",
		// Is female! Mapped correctly
		"YB TUAN MICHELLE NG MEI SZE":        "yb-puan-michelle-ng-mei-sze",
		"YB TUAN NG SZE HAN":                 "yb-tuan-ng-sze-han",
		"YB PUAN HANIZA BINTI MOHAMED TALHA": "yb-puan-haniza-binti-mohamed-talha",
		"YB TUAN ADHIF SYAN BIN ABDULLAH":    "yb-tuan-adhif-syan-bin-abdullah",
		// Correct female mapping; Binti!
		"YB PUAN HANIZA BIN MOHAMED TALHA":          "yb-puan-haniza-binti-mohamed-talha",
		"TUAN TIMBALAN SPEAKER":                     "tuan-timbalan-speaker",
		"YB PUAN LIM YI WEI":                        "yb-puan-lim-yi-wei",
		"YB TUAN SYAMSUL FIRDAUS BIN MOHAMED SUPRI": "yb-tuan-syamsul-firdaus-bin-mohamed-supri",
		"YB TUAN SALLEHUDIN BIN AMIRUDDIN":          "yb-tuan-sallehudin-bin-amiruddin",
		"YB PUAN WONG SIEW KI":                      "yb-puan-wong-siew-ki",
		"YB DATO TENG CHANG KIM":                    "yb-dato-teng-chang-kim",
		"YB TUAN RAJIV A/L RISHYAKARAN":             "yb-tuan-rajiv-a/l-rishyakaran",
		"YB TUAN HAJI SAARI BIN SUNGIB":             "yb-tuan-haji-saari-bin-sungib",
		"YB PUAN ELIZABETH WONG KEAT PING":          "yb-puan-elizabeth-wong-keat-ping",
		"YB DATO MOHD IMRAN BIN TAMRIN":             "yb-dato-mohd-imran-bin-tamrin",
		"YB TUAN MOHD NAJWAN BIN HALIMI":            "yb-tuan-mohd-najwan-bin-halimi",
		"YB TUAN MOHD SHAID BIN ROSLI":              "yb-tuan-mohd-shaid-bin-rosli",
		"YB PUAN DR DAROYAH BINTI ALWI":             "yb-puan-dr-daroyah-binti-alwi",
		"YB PUAN DR DAROYAH BT ALWI":                "yb-puan-dr-daroyah-binti-alwi",
		"YB TUAN GUNARAJAH A/L R GEORGE":            "yb-tuan-gunarajah-a/l-r-george",
		"YB DATO MOHD SHAMSUDIN BIN LIAS":           "yb-dato-mohd-shamsudin-bin-lias",
	}
	// NOTE: The typos need to be mapped and fixed
	repMap13 := map[string]akomantoso.RepresentativeID{
		"TUAN SPEAKER": "tuan-speaker",
		// Below gets removed
		//"PENCADANG (YAB DATO MENTERI BESAR)":        "pencadang-(yab-dato-menteri-besar)",
		//"PENYOKONG (YB DATO TENG CHANG KHIM)":       "penyokong-(yb-dato-teng-chang-khim)",
		// Eli Wong is female!
		"YB TUAN ELIZABETH WONG KEAT PING":    "yb-puan-elizabeth-wong-keat-ping",
		"YB TUAN LAI WAI CHONG":               "yb-tuan-lai-wai-chong",
		"YB PUAN DR SITI MARIAH BINTI MAHMUD": "yb-puan-dr-siti-mariah-binti-mahmud",
		"YB DATO DR AHMAD YUNUS BIN HAIRI":    "yb-dato-dr-ahmad-yunus-bin-hairi",
		"YB TUAN MOHD SANY BIN HAMZAN":        "yb-tuan-mohd-sany-bin-hamzan",
		"YB PUAN JAMALIAH BINTI JAMALUDDIN":   "yb-puan-jamaliah-binti-jamaluddin",
		"YB TUAN GANABATIRAU A/L VERAMAN":     "yb-tuan-ganabatirau-a/l-veraman",
		"YB TUAN HAJI SAARI BIN SUNGIB":       "yb-tuan-haji-saari-bin-sungib",
		"YB PUAN DR DAROYAH BINTI ALWI":       "yb-puan-dr-daroyah-binti-alwi",
		"YAB DATO MENTERI BESAR":              "yab-dato-menteri-besar",
		"YB TUAN RIZAM BIN ISMAIL":            "yb-tuan-rizam-bin-ismail",
		"YB PUAN RODZIAH BINTI ISMAIL":        "yb-puan-rodziah-binti-ismail",
		"YB DATO MOHD IMRAN BIN TAMRIN":       "yb-dato-mohd-imran-bin-tamrin",
		"YB TUAN HEE LOY SIAN":                "yb-tuan-hee-loy-sian",
		"YB TUAN DR IDRIS BIN AHMAD":          "yb-tuan-dr-idris-bin-ahmad",
		"YB TUAN ADHIF SYAN BIN ABDULLAH":     "yb-tuan-adhif-syan-bin-abdullah",
		"YB TUAN NG SZE HAN":                  "yb-tuan-ng-sze-han",
		"YB TUAN SAARI BIN SUNGIB":            "yb-tuan-saari-bin-sungib",
		"YB PUAN JUWAIRIYA BINTI ZULKIFLI":    "yb-puan-juwairiya-binti-zulkifli",
		"YB TUAN IR IZHAM BIN HASHIM":         "yb-tuan-ir-izham-bin-hashim",
		"YB TUAN HARUMAINI BIN HAJI OMAR":     "yb-tuan-harumaini-bin-haji-omar",
		"YB PUAN LIM YI WEI":                  "yb-puan-lim-yi-wei",
		"YB DATUK ABDUL RASHID BIN ASARI":     "yb-datuk-abdul-rashid-bin-asari",
		"YB TUAN HASNUL BIN BAHARUDDIN":       "yb-tuan-hasnul-bin-baharuddin",
		"YB TUAN LAU WENG SAN":                "yb-tuan-lau-weng-san",
		// Typoe fixed and mapped correctly
		"TUAN SEPAKER":                              "tuan-speaker",
		"YB DATO MOHD SHAMSUDIN BIN LIAS":           "yb-dato-mohd-shamsudin-bin-lias",
		"YB DATO TENG CHANG KHIM":                   "yb-dato-teng-chang-khim",
		"YB DATUK ROSNI BINTI SOHAR":                "yb-datuk-rosni-binti-sohar",
		"YB TUAN SYAMSUL FIRDAUS BIN MOHAMED SUPRI": "yb-tuan-syamsul-firdaus-bin-mohamed-supri",
		"YB TUAN SALLEHUDIN BIN AMIRUDDIN":          "yb-tuan-sallehudin-bin-amiruddin",
		"YB PUAN ELIZABETH WONG KEAT PING":          "yb-puan-elizabeth-wong-keat-ping",
		"YB TUAN MOHD NAJWAN BIN HALIMI":            "yb-tuan-mohd-najwan-bin-halimi",
		"TUAN TIMBALAN SPEAKER":                     "tuan-timbalan-speaker",
		"YB TUAN RAJIV A/L RISHYAKARAN":             "yb-tuan-rajiv-a/l-rishyakaran",
		"YB SALLEHUDDIN BIN AMIRUDDIN":              "yb-sallehuddin-bin-amiruddin",
		"YB TUAN MOHD FAKHRULRAZI BIN MOHD MOKHTAR": "yb-tuan-mohd-fakhrulrazi-bin-mohd-mokhtar",
		"YB TUAN MOHD SHAID BIN ROSLI":              "yb-tuan-mohd-shaid-bin-rosli",
		"YB TUAN LEONG TUCK CHEE":                   "yb-tuan-leong-tuck-chee",
		"YB DATO TENG CHANG KIM":                    "yb-dato-teng-chang-kim",
	}
	// Setup different PDFs and DPSs
	dps := DebateProcessorState{
		SectionMarkers: SectionMarkers{
			DatePageMarker: "13 JULAI 2020 (ISNIN)",
		},
		RepresentativesMap: repMap13,
	}
	// For test case 14th
	dps14 := DebateProcessorState{
		SectionMarkers: SectionMarkers{
			DatePageMarker: "14 JULAI 2020 (SELASA)",
		},
		RepresentativesMap: repMap14,
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case #1", args{
			pdfPath: "../../raw/StateAssembly/Hansard/HANSARD-13-JULAI-2020-1.pdf",
			allLines: []string{
				"13 JULAI 2020 (ISNIN) ",
				"15 ",
				"penyelia-penyelia PWB dan juga aktivis-aktivis masyarakat supaya mereka lebih peka ",
				"terhadap keganasan rumah tangga dan mereka tahu apakah yang patut dilakukan tapi ",
				"Y.B. DATO’ DR AHMAD YUNUS BIN HAIRI� : Soalan tambahan. ",
				"TUAN SPEAKER : Sijangkang. ",
				"Y.B. PUAN DR SITI MARIAH BINTI MAHMUD  : Terima kasih Yang ",
				"Berhormat Sijangkang.  Kes ini saya tidak dapat figure yang sebelumnya ya, sebelum ", "Y.B. TUAN MOHD SANY BIN HAMZAN : Tuan Speaker. ",
				"TUAN SPEAKER  : Taman Templer. ",
			},
			dps:         dps,
			currentPage: 15,
			//currentContent: []string{},
		}, true},
		{"case #2", args{
			pdfPath: "../../raw/StateAssembly/Hansard/HANSARD-14-JULAI-2020.pdf",
			allLines: []string{
				"14 JULAI 2020 (SELASA) ",
				"10 ",
				"Y.B. TUAN IR IZHAM BIN HASHIM : Terima kasih Yang Berhormat Sijangkang.  ",
				"selesaikan.  Begitu juga yang disebut oleh Yang Berhormat, Yang Berhormat tadi, kita ",
				"delay pun, seminima yang mungkin.  Itu yang saya boleh bagi jaminan di sini.  Terima ",
				"Y.B. TUAN LAU WENG SAN : Yang Berhormat, soalan tambahan. ",
				"TUAN SPEAKER : Banting. ",
				"Y.B. TUAN LAU WENG SAN : Terima kasih.  Saya ingin bertanya kepada Exco, ",
				"Kecil dan Besar itu.  Sebenarnya masalah dia adalah di Sungai kecil dan besar, yang ",
			},
			dps:         dps14,
			currentPage: 10,
			//currentContent: []string{},
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Extract out Section Metadata for attachment
			extractOptions := akomantoso.ExtractPDFOptions{
				StartPage:  tt.args.currentPage,
				NumPages:   1,
				MaxSampled: 10000,
			}
			pdfDocument, perr := akomantoso.NewPDFDocument(tt.args.pdfPath, &extractOptions)
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
			// Copy the structure
			currentDPS := tt.args.dps
			// Setup CurrentPage content and take only one page (or more??)
			// and any contentp prelude ..
			if err := DebateProcessSinglePage(pdfDocument.Pages[0].PDFTxtSameLines, &currentDPS); (err != nil) != tt.wantErr {
				t.Errorf("DebateProcessSinglePage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
