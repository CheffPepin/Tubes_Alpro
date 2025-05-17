package main

import "fmt"

var NMAX int = 6

type tabMinat [NMAX] string 

var (
	kuliner     = tabMinat{"Food Blogger / PR Kuliner", "Kepala Koki / Supervisor Dapur", "Quality Control Makanan", "Chef Inovatif", "Manajer Produksi Kuliner", "Food Stylist / Food Photographer"}
	TI          = tabMinat{"Technical Writer / IT Support", "IT Project Manager", "Programmer", "UI/UX Designer", "Data Analyst / Software Engineer", "Game Designer"}
	desain      = tabMinat{"Art Director", "Creative Designer", "Desainer Marketing Digital", "Animator / Motion Graphics Artist", "Desainer UI berbasis Data", "Graphic Designer"}
	kesehatan   = tabMinat{"Perawat", "Supervisor Medis", "Analis Kesehatan", "Terapis Seni", "Statistik Kesehatan", "Desainer Edukasi Kesehatan"}
	mekanik     = tabMinat{"Sales Engineer Otomotif", "Manajer Bengkel", "Analis Performa Mesin", "Desainer Otomotif", "Teknisi Otomotif", "Modifikator Kendaraan"}
	pariwisata  = tabMinat{"Customer Relations Manager", "Manajer Hotel", "Analis Pariwisata", "Event Organizer", "Analis Revenue Hotel", "Desainer Interior Hotel"}
	baca        = tabMinat{"Editor / Jurnalis", "Kepala Redaksi", "Peneliti Sastra", "Penulis", "Penulis Edukasi STEM", "Penyair Ilustratif"}
	kebun       = tabMinat{"Penyuluh Pertanian", "Manajer Kebun", "Agronomis", "Desainer Lanskap", "Perencana Irigasi", "Fotografer Alam"}
	olahraga    = tabMinat{"Instruktur Komunitas", "Pelatih Komunitas", "Analis Kinerja Atlet", "Pembuat Konten Olahraga", "Analis Statistik Olahraga", "Desainer Merchandise Olahraga"}
	berkarya    = tabMinat{"Kurator Galeri", "Manajer Studio Seni", "Sejarawan Seni", "Seniman Kontemporer", "Visualis Data Artistik", "Pelukis"}
	jelajah     = tabMinat{"Pemandu Wisata", "Manajer Tur", "Peneliti Budaya", "Travel Photographer", "Analis Data Wisata", "Travel Illustrator"}
)

func main() {
	var nama string
	var bakat1, bakat2, minat1, minat2 int
	var id, nMenu int

	menu()
	fmt.Print("what step would you like to do first? ")
	fmt.Scan(&nMenu)

	if nMenu == 1 {
		fmt.Println()
		inputData(&id, &nama, &minat1, &minat2, &bakat1, &bakat2)
	} else if nMenu == 2 {
		fmt.Println()
		
	}

}

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
	fmt.Println("5. Exit")
}

func inputData(nim *int, name *string, a1, a2, b1, b2 *int ) {
	fmt.Print("Nama: ")
	fmt.Scan(name)

	fmt.Print("Masukkan NIM: ")
	fmt.Scan(nim)

	daftarMinat()
	fmt.Scan(&a1, &a2)

	daftarBakat()
	fmt.Scan(&b1, &b2)
}

func daftarMinat() {
	fmt.Println("Dari daftar minat dan bakat dibawah ini, pilihlah 2 yang sesuai dengan Anda")
	fmt.Println("Minat: ")
	fmt.Println("a. Minat dalam bidang kuliner dan memasak")
	fmt.Println("b. Minat dalam bidang teknologi informasi dan pemroggraman")
	fmt.Println("c. Minat dalam desain grafis dan seni digital")
	fmt.Println("d. Minat dalam kebidanan atau perawatan Kesehatan")
	fmt.Println("e. Minat dalam mekanik dan perbaikan otomotif")
	fmt.Println("f. Minat dalam bidang pariwisatan dan perhotelan")
	fmt.Println("g. Minat dalam membaca buku dan menulis puisi secara hobi")
	fmt.Println("h. Minat dalam berkebun dan berkegiatan di alam terbuka")
	fmt.Println("i. Minat dalam bermain olahraga secara santai tanpa ambisi kompetisi")
	fmt.Println("j. Minat dalam melukis atau membuat karya seni sebagai hobi")
	fmt.Println("k. Minat dalam melakukan perjalanan dan menjelajahi tempat-tempat baru")
}

func daftarBakat() {
	fmt.Println("Bakat: ") 
	fmt.Println("1. Kemampuan komunikasi yang baik")
	fmt.Println("2. Kemampuan kepemimpinan")
	fmt.Println("3. Kemampuan analitis")
	fmt.Println("4. Kreativitas")
	fmt.Println("5. Bakat Matematika")
	fmt.Println("6. Bakat seni")

}

func rekomendasiKarir(job *string, nMinat string, nBakat int) {
	var daftar tabMinat

	if nMinat == "a" {
		daftar = kuliner
	} else if nMinat == "b" {
		daftar = TI
	} else if nMinat =="c" {
		daftar = desain
	} else if nMinat =="d"{
		daftar = kesehatan
	} else if nMinat =="e" {
		daftar = mekanik
	} else if nMinat =="f" {
		daftar = pariwisata
	} else if nMinat =="g" {
		daftar = baca
	} else if nMinat =="h" {
		daftar = kebun
	} else if nMinat =="i" {
		daftar = olahraga
	} else if nMinat == "j" {
		daftar = berkarya
	} else if nMinat =="k" {
		daftar = jelajah
	} else {
		*job = "Try Again"
	}

	if nBakat >= 1 && nBakat <= NMAX {
		*job = daftar[nBakat-1]
	}
}



