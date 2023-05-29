package main

import (
	"fmt"
	"sort"
)

type Mahasiswa struct {
	No        int
	Nama      string
	Nim       int
	Tgl_Lahir string
	Semester  int
	Ip        int
}

func menambahkanNamaBaruDidepan(nama_mhs []string) {

	result := append([]string{"Reyvaldo"}, nama_mhs...)

	fmt.Println("Nama baru di depan : ", result)

}

func menghapusNamaDiAkhir(nama_mhs []string) {
	var result []string
	if len(nama_mhs) > 0 {
		result = nama_mhs[:len(nama_mhs)-1]
	}

	fmt.Println(result)
}

func menghapusNamaDiAwal(nama_mhs []string) {
	var result []string
	if len(nama_mhs) > 0 {
		result = nama_mhs[1:]
	}

	fmt.Println(result)
}
func menghapusSemuaData(nama_mhs []string) {

	if len(nama_mhs) > 0 {
		nama_mhs = nama_mhs[:0]
	}

	fmt.Println(nama_mhs)
}

func menghapusNamaDitengah(nama_mhs []string) {
	var result []string
	if len(nama_mhs)%2 == 1 {
		length1 := len(nama_mhs) / 2

		result = append(nama_mhs[:length1], nama_mhs[length1+1:]...)
	} else {
		length1 := len(nama_mhs) / 2

		result = append(nama_mhs[:length1-1], nama_mhs[length1+1:]...)
	}

	fmt.Println(result)

}

func tampilkanDataMahasiswa(nama_mhs []string) {

	var mahasiswa = []Mahasiswa{}

	for index, mhs := range nama_mhs {
		currentMhs := []Mahasiswa{{No: index + 1, Nama: mhs}}
		mahasiswa = append(currentMhs, mahasiswa...)
	}

	// Menggunakan fungsi sort.Slice untuk mengurutkan data
	sort.Slice(mahasiswa, func(i, j int) bool {
		return mahasiswa[i].No < mahasiswa[j].No
	})

	fmt.Println("No 	     Nama")

	for _, mhs := range mahasiswa {
		fmt.Println(mhs.No, "	    ", mhs.Nama)
	}

}

func UrutkanDataMahasiswa(nama_mhs []string) {
	sort.Slice(nama_mhs, func(i, j int) bool {
		return nama_mhs[i] < nama_mhs[j]
	})

	fmt.Print(nama_mhs)
}

func tambahNamaBaru(nama_mhs []string, nama string) {
	result := append([]string{nama}, nama_mhs...)

	fmt.Println("before : ", nama_mhs)
	fmt.Println("after : ", result)

}

func main() {
	var nama_mhs = []string{"Irfan", "Satriyo", "Agra", "Aisyah", "Fatma", "aku", "hi", "hello"}

	fmt.Println(nama_mhs[0])
	// menghapusNamaDiAkhir(nama_mhs)
	// menghapusNamaDiAwal(nama_mhs)
	menghapusNamaDitengah(nama_mhs)
	// menghapusSemuaData(nama_mhs)
	// tampilkanDataMahasiswa(nama_mhs)
	// UrutkanDataMahasiswa(nama_mhs)
	// tambahNamaBaru(nama_mhs, "Satria")
}
