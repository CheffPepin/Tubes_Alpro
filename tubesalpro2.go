package main

import "fmt"

const NMAX int = 66

type jobbie struct {
	title    string
	kodeUnik int
	id       int
	gaji     int
}

type dataD struct {
	nama         string
	nim          int
	minat, bakat int
	karir        string
	cucok        float64
}

type tabProfile [NMAX]dataD
type tabMinat [NMAX]jobbie

//aray
var industri = []string{"Kuliner", "TI", "Desain", "Kesehatan", "Mekanik", "Pariwisata", "Baca", "Kebun", "Olahraga", "Berkarya", "Jelajah"}

var TI = []jobbie{{"IT_Support", 25, 22, 5804075}, {"IT_Project_Manager", 26, 22, 10710269}, {"Programmer", 27, 22, 10000000}, {"UI/UX_Designer", 28, 22, 5000000}, {"Software_Engineer", 29, 22, 14750000}, {"Game_Designer", 30, 22, 7950000}}
var desain = []jobbie{{"Art_Director", 37, 33, 20025000}, {"Creative_Designer", 38, 33, 4588523}, {"Desainer_Marketing_Digital", 39, 33, 4500000}, {"Animator", 40, 33, 6190000}, {"Desainer_UI_bebasis_Data", 41, 33, 5000000}, {"Graphic_Designer", 42, 33, 6190000}}
var kesehatan = []jobbie{{"Perawat", 49, 44, 5500000}, {"Supervisor_Medis", 50, 44, 9000000}, {"Analis_Kesehatan", 51, 44, 7000000}, {"Terapis_Seni", 52, 44, 6000000}, {"Statistik_Kesehatan", 53, 44, 7500000}, {"Desainer_Edukasi_Kesehatan", 54, 44, 5500000}}
var mekanik = []jobbie{{"Sales_Engineer_Otomotif", 61, 55, 7500000}, {"Manajer_Bengkel", 62, 55, 11000000}, {"Analis_Performa_Mesin", 63, 55, 9000000}, {"Desainer_Otomotif", 64, 55, 7000000}, {"Teknisi_Otomotif", 65, 55, 6000000}, {"Modifikator_Kendaraan", 66, 55, 7500000}}
var pariwisata = []jobbie{{"Customer_Relations_Manager", 73, 66, 9000000}, {"Manajer_Hotel", 74, 66, 15000000}, {"Analis_Pariwisata", 75, 66, 7000000}, {"Event_Organizer", 76, 66, 6000000}, {"Analis_Revenue_Hotel", 77, 66, 9000000}, {"Desainer_Interior_Hotel", 78, 66, 7000000}}
var baca = []jobbie{{"Jurnalis", 85, 77, 6000000}, {"Kepala_Redaksi", 86, 77, 9500000}, {"Peneliti_Sastra", 87, 77, 7000000}, {"Penulis", 88, 77, 4500000}, {"Penulis_Edukasi_STEM", 89, 77, 5500000}, {"Penyair_Ilustratif", 90, 77, 4725000}}
var kebun = []jobbie{{"Penyuluh_Pertanian", 97, 88, 4000000}, {"Manajer_Kebun", 98, 88, 16000000}, {"Agronomis", 99, 88, 5280000}, {"Desainer_Lanskap", 100, 88, 9019437}, {"Perencana_Irigasi", 101, 88, 6095000}, {"Fotografer_Alam", 102, 88, 5125000}}
var olahraga = []jobbie{{"Instruktur_Komunitas", 109, 99, 4720820}, {"Pelatih_Komunitas", 110, 99, 5375000}, {"Analis_Kinerja_Atlet", 111, 99, 10145563}, {"Pembuat_Konten_Olahraga", 112, 99, 4700000}, {"Analis_Statistik_Olahraga", 123, 99, 10145563}, {"Desainer_Merchandise_Olahraga", 124, 99, 6190000}}
var berkarya = []jobbie{{"Kurator_Galeri", 122, 110, 12500000}, {"Manajer_Studio_Seni", 123, 110, 21000000}, {"Sejarawan_Seni", 124, 110, 7500000}, {"Seniman_Kontemporer", 125, 110, 7500000}, {"Visualis_Data_Artistik", 126, 110, 6190000}, {"Pelukis", 127, 110, 7500000}}
var jelajah = []jobbie{{"Pemandu_Wisata", 132, 120, 4000000}, {"Manajer_Tur", 133, 120, 21000000}, {"Peneliti_Budaya", 134, 120, 7500000}, {"Travel_Photographer", 135, 120, 5125000}, {"Analis_Data_Wisata", 136, 120, 6190000}, {"Travel_Illustrator", 137, 120, 4725000}}
var kuliner = []jobbie{{"Food_Blogger", 13, 11, 4500000}, {"Kepala_koki", 14, 11, 9250000}, {"Quality_Control_Makanan", 15, 11, 10000000}, {"Chef", 16, 11, 11500000}, {"Manajer_Produksi_Kuliner", 17, 11, 12500000}, {"Food_Photographer", 18, 11, 13500000}}

func main() {
	var ans4 string
	var pekerjaan, dataPekerjaan tabMinat
	var profile tabProfile
	var nMenu, i, ans3, min, max, N, ans5, ans6 int
	var nPekerjaan int

	menu()
	fmt.Print("What step would you like to do first? ")
	fmt.Scan(&nMenu)

	for nMenu != 7 {
		switch nMenu {
		case 1:
			fmt.Println()
			fmt.Println("INPUT YOUR DATA")
			fmt.Println("----------------------")
			inputData(&profile)

		case 2:
			fmt.Println()
			fmt.Println("Berikut Daftar rekomendasi pekerjaan berdasarkan minat dan bakat anda!")
			fmt.Println("-----------------------------------------------------------------------")
			rekomendasiKarir(&pekerjaan[0], profile[0].minat, profile[0].bakat)
			rekomendasiKarir(&pekerjaan[1], profile[1].minat, profile[1].bakat)
			rekomendasiKarir(&pekerjaan[2], profile[0].minat, profile[1].bakat)
			rekomendasiKarir(&pekerjaan[3], profile[1].minat, profile[0].bakat)
			nPekerjaan = 4

			for i = 0; i < nPekerjaan; i++ {
				profile[i].karir = pekerjaan[i].title
			}

			fmt.Printf("| %-3s | %-25s |\n", "No", "Rekomendasi Pekerjaan")
			fmt.Println("|-----|---------------------------|")

			for i = 0; i < nPekerjaan; i++ {
				fmt.Printf("| %-3d | %-25s |\n", i+1, pekerjaan[i].title)
			}

		case 3:
			fmt.Println("Kategori Pekerjaan")
			fmt.Println("------------------")
			fmt.Println("1. Nama")
			fmt.Println("2. Industri")
			fmt.Println("------------------")
			fmt.Print("Berdasarkan apa Anda ingin mengkategorikan pekerjaan yang dicari? ")
			fmt.Scan(&ans3)
			if ans3 == 1 {
				var result jobbie
				fmt.Println()
				fmt.Println("Kategori pekerjaan by Nama")
				fmt.Println("---------------------------")
				fmt.Print("Pekerjaan apa yang ingin anda jelajahi hari ini?")
				fmt.Scan(&ans4)
				fmt.Println()

				sequentialSearch(ans4, &result)

				if result.title == "error" {
					fmt.Println("Maaf pekerjaan yang anda cari tidak ada dalam database kami")
				} else {
					fmt.Println("Berikut adalah data dari pekerjaan yang anda cari!")
					fmt.Println("--------------------------------------------------")
					fmt.Printf("Nama pekerjaan: %s\n", result.title)
					fmt.Printf("Rata-Rata Gaji: Rp %d\n", result.gaji)
				}
			} else if ans3 == 2 {
				fmt.Println()
				selectionSortIndustri()
				binSearch()
			}
		case 4:
			fmt.Println("Daftar Rekomendasi Karir")
			fmt.Println("-------------------------")
			fmt.Println("1. Kecocokan minat dan bakat")
			fmt.Println("2. Gaji")
			fmt.Print("Berdasarkan apa Anda ingin membuat list pekerjaan yang cocok dengan anda?")
			fmt.Scan(&ans5)

			if ans5 == 1 {
				fmt.Println()
				fmt.Println("Rekomendasi Karir by Minat dan Bakat")
				fmt.Println("-------------------------------------")
				kecocokanUser(&pekerjaan, nPekerjaan)
				selectionSort(&pekerjaan, nPekerjaan)
				fmt.Println("Berikut daftar pekerjaan yang paling cocok denganmu! ")
				fmt.Printf("| %-3s | %-25s |\n", "No", "Rekomendasi Pekerjaan")
				fmt.Println("|-----|---------------------------|")

				for i = 0; i < nPekerjaan; i++ {
					fmt.Printf("| %-3d | %-25s |\n", i+1, pekerjaan[i].title)
				}

			} else if ans5 == 2 {
				fmt.Println()
				fmt.Println("Rekomendasi Karir by Gaji")
				fmt.Println("----------------------------------------")
				fmt.Println("Masukkan range gaji yang ingin anda cari")
				fmt.Print("Range gaji minimum: ")
				fmt.Scan(&min)
				fmt.Print("Range gaji maximum: ")
				fmt.Scan(&max)

				isInRange(&dataPekerjaan, &N, min, max)
				InsertionSort(&dataPekerjaan, N)

				fmt.Println("Berikut seluruh data pekerjaan yang sesuai dengan range gaji yang anda inginkan!")
				fmt.Printf("| %-2s | %-30s | %-12s |\n", "No", "Pekerjaan", "Gaji")
				fmt.Println("------------------------------------------------------")
				for i = 0; i < N; i++ {
					fmt.Printf("| %-2d | %-30s | %-12d |\n", i+1, dataPekerjaan[i].title, dataPekerjaan[i].gaji)
				}
				fmt.Println("------------------------------------------------------")
			}
		case 5:
			statistik(profile, nPekerjaan, pekerjaan)

		case 6:
			fmt.Println("Apakah Anda ingin mengubah data rekomendasi Pekerjaan Anda atau mereset?")
			fmt.Println("1. Mengubah")
			fmt.Println("2. Reset")
			fmt.Print("Jawabanmu: ")
			fmt.Scan(&ans6)
			fmt.Println()

			if ans6 == 1 {
				editPekerjaan(&profile, &nPekerjaan)
				updatePekerjaan(profile, nPekerjaan, &pekerjaan)

			} else {
				reset(&profile, &nPekerjaan, &pekerjaan)
			}

			for i = 0; i < nPekerjaan; i++ {
				fmt.Printf("%d. %s\n", i+1, profile[i].karir)
			}

		case 7:
			hasilQ(pekerjaan, profile, nPekerjaan)
		}

		fmt.Println()
		menu()
		fmt.Print("Apa langkah selanjutnya? ")
		fmt.Scan(&nMenu)
		fmt.Println()
	}

}

//untuk menu 1

func menu() {
	fmt.Println("_______________________________\n")
	fmt.Println("LETS FIND YOUR SUITABLE CARRER")
	fmt.Println("_______________________________")
	fmt.Println("Pilih langkah Anda:")
	fmt.Println("1. Masukkan data Anda")
	fmt.Println("2. Temukan rekomendasi pekerjaan Anda")
	fmt.Println("3. Cari pekerjaan impian Anda")
	fmt.Println("4. Daftar rekomendasi karier Anda")
	fmt.Println("5. Statistik Anda")
	fmt.Println("6. Ubah data Anda")
	fmt.Println("7. Exit")
}

func inputData(D *tabProfile) {
	var i int

	fmt.Print("Nama: ")
	fmt.Scan(&D[i].nama)

	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&D[i].nim)

	daftarMinat()
	fmt.Print("Jawabanmu: ")
	//for i = 0; i < 2; i++ {
	fmt.Scan(&D[0].minat, &D[1].minat)
	//}

	daftarBakat()
	fmt.Print("Jawabanmu: ")
	//for i = 0; i < 2; i++ {
	fmt.Scan(&D[0].bakat, &D[1].bakat)
	//}
}

func daftarMinat() {
	fmt.Println()
	fmt.Println("Dari daftar minat dan bakat dibawah ini, pilihlah 2 yang sesuai dengan Anda")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("Minat: ")
	fmt.Println("1.  Minat dalam bidang kuliner dan memasak")
	fmt.Println("2.  Minat dalam bidang teknologi informasi dan pemroggraman")
	fmt.Println("3.  Minat dalam desain grafis dan seni digital")
	fmt.Println("4.  Minat dalam kebidanan atau perawatan Kesehatan")
	fmt.Println("5.  Minat dalam mekanik dan perbaikan otomotif")
	fmt.Println("6.  Minat dalam bidang pariwisatan dan perhotelan")
	fmt.Println("7.  Minat dalam membaca buku dan menulis puisi secara hobi")
	fmt.Println("8.  Minat dalam berkebun dan berkegiatan di alam terbuka")
	fmt.Println("9.  Minat dalam bermain olahraga secara santai tanpa ambisi kompetisi")
	fmt.Println("10. Minat dalam melukis atau membuat karya seni sebagai hobi")
	fmt.Println("11. Minat dalam melakukan perjalanan dan menjelajahi tempat-tempat baru")
}

func daftarBakat() {
	fmt.Println()
	fmt.Println("Bakat: ")
	fmt.Println("1. Kemampuan komunikasi yang baik")
	fmt.Println("2. Kemampuan kepemimpinan")
	fmt.Println("3. Kemampuan analitis")
	fmt.Println("4. Kreativitas")
	fmt.Println("5. Bakat Matematika")
	fmt.Println("6. Bakat seni")

}

//untuk menu 2

func rekomendasiKarir(job *jobbie, nMinat, nBakat int) {

	fmt.Println(nMinat, nBakat, kuliner[nBakat-1])

	if nMinat == 1 {
		*job = kuliner[nBakat-1]
	} else if nMinat == 2 {
		*job = TI[nBakat-1]
	} else if nMinat == 3 {
		*job = desain[nBakat-1]
	} else if nMinat == 4 {
		*job = kesehatan[nBakat-1]
	} else if nMinat == 5 {
		*job = mekanik[nBakat-1]
	} else if nMinat == 6 {
		*job = pariwisata[nBakat-1]
	} else if nMinat == 7 {
		*job = baca[nBakat-1]
	} else if nMinat == 8 {
		*job = kebun[nBakat-1]
	} else if nMinat == 9 {
		*job = olahraga[nBakat-1]
	} else if nMinat == 10 {
		*job = berkarya[nBakat-1]
	} else if nMinat == 11 {
		*job = jelajah[nBakat-1]
	} else {
		job.title = "Try Again"
	}

}

//untuk menu 3
// sequential search untuk jalur karir

func sequentialSearch(n string, idx *jobbie) {
	var i int

	idx.title = "error"

	for i = 0; i < 6; i++ {
		if kuliner[i].title == n {
			*idx = kuliner[i]
		}
	}

	if idx.title == "error" {
		for i = 0; i < 6; i++ {
			if TI[i].title == n {
				*idx = TI[i]
			}
		}
	}

	if idx.title == "error" {
		for i = 0; i < 6; i++ {
			if desain[i].title == n {
				*idx = desain[i]
			}
		}
	}

	if idx.title == "error" {
		for i = 0; i < 6; i++ {
			if kesehatan[i].title == n {
				*idx = kesehatan[i]
			}
		}
	}

	if idx.title == "error" {
		for i = 0; i < 6; i++ {
			if mekanik[i].title == n {
				*idx = mekanik[i]
			}
		}
	}

	if idx.title == "error" {
		for i = 0; i < 6; i++ {
			if pariwisata[i].title == n {
				*idx = pariwisata[i]
			}
		}
	}

	if idx.title == "error" {
		for i = 0; i < 6; i++ {
			if kebun[i].title == n {
				*idx = kebun[i]
			}
		}
	}

	if idx.title == "error" {
		for i = 0; i < 6; i++ {
			if olahraga[i].title == n {
				*idx = olahraga[i]
			}
		}
	}

	if idx.title == "error" {
		for i = 0; i < 6; i++ {
			if berkarya[i].title == n {
				*idx = berkarya[i]
			}
		}
	}

	if idx.title == "error" {
		for i = 0; i < 6; i++ {
			if jelajah[i].title == n {
				*idx = jelajah[i]
			}
		}
	}
}

//sorting list industri
func selectionSortIndustri() {
	var idxMin, i, pass int
	var temp string

	pass = 0
	for pass < 10 {
		idxMin = pass
		i = pass + 1

		for i < 11 {
			if industri[i] < industri[idxMin] {
				idxMin = i
			}
			i++
		}

		temp = industri[idxMin]
		industri[idxMin] = industri[pass]
		industri[pass] = temp

		pass++
	}
}

//binary search untuk pekerjaan sesuai industri
func binSearch() {
	var left, mid, right int
	var i, j int
	var d string

	fmt.Println("Kategori pekerjaan by Industri")
	fmt.Println("------------------------------")
	fmt.Println("Berikut daftar Industri:")
	for i < 11 {
		fmt.Printf("%d. %s\n", i+1, industri[i])
		i++
	}
	fmt.Print("Apa industri yang Anda cari (contoh: Baca)? ")
	fmt.Scan(&d)
	fmt.Println()

	left = 0
	right = 10
	mid = (left + right) / 2

	for left <= right && industri[mid] != d {
		if d < industri[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}

	if mid > -1 && industri[mid] == d {
		fmt.Printf("Industri yang anda pilih adalah %s, berikut pekerjaannya:\n", d)
		i = 0
		if d == "Kuliner" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, kuliner[i].title)
				i++
			}
		} else if d == "TI" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, TI[i].title)
				i++
			}
			fmt.Printf("%d. %s\n", i+1, jelajah[4].title)
		} else if d == "Desain" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, desain[i].title)
				i++
			}
			for j = 3; j <= 5; j++ {
				fmt.Printf("%d. %s\n", i+1, berkarya[j].title)
				i++
			}
			fmt.Printf("%d. %s\n", i+1, jelajah[5].title)
		} else if d == "Kesehatan" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, kesehatan[i].title)
				i++
			}
		} else if d == "Mekanik" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, mekanik[i].title)
				i++
			}
		} else if d == "Pariwisata" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, pariwisata[i].title)
				i++
			}
			for j = 1; j < 4; j++ {
				fmt.Printf("%d. %s\n", i+1, berkarya[j].title)
				i++
			}
		} else if d == "Baca" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, baca[i].title)
				i++
			}
		} else if d == "Kebun" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, kebun[i].title)
				i++
			}
		} else if d == "Olahraga" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, olahraga[i].title)
				i++
			}
		} else if d == "Berkarya" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, berkarya[i].title)
				i++
			}
			fmt.Printf("%d. %s\n", i+1, jelajah[3].title)

		} else if d == "Jelajah" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, jelajah[i].title)
				i++
			}
		}
	} else {
		fmt.Println("Industri tidak ditemukan.")
	}
}

//untuk menu 4
//selection sort untuk berdasarkan kecocokan, binary untuk gaji

func kecocokanUser(job *tabMinat, y int) {
	var i, n, x int

	n = y

	fmt.Println("Dari keempat rekomendasi pekerjaan yang kami berikan, mana yang paling anda minati? ")

	for i = 0; i < y; i++ {
		fmt.Printf("%d. %s\n", i+1, job[i].title)
	}

	for i = 0; i < y; i++ {
		fmt.Printf("Urut no %d? ", i+1)
		fmt.Scan(&x)
		job[x-1].id = job[x-1].id + n
		n--
	}

}

func selectionSort(job *tabMinat, n int) {
	var idxMax, i, pass int
	var temp jobbie

	pass = 1

	for pass < n {
		idxMax = pass - 1
		i = pass

		for i < n {
			if (job[i].kodeUnik - job[i].id) > (job[idxMax].kodeUnik - job[idxMax].id) {
				idxMax = i
			}
			i++
		}

		temp = job[idxMax]
		job[idxMax] = job[pass-1]
		job[pass-1] = temp

		pass++

	}

}

//insertion sort untuk memilah gaji yang masuk range dan mengurutkannya

func isInRange(masukRange *tabMinat, n *int, batasBawah, batasAtas int) {
	var i int
	*n = 0

	for i = 0; i < 6; i++ {
		if kuliner[i].gaji >= batasBawah && kuliner[i].gaji < batasAtas {
			masukRange[*n] = kuliner[i]
			*n++
		}
	}

	for i = 0; i < 6; i++ {
		if TI[i].gaji >= batasBawah && TI[i].gaji < batasAtas {
			masukRange[*n] = TI[i]
			*n++
		}
	}

	for i = 0; i < 6; i++ {
		if desain[i].gaji >= batasBawah && desain[i].gaji < batasAtas {
			masukRange[*n] = desain[i]
			*n++
		}
	}

	for i = 0; i < 6; i++ {
		if kesehatan[i].gaji >= batasBawah && kesehatan[i].gaji < batasAtas {
			masukRange[*n] = kesehatan[i]
			*n++
		}
	}

	for i = 0; i < 6; i++ {
		if mekanik[i].gaji >= batasBawah && mekanik[i].gaji < batasAtas {
			masukRange[*n] = mekanik[i]
			*n++
		}
	}

	for i = 0; i < 6; i++ {
		if pariwisata[i].gaji >= batasBawah && pariwisata[i].gaji < batasAtas {
			masukRange[*n] = pariwisata[i]
			*n++
		}
	}

	for i = 0; i < 6; i++ {
		if olahraga[i].gaji >= batasBawah && olahraga[i].gaji < batasAtas {
			masukRange[*n] = olahraga[i]
			*n++
		}
	}

	for i = 0; i < 6; i++ {
		if baca[i].gaji >= batasBawah && baca[i].gaji < batasAtas {
			masukRange[*n] = baca[i]
			*n++
		}
	}

	for i = 0; i < 6; i++ {
		if jelajah[i].gaji >= batasBawah && jelajah[i].gaji < batasAtas {
			masukRange[*n] = jelajah[i]
			*n++
		}
	}

	for i = 0; i < 6; i++ {
		if kebun[i].gaji >= batasBawah && kebun[i].gaji < batasAtas {
			masukRange[*n] = kebun[i]
			*n++
		}
	}

	for i = 0; i < 6; i++ {
		if berkarya[i].gaji >= batasBawah && berkarya[i].gaji < batasAtas {
			masukRange[*n] = berkarya[i]
			*n++
		}
	}
}

func InsertionSort(masukRange *tabMinat, n int) {
	var i, pass int
	var temp jobbie

	pass = 1

	for pass < n {
		i = pass
		temp = masukRange[pass]

		for i > 0 && masukRange[i-1].gaji > temp.gaji {
			masukRange[i] = masukRange[i-1]
			i--
		}

		masukRange[i] = temp
		pass++
	}

}

func editPekerjaan(p *tabProfile, n *int) {
	var ansD int
	fmt.Println("-------------------")
	fmt.Println("Edit Data Pekerjaan")
	fmt.Println("-------------------")
	fmt.Println("1. Menghapus")
	fmt.Println("2. Menambahkan")
	fmt.Print("Apa yang Anda ingin lakukan kepada list rekomendasi pekerjaan?")
	fmt.Scan(&ansD)
	if ansD == 1 {
		hapusP(p, n)
	} else if ansD == 2 {
		tambahP(p, n)
	}
}

func hapusP(p *tabProfile, n *int) {
	var i, hapus int

	fmt.Println("Pekerjaan berapa yang Anda ingin hapus?")

	for i = 0; i < *n; i++ {
		fmt.Printf("%d. %s\n", i+1, p[i].karir)
	}

	fmt.Scan(&hapus)

	if hapus < *n {
		i = hapus - 1
		for i < *n {
			p[i] = p[i+1]
			i++
		}
		*n = *n - 1
	} else {
		*n = *n - 1
	}
}

func tambahP(p *tabProfile, n *int) {
	var i, j int
	var apyh bool
	var newJob string
	fmt.Println("Masukkan maximal 3 pekerjaan baru, jika sudah cukup masukkan #")

	i = 1

	fmt.Scan(&newJob)
	for newJob != "#" && i <= 2 {
		j = 0
		apyh = false
		for j < *n && apyh == false {
			apyh = newJob != p[j].karir
			j++
		}

		if apyh == true {
			p[*n].karir = newJob
			*n = *n + 1

			i++

			fmt.Scan(&newJob)
		}

	}
}

// menu 6

func updatePekerjaan(p tabProfile, n int, job *tabMinat) {
	var i, j int
	var temp jobbie

	temp.title = "not found"
	i = 0

	for i < n {
		for j = 0; j < 6; j++ {
			if p[i].karir == kuliner[j].title {
				temp = kuliner[j]
				fmt.Println(temp)
			}
		}

		if temp.title == "not found" {
			for j = 0; j < 6; j++ {
				if p[i].karir == TI[j].title {
					temp = TI[j]
				}
			}
		}

		if temp.title == "not found" {
			for j = 0; j < 6; j++ {
				if p[i].karir == desain[j].title {
					temp = desain[j]
				}
			}
		}

		if temp.title == "not found" {
			for j = 0; j < 6; j++ {
				if p[i].karir == kesehatan[j].title {
					temp = kesehatan[j]
				}
			}
		}

		if temp.title == "not found" {
			for j = 0; j < 6; j++ {
				if p[i].karir == mekanik[j].title {
					temp = mekanik[j]
				}
			}
		}

		if temp.title == "not found" {
			for j = 0; j < 6; j++ {
				if p[i].karir == pariwisata[j].title {
					temp = pariwisata[j]
				}
			}
		}

		if temp.title == "not found" {
			for j = 0; j < 6; j++ {
				if p[i].karir == baca[j].title {
					temp = baca[j]
				}
			}
		}

		if temp.title == "not found" {
			for j = 0; j < 6; j++ {
				if p[i].karir == kebun[j].title {
					temp = kebun[j]
				}
			}
		}

		if temp.title == "not found" {
			for j = 0; j < 6; j++ {
				if p[i].karir == olahraga[j].title {
					temp = olahraga[j]
				}
			}
		}

		if temp.title == "not found" {
			for j = 0; j < 6; j++ {
				if p[i].karir == berkarya[j].title {
					temp = berkarya[j]
				}
			}
		}

		if temp.title == "not found" {
			for j = 0; j < 6; j++ {
				if p[i].karir == jelajah[j].title {
					temp = jelajah[j]
				}
			}
		}

		if temp.title != "not found" {
			job[i] = temp
		}

		temp.title = "not found"
		i++
	}

}

func reset(p *tabProfile, n *int, job *tabMinat) {
	var i int
	for i = 0; i < *n; i++ {

		p[i].karir = " "
		job[i].title = " "
	}

	for i = 0; i < *n; i++ {

		job[i].id = 0
		job[i].kodeUnik = 0
		job[i].gaji = 0
		*n = 0
	}

}

//menu 5

func statistik(p tabProfile, n int, job tabMinat) {
	var i, pass, ans, j, total int
	var temp jobbie
	var arrUrut tabMinat

	fmt.Println("------------------------------")
	fmt.Println("Statistik Kecocokan Pekerjaan")
	fmt.Println("------------------------------")

	for i = 0; i < n; i++ {
		arrUrut[i] = job[i]
	}

	//sorting
	pass = 1

	for pass < n {
		i = pass
		temp = arrUrut[pass]

		for i > 0 && arrUrut[i-1].gaji > temp.gaji {
			arrUrut[i] = arrUrut[i-1]
			i--
		}

		arrUrut[i] = temp
		pass++
	}

	// ngitung
	for i = 0; i < n; i++ {
		j = 0
		for j < n {
			if arrUrut[i].title == p[j].karir {
				p[j].cucok = (float64(job[j].kodeUnik) - float64(job[j].id)) + float64(i+1)
			}
			j++
		}

	}

	fmt.Println("Apakah anda sudah membuat daftar rekomendasi pekerjaan berdasarkan minat dan bakat?")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Print("Jawabanmu: ")
	fmt.Scan(&ans)

	if ans == 1 {
		if p[0].minat > p[1].minat {
			if p[0].bakat > p[1].bakat {
				total = p[0].minat + p[0].bakat + (n * 2)
			} else {
				total = p[0].minat + p[1].bakat + (n * 2)
			}
		} else {
			if p[1].bakat > p[0].bakat {
				total = p[1].minat + p[0].bakat + (n * 2)
			} else {
				total = p[1].minat + p[1].bakat + (n * 2)
			}
		}

	} else {
		if p[0].minat > p[1].minat {
			if p[0].bakat > p[1].bakat {
				total = p[0].minat + p[0].bakat + n
			} else {
				total = p[0].minat + p[1].bakat + n
			}
		} else {
			if p[1].bakat > p[0].bakat {
				total = p[1].minat + p[0].bakat + n
			} else {
				total = p[1].minat + p[1].bakat + n
			}
		}

	}

	for i = 0; i < n; i++ {
		p[i].cucok = (p[i].cucok / float64(total)) * 100
	}

	fmt.Printf("Nilai total = %d\n", total)
	fmt.Println()
	fmt.Println("-----------------------------------")
	fmt.Println("List Persentase Kecocokan Pekerjaan")
	fmt.Println("-----------------------------------")
	fmt.Println()
	fmt.Printf("| %-3s | %-30s | %-9ss |\n", "No", "Pekerjaan", "Persentase")
	fmt.Println("------------------------------------------------------")
	for i = 0; i < n; i++ {
		fmt.Printf("| %-3d | %-30s | %-9.2f%% |\n", i+1, arrUrut[i].title, p[i].cucok)
	}
	fmt.Println("------------------------------------------------------")
}

func hasilQ(j tabMinat, d tabProfile, n int) {
	var i int

	fmt.Println("Summary")
	fmt.Println("--------")
	fmt.Println("Name :", d[i].nama)
	fmt.Println("NIM :", d[i].nim)
	fmt.Println("Your Job Recommendation: ")
	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, j[i].title)
	}

}
