package parliament

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

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
		{"happy  #1", fields{"../../raw/Parliament/Hansard/DR-28072020.pdf"}, nil,
			[]akomantoso.Representative{
				akomantoso.Representative{},
			},
		},
		{"happy  #2", fields{"../../raw/Parliament/Hansard/DR-18052020.pdf"}, nil,
			[]akomantoso.Representative{
				akomantoso.Representative{},
			},
		},
		{"happy  #3", fields{"../../raw/Parliament/Hansard/DR-13072020 New 1.pdf"}, nil,
			[]akomantoso.Representative{
				akomantoso.Representative{},
			},
		},
		{"happy  #4", fields{"../../raw/Parliament/Hansard/DR-26082020.pdf"}, nil,
			[]akomantoso.Representative{
				akomantoso.Representative{},
			},
		},
		{"happy  #5", fields{"../../raw/Parliament/Hansard/DR-27082020.pdf"}, nil,
			[]akomantoso.Representative{
				akomantoso.Representative{},
			},
		},
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

func TestDebateProcessSinglePage(t *testing.T) {
	type args struct {
		allLines       []string
		pdfPath        string
		dps            DebateProcessorState
		currentPage    int
		currentContent []string
	}
	// Samples of RepMap
	repMap26082020 := map[string]akomantoso.RepresentativeID{
		"Datuk Hajah Azizah binti Mohd Dun [Beaufort]":                                  "datuk-hajah-azizah-binti-mohd-dun-beaufort",
		"Timbalan Menteri Pengajian Tinggi [Dato Mansor bin Othman]":                    "timbalan-menteri-pengajian-tinggi-dato-mansor-bin-othman",
		"Dato Hajah Azizah binti Mohd Dun [Beaufort]":                                   "dato-hajah-azizah-binti-mohd-dun-beaufort",
		"Dato Mansor bin Othman":                                                        "dato-mansor-bin-othman",
		"Tuan Hassan bin Abdul Karim [Pasir Gudang]":                                    "tuan-hassan-bin-abdul-karim-pasir-gudang",
		"Datuk Haji Hasanuddin bin Mohd Yunus [Hulu Langat]":                            "datuk-haji-hasanuddin-bin-mohd-yunus-hulu-langat",
		"Tuan Loke Siew Fook [Seremban]":                                                "tuan-loke-siew-fook-seremban",
		"Menteri Pengangkutan [Datuk Seri Ir Dr Wee Ka Siong]":                          "menteri-pengangkutan-datuk-seri-ir-dr-wee-ka-siong",
		"Datuk Seri Ir Dr Wee Ka Siong":                                                 "datuk-seri-ir-dr-wee-ka-siong",
		"Timbalan Yang di-Pertua [Dato Mohd Rashid Hasnon]":                             "timbalan-yang-di-pertua-dato-mohd-rashid-hasnon",
		"Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling]":                      "datuk-seri-panglima-abdul-azeez-bin-abdul-rahim-baling",
		"Tuan Ahmad Fadhli bin Shaari [Pasir Mas]":                                      "tuan-ahmad-fadhli-bin-shaari-pasir-mas",
		"Timbalan Menteri Kewangan I [Datuk Abd Rahim bin Bakri]":                       "timbalan-menteri-kewangan-i-datuk-abd-rahim-bin-bakri",
		"Tuan Oscar Ling Chai Yew [Sibu]":                                               "tuan-oscar-ling-chai-yew-sibu",
		"Datuk Abd Rahim bin Bakri":                                                     "datuk-abd-rahim-bin-bakri",
		"Datuk Seri Saifuddin Nasution bin Ismail [Kulim-Bandar Baharu]":                "datuk-seri-saifuddin-nasution-bin-ismail-kulim-bandar-baharu",
		"Tuan Abdul Latiff bin Abdul Rahman [Kuala Krai]":                               "tuan-abdul-latiff-bin-abdul-rahman-kuala-krai",
		"Tuan Nik Nazmi bin Nik Ahmad [Setiawangsa]":                                    "tuan-nik-nazmi-bin-nik-ahmad-setiawangsa",
		"Datuk Seri Dr Mujahid Yusof Rawa [Parit Buntar]":                               "datuk-seri-dr-mujahid-yusof-rawa-parit-buntar",
		"Datuk Seri Dr Mujahid bin Yusof Rawa [Parit Buntar]":                           "datuk-seri-dr-mujahid-bin-yusof-rawa-parit-buntar",
		"Dato Hajah Siti Zailah binti Mohd Yusoff":                                      "dato-hajah-siti-zailah-binti-mohd-yusoff",
		"Dato Sri Hajah Rohani binti Abdul Karim [Batang Lupar]":                        "dato-sri-hajah-rohani-binti-abdul-karim-batang-lupar",
		"Puan Nor Azrina binti Surip [Merbok]":                                          "puan-nor-azrina-binti-surip-merbok",
		"Tuan Ahmad Tarmizi bin Sulaiman [Sik]":                                         "tuan-ahmad-tarmizi-bin-sulaiman-sik",
		"Datuk Seri Haji Ahmad bin Hamzah":                                              "datuk-seri-haji-ahmad-bin-hamzah",
		"Puan Hannah Yeoh [Segambut]":                                                   "puan-hannah-yeoh-segambut",
		"Datuk Rozman bin Isli (Labuan) tidak hadir]":                                   "datuk-rozman-bin-isli-labuan-tidak-hadir",
		"Datuk Mohamad bin Alamin [Kimanis]":                                            "datuk-mohamad-bin-alamin-kimanis",
		"Timbalan Menteri Pelancongan, Seni dan Budaya [Datuk Dr Jeffrey Kitingan]":     "timbalan-menteri-pelancongan,-seni-dan-budaya-datuk-dr-jeffrey-kitingan",
		"Datuk Dr Jeffrey Kitingan":                                                     "datuk-dr-jeffrey-kitingan",
		"Tuan Chan Foong Hin [Kota Kinabalu]":                                           "tuan-chan-foong-hin-kota-kinabalu",
		"Dato Haji Salim Sharif [Jempol]":                                               "dato-haji-salim-sharif-jempol",
		"Tuan Gobind Singh Deo [Puchong]":                                               "tuan-gobind-singh-deo-puchong",
		"Timbalan Menteri Komunikasi dan Multimedia [Datuk Zahidi bin Zainul Abidin]":   "timbalan-menteri-komunikasi-dan-multimedia-datuk-zahidi-bin-zainul-abidin",
		"Datuk Zahidi bin Zainul Abidin":                                                "datuk-zahidi-bin-zainul-abidin",
		"Datuk Seri Dr Haji Dzulkefly bin Ahmad [Kuala Selangor]":                       "datuk-seri-dr-haji-dzulkefly-bin-ahmad-kuala-selangor",
		"Tuan Yamani Hafez bin Musa [Sipitang]":                                         "tuan-yamani-hafez-bin-musa-sipitang",
		"Tuan Haji Yamani Hafez bin Musa [Sipitang]":                                    "tuan-haji-yamani-hafez-bin-musa-sipitang",
		"Tuan Nga Kor Ming [Teluk Intan]":                                               "tuan-nga-kor-ming-teluk-intan",
		"Timbalan Menteri Pendidikan II [Tuan Muslimin bin Yahaya]":                     "timbalan-menteri-pendidikan-ii-tuan-muslimin-bin-yahaya",
		"Tuan Muslimin bin Yahaya":                                                      "tuan-muslimin-bin-yahaya",
		"Datuk Seri Haji Ahmad bin Haji Maslan [Pontian]":                               "datuk-seri-haji-ahmad-bin-haji-maslan-pontian",
		"Menteri Perumahan dan Kerajaan Tempatan [Puan Hajah Zuraida binti Kamaruddin]": "menteri-perumahan-dan-kerajaan-tempatan-puan-hajah-zuraida-binti-kamaruddin",
		"Dato Takiyuddin bin Hassan":                                                    "dato-takiyuddin-bin-hassan",
		"pg Menteri Pengangkutan [Datuk Seri Ir Dr Wee Ka Siong]":                       "menteri-pengangkutan-datuk-seri-ir-dr-wee-ka-siong",
		"Tuan Yang di-Pertua":                                                           "tuan-yang-di-pertua",
		"Menteri Pengajian Tinggi [Dato Dr Noraini Ahmad]":                              "menteri-pengajian-tinggi-dato-dr-noraini-ahmad",
		"tgh Tuan Loke Siew Fook [Seremban]":                                            "tuan-loke-siew-fook-seremban",
		"Dato Dr Xavier Jayakumar a/l Arulanandam [Kuala Langat]":                       "dato-dr-xavier-jayakumar-a/l-arulanandam-kuala-langat",
		"Tuan Teh Kok Lim [Taiping]":                                                    "tuan-teh-kok-lim-taiping",
		"tgh Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling]":                  "tgh-datuk-seri-panglima-abdul-azeez-bin-abdul-rahim-baling",
		"Tuan Sabri bin Azit [Jerai]":                                                   "tuan-sabri-bin-azit-jerai",
		"tgh Tuan Mohamed Hanipa bin Maidin [Sepang]":                                   "tgh-tuan-mohamed-hanipa-bin-maidin-sepang",
		"Tuan Mohamed Hanipa bin Maidin [Sepang]":                                       "tuan-mohamed-hanipa-bin-maidin-sepang",
		"Tuan Che Alias bin Hamid [Kemaman]":                                            "tuan-che-alias-bin-hamid-kemaman",
		"Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong]":                            "tuan-sanisvara-nethaji-rayer-a/l-rajaji-jelutong",
		"tgh Tuan Ahmad Fadhli bin Shaari [Pasir Mas]":                                  "tuan-ahmad-fadhli-bin-shaari-pasir-mas",
		"Tuan Khoo Poay Tiong [Kota Melaka]":                                            "tuan-khoo-poay-tiong-kota-melaka",
		"Tan Sri Noh bin Haji Omar [Tanjong Karang]":                                    "tan-sri-noh-bin-haji-omar-tanjong-karang",
		"tgh Tan Sri Noh bin Haji Omar [Tanjong Karang]":                                "tan-sri-noh-bin-haji-omar-tanjong-karang",
		"Datuk Wira Dr Mohd Hatta bin Md Ramli [Lumut]":                                 "datuk-wira-dr-mohd-hatta-bin-md-ramli-lumut",
		"Dato Sri Bung Moktar bin Radin [Kinabatangan]":                                 "dato-sri-bung-moktar-bin-radin-kinabatangan",
		"Timbalan Yang di-Pertua [Dato Sri Azalina Othman Said]":                        "timbalan-yang-di-pertua-dato-sri-azalina-othman-said",
		"Tuan Kesavan a/l Subramaniam [Sungai Siput]":                                   "tuan-kesavan-a/l-subramaniam-sungai-siput",
		"ptg Tuan Syed Ibrahim bin Syed Noh [Ledang]":                                   "tuan-syed-ibrahim-bin-syed-noh-ledang",
		"Tuan Syed Ibrahim bin Syed Noh [Ledang]":                                       "tuan-syed-ibrahim-bin-syed-noh-ledang",
		"ptg Tuan Wong Hon Wai [Bukit Bendera]":                                         "tuan-wong-hon-wai-bukit-bendera",
		"Tuan Wong Hon Wai [Bukit Bendera]":                                             "tuan-wong-hon-wai-bukit-bendera",
		"Dato Ngeh Koo Ham [Beruas]":                                                    "dato-ngeh-koo-ham-beruas",
		"ptg Tuan Hassan bin Abdul Karim [Pasir Gudang]":                                "tuan-hassan-bin-abdul-karim-pasir-gudang",
		"ptg Tuan Shaharizukirnain bin Abd Kadir [Setiu]":                               "tuan-shaharizukirnain-bin-abd-kadir-setiu",
		"ptg Dato Ngeh Koo Ham [Beruas]":                                                "dato-ngeh-koo-ham-beruas",
		"ptg Dato Seri Dr Shahidan bin Kassim [Arau]":                                   "dato-seri-dr-shahidan-bin-kassim-arau",
		"Dato Seri Tiong King Sing [Bintulu]":                                           "dato-seri-tiong-king-sing-bintulu",
		"Dato Jalaluddin bin Alias [Jelebu]":                                            "dato-jalaluddin-bin-alias-jelebu",
		"Dato Seri Dr Shahidan bin Kassim [Arau]":                                       "dato-seri-dr-shahidan-bin-kassim-arau",
		"Tuan Mamun bin Sulaiman [Kalabakan]":                                           "tuan-mamun-bin-sulaiman-kalabakan",
		"ptg Tuan Steven Choong Shiau Yoon [Tebrau]":                                    "tuan-steven-choong-shiau-yoon-tebrau",
		"Tuan Steven Choong Shiau Yoon [Tebrau]":                                        "tuan-steven-choong-shiau-yoon-tebrau",
		"ptg Tuan Lukanisman bin Awang Sauni [Sibuti]":                                  "tuan-lukanisman-bin-awang-sauni-sibuti",
		"Tuan Lukanisman bin Awang Sauni [Sibuti]":                                      "tuan-lukanisman-bin-awang-sauni-sibuti",
		"ptg Tuan Che Alias bin Hamid [Kemaman]":                                        "tuan-che-alias-bin-hamid-kemaman",
		"ptg Tuan Karupaiya a/l Mutusami [Padang Serai]":                                "tuan-karupaiya-a/l-mutusami-padang-serai",
		"Tuan Shaharizukirnain bin Abd Kadir [Setiu]":                                   "tuan-shaharizukirnain-bin-abd-kadir-setiu",
		"ptg Datuk Wilson Ugak anak Kumbong [Hulu Rajang]":                              "datuk-wilson-ugak-anak-kumbong-hulu-rajang",
		"ptg Menteri Pengangkutan [Datuk Seri Ir Dr Wee Ka Siong]":                      "menteri-pengangkutan-datuk-seri-ir-dr-wee-ka-siong",
		"Tuan Karupaiya a/l Mutusami [Padang Serai]":                                    "tuan-karupaiya-a/l-mutusami-padang-serai",
		"Tuan Pengerusi [Dato Mohd Rashid Hasnon]":                                      "tuan-pengerusi-dato-mohd-rashid-hasnon",
		"Tuan Pengerusi":                      "tuan-pengerusi",
		"Tuan Sivarasa Rasiah [Sungai Buloh]": "tuan-sivarasa-rasiah-sungai-buloh",
		"Menteri Komunikasi dan Multimedia [Dato Saifuddin Abdullah]":                           "menteri-komunikasi-dan-multimedia-dato-saifuddin-abdullah",
		"ptg Tan Sri Noh bin Haji Omar [Tanjong Karang]":                                        "tan-sri-noh-bin-haji-omar-tanjong-karang",
		"Tuan Noor Amin bin Ahmad [Kangar]":                                                     "tuan-noor-amin-bin-ahmad-kangar",
		"Tuan M Kulasegaran [Ipoh Barat]":                                                       "tuan-m-kulasegaran-ipoh-barat",
		"Dato Sri Hasan bin Arifin [Rompin]":                                                    "dato-sri-hasan-bin-arifin-rompin",
		"Tuan Ramkarpal Singh a/l Karpal Singh [Bukit Gelugor]":                                 "tuan-ramkarpal-singh-a/l-karpal-singh-bukit-gelugor",
		"ptg Datuk Liew Vui Keong [Batu Sapi]":                                                  "datuk-liew-vui-keong-batu-sapi",
		"ptg Tuan Haji Wan Hassan bin Mohd Ramli [Dungun]":                                      "tuan-haji-wan-hassan-bin-mohd-ramli-dungun",
		"ptg Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong]":                                "tuan-sanisvara-nethaji-rayer-a/l-rajaji-jelutong",
		"ptg Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling]":                          "datuk-seri-panglima-abdul-azeez-bin-abdul-rahim-baling",
		"ptg Tuan Ramkarpal Singh a/l Karpal Singh [Bukit Gelugor]":                             "tuan-ramkarpal-singh-a/l-karpal-singh-bukit-gelugor",
		"mlm Tuan Sivarasa Rasiah [Sungai Buloh]":                                               "tuan-sivarasa-rasiah-sungai-buloh",
		"mlm Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong]":                                "tuan-sanisvara-nethaji-rayer-a/l-rajaji-jelutong",
		"mlm Tuan Hassan bin Abdul Karim [Pasir Gudang]":                                        "tuan-hassan-bin-abdul-karim-pasir-gudang",
		"mlm Tuan Chan Ming Kai [Alor Setar]":                                                   "tuan-chan-ming-kai-alor-setar",
		"mlm Tuan Haji Wan Hassan bin Mohd Ramli [Dungun]":                                      "tuan-haji-wan-hassan-bin-mohd-ramli-dungun",
		"Tuan Haji Wan Hassan bin Mohd Ramli [Dungun]":                                          "tuan-haji-wan-hassan-bin-mohd-ramli-dungun",
		"Perumahan dan Kerajaan Tempatan (Puan Hajah Zuraida binti Kamaruddin) dan diluluskan]": "perumahan-dan-kerajaan-tempatan-(puan-hajah-zuraida-binti-kamaruddin)-dan-diluluskan",
	}
	spew.Dump(repMap26082020)
	repMap27082020 := map[string]akomantoso.RepresentativeID{
		"pg Timbalan Menteri Kesihatan II [Datuk Aaron Ago Dagang]":                        "timbalan-menteri-kesihatan-ii-datuk-aaron-ago-dagang",
		"Timbalan Yang di-Pertua [Dato Mohd Rashid Hasnon]":                                "timbalan-yang-di-pertua-dato-mohd-rashid-hasnon",
		"Dato Haji Jalaluddin bin Haji Alias [Jelebu]":                                     "dato-haji-jalaluddin-bin-haji-alias-jelebu",
		"Menteri Pertanian dan Industri Makanan [Datuk Seri Dr Ronald Kiandee]":            "menteri-pertanian-dan-industri-makanan-datuk-seri-dr-ronald-kiandee",
		"Datuk Seri Dr Ronald Kiandee":                                                     "datuk-seri-dr-ronald-kiandee",
		"Datuk Seri Saifuddin Nasution bin Ismail [Kulim-Bandar Baharu]":                   "datuk-seri-saifuddin-nasution-bin-ismail-kulim-bandar-baharu",
		"Tuan Chang Lih Kang [Tanjong Malim]":                                              "tuan-chang-lih-kang-tanjong-malim",
		"Menteri Perumahan dan Kerajaan Tempatan [Puan Hajah Zuraida binti Kamaruddin]":    "menteri-perumahan-dan-kerajaan-tempatan-puan-hajah-zuraida-binti-kamaruddin",
		"Puan Hajah Zuraida binti Kamaruddin":                                              "puan-hajah-zuraida-binti-kamaruddin",
		"Tuan Wong Hon Wai [Bukit Bendera]":                                                "tuan-wong-hon-wai-bukit-bendera",
		"Menteri Pengangkutan [Datuk Seri Ir Dr Wee Ka Siong]":                             "menteri-pengangkutan-datuk-seri-ir-dr-wee-ka-siong",
		"Datuk Seri Ir Dr Wee Ka Siong":                                                    "datuk-seri-ir-dr-wee-ka-siong",
		"Tuan Chow Kon Yeow [Tanjong]":                                                     "tuan-chow-kon-yeow-tanjong",
		"Datuk Zakaria bin Mohd Edris @ Tubau [Libaran]":                                   "datuk-zakaria-bin-mohd-edris-@-tubau-libaran",
		"Dato Sri Dr Haji Ismail bin Haji Mohamed Said":                                    "dato-sri-dr-haji-ismail-bin-haji-mohamed-said",
		"Tuan Fong Kui Lun [Bukit Bintang]":                                                "tuan-fong-kui-lun-bukit-bintang",
		"Datuk Seri Haji Salahuddin bin Ayub [Pulai]":                                      "datuk-seri-haji-salahuddin-bin-ayub-pulai",
		"Dato Haji Che Abdullah bin Mat Nawi":                                              "dato-haji-che-abdullah-bin-mat-nawi",
		"Tuan Sabri bin Azit [Jerai]":                                                      "tuan-sabri-bin-azit-jerai",
		"Dato Sri Hasan bin Arifin [Rompin]":                                               "dato-sri-hasan-bin-arifin-rompin",
		"Dato Sri Haji Abdul Rahman bin Mohamad":                                           "dato-sri-haji-abdul-rahman-bin-mohamad",
		"Tuan Kesavan a/l Subramaniam [Sungai Siput]":                                      "tuan-kesavan-a/l-subramaniam-sungai-siput",
		"Puan Wong Shu Qi [Kluang]":                                                        "puan-wong-shu-qi-kluang",
		"Timbalan Menteri Kewangan II [Tuan Mohd Shahar bin Abdullah]":                     "timbalan-menteri-kewangan-ii-tuan-mohd-shahar-bin-abdullah",
		"Tuan Mohd Shahar bin Abdullah":                                                    "tuan-mohd-shahar-bin-abdullah",
		"Datuk Seri Haji Ahmad bin Haji Maslan [Pontian]":                                  "datuk-seri-haji-ahmad-bin-haji-maslan-pontian",
		"Timbalan Yang di-Pertua Parlimen [Dato Mohd Rashid Hasnon]":                       "timbalan-yang-di-pertua-parlimen-dato-mohd-rashid-hasnon",
		"Datuk Wilson Ugak anak Kumbong [Hulu Rajang]":                                     "datuk-wilson-ugak-anak-kumbong-hulu-rajang",
		"Timbalan Menteri Kesihatan II [Datuk Aaron Ago Dagang]":                           "timbalan-menteri-kesihatan-ii-datuk-aaron-ago-dagang",
		"Datuk Aaron Ago Dagang":                                                           "datuk-aaron-ago-dagang",
		"Tuan Anyi Ngau [Baram]":                                                           "tuan-anyi-ngau-baram",
		"Puan Noorita binti Sual [Tenom]":                                                  "puan-noorita-binti-sual-tenom",
		"Timbalan Menteri Pendidikan I [Dato Dr Mah Hang Soon]":                            "timbalan-menteri-pendidikan-i-dato-dr-mah-hang-soon",
		"Dato Dr Mah Hang Soon":                                                            "dato-dr-mah-hang-soon",
		"Tuan Jugah a/k Muyang @ Tambat [Lubok Antu]":                                      "tuan-jugah-a/k-muyang-@-tambat-lubok-antu",
		"Tuan Yusuf bin Abd Wahab [Tanjong Manis]":                                         "tuan-haji-yusuf-bin-abd-wahab-tanjong-manis",
		"Tuan Haji Yusuf bin Abd Wahab [Tanjong Manis]":                                    "tuan-haji-yusuf-bin-abd-wahab-tanjong-manis",
		"Tuan Haji Ahmad Johnie bin Zawawi [Igan]":                                         "tuan-haji-ahmad-johnie-bin-zawawi-igan",
		"Tuan Karupaiya a/l Mutusami [Padang Serai]":                                       "tuan-karupaiya-a/l-mutusami-padang-serai",
		"Timbalan Menteri Kesihatan I [Dato Dr Haji Noor Azmi bin Ghazali]":                "timbalan-menteri-kesihatan-i-dato-dr-haji-noor-azmi-bin-ghazali",
		"Dato Dr Haji Noor Azmi bin Ghazali":                                               "dato-dr-haji-noor-azmi-bin-ghazali",
		"Tuan Abdul Latiff bin Abdul Rahman [Kuala Krai]":                                  "tuan-abdul-latiff-bin-abdul-rahman-kuala-krai",
		"YB Dato Hasbullah bin Osman (Gerik) tidak hadir]":                                 "yb-dato-hasbullah-bin-osman-(gerik)-tidak-hadir",
		"Tuan Tony Pua Kiam Wee [Damansara]":                                               "tuan-tony-pua-kiam-wee-damansara",
		"Menteri di Jabatan Perdana Menteri (Ekonomi) [Dato Sri Mustapa bin Mohamed]":      "menteri-di-jabatan-perdana-menteri-ekonomi-dato-sri-mustapa-bin-mohamed",
		"Dato Sri Mustapa bin Mohamed":                                                     "dato-sri-mustapa-bin-mohamed",
		"Timbalan Menteri Pembangunan Luar Bandar II [Dato Henry Sum Agong]":               "timbalan-menteri-pembangunan-luar-bandar-ii-dato-henry-sum-agong",
		"Dato Henry Sum Agong":                                                             "dato-henry-sum-agong",
		"Tuan Syed Ibrahim bin Syed Noh [Ledang]":                                          "tuan-syed-ibrahim-bin-syed-noh-ledang",
		"Timbalan Menteri Pengangkutan [Tuan Haji Hasbi bin Habibollah]":                   "timbalan-menteri-pengangkutan-tuan-haji-hasbi-bin-habibollah",
		"Tuan Haji Hasbi bin Habibollah":                                                   "tuan-haji-hasbi-bin-habibollah",
		"Datuk Mohamad bin Alamin [Kimanis]":                                               "datuk-mohamad-bin-alamin-kimanis",
		"Tuan Sim Tze Tzin [Bayan Baru]":                                                   "tuan-sim-tze-tzin-bayan-baru",
		"Datuk Liew Vui Keong [Batu Sapi]":                                                 "datuk-liew-vui-keong-batu-sapi",
		"Dato Takiyuddin bin Hassan":                                                       "dato-takiyuddin-bin-hassan",
		"Menteri Sumber Manusia [Datuk Seri M Saravanan]":                                  "menteri-sumber-manusia-datuk-seri-m-saravanan",
		"tgh Datuk Liew Vui Keong [Batu Sapi]":                                             "tgh-datuk-liew-vui-keong-batu-sapi",
		"Timbalan Yang di-Pertua [Dato Sri Azalina Othman Said]":                           "timbalan-yang-di-pertua-dato-sri-azalina-othman-said",
		"Timbalan Menteri di Jabatan Perdana Menteri (Ekonomi) [Tuan Arthur Joseph Kurup]": "timbalan-menteri-di-jabatan-perdana-menteri-ekonomi-tuan-arthur-joseph-kurup",
		"Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling]":                         "datuk-seri-panglima-abdul-azeez-bin-abdul-rahim-baling",
		"tgh Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling]":                     "tgh-datuk-seri-panglima-abdul-azeez-bin-abdul-rahim-baling",
		"tgh Tuan Lim Guan Eng [Bagan]":                                                    "tgh-tuan-lim-guan-eng-bagan",
		"Tuan Lim Guan Eng [Bagan]":                                                        "tuan-lim-guan-eng-bagan",
		"Tuan Mohamed Hanipa bin Maidin [Sepang]":                                          "tuan-mohamed-hanipa-bin-maidin-sepang",
		"Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong]":                               "tuan-sanisvara-nethaji-rayer-a/l-rajaji-jelutong",
		"tgh Timbalan Menteri Kewangan II [Tuan Mohd Shahar bin Abdullah]":                 "tgh-timbalan-menteri-kewangan-ii-tuan-mohd-shahar-bin-abdullah",
		"Dato Seri Anwar bin Ibrahim [Port Dickson]":                                       "dato-seri-anwar-bin-ibrahim-port-dickson",
		"Datuk Seri Dr Haji Dzulkefly bin Ahmad [Kuala Selangor]":                          "datuk-seri-dr-haji-dzulkefly-bin-ahmad-kuala-selangor",
		"Timbalan Menteri Sumber Manusia [Tuan Haji Awang bin Hashim]":                     "timbalan-menteri-sumber-manusia-tuan-haji-awang-bin-hashim",
		"Datuk Wira Dr Mohd Hatta bin Md Ramli [Lumut]":                                    "datuk-wira-dr-mohd-hatta-bin-md-ramli-lumut",
		"tgh Datuk Wira Dr Mohd Hatta bin Md Ramli [Lumut]":                                "tgh-datuk-wira-dr-mohd-hatta-bin-md-ramli-lumut",
		"tgh Tuan Mohd Shahar bin Abdullah":                                                "tgh-tuan-mohd-shahar-bin-abdullah",
		"Dato Johari bin Abdul [Sungai Petani]":                                            "dato-johari-bin-abdul-sungai-petani",
		"Tuan Yang di-Pertua":                                                              "tuan-yang-di-pertua",
		"ptg Timbalan Menteri Kewangan II [Tuan Mohd Shahar bin Abdullah]":                 "timbalan-menteri-kewangan-ii-tuan-mohd-shahar-bin-abdullah",
		"Tuan Ramli bin Dato Mohd Nor [Cameron Highlands]":                                 "tuan-ramli-bin-dato-mohd-nor-cameron-highlands",
		"Datuk Mohamad bin Alamin [ Kimanis]":                                              "datuk-mohamad-bin-alamin--kimanis",
		"Tuan Sivarasa Rasiah [Sungai Buloh]":                                              "tuan-sivarasa-rasiah-sungai-buloh",
		"Dato Seri Dr Shahidan bin Kassim [Arau]":                                          "dato-seri-dr-shahidan-bin-kassim-arau",
		"ptg Dato Seri Dr Shahidan bin Kassim [Arau]":                                      "dato-seri-dr-shahidan-bin-kassim-arau",
		"Tuan Nik Nazmi bin Nik Ahmad [Setiawangsa]":                                       "tuan-nik-nazmi-bin-nik-ahmad-setiawangsa",
		"Tuan Charles Anthony Santiago [Klang]":                                            "tuan-charles-anthony-santiago-klang",
		"Tuan Pang Hok Liong [Labis]":                                                      "tuan-pang-hok-liong-labis",
		"Tuan Cha Kee Chin [Rasah]":                                                        "tuan-cha-kee-chin-rasah",
		"Dato Dr Xavier Jayakumar a/l Arulanandam [Kuala Langat]":                          "dato-dr-xavier-jayakumar-a/l-arulanandam-kuala-langat",
		"Datuk Dr Hasan bin Bahrom [Tampin]":                                               "datuk-dr-hasan-bin-bahrom-tampin",
		"Dato Haji Salim Sharif [Jempol]":                                                  "dato-haji-salim-sharif-jempol",
		"ptg Dato Seri Anwar bin Ibrahim [Port Dickson]":                                   "dato-seri-anwar-bin-ibrahim-port-dickson",
		"Tuan Ahmad Tarmizi bin Sulaiman [Sik]":                                            "tuan-ahmad-tarmizi-bin-sulaiman-sik",
		"ptg Tuan Mohamed Hanipa bin Maidin [Sepang]":                                      "tuan-mohamed-hanipa-bin-maidin-sepang",
		"Tuan Steven Choong Shiau Yoon [Tebrau]":                                           "tuan-steven-choong-shiau-yoon-tebrau",
		"ptg Datuk Seri Haji Ahmad bin Haji Maslan [Pontian]":                              "datuk-seri-haji-ahmad-bin-haji-maslan-pontian",
		"ptg Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling]":                     "datuk-seri-panglima-abdul-azeez-bin-abdul-rahim-baling",
		"ptg Dato Sri Haji Tajuddin bin Abdul Rahman [Pasir Salak]":                        "dato-sri-haji-tajuddin-bin-abdul-rahman-pasir-salak",
		"Dato Sri Bung Moktar bin Radin [Kinabatangan]":                                    "dato-sri-bung-moktar-bin-radin-kinabatangan",
		"ptg Dato Sri Bung Moktar bin Radin [Kinabatangan]":                                "dato-sri-bung-moktar-bin-radin-kinabatangan",
		"Dato Seri Tiong King Sing [Bintulu]":                                              "dato-seri-tiong-king-sing-bintulu",
		"Tuan Ahmad Fahmi bin Mohamed Fadzil [Lembah Pantai]":                              "tuan-ahmad-fahmi-bin-mohamed-fadzil-lembah-pantai",
		"Dato Mahfuz bin Haji Omar [Pokok Sena]":                                           "dato-mahfuz-bin-haji-omar-pokok-sena",
		"Puan Kasthuriraani a/p Patto [Batu Kawan]":                                        "puan-kasthuriraani-a/p-patto-batu-kawan",
		"Puan Teo Nie Ching [Kulai]":                                                       "puan-teo-nie-ching-kulai",
		"Puan Hannah Yeoh [Segambut]":                                                      "puan-hannah-yeoh-segambut",
		"Timbalan Menteri Wilayah Persekutuan [Dato Seri Dr Santhara]":                     "timbalan-menteri-wilayah-persekutuan-dato-seri-dr-santhara",
		"Dato Hasbullah bin Osman [Gerik]":                                                 "dato-hasbullah-bin-osman-gerik",
	}
	spew.Dump(repMap27082020)
	// Setup different PDFs and DPSs
	dps := DebateProcessorState{
		SectionMarkers: SectionMarkers{
			DatePageMarker: "august 26 2020",
		},
		RepresentativesMap: repMap26082020,
	}
	// For test case 14th
	dps14 := DebateProcessorState{
		SectionMarkers: SectionMarkers{
			DatePageMarker: "August 27 2020",
		},
		RepresentativesMap: repMap27082020,
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case #1", args{
			pdfPath: "../../raw/Parliament/Hansard/DR-26082020.pdf",
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
			pdfPath: "../../raw/Parliament/Hansard/DR-27082020.pdf",
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

func TestDebateProcessPages(t *testing.T) {
	type args struct {
		pdfDocument *akomantoso.PDFDocument
		dps         DebateProcessorState
	}
	// Samples of RepMap
	repMap28072020 := map[string]akomantoso.RepresentativeID{
		"Dato Sri Hajah Rohani binti Abdul Karim [Batang Lupar]":                        "dato-sri-hajah-rohani-binti-abdul-karim-batang-lupar",
		"Menteri Perumahan dan Kerajaan Tempatan [Puan Hajah Zuraida binti Kamaruddin]": "menteri-perumahan-dan-kerajaan-tempatan-puan-hajah-zuraida-binti-kamaruddin",
		"Tuan Yang di-Pertua":                                                               "tuan-yang-di-pertua",
		"Puan Hajah Zuraida binti Kamaruddin":                                               "puan-hajah-zuraida-binti-kamaruddin",
		"Datuk Haji Hasanuddin bin Mohd Yunus [Hulu Langat]":                                "datuk-haji-hasanuddin-bin-mohd-yunus-hulu-langat",
		"Tuan Wong Hon Wai [Bukit Bendera]":                                                 "tuan-wong-hon-wai-bukit-bendera",
		"Tan Sri Datuk Seri Panglima Haji Annuar bin Haji Musa":                             "tan-sri-datuk-seri-panglima-haji-annuar-bin-haji-musa",
		"Tuan Abdul Latiff bin Abdul Rahman [Kuala Krai]":                                   "tuan-abdul-latiff-bin-abdul-rahman-kuala-krai",
		"Menteri Kesihatan [Datuk Seri Dr Adham bin Baba]":                                  "menteri-kesihatan-datuk-seri-dr-adham-bin-baba",
		"Tuan Chong Chieng Jen [Stampin]":                                                   "tuan-chong-chieng-jen-stampin",
		"Datuk Seri Dr Adham bin Baba":                                                      "datuk-seri-dr-adham-bin-baba",
		"Dato Hajah Azizah binti Mohd Dun [Beaufort]":                                       "dato-hajah-azizah-binti-mohd-dun-beaufort",
		"Tuan Khoo Poay Tiong [Kota Melaka]":                                                "tuan-khoo-poay-tiong-kota-melaka",
		"Tuan Ramli bin Dato Mohd Nor [Cameron Highlands]":                                  "tuan-ramli-bin-dato-mohd-nor-cameron-highlands",
		"Tuan Hassan bin Abdul Karim [Pasir Gudang]":                                        "tuan-hassan-bin-abdul-karim-pasir-gudang",
		"Puan Hajah Fuziah binti Salleh [Kuantan]":                                          "puan-hajah-fuziah-binti-salleh-kuantan",
		"Datuk Dr Haji Zulkifli Mohamad Al-Bakri":                                           "datuk-dr-haji-zulkifli-mohamad-al-bakri",
		"Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling]":                          "datuk-seri-panglima-abdul-azeez-bin-abdul-rahim-baling",
		"Dato Hasbullah bin Osman [Gerik]":                                                  "dato-hasbullah-bin-osman-gerik",
		"Timbalan Menteri Perusahaan Perladangan dan Komoditi II [Tuan Willie anak Mongin]": "timbalan-menteri-perusahaan-perladangan-dan-komoditi-ii-tuan-willie-anak-mongin",
		"Tuan Willie anak Mongin":                                                           "tuan-willie-anak-mongin",
		"Tuan Karupaiya a/l Mutusami [Padang Serai]":                                        "tuan-karupaiya-a/l-mutusami-padang-serai",
		"Datuk Zakaria bin Mohd Edris @ Tubau [Libaran]":                                    "datuk-zakaria-bin-mohd-edris-@-tubau-libaran",
		"Dato Jalaluddin bin Alias [Jelebu]":                                                "dato-jalaluddin-bin-alias-jelebu",
		"Datuk Wilson Ugak anak Kumbong [Hulu Rajang]":                                      "datuk-wilson-ugak-anak-kumbong-hulu-rajang",
		"Datuk Liew Vui Keong [Batu Sapi]":                                                  "datuk-liew-vui-keong-batu-sapi",
		"Timbalan Menteri Kesihatan II [Datuk Aaron Ago Dagang]":                            "timbalan-menteri-kesihatan-ii-datuk-aaron-ago-dagang",
		"Timbalan Menteri Kesihatan I [Dato Dr Haji Noor Azmi bin Ghazali]":                 "timbalan-menteri-kesihatan-i-dato-dr-haji-noor-azmi-bin-ghazali",
		"Dato Dr Haji Noor Azmi bin Ghazali":                                                "dato-dr-haji-noor-azmi-bin-ghazali",
		"Tuan Che Alias bin Hamid [Kemaman]":                                                "tuan-che-alias-bin-hamid-kemaman",
		"Datuk Aaron Ago Dagang":                                                            "datuk-aaron-ago-dagang",
		"Tuan Haji Ahmad bin Hassan [Papar]":                                                "tuan-haji-ahmad-bin-hassan-papar",
		"Dato Seri Dr Wan Azizah Wan Ismail [Pandan]":                                       "dato-seri-dr-wan-azizah-wan-ismail-pandan",
		"Puan Teo Nie Ching [Kulai]":                                                        "puan-teo-nie-ching-kulai",
		"Dato Mohd Nizar bin Haji Zakaria [Parit]":                                          "dato-mohd-nizar-bin-haji-zakaria-parit",
		"Menteri Sains Teknologi dan Inovasi [Tuan Khairy Jamaluddin Abu Bakar]":            "menteri-sains-teknologi-dan-inovasi-tuan-khairy-jamaluddin-abu-bakar",
		"Tuan Khairy Jamaluddin Abu Bakar":                                                  "tuan-khairy-jamaluddin-abu-bakar",
		"Datuk Seri Saifuddin Nasution bin Ismail [Kulim-Bandar Baharu]":                    "datuk-seri-saifuddin-nasution-bin-ismail-kulim-bandar-baharu",
		"Dato Wira Haji Amiruddin bin Haji Hamzah [Kubang Pasu]":                            "dato-wira-haji-amiruddin-bin-haji-hamzah-kubang-pasu",
		"Timbalan Menteri Perdagangan Antarabangsa dan Industri [Datuk Lim Ban Hong]":       "timbalan-menteri-perdagangan-antarabangsa-dan-industri-datuk-lim-ban-hong",
		"Datuk Lim Ban Hong":                                                                "datuk-lim-ban-hong",
		"Dato Sri Hassan bin Ariffin [Rompin]":                                              "dato-sri-hassan-bin-ariffin-rompin",
		"Datuk Seri Haji Ahmad bin Haji Maslan [Pontian]":                                   "datuk-seri-haji-ahmad-bin-haji-maslan-pontian",
		"Tuan Haji Yusuf bin Abd Wahab [Tanjong Manis]":                                     "tuan-haji-yusuf-bin-abd-wahab-tanjong-manis",
		"Datuk Wira Dr Mohd Hatta bin Md Ramli [Lumut]":                                     "datuk-wira-dr-mohd-hatta-bin-md-ramli-lumut",
		"Datuk Mohamad bin Alamin [Kimanis]":                                                "datuk-mohamad-bin-alamin-kimanis",
		"Puan Hannah Yeoh [Segambut]":                                                       "puan-hannah-yeoh-segambut",
		"Tuan P Prabakaran [Batu]":                                                          "tuan-p-prabakaran-batu",
		"Dato Haji Mohd Fasiah bin Haji Mohd Fakeh [Sabak Bernam]":                          "dato-haji-mohd-fasiah-bin-haji-mohd-fakeh-sabak-bernam",
		"Dato Haji Mohd Fasiah bin Mohd Fakeh [Sabak Bernam]":                               "dato-haji-mohd-fasiah-bin-mohd-fakeh-sabak-bernam",
		"tgh Dato Mohd Nizar bin Haji Zakaria [Parit]":                                      "dato-mohd-nizar-bin-haji-zakaria-parit",
		"Timbalan Yang di-Pertua [Dato Mohd Rashid Hasnon]":                                 "timbalan-yang-di-pertua-dato-mohd-rashid-hasnon",
		"tgh Tuan Chang Lih Kang [Tanjong Malim]":                                           "tuan-chang-lih-kang-tanjong-malim",
		"Tuan Chang Lih Kang [Tanjong Malim]":                                               "tuan-chang-lih-kang-tanjong-malim",
		"Dato Seri Dr Shahidan bin Kassim [Arau]":                                           "dato-seri-dr-shahidan-bin-kassim-arau",
		"tgh Tuan Gobind Singh Deo [Puchong]":                                               "tuan-gobind-singh-deo-puchong",
		"Tuan Mohamed Hanipa bin Maidin [Sepang]":                                           "tuan-mohamed-hanipa-bin-maidin-sepang",
		"Tuan Gobind Singh Deo [Puchong]":                                                   "tuan-gobind-singh-deo-puchong",
		"tgh Tuan Syed Ibrahim bin Syed Noh [Ledang]":                                       "tuan-syed-ibrahim-bin-syed-noh-ledang",
		"Tuan Syed Ibrahim bin Syed Noh [Ledang]":                                           "tuan-syed-ibrahim-bin-syed-noh-ledang",
		"Timbalan Yang di-Pertua [Dato Sri Azalina Othman Said]":                            "timbalan-yang-di-pertua-dato-sri-azalina-othman-said",
		"ptg Tuan Syed Ibrahim bin Syed Noh [Ledang]":                                       "tuan-syed-ibrahim-bin-syed-noh-ledang",
		"ptg Tuan Baru Bian [Selangau]":                                                     "tuan-baru-bian-selangau",
		"ptg Tuan Sabri bin Azit [Jerai]":                                                   "tuan-sabri-bin-azit-jerai",
		"Datuk Dr Hasan bin Bahrom [Tampin]":                                                "datuk-dr-hasan-bin-bahrom-tampin",
		"Tuan Sabri bin Azit [Jerai]":                                                       "tuan-sabri-bin-azit-jerai",
		"Tuan Syed Saddiq bin Syed Abdul Rahman [Muar]":                                     "tuan-syed-saddiq-bin-syed-abdul-rahman-muar",
		"Tuan Nik Nazmi bin Nik Ahmad [Setiawangsa]":                                        "tuan-nik-nazmi-bin-nik-ahmad-setiawangsa",
		"Datuk Seri Dr Mujahid Yusof Rawa [Parit Buntar]":                                   "datuk-seri-dr-mujahid-yusof-rawa-parit-buntar",
		"Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong]":                                "tuan-sanisvara-nethaji-rayer-a/l-rajaji-jelutong",
		"Dato Haji Salim Sharif [Jempol]":                                                   "dato-haji-salim-sharif-jempol",
		"ptg Datuk Liew Vui Keong [Batu Sapi]":                                              "datuk-liew-vui-keong-batu-sapi",
		"Dato Ngeh Koo Ham [Beruas]":                                                        "dato-ngeh-koo-ham-beruas",
		"Puan Noorita binti Sual [Tenom]":                                                   "puan-noorita-binti-sual-tenom",
		"ptg Puan Rusnah binti Aluai [Tangga Batu]":                                         "puan-rusnah-binti-aluai-tangga-batu",
		"Puan Rusnah binti Aluai [Tangga Batu]":                                             "puan-rusnah-binti-aluai-tangga-batu",
		"ptg Tuan Su Keong Siong [Kampar]":                                                  "tuan-su-keong-siong-kampar",
		"Tuan Su Keong Siong [Kampar]":                                                      "tuan-su-keong-siong-kampar",
		"ptg Puan Hajah Natrah Ismail [Sekijang]":                                           "puan-hajah-natrah-ismail-sekijang",
		"Puan Hajah Natrah Ismail [Sekijang]":                                               "puan-hajah-natrah-ismail-sekijang",
		"Menteri Pembangunan Luar Bandar [Datuk Dr Haji Abd Latiff bin Ahmad]":              "menteri-pembangunan-luar-bandar-datuk-dr-haji-abd-latiff-bin-ahmad",
		"ptg Tuan Larry Soon @ Larry Sng Wei Shien [Julau]":                                 "tuan-larry-soon-@-larry-sng-wei-shien-julau",
		"Tuan Larry Soon @ Larry Sng Wei Shien [Julau]":                                     "tuan-larry-soon-@-larry-sng-wei-shien-julau",
		"ptg Puan Kasthuriraani a/p Patto [Batu Kawan]":                                     "puan-kasthuriraani-a/p-patto-batu-kawan",
		"Puan Teresa Kok Suh Sim [Seputeh]":                                                 "puan-teresa-kok-suh-sim-seputeh",
		"ptg Tuan Kesavan a/l Subramaniam [Sungai Siput]":                                   "tuan-kesavan-a/l-subramaniam-sungai-siput",
	}
	// DEBUG
	//spew.Dump(repMap28072020)
	repMap18052020 := map[string]akomantoso.RepresentativeID{
		"Tuan Yang di-Pertua": "tuan-yang-di-pertua",
	}
	spew.Dump(repMap18052020)
	repMap13072020 := map[string]akomantoso.RepresentativeID{
		"Tuan Yang di-Pertua":                                                             "tuan-yang-di-pertua",
		"Dato Seri Anwar bin Ibrahim [Port Dickson]":                                      "dato-seri-anwar-bin-ibrahim-port-dickson",
		"Dato Sri Bung Moktar bin Radin [Kinabatangan]":                                   "dato-sri-bung-moktar-bin-radin-kinabatangan",
		"Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling]":                        "datuk-seri-panglima-abdul-azeez-bin-abdul-rahim-baling",
		"Puan Rubiah binti Haji Wang [Kota Samarahan]":                                    "puan-rubiah-binti-haji-wang-kota-samarahan",
		"Perdana Menteri [Tan Sri Dato Sri Haji Muhyiddin bin Md Yassin]":                 "perdana-menteri-tan-sri-dato-sri-haji-muhyiddin-bin-md-yassin",
		"Puan Rubiah binti Wang [Kota Samarahan]":                                         "puan-rubiah-binti-wang-kota-samarahan",
		"Tan Sri Dato Sri Haji Muhyiddin bin Md Yassin":                                   "tan-sri-dato-sri-haji-muhyiddin-bin-md-yassin",
		"Tuan Khalid bin Abd Samad [Shah Alam]":                                           "tuan-khalid-bin-abd-samad-shah-alam",
		"Tuan Sim Tze Tzin [Bayan Baru]":                                                  "tuan-sim-tze-tzin-bayan-baru",
		"Tuan Haji Akmal Nasrullah bin Mohd Nasir [Johor Bahru]":                          "tuan-haji-akmal-nasrullah-bin-mohd-nasir-johor-bahru",
		"Tuan Nik Nazmi bin Nik Ahmad [Setiawangsa]":                                      "tuan-nik-nazmi-bin-nik-ahmad-setiawangsa",
		"Datuk Seri Haji Ahmad bin Haji Maslan [Pontian]":                                 "datuk-seri-haji-ahmad-bin-haji-maslan-pontian",
		"Tuan Chang Lih Kang [Tanjong Malim]":                                             "tuan-chang-lih-kang-tanjong-malim",
		"Datuk Seri Panglima Haji Mohd Shafie bin Haji Apdal [Semporna]":                  "datuk-seri-panglima-haji-mohd-shafie-bin-haji-apdal-semporna",
		"Dato Seri Dr Wan Azizah Wan Ismail [Pandan]":                                     "dato-seri-dr-wan-azizah-wan-ismail-pandan",
		"Datuk Seri Mohd Redzuan bin Md Yusof":                                            "datuk-seri-mohd-redzuan-bin-md-yusof",
		"Dato Haji Salim Sharif [Jempol]":                                                 "dato-haji-salim-sharif-jempol",
		"Datuk Seri Dr Haji Dzulkefly bin Ahmad [Kuala Selangor]":                         "datuk-seri-dr-haji-dzulkefly-bin-ahmad-kuala-selangor",
		"Dato Seri Dr Shahidan bin Kassim [Arau]":                                         "dato-seri-dr-shahidan-bin-kassim-arau",
		"Tuan Syed Ibrahim bin Syed Noh [Ledang]":                                         "tuan-syed-ibrahim-bin-syed-noh-ledang",
		"Tuan Fong Kui Lun [Bukit Bintang]":                                               "tuan-fong-kui-lun-bukit-bintang",
		"Tan Sri Datuk Seri Panglima Haji Annuar bin Haji Musa":                           "tan-sri-datuk-seri-panglima-haji-annuar-bin-haji-musa",
		"Tuan Nik Mohamad Abduh bin Nik Abdul Aziz [Bachok]":                              "tuan-nik-mohamad-abduh-bin-nik-abdul-aziz-bachok",
		"Menteri Kesihatan [Datuk Seri Dr Adham bin Baba]":                                "menteri-kesihatan-datuk-seri-dr-adham-bin-baba",
		"Datuk Seri Dr Adham bin Baba [Tenggara]":                                         "datuk-seri-dr-adham-bin-baba-tenggara",
		"Puan Nurul Izzah binti Anwar [Permatang Pauh]":                                   "puan-nurul-izzah-binti-anwar-permatang-pauh",
		"Datuk Seri Dr Adham bin Baba":                                                    "datuk-seri-dr-adham-bin-baba",
		"Datuk Wira Dr Mohd Hatta bin Md Ramli [Lumut]":                                   "datuk-wira-dr-mohd-hatta-bin-md-ramli-lumut",
		"Dato Dr Mohd Khairuddin bin Aman Razali":                                         "dato-dr-mohd-khairuddin-bin-aman-razali",
		"Dato Mohd Nizar bin Haji Zakaria [Parit]":                                        "dato-mohd-nizar-bin-haji-zakaria-parit",
		"Dato Hasbullah bin Osman [Gerik]":                                                "dato-hasbullah-bin-osman-gerik",
		"Tuan Haji Wan Hassan bin Mohd Ramli [Dungun]":                                    "tuan-haji-wan-hassan-bin-mohd-ramli-dungun",
		"Menteri Sains Teknologi dan Inovasi [Tuan Khairy Jamaluddin Abu Bakar]":          "menteri-sains-teknologi-dan-inovasi-tuan-khairy-jamaluddin-abu-bakar",
		"Timbalan Yang di-Pertua [Dato Mohd Rashid Hasnon]":                               "timbalan-yang-di-pertua-dato-mohd-rashid-hasnon",
		"Tuan Khairy Jamaluddin Abu Bakar":                                                "tuan-khairy-jamaluddin-abu-bakar",
		"Puan Nor Azrina binti Surip [Merbok]":                                            "puan-nor-azrina-binti-surip-merbok",
		"Datuk Seri Saifuddin Nasution bin Ismail [Kulim-Bandar Baharu]":                  "datuk-seri-saifuddin-nasution-bin-ismail-kulim-bandar-baharu",
		"Datuk Ignatius Darell Leiking [Penampang]":                                       "datuk-ignatius-darell-leiking-penampang",
		"Menteri Perdagangan Antarabangsa dan Industri [Dato Seri Mohamed Azmin bin Ali]": "menteri-perdagangan-antarabangsa-dan-industri-dato-seri-mohamed-azmin-bin-ali",
		"Datuk Ignatius Dorell Leiking [Penampang]":                                       "datuk-ignatius-dorell-leiking-penampang",
		"Dato Seri Mohamed Azmin bin Ali":                                                 "dato-seri-mohamed-azmin-bin-ali",
		"Dato Sri Richard Riot anak Jaem [Serian]":                                        "dato-sri-richard-riot-anak-jaem-serian",
		"pg Menteri Dalam Negeri [Dato Seri Hamzah bin Zainudin]":                         "menteri-dalam-negeri-dato-seri-hamzah-bin-zainudin",
		"pg Dato Seri Anwar bin Ibrahim [Port Dickson]":                                   "dato-seri-anwar-bin-ibrahim-port-dickson",
		"Datuk Mohd Azis bin Jamman [Sepanggar]":                                          "datuk-mohd-azis-bin-jamman-sepanggar",
		"Tuan Mohamad bin Sabu [Kota Raja]":                                               "tuan-mohamad-bin-sabu-kota-raja",
		"Dato Seri Haji Salahuddin bin Ayub [Pulai]":                                      "dato-seri-haji-salahuddin-bin-ayub-pulai",
		"Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong]":                              "tuan-sanisvara-nethaji-rayer-a/l-rajaji-jelutong",
		"Dato Jalaluddin bin Alias [Jelebu]":                                              "dato-jalaluddin-bin-alias-jelebu",
		"Tuan Gobind Singh Deo [Puchong]":                                                 "tuan-gobind-singh-deo-puchong",
		"Timbalan Menteri Sumber Manusia [Tuan Haji Awang bin Hashim]":                    "timbalan-menteri-sumber-manusia-tuan-haji-awang-bin-hashim",
		"Datuk Liew Vui Keong [Batu Sapi]":                                                "datuk-liew-vui-keong-batu-sapi",
		"Dato Takiyuddin bin Hassan":                                                      "dato-takiyuddin-bin-hassan",
		"Tuan Sivakumar Varatharaju Naidu [Batu Gajah]":                                   "tuan-sivakumar-varatharaju-naidu-batu-gajah",
		"Datuk Seri Dr Mujahid Yusof Rawa [Parit Buntar]":                                 "datuk-seri-dr-mujahid-yusof-rawa-parit-buntar",
		"Tuan Sim Chee Keong [Bukit Mertajam]":                                            "tuan-sim-chee-keong-bukit-mertajam",
		"Tuan Mohamed Hanipa bin Maidin [Sepang]":                                         "tuan-mohamed-hanipa-bin-maidin-sepang",
		"Dato Ngeh Koo Ham [Beruas]":                                                      "dato-ngeh-koo-ham-beruas",
		"Dato Sri Dr Haji Ismail bin Haji Mohamed Said":                                   "dato-sri-dr-haji-ismail-bin-haji-mohamed-said",
		"Tuan Ramkarpal Singh a/l Karpal Singh [Bukit Gelugor]":                           "tuan-ramkarpal-singh-a/l-karpal-singh-bukit-gelugor",
		"Dato Sri Haji Tajuddin bin Abdul Rahman [Pasir Salak]":                           "dato-sri-haji-tajuddin-bin-abdul-rahman-pasir-salak",
		"Tuan Abdul Latiff bin Abdul Rahman [Kuala Krai]":                                 "tuan-abdul-latiff-bin-abdul-rahman-kuala-krai",
		"Tan Sri Noh bin Haji Omar [Tanjong Karang]":                                      "tan-sri-noh-bin-haji-omar-tanjong-karang",
		"Tuan Ahmad Fahmi bin Mohamed Fadzil [Lembah Pantai]":                             "tuan-ahmad-fahmi-bin-mohamed-fadzil-lembah-pantai",
		"Tuan Mamun bin Sulaiman [Kalabakan]":                                             "tuan-mamun-bin-sulaiman-kalabakan",
		"Tun Dr Mahathir bin Mohamad [Langkawi]":                                          "tun-dr-mahathir-bin-mohamad-langkawi",
		"Dato Seri Mohamed Nazri bin Abdul Aziz [Padang Rengas]":                          "dato-seri-mohamed-nazri-bin-abdul-aziz-padang-rengas",
		"Menteri Tenaga dan Sumber Asli [Dato Dr Shamsul Anuar bin Nasarah]":              "menteri-tenaga-dan-sumber-asli-dato-dr-shamsul-anuar-bin-nasarah",
		"Dato Dr Shamsul Anuar bin Nasarah":                                               "dato-dr-shamsul-anuar-bin-nasarah",
		"Tuan Syed Saddiq bin Syed Abdul Rahman [Muar]":                                   "tuan-syed-saddiq-bin-syed-abdul-rahman-muar",
		"Dato Mahfuz bin Haji Omar [Pokok Sena]":                                          "dato-mahfuz-bin-haji-omar-pokok-sena",
		"Dato Seri Utama Haji Mukhriz Tun Dr Mahathir [Jerlun]":                           "dato-seri-utama-haji-mukhriz-tun-dr-mahathir-jerlun",
		"Menteri Dalam Negeri [Dato Seri Hamzah bin Zainudin]":                            "menteri-dalam-negeri-dato-seri-hamzah-bin-zainudin",
		"Tuan Sivarasa Rasiah [Sungai Buloh]":                                             "tuan-sivarasa-rasiah-sungai-buloh",
		"Tuan M Kulasegaran [Ipoh Barat]":                                                 "tuan-m-kulasegaran-ipoh-barat",
		"Tuan Wong Kah Woh [Ipoh Timur]":                                                  "tuan-wong-kah-woh-ipoh-timur",
		"Tuan Wong Hon Wai [Bukit Bendera]":                                               "tuan-wong-hon-wai-bukit-bendera",
		"Tan Sri Dato Haji Muhyiddin bin Md Yassin [Pagoh]":                               "tan-sri-dato-haji-muhyiddin-bin-md-yassin-pagoh",
		"Datuk Haji Hasanuddin bin Mohd Yunus [Hulu Langat]":                              "datuk-haji-hasanuddin-bin-mohd-yunus-hulu-langat",
		"Tuan Haji Awang bin Hashim":                                                      "tuan-haji-awang-bin-hashim",
		"Puan Kasthuriraani a/p Patto [Batu Kawan]":                                       "puan-kasthuriraani-a/p-patto-batu-kawan",
		"Tuan Lim Guan Eng [Bagan]":                                                       "tuan-lim-guan-eng-bagan",
		"Menteri Sumber Manusia [Datuk Seri M Saravanan]":                                 "menteri-sumber-manusia-datuk-seri-m-saravanan",
		"Tuan P Prabakaran [Batu]":                                                        "tuan-p-prabakaran-batu",
		"Timbalan Menteri Kesihatan I [Dato Dr Haji Noor Azmi bin Ghazali]":               "timbalan-menteri-kesihatan-i-dato-dr-haji-noor-azmi-bin-ghazali",
	}
	// DEBUG
	//spew.Dump(repMap13072020)

	// Setup different PDFs and DPSs
	// Extract out Section Metadata for attachment
	extractOptions := akomantoso.ExtractPDFOptions{
		StartPage:  2,
		NumPages:   5,
		MaxSampled: 10000,
	}
	pdfDocument, perr := akomantoso.NewPDFDocument("../../raw/Parliament/Hansard/DR-28072020.pdf", &extractOptions)
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
			DatePageMarker:         "26 aug 2020",
			SessionStartMarkerLine: 2,
		},
		RepresentativesMap: repMap28072020,
	}

	// For 15th test case
	pdfDocument15, perr := akomantoso.NewPDFDocument("../../raw/Parliament/Hansard/DR-13072020 New 1.pdf", &extractOptions)
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
			DatePageMarker:         "27 aug 2020",
			SessionStartMarkerLine: 2,
		},
		RepresentativesMap: repMap13072020,
	}

	tests := []struct {
		name string
		args args
		want ParliamentDebateContent
	}{
		{"case #1", args{pdfDocument, dps}, ParliamentDebateContent{}},
		{"case #2", args{pdfDocument15, dps15}, ParliamentDebateContent{}},
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
