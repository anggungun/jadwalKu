//kelompok 7 (anggun, nares, elsa)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Jadwal struct {
	Kode       string
	Matkul     string
	Dosen      string
	Hari       string
	JamMulai   string
	JamSelesai string
	Ruangan    string
}

var dataJadwal []Jadwal
var input = bufio.NewReader(os.Stdin)

func main() {
	var pilih int

	for pilih != 10 {
		fmt.Println("\n===== JADWALKU =====")
		fmt.Println("1. Tambah Jadwal")
		fmt.Println("2. Tampilkan Jadwal")
		fmt.Println("3. Ubah Jadwal")
		fmt.Println("4. Hapus Jadwal")
		fmt.Println("5. Sequential Search")
		fmt.Println("6. Binary Search")
		fmt.Println("7. Selection Sort")
		fmt.Println("8. Insertion Sort")
		fmt.Println("9. Statistik")
		fmt.Println("10. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			tambahJadwal()

		case 2:
			tampilkanJadwal()

		case 3:
			ubahJadwal()

		case 4:
			hapusJadwal()

		case 5:
			sequentialSearch()

		case 6:
			binarySearch()

		case 7:
			selectionSort()

		case 8:
			insertionSort()

		case 9:
			statistik()

		case 10:
			fmt.Println("Program selesai")

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func inputString() string { //biar bisa make spasi

	text, _ := input.ReadString('\n')

	return strings.TrimSpace(text)
}

func tambahJadwal() { //nambahd ata
	var data Jadwal

	fmt.Print("Kode Matkul : ")
	fmt.Scanln(&data.Kode)

	fmt.Print("Nama Matkul : ")
	data.Matkul = inputString()

	fmt.Print("Nama Dosen : ")
	data.Dosen = inputString()

	fmt.Print("Hari : ")
	fmt.Scanln(&data.Hari)
	fmt.Print("Jam Mulai : ")
	fmt.Scanln(&data.JamMulai)

	fmt.Print("Jam Selesai : ")
	fmt.Scanln(&data.JamSelesai)

	fmt.Print("Ruangan : ")
	fmt.Scanln(&data.Ruangan)

	dataJadwal = append(dataJadwal, data)

	fmt.Println("Data berhasil ditambahkan")
}

func tampilkanJadwal() { //view
	if len(dataJadwal) == 0 {

		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < len(dataJadwal); i++ {
		fmt.Println("\n----------------")
		fmt.Println("Kode   :", dataJadwal[i].Kode)
		fmt.Println("Matkul :", dataJadwal[i].Matkul)
		fmt.Println("Dosen  :", dataJadwal[i].Dosen)
		fmt.Println("Hari   :", dataJadwal[i].Hari)
		fmt.Println("Jam Mulai   :", dataJadwal[i].JamMulai)
		fmt.Println("Jam Selesai :", dataJadwal[i].JamSelesai)
		fmt.Println("Ruangan     :", dataJadwal[i].Ruangan)
	}
}

func ubahJadwal() { //edit data
	var kode string

	fmt.Print("Masukkan kode matkul yang ingin diubah : ")
	fmt.Scanln(&kode)

	for i := 0; i < len(dataJadwal); i++ { // buat cari data ngambil dari kode
		if dataJadwal[i].Kode == kode {
			fmt.Print("Nama Matkul Baru : ")
			dataJadwal[i].Matkul = inputString()

			fmt.Print("Nama Dosen Baru : ")
			dataJadwal[i].Dosen = inputString()

			fmt.Print("Hari Baru : ")
			fmt.Scanln(&dataJadwal[i].Hari)

			fmt.Print("Jam Mulai Baru : ")
			fmt.Scanln(&dataJadwal[i].JamMulai)

			fmt.Print("Jam Selesai Baru : ")
			fmt.Scanln(&dataJadwal[i].JamSelesai)

			fmt.Print("Ruangan Baru : ")
			fmt.Scanln(&dataJadwal[i].Ruangan)

			fmt.Println("Data berhasil diubah")
			return
		}
	}
	fmt.Println("Data tidak ditemukan")
}

func hapusJadwal() { //hapus data
	var kode string
	fmt.Print("Masukkan kode matkul yang ingin dihapus : ")
	fmt.Scanln(&kode)

	for i := 0; i < len(dataJadwal); i++ { // buat cek data
		if dataJadwal[i].Kode == kode {
			dataJadwal = append(dataJadwal[:i], dataJadwal[i+1:]...)
			fmt.Println("Data berhasil dihapus")
			return
		}
	}
	fmt.Println("Data tidak ditemukan")
}

func sequentialSearch() {

	var keyword string

	fmt.Print("Masukkan nama matkul : ")
	keyword = inputString()

	for i := 0; i < len(dataJadwal); i++ {

		if strings.ToLower(dataJadwal[i].Matkul) == strings.ToLower(keyword) {

			fmt.Println("\nData ditemukan")
			fmt.Println("Kode :", dataJadwal[i].Kode)
			fmt.Println("Matkul :", dataJadwal[i].Matkul)
			fmt.Println("Dosen :", dataJadwal[i].Dosen)
			fmt.Println("Hari :", dataJadwal[i].Hari)

			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func selectionSort() {

	n := len(dataJadwal)

	for i := 0; i < n-1; i++ {

		min := i

		for j := i + 1; j < n; j++ {

			if dataJadwal[j].JamMulai < dataJadwal[min].JamMulai {

				min = j
			}
		}

		dataJadwal[i], dataJadwal[min] = dataJadwal[min], dataJadwal[i]
	}

	fmt.Println("Data berhasil diurutkan menggunakan Selection Sort")
	tampilkanJadwal()
}

func insertionSort() {

	n := len(dataJadwal)

	for i := 1; i < n; i++ {

		key := dataJadwal[i]

		j := i - 1

		for j >= 0 && dataJadwal[j].JamMulai > key.JamMulai {

			dataJadwal[j+1] = dataJadwal[j]

			j--
		}

		dataJadwal[j+1] = key
	}

	fmt.Println("Data berhasil diurutkan menggunakan Insertion Sort")
	tampilkanJadwal()
}

//test git
