package main

import (
	"fmt"
)

// Fungsi utama
func main() {
	defer fmt.Println(Red + "Terima kasih telah menggunakan aplikasi ini!" + Reset) // defer

	for {
		fmt.Println(Red + "===== Aplikasi CLI Kalkulator Saham =====" + Reset)
		fmt.Println(Green + "1. Hitung EPS Saham" + Reset)
		fmt.Println(Green + "2. Hitung Price Earning Ratio" + Reset)
		fmt.Println(Green + "3. Hitung Dividen Yield" + Reset)
		fmt.Println(Green + "4. Hitung PBV Ratio" + Reset)
		fmt.Println(Green + "5. Hitung Roe" + Reset)
		fmt.Println(Green + "6. Hitung Der" + Reset)
		fmt.Println(Yellow + "0. Keluar" + Reset)
		fmt.Println(Red + "=========================================" + Reset)
		fmt.Print(Green + "Pilih menu: " + Reset)

		var pilihMenu int
		_, err := fmt.Scan(&pilihMenu)
		if err != nil {
			fmt.Println(Red + "Input tidak valid. Silakan coba lagi." + Reset)
			continue
		}

		switch pilihMenu {
		case 1:
			hitungEps()
		case 2:
			hitungPer()
		case 3:
			hitungDividenYield()
		case 4:
			hitungPbv()
		case 5:
			hitungRoe()
		case 6:
			hitungDer()
		case 0:
			fmt.Println(Red + "Keluar dari aplikasi..." + Reset)
			return
		default:
			fmt.Println(Red + "Menu tidak valid. Silakan coba lagi." + Reset)
			continue
		}

		//inputan mau lanjut atau berhenti
		var lagi string
		fmt.Print(Green +"Apakah Anda ingin melakukan perhitungan fundamental lagi? (ya/tidak): " + Reset)
		fmt.Scan(&lagi)

		if lagi != "ya" {
			fmt.Println(Green + "Keluar dari aplikasi..." + Reset)
			return
		}
	}
}
