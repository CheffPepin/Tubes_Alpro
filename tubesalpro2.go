package main

import "fmt"

const NMAX int = 6

type jobbie struct {
	title    string
	id       int
	kodeUnik int
}

type tabMinat [NMAX]jobbie

//aray untuk pekerjaan berdasarkan minat dan bakat
var industri = []string {"Kuliner", "TI", "Desain", "Kesehatan", "Mekanik", "Pariwisata", "Baca", "Kebun", "Olahraga", "Berkarya", "Jelajah"}
var kuliner = []string{"Food Blogger", "kepala koki", "Quality Control Makanan", "Chef", "Manajer Produksi Kuliner", "Food Photographer"}
var TI = []string{"IT Support", "IT Project Manager", "Programmer", "UI/UX Designer", "Software Engineer", "Game Designer"}
var desain = []string{"Art Director", "Creative Designer", "Desainer Marketing Digital", "Animator", "Desainer UI bebasis Data", "Graphic Designer"}
var kesehatan = []string{"Perawat", "Supervisor Medis", "Analis Kesehatan", "Terapis Seni", "Statistik Kesehatan", "Desainer Edukasi Kesehatan"}
var mekanik = []string{"Sales Engineer Otomotif", "Manajer Bengkel", "Analis Performa Mesin", "Desainer Otomotif", "Teknisi Otomotif", "Modifikator Kendaraan"}
var pariwisata = []string{"Customer Relations Manager", "Manajer Hotel", "Analis Pariwisata", "Event Organizer", "Analis Revenue Hotel", "Desainer Interior Hotel"}
var baca = []string{"Jurnalis", "Kepala Redaksi", "Peneliti Sastra", "Penulis", "Penulis Edukasi STEM", "Penyair Ilustratif"}
var kebun = []string{"Penyuluh Pertanian", "Manajer Kebun", "Agronomis", "Desainer Lanskap", "Perencana Irigasi", "Fotografer Alam"}
var olahraga = []string{"Instruktur Komunitas", "Pelatih Komunitas", "Analis Kinerja Atlet", "Pembuat Konten Olahraga", "Analis Statistik Olahraga", "Desainer Merchandise Olahraga"}
var berkarya = []string{"Kurator Galeri", "Manajer Studio Seni", "Sejarawan Seni", "Seniman Kontemporer", "Visualis Data Artistik", "Pelukis"}
var jelajah = []string{"Pemandu Wisata", "Manajer Tur", "Peneliti Budaya", "Travel Photographer", "Analis Data Wisata", "Travel Illustrator"}

//array

func main() {
	var nama, ans4 string
	var pekerjaan tabMinat
	var bakat1, bakat2, minat1, minat2 int
	var id, nMenu, i int

	menu()
	fmt.Print("what step would you like to do first? ")
	fmt.Scan(&nMenu)

	for nMenu != 6 {
		switch nMenu {
		case 1:
			fmt.Println()
			inputData(&id, &nama, &minat1, &minat2, &bakat1, &bakat2)

		case 2:
			fmt.Println()
			fmt.Println("berikut Daftar rekomendasi pekerjaan berdasarkan minat dan bakat anda!")
			rekomendasiKarir(&pekerjaan[0], minat1, bakat1)
			rekomendasiKarir(&pekerjaan[1], minat2, bakat2)
			rekomendasiKarir(&pekerjaan[2], minat1, bakat2)
			rekomendasiKarir(&pekerjaan[3], minat2, bakat1)

			for i = 0; i < 4; i++ {
				fmt.Printf("%d. %s\n", i+1, pekerjaan[i].title)
			}

		case 3:
			fmt.Println("Pekerjaan apa yang ingin anda jelajahi hari ini?")
			fmt.Scan(&ans4)

			if sequentialSearch(ans4) == "error" {
				fmt.Println("maaf pekerjaan yang anda cari tidak ada dalam database kami")
			} else {
				fmt.Println(sequentialSearch(ans4))
			}

		case 4:
			kecocokanUser(&pekerjaan)
			selectionSort(&pekerjaan)
			fmt.Println("berikut daftar pekerjaan yang apling cocok denganmu! ")
			for i = 0; i < 4; i++ {
				fmt.Printf("%d. %s\n", i+1, pekerjaan[i].title)
			}
		}

		menu()
		fmt.Print("whats your next step? ")
		fmt.Scan(&nMenu)

	}

}

//untuk menu 1

func menu() {
	fmt.Println("_______________________________\n")
	fmt.Println(" LETS FIND YOUR SUITABLE CAREER")
	fmt.Println("_______________________________")
	fmt.Println("Pick your step")
	fmt.Println("1. Input  your data")
	fmt.Println("2. Find your recomended job")
	fmt.Println("3. Search for your dream job")
	fmt.Println("4. List your recommended carrer")
	fmt.Println("4. Your Statistic")
	fmt.Println("5. Edit Your Data")
	fmt.Println("6. Exit")
}

func inputData(nim *int, name *string, a1, a2, b1, b2 *int) {
	fmt.Print("Nama: ")
	fmt.Scan(name)

	fmt.Print("Masukkan NIM: ")
	fmt.Scan(nim)

	daftarMinat()
	fmt.Print("Your Answer: ")
	fmt.Scan(a1, a2)

	daftarBakat()
	fmt.Print("Your Answer: ")
	fmt.Scan(b1, b2)
}

func daftarMinat() {
	fmt.Println()
	fmt.Println("Dari daftar minat dan bakat dibawah ini, pilihlah 2 yang sesuai dengan Anda")
	fmt.Println("Minat: ")
	fmt.Println("1. Minat dalam bidang kuliner dan memasak")
	fmt.Println("2. Minat dalam bidang teknologi informasi dan pemroggraman")
	fmt.Println("3. Minat dalam desain grafis dan seni digital")
	fmt.Println("4. Minat dalam kebidanan atau perawatan Kesehatan")
	fmt.Println("5. Minat dalam mekanik dan perbaikan otomotif")
	fmt.Println("6. Minat dalam bidang pariwisatan dan perhotelan")
	fmt.Println("7. Minat dalam membaca buku dan menulis puisi secara hobi")
	fmt.Println("8. Minat dalam berkebun dan berkegiatan di alam terbuka")
	fmt.Println("9. Minat dalam bermain olahraga secara santai tanpa ambisi kompetisi")
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

	if nMinat == 1 {
		job.title = kuliner[nBakat-1]
		job.kodeUnik = 11
	} else if nMinat == 2 {
		job.title = TI[nBakat-1]
		job.kodeUnik = 22
	} else if nMinat == 3 {
		job.title = desain[nBakat-1]
		job.kodeUnik = 33
	} else if nMinat == 4 {
		job.title = kesehatan[nBakat-1]
		job.kodeUnik = 44
	} else if nMinat == 5 {
		job.title = mekanik[nBakat-1]
		job.kodeUnik = 55
	} else if nMinat == 6 {
		job.title = pariwisata[nBakat-1]
		job.kodeUnik = 66
	} else if nMinat == 7 {
		job.title = baca[nBakat-1]
		job.kodeUnik = 77
	} else if nMinat == 8 {
		job.title = kebun[nBakat-1]
		job.kodeUnik = 88
	} else if nMinat == 9 {
		job.title = olahraga[nBakat-1]
		job.kodeUnik = 99
	} else if nMinat == 10 {
		job.title = berkarya[nBakat-1]
		job.kodeUnik = 110
	} else if nMinat == 11 {
		job.title = jelajah[nBakat-1]
		job.kodeUnik = 120
	} else {
		job.title = "Try Again"
	}

	job.id = job.kodeUnik + nMinat + nBakat

}

//untuk menu 3
// sequential search untuk jalur karir

func sequentialSearch(n string) string {
	var i int
	var idx string = "error"

	for i = 0; i < 6; i++ {
		if kuliner[i] == n {
			idx = kuliner[i]
		}
	}

	if idx == "error" {
		for i = 0; i < 6; i++ {
			if TI[i] == n {
				idx = TI[i]
			}
		}
	}

	if idx == "error" {
		for i = 0; i < 6; i++ {
			if desain[i] == n {
				idx = desain[i]
			}
		}
	}

	if idx == "error" {
		for i = 0; i < 6; i++ {
			if kesehatan[i] == n {
				idx = kesehatan[i]
			}
		}
	}

	if idx == "error" {
		for i = 0; i < 6; i++ {
			if mekanik[i] == n {
				idx = mekanik[i]
			}
		}
	}

	if idx == "error" {
		for i = 0; i < 6; i++ {
			if pariwisata[i] == n {
				idx = pariwisata[i]
			}
		}
	}

	if idx == "error" {
		for i = 0; i < 6; i++ {
			if kebun[i] == n {
				idx = kebun[i]
			}
		}
	}

	if idx == "error" {
		for i = 0; i < 6; i++ {
			if olahraga[i] == n {
				idx = olahraga[i]
			}
		}
	}

	if idx == "error" {
		for i = 0; i < 6; i++ {
			if berkarya[i] == n {
				idx = berkarya[i]
			}
		}
	}

	if idx == "error" {
		for i = 0; i < 6; i++ {
			if jelajah[i] == n {
				idx = jelajah[i]
			}
		}
	}
	return idx
}

//untuk menu 4
//selection sort untuk berdasarkan kecocokan, binary untuk gaji

func kecocokanUser(job *tabMinat) {
	var i, n, x int

	n = 4

	fmt.Println()
	fmt.Println("dari keempat rekomendasi pekerjaan yang kami berikan, mana yang paling anda minati? ")

	for i = 0; i < 4; i++ {
		fmt.Printf("%d. %s\n", i+1, job[i].title)
	}

	for i = 0; i < 4; i++ {
		fmt.Printf("urut no %d?", i+1)
		fmt.Scan(&x)
		job[x-1].id = job[x-1].id + n
		n--
	}

}

func selectionSort(job *tabMinat) {
	var idxMax, i, pass int
	var temp jobbie

	pass = 1

	for pass < 4 {
		idxMax = pass - 1
		i = pass

		for i < 4 {
			if (job[i].id - job[i].kodeUnik) > (job[idxMax].id - job[idxMax].kodeUnik) {
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
func binSearch(){
	var left, mid, right int 
	var i, j int
	var d string 
	
	fmt.Println("Berikut daftar Industri:")
	for i < 11 {
		fmt.Printf("%d. %s\n", industriN[i], industri[i])
		i++
	}
	fmt.Print("Apa industri yang Anda cari? ")
	fmt.Scan(&d)
	
	
	left = 0
	right = 10 
	mid = (left+right)/2
	
	for left <= right && industri[mid] != d {
		if d < industri[mid] {
			right = mid -1
		} else {
			left = mid + 1
		}
		mid = (left+right)/2
	}
	 
	if mid > -1 && industri[mid] == d {
		fmt.Printf("Industri yang anda pilih adalah %s, berikut pekerjaannya:\n", d)
		i = 0
		if d == "Kuliner" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, kuliner[i])
				i++
			}
		} else if d == "TI" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, TI[i])
				i++
			}
			fmt.Printf("%d. %s\n", i+1, jelajah[4])
		} else if d == "Desain" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, desain[i])
				i++
			}
			for j = 3; j <= 5;j++ {
				fmt.Printf("%d. %s\n", i+1, berkarya[j])
				i++
			}
			fmt.Printf("%d. %s\n", i+1, jelajah[5])
		} else if d == "Kesehatan" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, kesehatan[i])
				i++
			}
		} else if d == "Mekanik" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, mekanik[i])
				i++
			}
		} else if d == "Pariwisata" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, pariwisata[i])
				i++
			}
			for j = 1; j < 4;j++ {
				fmt.Printf("%d. %s\n", i+1, berkarya[j])
				i++
			}
		} else if d == "Baca" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, baca[i])
				i++
			}
		} else if d == "Kebun" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, kebun[i])
				i++
			}
		} else if d == "Olahraga" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, olahraga[i])
				i++
			}
		} else if d == "Berkarya" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, berkarya[i])
				i++
			}
			fmt.Printf("%d. %s\n", i+1, jelajah[3])
			
		} else if d == "Jelajah" {
			for i <= 5 {
				fmt.Printf("%d. %s\n", i+1, jelajah[i])
				i++
			}
		}
	} else {
		fmt.Println("Industri tidak ditemukan.")
	}
	
}



