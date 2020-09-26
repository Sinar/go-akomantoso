package main

import (
	"fmt"
	"strings"
	"unicode"

	state_assembly "github.com/Sinar/go-akomantoso/internal/state-assembly"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
)

type SayItCmd struct {
	ID            int    `flag:"-"`
	DebateRawFile string `help:"Where is raw?" flag:"source"`
	DebateType    string `help:"Debate Type? dun,par?"`
	OutputFile    string `help:"Where store output? Default prefix source." flag:"output"`
}

func NewSayItCmd(conf Config) *SayItCmd { return &SayItCmd{DebateType: "dun"} }

func (m *SayItCmd) Run() error {
	if m.DebateRawFile == "" {
		return fmt.Errorf("Select filename plz!!")
	}
	fmt.Println("SAYIT! ..")
	repMap15 := loadRepresentativeMapping("15-julai-2020")
	dps15 := state_assembly.DebateProcessorState{
		SectionMarkers: state_assembly.SectionMarkers{
			DatePageMarker:         "SayIT Test: 15 JULAI 2020",
			SessionStartMarkerLine: 7,
		},
		RepresentativesMap: repMap15,
	}
	// split out the name from file
	//dir, fileName := filepath.Split(m.DebateRawFile)
	// Copy the structure
	currentDPS := dps15
	// For 15th test case
	// Setup different PDFs and DPSs
	// Extract out Section Metadata for attachment
	extractOptions := akomantoso.ExtractPDFOptions{
		StartPage: dps15.SectionMarkers.SessionStartMarkerLine,
		//NumPages:   5,
		MaxSampled: 10000,
	}
	pdfDocument15, perr := akomantoso.NewPDFDocument(m.DebateRawFile, &extractOptions)
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
	currentDPS.CurrentContents = state_assembly.DebateProcessPages(pdfDocument15, currentDPS)
	//saDebateContent := state_assembly.DebateProcessPages(pdfDocument15, currentDPS)
	// DEBUG
	//spew.Dump(saDebateContent)
	// Output it to yaml; here; depending on options?
	outputAsSayItFormat(currentDPS, m.OutputFile)
	return nil
}

func loadRepresentativeMapping(label string) map[string]akomantoso.RepresentativeID {
	switch label {
	case "15-julai-2020":
		return map[string]akomantoso.RepresentativeID{
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
	case "16-julai-2020":
		return map[string]akomantoso.RepresentativeID{
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
	default:
		return map[string]akomantoso.RepresentativeID{}
	}
}
func outputAsSayItFormat(currentDPS state_assembly.DebateProcessorState, output string) error {
	// Output templates look like ..
	var allRepMetadata string
	var allParagraphs string
	// Loop through each Representative -- Example ..
	//<TLCPerson id="tuan-speaker" showAs="TUAN SPEAKER"/>
	//<TLCPerson id="yb-dato-seri-mohamed-azmin-bin-ali" showAs="YB DATO SERI MOHAMED AZMIN BIN ALI"/>
	//<TLCPerson id="yb-tuan-haji-saari-bin-sungib" showAs="YB TUAN HAJI SAARI BIN SUNGIB"/>
	//<TLCPerson id="yb-puan-dr-siti-mariah-bt-mahmud" showAs="YB PUAN DR SITI MARIAH BT MAHMUD"/>
	for singleRepDisplay, singleRepID := range currentDPS.RepresentativesMap {
		allRepMetadata += fmt.Sprintf("<TLCPerson href=\"/ontology/person/staging-dundocs.sayit.mysociety.org/%s\" id=\"%s\" showAs=\"%s\" />\n", singleRepID, singleRepID, singleRepDisplay)
	}
	// Loop through each paragraph Content
	for _, singleParagraph := range currentDPS.CurrentContents {
		// DEBUG
		//spew.Dump(singleParagraph)
		bodyOutput := fmt.Sprintf("<p>\n%s</p>\n", removeNonASCII(strings.ReplaceAll(singleParagraph.RawContent, "\n", "<br/>")))
		fileOutput := fmt.Sprintf("<speech by=\"#%s\">\n%s</speech>\n", singleParagraph.RepresentativeID, bodyOutput)
		allParagraphs += fileOutput
	}
	outputComplete(allRepMetadata, currentDPS.SectionMarkers.DatePageMarker, allParagraphs)
	return nil
}

func outputComplete(allRepMetadata string, headingSession string, allParagraphs string) {
	outputHeader := fmt.Sprintf("%s%s%s%s%s", `
<akomaNtoso>
    <debate>
        <meta>
            <references>
`, allRepMetadata,
		`
            </references>
        </meta>
        <debateBody>
            <debateSection>
                <heading>
`, headingSession, "</heading>")

	outputFooter := `
            </debateSection>
        </debateBody>
    </debate>
</akomaNtoso>
`
	fmt.Println(fmt.Sprintf("%s\n%s\n%s", outputHeader, allParagraphs, outputFooter))
}

func removeNonASCII(line string) string {
	// https://programming-idioms.org/idiom/147/remove-all-non-ascii-characters/1848/go
	return strings.Map(func(r rune) rune {
		if r > unicode.MaxASCII {
			return -1
		}
		return r
	}, line)
}
