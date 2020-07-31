package parliament

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
)

func Test_extractDebaters(t *testing.T) {
	type args struct {
		allLines []string
	}
	tests := []struct {
		name string
		args args
		want []akomantoso.Representative
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractDebaters(tt.args.allLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractDebaters() = %v, want %v", got, tt.want)
			}
		})
	}
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
		{"case #1", args{line: "   JAWAPAN-JAWAPAN MENTERI BAGI PERTANYAAN-PERTANYAAN  1.  Dato Sri Hajah Rohani binti Abdul Karim [Batang Lupar]  "}, true, "Dato Sri Hajah Rohani binti Abdul Karim [Batang Lupar]"},
		{"case #2", args{line: " Tuan Yang di-Pertua:  "}, true, "Tuan Yang di-Pertua"},
		{"case #3", args{line: " Puan Hajah Zuraida binti Kamaruddin:  "}, true, "Puan Hajah Zuraida binti Kamaruddin"},
		{"case #4", args{line: "  2. Tuan Wong Hon Wai [Bukit Bendera]  "}, true, "Tuan Wong Hon Wai [Bukit Bendera]"},
		{"case #5", args{line: "  Menteri Wilayah Persekutuan [Tan Sri Datuk Seri Panglima Haji Annuar bin Haji Musa]:  "}, true, "Menteri Wilayah Persekutuan [Tan Sri Datuk Seri Panglima Haji Annuar bin Haji Musa]"},
		//{"case #6", args{line: "   3. Tuan Chong Chieng Jen [Stampin "}, true, "Tuan Chong Chieng Jen [Stampin]"},
		{"case #7", args{line: " Menteri Kesihatan [Datuk Seri Dr. Adham bin Baba]: "}, true, "Menteri Kesihatan [Datuk Seri Dr Adham bin Baba]"},
		{"case #8", args{line: " ■�1030 Tuan Chong Chieng Jen [Stampin]: "}, true, "Tuan Chong Chieng Jen [Stampin]"},
		{"case #9", args{line: " Timbalan Menteri di Jabatan Perdana Menteri (Tugas-tugas Khas) [Datin Mastura binti Mohd Yazid]: "}, true, "Timbalan Menteri di Jabatan Perdana Menteri (Tugas-tugas Khas) [Datin Mastura binti Mohd Yazid]"},
		{"case #10", args{line: " Datin Mastura binti Mohd Yazid: "}, true, "Datin Mastura binti Mohd Yazid"},
		{"case #11", args{line: " Datuk Dr. Haji Zulkifli Mohamad Al-Bakri:  "}, true, "Datuk Dr Haji Zulkifli Mohamad Al-Bakri"},
		{"case #12", args{line: "  2. Puan Hajah Fuziah binti Salleh [Kuantan] "}, true, "Puan Hajah Fuziah binti Salleh [Kuantan]"},
		{"case #13", args{line: " Tuan Hassan bin Abdul Karim [Pasir Gudang]: "}, true, "Tuan Hassan bin Abdul Karim [Pasir Gudang]"},
		{"case #14", args{line: "  Tan Sri Datuk Seri Panglima Haji Annuar bin Haji Musa:  "}, true, "Tan Sri Datuk Seri Panglima Haji Annuar bin Haji Musa"},
		{"negative #1", args{line: "   Satu lagi.  Tuan Yang di-Pertua, ini berkait dengan COVID-19 ini, penting.   "}, false, ""},
		{"negative #2", args{line: " meningkatkan penawaran dan permintaan getah di pasaran dunia melalui tiga mekanisme utama iaitu : (i)  "}, false, ""},
		{"negative #3", args{line: " Adakah tambahan? Oleh sebab COVID ini kita tidak tahu sampai bila Yang Berhormat Menteri. Terima kasih Yang Berhormat Menteri.  "}, false, ""},
		{"negative #4", args{line: " [Sesi Waktu Pertanyaan-pertanyaan Menteri tamat] "}, false, ""},
		{"negative #5", args{line: "  Terima kasih Yang Berhormat Kuala Krai. Silakan Yang Berhormat Menteri.   "}, false, ""},
		{"negative #6", args{line: "  Tuan Yang di-Pertua, KKM Petrajaya juga memobilisasikan pasukan bantuan penasihat teknikal ke Sarawak untuk meninjau dan membantu memberikan pandangan mengenai penularan serta pengurusan jangkitan COVID-19 di sana."}, false, ""},
		{"negative #7", args{line: "  Jadi saya mencadangkan kepada Yang Berhormat Menteri, apalah salah kiranya Yang Berhormat Menteri panggil semua yayasan yang ada, zakat negeri, zakat pusat, YaPEIM, YADIM, JAKIM, MAWIP, Yayasan Dakwah yang semua itu di bawah kepimpinan Yang Berhormat Menteri."}, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := looksLikeRep(strings.Trim(tt.args.line, " "))
			if got != tt.want {
				t.Errorf("looksLikeRep() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("looksLikeRep() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
