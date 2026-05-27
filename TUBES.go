package main

import (
	"fmt"
	"strings"
)

type SideHustle struct {
	ID           int
	Nama         string
	Kategori     string
	Pemasukan    int
	Pengeluaran  int
	Profit       int
	Status       string
}

var dataHustle []SideHustle
var nextID int = 1

func main() {

	var menu int

	for {
		tampilkanDashboard()

		fmt.Println("\n========== MENU ==========")
		fmt.Println("1. Tambah Side Hustle")
		fmt.Println("2. Lihat Semua Data")
		fmt.Println("3. Cari Hustle")
		fmt.Println("4. Update Data")
		fmt.Println("5. Hapus Data")
		fmt.Println("6. Urutkan Profit")
		fmt.Println("7. Statistik")
		fmt.Println("8. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)

		switch menu {

		case 1:
			tambahData()

		case 2:
			tampilkanData()

		case 3:
			cariData()

		case 4:
			updateData()

		case 5:
			hapusData()

		case 6:
			urutkanProfit()

		case 7:
			statistik()

		case 8:
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Menu tidak tersedia!")
		}
	}
}

// ================= DASHBOARD =================

func tampilkanDashboard() {

	total := 0

	for _, item := range dataHustle {
		total += item.Profit
	}

	fmt.Println("\n======================================")
	fmt.Println("     SIDE HUSTLE TRACKER SYSTEM")
	fmt.Println("======================================")
	fmt.Println("Jumlah Hustle :", len(dataHustle))
	fmt.Println("Total Profit  :", total)
	fmt.Println("======================================")
}

// ================= TAMBAH DATA =================

func tambahData() {

	var hustle SideHustle

	fmt.Println("\n===== TAMBAH SIDE HUSTLE =====")

	hustle.ID = nextID

	fmt.Print("Nama Hustle      : ")
	fmt.Scan(&hustle.Nama)

	fmt.Print("Kategori         : ")
	fmt.Scan(&hustle.Kategori)

	fmt.Print("Total Pemasukan  : ")
	fmt.Scan(&hustle.Pemasukan)

	fmt.Print("Total Pengeluaran: ")
	fmt.Scan(&hustle.Pengeluaran)

	hustle.Profit = hustle.Pemasukan - hustle.Pengeluaran

	if hustle.Profit > 0 {
		hustle.Status = "Untung"
	} else if hustle.Profit < 0 {
		hustle.Status = "Rugi"
	} else {
		hustle.Status = "Impas"
	}

	dataHustle = append(dataHustle, hustle)

	nextID++

	fmt.Println("Data berhasil ditambahkan!")
}

// ================= TAMPILKAN DATA =================

func tampilkanData() {

	if len(dataHustle) == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	fmt.Println("\n=========== DATA SIDE HUSTLE ===========")

	for _, item := range dataHustle {

		fmt.Println("----------------------------------")
		fmt.Println("ID          :", item.ID)
		fmt.Println("Nama        :", item.Nama)
		fmt.Println("Kategori    :", item.Kategori)
		fmt.Println("Pemasukan   :", item.Pemasukan)
		fmt.Println("Pengeluaran :", item.Pengeluaran)
		fmt.Println("Profit      :", item.Profit)
		fmt.Println("Status      :", item.Status)
	}
}

// ================= CARI DATA =================

func cariData() {

	var keyword string
	var ditemukan bool

	fmt.Print("Masukkan nama hustle: ")
	fmt.Scan(&keyword)

	keyword = strings.ToLower(keyword)

	for _, item := range dataHustle {

		if strings.ToLower(item.Nama) == keyword {

			fmt.Println("\nData ditemukan:")
			fmt.Println("---------------------------")
			fmt.Println("ID       :", item.ID)
			fmt.Println("Nama     :", item.Nama)
			fmt.Println("Profit   :", item.Profit)
			fmt.Println("Status   :", item.Status)

			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}

// ================= UPDATE DATA =================

func updateData() {

	var id int
	var ditemukan bool

	tampilkanData()

	fmt.Print("\nMasukkan ID yang ingin diupdate: ")
	fmt.Scan(&id)

	for i := 0; i < len(dataHustle); i++ {

		if dataHustle[i].ID == id {

			fmt.Print("Nama baru        : ")
			fmt.Scan(&dataHustle[i].Nama)

			fmt.Print("Kategori baru    : ")
			fmt.Scan(&dataHustle[i].Kategori)

			fmt.Print("Pemasukan baru   : ")
			fmt.Scan(&dataHustle[i].Pemasukan)

			fmt.Print("Pengeluaran baru : ")
			fmt.Scan(&dataHustle[i].Pengeluaran)

			dataHustle[i].Profit =
				dataHustle[i].Pemasukan - dataHustle[i].Pengeluaran

			if dataHustle[i].Profit > 0 {
				dataHustle[i].Status = "Untung"
			} else if dataHustle[i].Profit < 0 {
				dataHustle[i].Status = "Rugi"
			} else {
				dataHustle[i].Status = "Impas"
			}

			fmt.Println("Data berhasil diupdate!")
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("ID tidak ditemukan.")
	}
}

// ================= HAPUS DATA =================

func hapusData() {

	var id int
	var index int = -1

	tampilkanData()

	fmt.Print("\nMasukkan ID yang ingin dihapus: ")
	fmt.Scan(&id)

	for i := 0; i < len(dataHustle); i++ {

		if dataHustle[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	dataHustle =
		append(dataHustle[:index], dataHustle[index+1:]...)

	fmt.Println("Data berhasil dihapus.")
}

// ================= SORTING PROFIT =================

func urutkanProfit() {

	if len(dataHustle) == 0 {
		fmt.Println("Data kosong.")
		return
	}

	for i := 0; i < len(dataHustle)-1; i++ {

		for j := 0; j < len(dataHustle)-i-1; j++ {

			if dataHustle[j].Profit <
				dataHustle[j+1].Profit {

				dataHustle[j],
					dataHustle[j+1] =
					dataHustle[j+1],
					dataHustle[j]
			}
		}
	}

	fmt.Println("\nData berhasil diurutkan berdasarkan profit terbesar.")
	tampilkanData()
}

// ================= STATISTIK =================

func statistik() {

	if len(dataHustle) == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	totalProfit := 0
	terbaik := dataHustle[0]

	for _, item := range dataHustle {

		totalProfit += item.Profit

		if item.Profit > terbaik.Profit {
			terbaik = item
		}
	}

	rataRata := totalProfit / len(dataHustle)

	fmt.Println("\n========== STATISTIK ==========")
	fmt.Println("Jumlah Hustle :", len(dataHustle))
	fmt.Println("Total Profit  :", totalProfit)
	fmt.Println("Rata-rata     :", rataRata)
	fmt.Println("Hustle Terbaik:", terbaik.Nama)
	fmt.Println("Profit Tertinggi:", terbaik.Profit)
}