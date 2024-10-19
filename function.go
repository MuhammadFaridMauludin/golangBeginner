package main

import "fmt"

// Hitung EPS
func hitungEps() {
	fmt.Println(Green + "Saham apa yang akan kamu hitung EPS nya???" + Reset)
	emiten := inputEmiten()
	saham := &Saham{
		emiten:      emiten,
		sahamBeredar: inputSahamBeredar(),
		labaBersih:   inputLabaBersih(),
	}
	eps := saham.labaBersih / saham.sahamBeredar
	fmt.Printf(Green+"EPS untuk emiten %s yaitu : %.2f\n"+Reset, saham.emiten, eps)
}

// Hitung PER
func hitungPer() {
	fmt.Println(Green + "Saham apa yang akan kamu hitung PER nya???" + Reset)
	emiten := inputEmiten()
	saham := &Saham{
		emiten:      emiten,
		eps: inputEps(),
		hargaSaham: inputHargaSaham(),
	}
	per := saham.hargaSaham / saham.eps
	fmt.Printf(Green+"PER untuk emiten %s yaitu : %.2f\n"+Reset, saham.emiten, per)
}

// Hitung Dividen Yield
func hitungDividenYield() {
	fmt.Println(Green + "Saham apa yang akan kamu hitung Dividen Yield nya???" + Reset)
	emiten := inputEmiten()
	saham := &Saham{
		emiten:      emiten,
		dividen: inputDividen(),
		hargaSaham: inputHargaSaham(),
	}
	divyield := (saham.dividen / saham.hargaSaham) * 100
	fmt.Printf(Green+"Dividen Yield untuk emiten %s yaitu : %.2f\n"+Reset, saham.emiten, divyield)
}

// Hitung PBV
func hitungPbv() {
	fmt.Println(Green + "Saham apa yang akan kamu hitung PBV nya???" + Reset)
	emiten := inputEmiten()
	saham := &Saham{
		emiten:      emiten,
		nilaiBukupersaham: inputNilaiBuku(),
		hargaSaham: inputHargaSaham(),
	}
	nilaibuku := saham.hargaSaham / saham.nilaiBukupersaham
	fmt.Printf(Green+"Nilai buku per saham untuk emiten %s yaitu : %.2f\n"+Reset, saham.emiten, nilaibuku)
}

// Hitung ROE
func hitungRoe() {
	fmt.Println(Green + "Saham apa yang akan kamu hitung ROE nya???" + Reset)
	emiten := inputEmiten()
	saham := &Saham{
		emiten:      emiten,
		labaBersih: inputLabaBersih(),
		ekuitas: inputEkuitas(),
	}
	roe := saham.labaBersih / saham.ekuitas
	fmt.Printf(Green+"ROE untuk emiten %s yaitu : %.2f\n"+Reset, saham.emiten, roe)
}

// Hitung DER
func hitungDer() {
	fmt.Println(Green + "Saham apa yang akan kamu hitung DER nya???" + Reset)
	emiten := inputEmiten()
	saham := &Saham{
		emiten:      emiten,
		utang: inputUtang(),
		ekuitas: inputEkuitas(),
	}
	der := saham.utang / saham.ekuitas
	fmt.Printf(Green+"DER untuk emiten %s yaitu : %.2f\n"+Reset, saham.emiten, der)
}
