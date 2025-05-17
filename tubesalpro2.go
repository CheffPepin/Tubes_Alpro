package main

import "fmt"

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

func rekomendasiKarir(job *string, nMinat, nBakat int) {

	if nMinat == a {
		job = kuliner[nBakat-1]
	} else if nMinat == b {
		job = TI[nBakat-1]
	}else if nMinat == c {
		job = desain[nBakat-1]
	}else if nMinat == d {
		job = kesehatan[nBakat-1]
	}else if nMinat == e {
		job = mekanik[nBakat-1]
	}else if nMinat == f {
		job = pariwisata[nBakat-1]
	}else if nMinat == g {
		job = baca[nBakat-1]
	}else if nMinat == h {
		job = kebun[nBakat-1]
	}else if nMinat == i {
		job = olahraga[nBakat-1]
	}else if nMinat == j {
		job = berkarya[nBakat-1]
	} else if nMinat == k {
		job = jelajah[nBakat-1]
	} else {
		job = "Try Again"
	}

	fmt.Println(job)

}


