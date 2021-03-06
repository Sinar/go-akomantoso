package parliament

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
)

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

func Test_cleanExtractedDebaters(t *testing.T) {
	type args struct {
		normalizedReps []string
	}
	tests := []struct {
		name string
		args args
		want []akomantoso.Representative
	}{
		{"case #1", args{[]string{
			" Seorang Ahli ",
			" Tuan Sim Chee Keong [Bukit Mertajam] ",
			" Tuan Gobind Singh Deo [Puchong] ",
			" point of order [Sambil buku Peraturan Mesyuarat] ",
			" Dato’ Seri Anwar bin Ibrahim [Port Dicks�on] ",
			" Tuan Gobind Singh Deo [Puchong] ",
			" Timbalan Yang di-Pertua [Dato’� Mohd Rashid Hasnon] ",
			" Tuan Mohamad bin Sabu [Kota Raja] ",
			" Dato’ Ngeh Koo� Ham [Beruas] ",
			" Tuan Sim Chee Keong [Bukit Mertajam] ",
			" Timbalan Menteri Dalam Negeri I [Dato’ Sri Dr Haji Ismail bin Haji Moham�ed Said] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Timbalan Yang di-Pertua [Dato’ M�ohd Rashid Hasnon] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
		}}, []akomantoso.Representative{},
		},
		{"case #2", args{[]string{
			" Tuan Yang di-Pertua ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Tuan M Kulasegaran [Ipoh Barat] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan M Kulasegaran [Ipoh Barat] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Yang di-Pertua ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Beberapa Ahli ",
			" Tuan Yang di-Pertua ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Puan Kasthuriraani a/p Patto [Batu Kawan] ",
			" Tuan Nik Nazmi bin Nik Ahmad [Setiawangsa] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Tuan Yang di-Pertua ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Tuan Yang di-Pertua ",
			" Tuan Nik Nazmi bin Nik Ahmad [Setiawangsa] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Yang di-Pertua ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" you cannot allow that kind of language [Pembesar suara dimatikan] [Dewan riuh] ",
			" Seorang Ahli ",
			" Tuan Yang di-Pertua ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Lim Guan Eng [Bagan] ",
			" Tuan Yang di-Pertua ",
			" Puan Kasthuriraani a/p Patto [Batu Kawan] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Tuan Yang di-Pertua ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Puan Kasthuriraani a/p Patto [Batu Kawan] ",
			" Berucap tanpa menggunakan pembesar suara] ",
			" Tuan P Prabakaran [Batu] ",
			" Tuan Yang di-Pertua ",
			" Tuan Yang di-Pertua ",
			" Tan Sri Noh bin Haji Omar [Tanjong Karang] ",
			" Seorang Ahli ",
			" Beberapa Ahli ",
			" Puan Kasthuriraani a/p Patto [Batu Kawan] ",
			" Tuan P Prabakaran [Batu] ",
			" Racist remark [Dewan riuh] ",
			" Seorang Ahli ",
			" you have to make your ruling on that… �[Dewan riuh] ",
		}}, []akomantoso.Representative{},
		},
		{"case #3", args{[]string{
			" Tan Sri Datuk Seri Panglima Haji Annuar bin Haji Musa ",
			" Tuan Nik Mohamad Abduh bin Nik Abdul Aziz [Bachok] ",
			" Menteri Kesihatan [Datuk Seri Dr Adham bin Baba]  ",
			" Tuan Yang di-Pertua ",
			" Tuan Nik Mohamad Abduh bin Nik Abdul Aziz [Bachok] ",
			" Datuk Seri Dr Adham bin Baba [Tenggara] ",
			" Puan Nurul Izzah binti Anwar [Permatang Pauh] ",
			" Datuk Wira Dr Mohd Hatta bin Md Ramli [Lumut] ",
			" Menteri Perusahaan Perladangan dan Komoditi [Dato’ Dr Mohd Khairuddin �bin Aman Razali] ",
			" Bismillahi Rahmani Rahim, alhamdulillah  [Membaca sepotong doa] ",
			" Datuk Wira Dr Mohd Hatta bin Md Ramli [Lumut] ",
			" Datuk Seri Shamsul Iskandar @ Yusre bin haji Mohd Akin [Hang Tuah Jaya] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Haji Wan Hassan bin Mohd Ramli [Dungun] ",
			" Menteri Sains, Teknologi dan Inovasi [Tuan Khairy Jamaluddin Abu Bakar] ",
			" ] ",
			" Tuan Haji Wan Hassan bin Mohd Ramli [Dungun] ",
			" Tuan Khairy Jamaluddin Abu Bakar ",
			" Puan Nor Azrina binti Surip [Merbok] ",
			" Datuk Seri Saifuddin Nasution bin Ismail [Kulim-Bandar Baharu] ",
			" Tuan Khairy Jamaluddin Abu Bakar ",
			" Datuk Ignatius Darell Leiking [Penampang] ",
			" Menteri Perdagangan Antarabangsa dan Industri [Dato’ Seri Mohamed �Azmin bin Ali] ",
			" Datuk Ignatius Dorell Leiking [Penampang] ",
			" Dato' Sri Richard Riot anak Jaem [Serian] ",
			" Timbalan Yang di-Pertua [Dato’ �Mohd Rashid Hasnon] ",
			" Menteri di Jabatan Perdana Menteri (Parlimen dan Undang-undang) [Dato’ �Takiyuddin bin Hassan] ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Rashid Hasnon�] ",
			" pg  Perdana Menteri [Tan Sri Dato’ �Sri Haji Muhyiddin bin Md Yassin] ",
			" Menteri Perdagangan Antarabangsa dan Industri [Dato’ Seri Mohamed �Azmin bin Ali] ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Rashid Hasnon�] ",
			" Beberapa Ahli ",
			" Tuan Khalid bin Abd Samad [Shah Alam] ",
			" Dato’ Seri An�war bin Ibrahim [Port Dickson] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Rashid Ha�snon] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Chang Lih Kang [Tanjong Malim] ",
			" Tuan Chang Lih Kang [Tanjong Malim] ",
			" Dr Nik Muhammad Zawawi bin Haji Salleh [Pasir Puteh] ",
			" Seorang Ahli ",
			" Dato’ Seri Anwar bin Ibrahim [Po�rt Dickson] ",
			" Tuan Gobind Singh Deo [Puchong] ",
			" Timbalan Menteri Sumber Manusia [Tuan Haji Awang bin Hashim] ",
			" Tuan Gobind Singh Deo [Puchong] ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Ras�hid Hasnon] ",
			" Datuk Liew Vui Keong [Batu Sapi] ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Rashid H�asnon] ",
			" Menteri di Jabatan Perdana Menteri (Parlimen dan Undang-undang) [Dato’ �Takiyuddin bin Hassan] ",
			" Tuan Sim Chee Keong [Bukit Mertajam] ",
			" Dr Nik Muhammad Zawawi bin Haji Salleh [Pasir Puteh] ",
			" Dato’� Seri Anwar bin Ibrahim [Port Dickson] ",
			" Timbalan Yang di-Pertua [Dato’� Mohd Rashid Hasnon] ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Ras�hid Hasnon] ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Rashid Ha�snon] ",
			" Dato’ Sri� Haji Tajuddin bin Abdul Rahman [Pasir Salak] ",
			" Dato’ Sri Haji Tajudd�in bin Abdul Rahman [Pasir Salak] ",
			" Timbalan Yang di-Pertua [Dato’ M�ohd Rashid Hasnon] ",
			" Dato’ Sri Haji Tajuddin� bin Abdul Rahman [Pasir Salak] ",
			" Tuan Ramkarpal Singh a/l Karpal Singh [Bukit Gelugor] ",
			" Dato’ Seri Anwar bin Ibrahim [Po�rt Dickson] ",
			" Tan Sri Noh bin Haji Omar [Tanjong Karang] ",
			" Tuan Abdul Latiff bin Abdul Rahman [Kuala Krai] ",
			" tgh Dr Nik Muhammad Zawawi bin Haji Salleh [Pasir Puteh] ",
			" Dato’ Seri Mohamed Nazri b�in Abdul Aziz [Padang Rengas] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Rashid Hasnon]� ",
			" Datuk Seri Panglima Haji Mohd Shafie bin Haji Apdal [Semporna] ",
			" Timbalan Yang di-Pertua [Dato’ M�ohd Rashid Hasnon] ",
			" Datuk Seri Panglima Haji Mohd Shafie bin Haji Apdal [Semporna] ",
			" Dato’ Dr Shamsul Anuar bin Nas�arah ",
			" Timbalan Menteri di Jabatan Perdana Menteri (Hal Ehwal Agama) [Ustaz Haji Ahmad Marzuk bin Shaary] ",
			" Tun Dr Mahathir bin Mohamad [Langkawi] ",
			" Tuan Syed Saddiq bin Syed Abdul Rahman [Muar] ",
			" Dato’ Seri Dr� Shahidan bin Kassim [Arau] ",
			" Tun Dr Mahathir bin Mohamad [Langkawi] ",
			" Tuan Khalid bin Abd Samad [Shah Alam] ",
			" Dato’ Sri Bung Moktar bin Radin [Kinabatanga�n] ",
			" Dato’ Seri Dr Shahidan bin Kassim [A�rau] ",
			" Dato’ Seri Dr Shahidan bin Kassim� [Arau] ",
			" Beberapa Ahli ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Dato’ Seri Dr Shahid�an bin Kassim [Arau] ",
			" Dato’� Seri Dr Shahidan bin Kassim [Arau] ",
			" Dato’ Seri Dr Shahidan bin Kassi�m [Arau] ",
			" Tuan Khalid bin Abd Samad [Shah Alam] ",
			" Dato’� Seri Dr Shahidan bin Kassim [Arau] ",
			" Dato’ Seri Dr Shahidan bin Kassim [Ara�u] ",
			" Seorang Ahli ",
			" Dato’ Sri Bung Moktar bin R�adin [Kinabatangan] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tun Dr Mahathir bin Mohamad [Langkawi] ",
			" Dato’ Se�ri Dr Shahidan bin Kassim [Arau] ",
			" Tun Dr Mahathir bin Mohamad [Langkawi] ",
			" Dato’� Mahfuz bin Haji Omar [Pokok Sena] ",
			" Dato’ Sri Bung Moktar bin Radin [Ki�nabatangan] ",
			" Tuan Khalid bin Abd Samad [Shah Alam] ",
			" Dato’ Seri Dr� Shahidan bin Kassim [Arau] ",
			" Dato’ Seri Dr Shahidan bin Kassim [A�rau] ",
			" Dato’ Se�ri Dr Shahidan bin Kassim [Arau] ",
			" Tuan Khalid bin Abd Samad [Shah Alam] ",
			" Dato’ Seri Dr Shahidan b�in Kassim [Arau] ",
			" Dato’ Seri Dr Shahidan bin Kassim� [Arau] ",
			" Tuan Sim Tze Tzin [Bayan Baru] ",
			" Tuan Nik Nazmi bin Nik Ahmad [Setiawangsa] ",
			" Dato’ Seri Dr� Shahidan bin Kassim [Arau] ",
			" tgh Menteri di Jabatan Perdana Menteri (Parlimen dan Undang-undang) [Dato’ �Takiyuddin bin Hassan] ",
			" Dato’ N�geh Koo Ham [Beruas] ",
			" Dato’ Se�ri Utama Haji Mukhriz Tun Dr Mahathir [Jerlun] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Gobind Singh Deo [Puchong] ",
			" Dato’ Ta�kiyuddin bin Hassan ",
			" Tuan Sim Chee Keong [Bukit Mertajam] ",
			" Dato’� Takiyuddin bin Hassan ",
			" Tuan Sim Chee Keong [Bukit Mertajam] ",
			" Tuan Gobind Singh Deo [Puchong] ",
			" Dato’ Ngeh Koo Ha�m [Beruas] ",
			" tgh Menteri di Jabatan Perdana Menteri (Parlimen dan Undang-undang) [Dato’ �Takiyuddin bin Hassan] ",
			" ] ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Rashid Hasnon�] ",
			" Setiausaha ",
			" Dato’ Ngeh Koo Ham [Beruas]� ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Rashid Hasnon�] ",
			" Tuan Sivarasa Rasiah [Sungai Buloh] ",
			" Menteri di Jabatan Perdana Menteri (Parlimen dan Undang-undang) [Dato’ �Takiyuddin bin Hassan] ",
			" Tuan Sivarasa Rasiah [Sungai Buloh] ",
			" Dato’ Takiyuddin bin Hassan� ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Rashid Hasnon�] ",
			" Dato’ Takiyuddin bin Hassan� ",
			" Tuan Chang Lih Kang [Tanjong Malim] ",
			" Dato’ Ngeh Koo Ham [Beruas]� ",
			" Dato’ Takiyuddin bin Hassan� ",
			" Dato’ Ngeh Koo� Ham [Beruas] ",
			" Dato’ Takiyuddin bin Hassan� ",
			" Dato’ Ngeh Koo Ham [Beruas]� ",
			" Dato’ Takiyuddin b�in Hassan ",
			" Dato’ Takiyuddin bin Hassan� ",
			" Tuan M Kulasegaran [Ipoh Barat] ",
			" Timbalan Yang di-Pertua [Dato’ Mohd Rashid Hasnon�] ",
			" � Setiausaha ",
			" Menteri Perdagangan Antarabangsa dan Industri [Dato’ �Seri Mohamed Azmin bin Ali] ",
			" Tuan Wong Hon Wai [Bukit Bendera] ",
			" Tuan Sanisvara Nethaji Rayer a/l Rajaji [Jelutong] ",
			" Tuan Wong Hon Wai [Bukit Bendera] ",
			" Setiausaha ",
			" masuk ke dalam Dewan dan dipakaikan jubah oleh Setiausaha, kemudian menuju ke Kerusi] ",
			" Tuan Wong Hon Wai [Bukit Bendera] ",
			" Dato’ Seri Dr Shahidan bin Kassi�m [Arau] ",
			" Tuan Wong Hon Wai [Bukit Bendera] ",
			" Dato’ Seri Dr� Shahidan bin Kassim [Arau] ",
			" Timbalan Menteri Sumber Manusia [Tuan Haji Awang bin Hashim] ",
			" Tuan Haji Awang bin Hashim ",
			" TUAN YANG DI-PERTUA MENGANGKAT SUMPAH Tuan Yang di-Pertua ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Tuan Sivakumar Varatharaju Naidu [Batu Gajah] ",
			" Tuan Chang Lih Kang [Tanjong Malim] ",
			" Tuan Sivakumar Varatharaju Naidu [Batu Gajah] ",
			" Tuan Yang di-Pertua ",
			" Puan Kasthuriraani a/p Patto [Batu Kawan] ",
			" Tuan Yang di-Pertua ",
			" Puan Kasthuriraani a/p Patto [Batu Kawan] ",
			" shame on you  [Dewan riuh] ",
			" Dato’ Sri Dr Haji Ismail bin Haji Mohamed Said� ",
			" Tuan Yang di-Pertua ",
			" Datuk Wira Dr Mohd Hatta bin Md Ramli [Lumut] ",
			" Timbalan Menteri Dalam Negeri I [Dato’ Sri Dr Haji I�smail bin Haji Mohamed Said] ",
			" Tuan Chang Lih Kang [Tanjong Malim] ",
			" You dare not face competition [Dewan riuh] ",
			" Tuan Yang di-Pertua ",
			" Tuan Khalid bin Abd Samad [Shah Alam] ",
			" Tuan Yang di-Pertua ",
			" Dato’ Sri Dr Haji Ismail bin Haji Mo�hamed Said ",
			" Tuan Mohamad bin Sabu [Kota Raja] ",
			" Tuan Chang Lih Kang [Tanjong Malim] ",
			" Datuk Seri Shamsul Iskandar @ Yusre bin haji Mohd Akin [Hang Tuah Jaya] ",
			" Dato’ Seri Dr Shahidan bin Kassim [A�rau] ",
			" Bye-bye  [Dewan riuh] ",
			" Tan Sri Dato' Sri Haji Muhyiddin bin Md Yassin ",
			" Menteri Perdagangan Antarabangsa dan Industri [Dato' Seri Mohamed Azmin bin Ali] ",
			" Tuan Yang di-Pertua ",
			" USUL MEMILIH TIMBALAN YANG DI-PERTUA  Setiausaha ",
			" Perdana Menteri [Tan Sri Dato’ �Sri Haji Muhyiddin bin Md Yassin] ",
			" Menteri Perdagangan Antarabangsa dan Industri [Dato’ Seri Mohamed �Azmin bin Ali] ",
			" Tuan Yang di-Pertua ",
			" Dato’ Nge�h Koo Ham [Beruas] ",
			" Tuan Yang di-Pertua ",
			" Dato’ Sri Bung Moktar� bin Radin [Kinabatangan] ",
			" Tuan Yang di-Pertua ",
			" Dato’ Sri Bung Moktar bin Radin [Kinabatang�an] ",
			" Datuk Seri Panglima Abdul Azeez bin Abdul Rahim [Baling] ",
			" Dato’� Sri Bung Moktar bin Radin [Kinabatangan] ",
			" USUL MENANGGUH MESYUARAT KAMAR KHAS  Menteri di Jabatan Perdana Menteri (Parlimen dan Undang-undang) [Dato’ �Takiyuddin bin Hassan] ",
			" Menteri Dalam Negeri [Dato’ Seri Hamzah bi�n Zainudin] ",
			" Tuan Yang di-Pertua ",
			" Menteri di Jabatan Perdana Menteri (Parlimen dan Undang-undang) [Dato’ �Takiyuddin bin Hassan] ",
			" Menteri Sumber Manusia [Datuk Seri M Saravanan] ",
			" You are the Speaker of the House You claim to be a Speaker of the House [Dewan riuh] [Pembesar suara dimatikan] ",
			" Menteri Perusahaan Perladangan dan Komoditi [Dato’ Dr Mohd Khairud�din bin Aman Razali] ",
			" Timbalan Menteri Kesihatan I [Dato’ Dr Haji Noor� Azmi bin Ghazali] ",
			" There is no justice for women in the House [Pembesar suara dimatikan] [Dewan riuh] ",
			" Menteri di Jabatan Perdana Menteri (Parlimen dan Undang-undang) [Dato’ �Takiyuddin bin Hassan] ",
			" You are being unfair on us… �[Dewan riuh] ",
			" Tuan P Prabakaran [Batu] ",
			" discrimination [Dewan riuh] ",
			" Tan Sri Noh bin Haji Omar [Tanjong Karang] ",
			" Puan Kasthuriraani a/p Patto [Batu Kawan] ",
			" ptg Menteri di Jabatan Perdana Menteri (Parlimen dan Undang-undang) [Dato’ �Takiyuddin bin Hassan] ",
			" Puan Kasthuriraani a/p Patto [Batu Kawan] ",
			" Beberapa Ahli ",
			" Seorang Ahli ",
			" Ketawa] ",
			" Tuan Yang di-Pertua ",
			" K A N D U N G A N  PEMASYHURAN TUAN YANG DI-PETUA ",
			" USUL MENTERI DALAM NEGERI ",
			" USUL-USUL PERDANA MENTERI ",
			" USUL-USUL ",
		}}, []akomantoso.Representative{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanExtractedDebaters(tt.args.normalizedReps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cleanExtractedDebaters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractDebaters(t *testing.T) {
	type args struct {
		allLines []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"case #1", args{[]string{
			"",
			"",
		}}, []string{
			"bob",
			"joe",
		}},
		{"case #2", args{[]string{
			"",
			"",
		}}, []string{
			"dude",
			"moo",
		}},
		{"case #3", args{[]string{
			"",
			"",
		}}, []string{
			"kim",
			"donald",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractDebaters(tt.args.allLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractDebaters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRepTitle(t *testing.T) {
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
			if got := isRepTitle(tt.args.line); got != tt.want {
				t.Errorf("isRepTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
