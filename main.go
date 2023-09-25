package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)

	for {
		fmt.Println("=============================================")
		fmt.Println("Masukkan pilihan yang ingin dipilih:")
		fmt.Println("1. Perkalian")
		fmt.Println("2. Pembagian")
		fmt.Println("3. Pertambahan")
		fmt.Println("4. Pengurangan")
		fmt.Println("=============================================")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Masukkan angka pertama:")
			var angka1 int
			fmt.Scan(&angka1)

			fmt.Println("Masukkan angka kedua:")
			var angka2 int
			fmt.Scan(&angka2)

			hasil := perkalian(angka1, angka2)
			fmt.Println("Hasil perkalian dari", angka1, "dan", angka2, "adalah =", hasil)
		case 2:
			fmt.Println("Masukkan angka pertama:")
			var angka1 int
			fmt.Scan(&angka1)

			fmt.Println("Masukkan angka kedua:")
			var angka2 int
			fmt.Scan(&angka2)

			hasil := pembagian(angka1, angka2)
			fmt.Println("Hasil pembagian dari", angka1, "dan", angka2, "adalah =", hasil)
		case 3:
			fmt.Println("Masukkan angka pertama:")
			var angka1 int
			fmt.Scan(&angka1)

			fmt.Println("Masukkan angka kedua:")
			var angka2 int
			fmt.Scan(&angka2)

			hasil := pertambahan(angka1, angka2)
			fmt.Println("Hasil pertambahan dari", angka1, "dan", angka2, "adalah =", hasil)
		case 4:
			fmt.Println("Masukkan angka pertama:")
			var angka1 int
			fmt.Scan(&angka1)

			fmt.Println("Masukkan angka kedua:")
			var angka2 int
			fmt.Scan(&angka2)

			hasil := perkurangan(angka1, angka2)
			fmt.Println("Hasil perkurangan dari", angka1, "dan", angka2, "adalah =", hasil)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		
		fmt.Print("Apakah ingin lanjut? (y/n): ")
		var lanjut string
		fmt.Scan(&lanjut)

		if lanjut != "y" {
			fmt.Println("Terima Kasih")
			break
		}
	}
}

func perkalian(num1, num2 int) int {
	return num1 * num2
}

func pembagian(num1, num2 int) int {
	return num1 / num2

}

func pertambahan(num1, num2 int) int {
	return num1 + num2
}

func perkurangan(num1, num2 int) int {
	return num1 - num2
}
