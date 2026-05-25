package main
import "fmt"

type Jadwal struct {
	Kode   string
	Matkul string
	Dosen  string
	Hari   string
}

var dataJadwal []Jadwal

func main() { 
	var pilih int

	for pilih != 5 {
		fmt.Println("\n===== JADWALKU =====")
		fmt.Println("1. Tambah Jadwal")
		fmt.Println("2. Tampilkan Jadwal")
		fmt.Println("3. Ubah Jadwal")
		fmt.Println("4. Hapus Jadwal")
		fmt.Println("5. Keluar")
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
			fmt.Println("Program selesai")

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func tambahJadwal() { //nambahd ata
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
	}
}

func ubahJadwal() { //edit data
	var kode string

	fmt.Print("Masukkan kode matkul yang ingin diubah : ")
	fmt.Scanln(&kode)

	for i := 0; i < len(dataJadwal); i++ { // buat cari data ngambil dari kode
		if dataJadwal[i].Kode == kode {
			fmt.Print("Nama Matkul Baru : ")
			fmt.Scanln(&dataJadwal[i].Matkul)

			fmt.Print("Nama Dosen Baru : ")
			fmt.Scanln(&dataJadwal[i].Dosen)

			fmt.Print("Hari Baru : ")
			fmt.Scanln(&dataJadwal[i].Hari)

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


//test git