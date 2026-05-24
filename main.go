package main

import "fmt"

// struct data jadwal (nad)
type Jadwal struct {
	Kode   string
	Matkul string
	Dosen  string
	Hari   string
}

var dataJadwal []Jadwal

// fungsi utama program (nad)
func main() {

	var pilih int

	for pilih != 3 {

		fmt.Println("===== JADWALKU =====")
		fmt.Println("1. Tambah Jadwal")
		fmt.Println("2. Tampilkan Jadwal")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilih)

		switch pilih {

		case 1:
			tambahJadwal()

		case 2:
			tampilkanJadwal()

		case 3:
			fmt.Println("Program selesai")

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

// fungsi tambah jadwal (nad)
func tambahJadwal() {

	var data Jadwal

	fmt.Print("Kode Matkul : ")
	fmt.Scanln(&data.Kode)

	fmt.Print("Nama Matkul : ")
	fmt.Scanln(&data.Matkul)

	fmt.Print("Nama Dosen : ")
	fmt.Scanln(&data.Dosen)

	fmt.Print("Hari : ")
	fmt.Scanln(&data.Hari)

	dataJadwal = append(dataJadwal, data)

	fmt.Println("Data berhasil ditambahkan")
}

// fungsi tampil jadwal (nad)
func tampilkanJadwal() {

	if len(dataJadwal) == 0 {

		fmt.Println("Data kosong")
		return
	}

	// perulangan untuk menampilkan data (nad)
	for i := 0; i < len(dataJadwal); i++ {

		fmt.Println("---------------")
		fmt.Println("Kode :", dataJadwal[i].Kode)
		fmt.Println("Matkul :", dataJadwal[i].Matkul)
		fmt.Println("Dosen :", dataJadwal[i].Dosen)
		fmt.Println("Hari :", dataJadwal[i].Hari)
	}
}