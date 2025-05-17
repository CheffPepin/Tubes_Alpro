package main

import "fmt"

func main() {
	var nama, bakat, minat string
	var id, nMenu int

	menu()
	fmt.Print("what step would you like to do first? ")
	fmt.Scan(&nMenu)

	if nMenu == 1 {
		fmt.Println()
		inputData(&id, &nama, &bakat, &minat)
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
}

func inputData(nim *int, name, talent, interest *string) {
	fmt.Print("Name: ")
	fmt.Scan(name)
	fmt.Print("Enter your NIM: ")
	fmt.Scan(nim)
	fmt.Print("Talent: ")
	fmt.Scan(talent)
	fmt.Print("Interest: ")
	fmt.Scan(interest)

}
