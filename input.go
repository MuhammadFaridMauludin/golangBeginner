package main

import (
	"fmt"
)

// Konstanta warna
const (
	Reset  = "\033[0m"  // Reset warna
	Red    = "\033[31m" // Merah
	Green  = "\033[32m" // Hijau
	Yellow = "\033[33m" // Kuning
	Blue   = "\033[34m" // Biru
)
// struct
type Saham struct {
	emiten              string
	labaBersih         float64
	sahamBeredar       float64
	hargaSaham         float64
	eps                 float64
	dividen             float64
	nilaiBukupersaham  float64
	ekuitas            float64
	utang              float64
}

// pointer
func (s *Saham) SetEmiten(emiten string) {
	s.emiten = emiten
}

// input emiten
func inputEmiten() string {
	var emiten string
	fmt.Print(Blue + "Masukkan nama emiten: " + Reset)
	fmt.Scan(&emiten)
	return emiten
}

// input saham beredar
func inputSahamBeredar() float64 {
	var jumlah float64
	for {
		fmt.Print(Blue + "Masukkan jumlah saham beredar: " + Reset)
		fmt.Scan(&jumlah)
		if jumlah <= 0 {
			fmt.Println(Red + "Jumlah saham beredar tidak boleh nol atau negatif. Silakan coba lagi." + Reset)
		} else {
			break
		}
	}
	return jumlah
}

// input Laba bersih
func inputLabaBersih() float64 {
	var jumlah float64
	for {
		fmt.Print(Blue + "Masukkan Laba bersih: " + Reset)
		fmt.Scan(&jumlah)
		if jumlah <= 0 {
			fmt.Println(Red + "Jumlah Laba-Bersih tidak boleh nol atau negatif. Silakan coba lagi." + Reset)
		} else {
			break
		}
	}
	return jumlah
}

// input Eps
func inputEps() float64 {
	var jumlah float64
	for {
		fmt.Print(Blue + "Masukkan Eps saham: " + Reset)
		fmt.Scan(&jumlah)
		if jumlah <= 0 {
			fmt.Println(Red + "Eps saham tidak boleh nol atau negatif. Silakan coba lagi." + Reset)
		} else {
			break
		}
	}
	return jumlah
}

// input Harga saham
func inputHargaSaham() float64 {
	var jumlah float64
	for {
		fmt.Print(Blue + "Masukkan Harga saham saat ini: " + Reset)
		fmt.Scan(&jumlah)
		if jumlah <= 0 {
			fmt.Println(Red + "Harga saham saat ini tidak boleh nol atau negatif. Silakan coba lagi." + Reset)
		} else {
			break
		}
	}
	return jumlah
}

// input dividen
func inputDividen() float64 {
	var jumlah float64
	for {
		fmt.Print(Blue + "Masukkan dividen per saham: " + Reset)
		fmt.Scan(&jumlah)
		if jumlah <= 0 {
			fmt.Println(Red + "Dividen tidak boleh nol atau negatif. Silakan coba lagi." + Reset)
		} else {
			break
		}
	}
	return jumlah
}

// input nilai buku per saham
func inputNilaiBuku() float64 {
	var jumlah float64
	for {
		fmt.Print(Blue + "Masukkan nilai buku per saham: " + Reset)
		fmt.Scan(&jumlah)
		if jumlah <= 0 {
			fmt.Println(Red + "Nilai buku per saham tidak boleh nol atau negatif. Silakan coba lagi." + Reset)
		} else {
			break
		}
	}
	return jumlah
}

// input ekuitas
func inputEkuitas() float64 {
	var jumlah float64
	for {
		fmt.Print(Blue + "Masukkan nilai ekuitas Saham: " + Reset)
		fmt.Scan(&jumlah)
		if jumlah <= 0 {
			fmt.Println(Red + "Nilai ekuitas tidak boleh nol atau negatif. Silakan coba lagi." + Reset)
		} else {
			break
		}
	}
	return jumlah
}

// input utang
func inputUtang() float64 {
	var jumlah float64
	for {
		fmt.Print(Blue + "Masukkan total utang Saham: " + Reset)
		fmt.Scan(&jumlah)
		if jumlah <= 0 {
			fmt.Println(Red + "Total utang tidak boleh nol atau negatif. Silakan coba lagi." + Reset)
		} else {
			break
		}
	}
	return jumlah
}
