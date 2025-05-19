package main

import "fmt"

const NMAX int = 6

type jobbie struct {
	title    string
	id       int
	kodeUnik int
	gaji     int
}

type tabMinat [NMAX]jobbie

//aray untuk pekerjaan berdasarkan minat dan bakat
var industri = []string{"Kuliner", "TI", "Desain", "Kesehatan", "Mekanik", "Pariwisata", "Baca", "Kebun", "Olahraga", "Berkarya", "Jelajah"}
var kuliner = []jobbie{{"Food_Blogger", 0, 11, 4500000}, {"kepala_koki", 0, 11, 9250000}, {"Quality_Control_Makanan", 0, 11, 10000000}, {"Chef", 0, 11, 11500000}, {"Manajer_Produksi_Kuliner", 0, 11, 12500000}, {"Food_Photographer", 0, 11, 13500000}}
var TI = []jobbie{{"IT_Support", 0, 22, 5804075}, {"IT_Project_Manager", 0, 22, 10710269}, {"Programmer", 0, 22, 10000000}, {"UI/UX_Designer", 0, 22, 5000000}, {"Software_Engineer", 0, 22, 14750000}, {"Game_Designer", 0, 22, 7950000}}
var desain = []jobbie{{"Art_Director", 0, 33, 20025000}, {"Creative_Designer", 0, 33, 4588523}, {"Desainer_Marketing_Digital", 0, 33, 4500000}, {"Animator", 0, 33, 6190000}, {"Desainer_UI_bebasis_Data", 0, 33, 5000000}, {"Graphic_Designer", 0, 33, 6190000}}
var kesehatan = []jobbie{{"Perawat", 0, 44, 5500000}, {"Supervisor_Medis", 0, 44, 9000000}, {"Analis_Kesehatan", 0, 44, 7000000}, {"Terapis_Seni", 0, 44, 6000000}, {"Statistik_Kesehatan", 0, 44, 7500000}, {"Desainer_Edukasi_Kesehatan", 0, 44, 5500000}}
var mekanik = []jobbie{{"Sales_Engineer_Otomotif", 0, 55, 7500000}, {"Manajer_Bengkel", 0, 55, 11000000}, {"Analis_Performa_Mesin", 0, 55, 9000000}, {"Desainer_Otomotif", 0, 55, 7000000}, {"Teknisi_Otomotif", 0, 55, 6000000}, {"Modifikator_Kendaraan", 0, 55, 7500000}}
var pariwisata = []jobbie{{"Customer_Relations_Manager", 0, 66, 9000000}, {"Manajer_Hotel", 0, 66, 15000000}, {"Analis_Pariwisata", 0, 66, 7000000}, {"Event_Organizer", 0, 66, 6000000}, {"Analis_Revenue_Hotel", 0, 66, 9000000}, {"Desainer_Interior_Hotel", 0, 66, 7000000}}
var baca = []jobbie{{"Jurnalis", 0, 77, 6000000}, {"Kepala_Redaksi", 0, 77, 9500000}, {"Peneliti_Sastra", 0, 77, 7000000}, {"Penulis", 0, 77, 4500000}, {"Penulis_Edukasi_STEM", 0, 77, 5500000}, {"Penyair_Ilustratif", 0, 77, 4725000}}
var kebun = []jobbie{{"Penyuluh_Pertanian", 0, 88, 4000000}, {"Manajer_Kebun", 0, 88, 16000000}, {"Agronomis", 0, 88, 5280000}, {"Desainer_Lanskap", 0, 88, 9019437}, {"Perencana_Irigasi", 0, 88, 6095000}, {"Fotografer_Alam", 0, 88, 5125000}}
var olahraga = []jobbie{{"Instruktur_Komunitas", 0, 99, 4720820}, {"Pelatih_Komunitas", 0, 99, 5375000}, {"Analis_Kinerja_Atlet", 0, 99, 10145563}, {"Pembuat_Konten_Olahraga", 0, 99, 4700000}, {"Analis_Statistik_Olahraga", 0, 99, 10145563}, {"Desainer_Merchandise_Olahraga", 0, 99, 6190000}}
var berkarya = []jobbie{{"Kurator_Galeri", 0, 110, 12500000}, {"Manajer_Studio_Seni", 0, 110, 21000000}, {"Sejarawan_Seni", 0, 110, 7500000}, {"Seniman_Kontemporer", 0, 110, 7500000}, {"Visualis_Data_Artistik", 0, 110, 6190000}, {"Pelukis", 0, 110, 7500000}}
var jelajah = []jobbie{{"Pemandu_Wisata", 0, 120, 4000000}, {"Manajer_Tur", 0, 120, 21000000}, {"Peneliti_Budaya", 0, 120, 7500000}, {"Travel_Photographer", 0, 120, 5125000}, {"Analis_Data_Wisata", 0, 120, 6190000}, {"Travel_Illustrator", 0, 120, 4725000}}

//array

func main() {
	var nama, ans4 string
	var pekerjaan tabMinat
	var bakat1, bakat2, minat1, minat2 int
	var id, nMenu, i, ans3 int

	menu()
	fmt.Print("What step would you like to do first? ")
	fmt.Scan(&nMenu)

	for nMenu != 6 {
		switch nMenu {
		case 1:
			fmt.Println()
			inputData(&id, &nama, &minat1, &minat2, &bakat1, &bakat2)

		case 2:
			fmt.Println()
			fmt.Println("Berikut Daftar rekomendasi pekerjaan berdasarkan minat dan bakat anda!")
			rekomendasiKarir(&pekerjaan[0], minat1, bakat1)
			rekomendasiKarir(&pekerjaan[1], minat2, bakat2)
			rekomendasiKarir(&pekerjaan[2], minat1, bakat2)
			rekomendasiKarir(&pekerjaan[3], minat2, bakat1)

			for i = 0; i < 4; i++ {
				fmt.Printf("%d. %s\n", i+1, pekerjaan[i].title)
			}

		case 3:
			fmt.Println("Berdasarkan apa Anda ingin mengkategorikan pekerjaan yang dicari?")
			fmt.Println("1. Nama")
			fmt.Println("2. Industri")
			fmt.Scan(&ans3)
			if ans3 == 1 {
				fmt.Println("Pekerjaan apa yang ingin anda jelajahi hari ini?")
				fmt.Scan(&ans4)

				if sequentialSearch(ans4) == "error" {
					fmt.Println("Maaf pekerjaan yang anda cari tidak ada dalam database kami")
				} else {
					fmt.Println(sequentialSearch(ans4))
				}
			} else if ans3 == 2 {
				selectionSortIndustri()
				binSearch()
			}
		case 4:
			kecocokanUser(&pekerjaan)
			selectionSort(&pekerjaan)
			fmt.Println("Berikut daftar pekerjaan yang apling cocok denganmu! ")
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


//untuk menu 4
//selection sort untuk berdasarkan kecocokan, binary untuk gaji

func kecocokanUser(job *tabMinat) {
	var i, n, x int

	n = 4

	fmt.Println()
	fmt.Println("Dari keempat rekomendasi pekerjaan yang kami berikan, mana yang paling anda minati? ")

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

