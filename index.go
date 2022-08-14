package main

import "fmt"

func main()  {
	var nilaiSiswa int
	var namaSiswa string
	var batasNilaiRemedial int
	var arrayNumber[3] int
	// var arr2 = [3]int{20, 11, 23}
	var arr3 = [5]int{1, 2, 3, 4, 5} 

	namaSiswa = "Budi"
	nilaiSiswa = 65
	batasNilaiRemedial = 70
	arrayNumber[0] = 10

	fmt.Printf("%s punya nilai %d \n", namaSiswa, nilaiSiswa)

	if (nilaiSiswa >= batasNilaiRemedial) {
		fmt.Printf("%s lulus \n", namaSiswa)
	} else {
		fmt.Printf("%s tidak lulus, kurang %d dari batas remedial \n", namaSiswa, SelisihRemedial(nilaiSiswa, batasNilaiRemedial))
	}

	var murid1 Murid
	murid1 = Murid{"imran", 20, 80}

	// FizzBuzz(arrayNumber[:])
	// FizzBuzz(arr2[:])
	CekBilanganGenap(arr3[:])
	CheckSiswa(murid1)
}

type Murid struct {
	nama string
	umur int
	nilai int
}

func FizzBuzz(n []int) {
	for i := 0; i < len(n); i++ {
		if n[i] % 2 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Buzz")
		}
	}
}

func CekBilanganGenap(arr []int) {
	for i:= 0; i < len(arr); i++ {
		if  arr[i] % 2 == 0 {
			fmt.Printf("%d merupakan bilangan genap \n", arr[i])
		} else {
			fmt.Printf("%d merupakan bilangan ganjil \n", arr[i])
		}
	}
}

func SelisihRemedial(nilai int, batasRemedial int) int {
	return batasRemedial - nilai
}

func CheckSiswa(murid Murid) {
	fmt.Println(murid.nama)
}