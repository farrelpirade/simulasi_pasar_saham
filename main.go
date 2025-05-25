package main

import (
	"fmt"
	"math/rand"
	"time"
)

// aplikasi simulasi pasar saham virtual

const nSaham int = 36
const nHistori int = 100

type saham struct {
	kode                 string
	nama                 string
	harga                float64
	perubahan_persentase float64
	volume               int
}

type arrTransaksi struct {
	kode_saham_transaksi         string
	nama_saham_transaksi         string
	jumlah_saham_transaksi       int
	harga_lembar_saham_transaksi float64
	harga_total_transaksi        float64
	jenis_transaksi              string // beli atau jual saham
}

var histori [nHistori]arrTransaksi
var hitungHistori int

// penggunaan alias untuk array jumlah saham yang dimiliki yang di fixed-size
type jumlah_owned_saham [nSaham]int

var ownedSaham jumlah_owned_saham

// penggunaan alias untuk array saham yang di fixed-size
type daftarSaham [nSaham]saham

// inisialisasi kode saham dan nama perusahaan
var listSaham daftarSaham

func initKodeNama() {
	var kode [nSaham]string
	var nama [nSaham]string

	kode = [nSaham]string{
		"ASII", "BBRI", "BMRI", "TLKM", "UNTR", "BBCA",
		"AMRT", "INDF", "GGRM", "BBNI", "AADI", "SMAR",
		"ICBP", "ANTM", "CPIN", "ERAA", "PGAS", "ISAT",
		"JPFA", "BYAN", "GIAA", "INKP", "HMSP", "DSSA",
		"GEMS", "PTBA", "AMMN", "INDY", "AKRA", "MEDC",
		"MAPI", "BRPT", "ITMG", "SMGR", "MYOR", "MDKA",
	}
	nama = [nSaham]string{
		"ASTRA_INTERNATIONAL", "BANK_RAKYAT_INDONESIA",
		"BANK_MANDIRI_(PERSERO)_TBK", "TELKOM_INDONESIA_(PERSERO)_TBK",
		"UNITED_TRACTORS", "BANK_CENTRAL_ASIA",
		"SUMBER_ALFARIA_TRIJAYA_TBK", "PT_INDOFOOD_SUKSES_MAKMUR_TBK",
		"GUDANG_GARAM_TBK", "PT_BANK_NEGARA_INDONESIA_(PERSERO)",
		"ADARO_ANDALAN_INDONESIA_TBK", "SINAR_MAS_AGRO_RESOURCES_TECHNOLOGY",
		"INDOFOOD_CBP_SUKSES_MAKMUR_TBK_PT", "ANEKA_TAMBANG",
		"CHAROEN_POKPHAND_INDONESIA", "ERAJAYA_TBK",
		"PERUSAHAAN_GAS_NEGARA_TBK", "INDOSAT_(PT_INDOSAT_TBK)",
		"JAPFA_COMFEED_INDONESIA", "BAYAN_RESOURCES_TBK",
		"GARUDA_INDONESIA_(PERSERO)_TBK", "INDAH_KIAT_PULP_&_PAPER",
		"HANJAYA_MANDALA_SAMPOERNA", "DIAN_SWASTATIKA_SENTOSA_TBK",
		"GOLDEN_ENERGY_MINES_TBK", "BUKIT_ASAM_TBK",
		"AMMAN_MINERAL_INTERNASIONAL_TBK", "INDIKA_ENERGY_TBK",
		"AKR_CORPORINDO_TBK", "MEDCO_ENERGI_INTERNATIONAL",
		"MITRA_ADIPERKASA_TBK", "BARITO_PACIFIC_TBK",
		"INDO_TAMBANGRAYA_MEGAH_TBK", "SEMEN_INDONESIA_(PERSERO)_TBK",
		"MAYORA_INDAH", "MERDEKA_COPPER_GOLD_TBK_PT",
	}

	var i int
	for i = 0; i < nSaham; i++ {
		listSaham[i].kode = kode[i]
		listSaham[i].nama = nama[i]
	}
}

const asciiArt = `
███████╗██╗███╗   ███╗██╗   ██╗██╗      █████╗ ███████╗██╗
██╔════╝██║████╗ ████║██║   ██║██║     ██╔══██╗██╔════╝██║
███████╗██║██╔████╔██║██║   ██║██║     ███████║███████╗██║
╚════██║██║██║╚██╔╝██║██║   ██║██║     ██╔══██║╚════██║██║
███████║██║██║ ╚═╝ ██║╚██████╔╝███████╗██║  ██║███████║██║
╚══════╝╚═╝╚═╝     ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝╚══════╝╚═╝
														  
██████╗  █████╗ ███████╗ █████╗ ██████╗                   
██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔══██╗                  
██████╔╝███████║███████╗███████║██████╔╝                  
██╔═══╝ ██╔══██║╚════██║██╔══██║██╔══██╗                  
██║     ██║  ██║███████║██║  ██║██║  ██║                  
╚═╝     ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝                  
														  
███████╗ █████╗ ██╗  ██╗ █████╗ ███╗   ███╗               
██╔════╝██╔══██╗██║  ██║██╔══██╗████╗ ████║               
███████╗███████║███████║███████║██╔████╔██║               
╚════██║██╔══██║██╔══██║██╔══██║██║╚██╔╝██║               
███████║██║  ██║██║  ██║██║  ██║██║ ╚═╝ ██║               
╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝               
`

func main() {
	// program main merupakan interface awal aplikasi saat dibuka
	rand.Seed(time.Now().UnixNano())
	randomisasi(&listSaham)
	initKodeNama()
	var jumSaldo int
	var pilihan int

	// selama pilihan tidak 7 (keluar), program akan terus berjalan
	for pilihan != 7 {
		displayMainMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			// menu saldo
			saldo(&jumSaldo)
		case 2:
			// menu daftar saham
			daftar_saham(&listSaham)
		case 3:
			// menu transaksi saham
			transaksi_saham(&listSaham, &jumSaldo)
		case 4:
			// menu portofolio
			portofolio()
		case 5:
			// menu histori transaksi
			histori_transaksi(&histori, hitungHistori)
		case 6:
			// menu bantuan
			var pilih_bantuan int
			bantuan(pilih_bantuan)
		}

		fmt.Println()
	}
	fmt.Println("Terima kasih telah menggunakan aplikasi. Sampai jumpa lagi!")
	fmt.Println()
}

func displayMainMenu() {
	// prosedur untuk menampilkan menu utama
	fmt.Print(asciiArt)
	fmt.Println()
	fmt.Println("Pilih menu")
	fmt.Println("1. Saldo anda")
	fmt.Println("2. Daftar saham")
	fmt.Println("3. Transaksi saham")
	fmt.Println("4. Portofolio")
	fmt.Println("5. Histori transaksi")
	fmt.Println("6. Bantuan")
	fmt.Println("7. Keluar")
	fmt.Print("Masukkan pilihan menu > ")
}

func randomisasi(A *daftarSaham) {
	var i int
	// fungsi randomisasi untuk mengacak harga, perubahan persentase, dan volume saham
	for i = 0; i < nSaham; i++ {
		// Randomisasi harga saham dengan range sampai 21000
		A[i].harga = float64(rand.Intn(21000)) + 1
		if A[i].harga < 1 {
			A[i].harga = 1 // harga saham tidak boleh kurang dari 1
		}

		//Randomisasi perubahan persentase saham
		A[i].perubahan_persentase = float64(rand.Intn(100)) + 1

		// Randomisasi volume saham dengan fluktuasi kurang lebih antara 100 hingga 3000000000
		A[i].volume = rand.Intn(3000000000) + 100
	}
}

func saldo(s *int) {
	// prosedur saldo untuk menampilkan saldo, top up, dan withdraw
	for {
		var pilihan int
		fmt.Printf("Saldo anda adalah %d \n", *s)
		fmt.Println("1. Top up")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Kembali")
		fmt.Print("Masukkan pilihan menu > ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			// melakukan top up
			handlerTopup(s)
		case 2:
			// melakukan withdraw
			handlerWithdraw(s)
		case 3:
			return
		}
	}
}

func handlerTopup(s *int) {
	fmt.Print("Masukkan jumlah top up > ")
	var topup int
	fmt.Scan(&topup)
	if topup > 0 {
		*s += topup
		fmt.Printf("Top up berhasil. Saldo anda sekarang adalah %d \n", *s)
	} else {
		fmt.Println("Nilai tidak valid")
	}
}

func handlerWithdraw(s *int) {
	var withdraw int
	fmt.Print("Masukkan jumlah withdraw > ")
	fmt.Scan(&withdraw)
	if withdraw > 0 && withdraw <= *s {
		*s -= withdraw
		fmt.Printf("Withdraw berhasil. Saldo anda sekarang adalah %d \n", *s)
	} else {
		fmt.Println("Saldo tidak cukup atau nilai tidak valid")
	}
}

func sequential_search(A daftarSaham, kode_saham string) int {
	var iterasi_kode_seq int
	var found bool = false
	for !found && iterasi_kode_seq < nSaham {
		found = A[iterasi_kode_seq].kode == kode_saham
		iterasi_kode_seq++
	}
	if found {
		return iterasi_kode_seq - 1
	} else {
		return -1
	}
}

func sel_sort_for_bin_search(A *daftarSaham) {
	// pengurutan daftar saham berdasarkan nama perusahaan menggunakan selection sort
	var iterasi_nama_bin1, iterasi_nama_bin2 int
	var temp saham
	var min int
	for iterasi_nama_bin1 = 0; iterasi_nama_bin1 < nSaham; iterasi_nama_bin1++ {
		min = iterasi_nama_bin1
		for iterasi_nama_bin2 = iterasi_nama_bin1 + 1; iterasi_nama_bin2 < nSaham; iterasi_nama_bin2++ {
			if A[iterasi_nama_bin2].nama < A[min].nama {
				min = iterasi_nama_bin2
			}
		}
		temp = A[iterasi_nama_bin1]
		A[iterasi_nama_bin1] = A[min]
		A[min] = temp
	}
}

func binary_search(A daftarSaham, nama_saham string) int {
	// menggunakan binary search untuk mencari nama perusahaan
	var left, right, mid int
	left = 0
	right = nSaham - 1
	for left <= right {
		mid = (left + right) / 2
		if nama_saham == A[mid].nama {
			return mid
		} else if nama_saham > A[mid].nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func daftar_saham(A *daftarSaham) {
	sel_sort_for_bin_search(A)

	// beberapa daftar saham yang bisa diakses
	for {
		displayDaftarSaham(A)
		var pilih int = pilihMenuDaftarSaham()
		switch pilih {
		case 1:
			// menu searching saham
			menuCariSaham(A)
		case 2:
			// menu sorting saham
			menuSortSaham(A)
		case 3:
			// kembali ke main menu
			return
		}
	}
}

func displayDaftarSaham(A *daftarSaham) {
	var i int
	fmt.Println("Daftar saham : ")
	fmt.Println("_______________________________________________________________________________________________________________________")
	fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
	fmt.Println("_______________________________________________________________________________________________________________________")

	for i = 0; i < nSaham; i++ {
		// Tampilan ke user
		fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", A[i].kode, A[i].nama, A[i].harga/1000, A[i].perubahan_persentase, A[i].volume)
		fmt.Println()
	}
	fmt.Println("_______________________________________________________________________________________________________________________")
	fmt.Println()
}

func pilihMenuDaftarSaham() int {
	fmt.Println("Pilih menu")
	fmt.Println("1. Cari saham")
	fmt.Println("2. Urutkan saham")
	fmt.Println("3. Kembali")
	fmt.Print("Masukkan pilihan menu > ")
	var pilih_daftar_menu_saham int
	fmt.Scan(&pilih_daftar_menu_saham)
	return pilih_daftar_menu_saham
}

func menuCariSaham(A *daftarSaham) {
	fmt.Println("Pilih metode pencarian")
	fmt.Println("1. Pencarian dengan kode saham")
	fmt.Println("2. Pencarian dengan nama saham")
	fmt.Println("3. Kembali")
	fmt.Print("Masukkan pilihan menu > ")
	var pilih_search int
	fmt.Scan(&pilih_search)

	switch pilih_search {
	case 1:
		// pencarian dengan kode saham dengan sequential search
		var stringKode string = inputString("Masukkan kode saham > ")
		var hasil int = sequential_search(*A, stringKode)
		displayHasilCari(A, hasil)
	case 2:
		// pencarian dengan nama perusahaan saham dengan binary search
		fmt.Println("Gunakan underscore (_) untuk spasi")
		var stringNama string = inputString("Masukkan nama perusahaan > ")
		var hasil int = binary_search(*A, stringNama)
		displayHasilCari(A, hasil)
	case 3:
		// kembali ke menu daftar saham
		return
	}
}

func displayHasilCari(A *daftarSaham, hasil int) {
	if hasil == -1 {
		fmt.Println("Kode saham tidak ditemukan")
	} else {
		fmt.Println("Berikut merupakan hasil pencarian : ")
		fmt.Println("_______________________________________________________________________________________________________________________")
		fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
		fmt.Println("_______________________________________________________________________________________________________________________")
		// Tampilan ke user
		fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", A[hasil].kode, A[hasil].nama, A[hasil].harga/1000, A[hasil].perubahan_persentase, A[hasil].volume)
		fmt.Println()
		fmt.Println("_______________________________________________________________________________________________________________________")
		fmt.Print("Ketik X untuk kembali > ")
		var kembali string
		fmt.Scan(&kembali)
		return
	}
}

func menuSortSaham(A *daftarSaham) {
	var pilih int = inputInt("Metode Pengurutan (1.Harga/2.Volume/3.Kembali) > ")

	switch pilih {
	case 1:
		// pengurutan dengan harga saham dengan selection sort
		subSortHarga(A)
	case 2:
		// pengurutan dengan volume saham dengan insertion sort
		subSortVolume(A)
	case 3:
		// kembali ke menu daftar saham
		return
	}
}

func subSortHarga(A *daftarSaham) {
	var urutan int = inputInt("Urutan (1.Tertinggi/2.Terendah) > ")
	if urutan == 1 {
		sortHargaDesc(A)
	} else if urutan == 2 {
		sortHargaAsc(A)
	} else {
		fmt.Println("Pilihan tidak valid")
		return
	}
	displaySaham(A)
}

func subSortVolume(A *daftarSaham) {
	var urutan int = inputInt("Urutan (1.Tertinggi/2.Terendah) > ")
	if urutan == 1 {
		sortVolumeDesc(A)
	} else if urutan == 2 {
		sortVolumeAsc(A)
	} else {
		fmt.Println("Pilihan tidak valid")
		return
	}
	displaySaham(A)
}

func sortHargaDesc(A *daftarSaham) {
	var temp saham
	var n int = nSaham
	var i, j, max int
	for i = 0; i < n-1; i++ {
		max = i
		for j = i + 1; j < n; j++ {
			if A[j].harga > A[max].harga {
				max = j
			}
		}
		temp = A[i]
		A[i] = A[max]
		A[max] = temp
	}
}

func sortHargaAsc(A *daftarSaham) {
	var temp saham
	var n int = nSaham
	var i, j, min int
	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if A[j].harga < A[min].harga {
				min = j
			}
		}
		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
}

func sortVolumeDesc(A *daftarSaham) {
	var n int = nSaham
	var i, j int
	for i = 1; i < n; i++ {
		var temp saham = A[i]
		j = i - 1
		for j >= 0 && A[j].volume < temp.volume {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

func sortVolumeAsc(A *daftarSaham) {
	var n int = nSaham
	var i, j int
	for i = 1; i < n; i++ {
		var temp saham = A[i]
		j = i - 1
		for j >= 0 && A[j].volume > temp.volume {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

func displaySaham(A *daftarSaham) {
	var i int
	fmt.Println("Berikut merupakan daftar saham setelah diurutkan : ")
	fmt.Println("_______________________________________________________________________________________________________________________")
	fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
	fmt.Println("_______________________________________________________________________________________________________________________")
	for i = 0; i < nSaham; i++ {
		// Tampilan ke user
		fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", A[i].kode, A[i].nama, A[i].harga/1000, A[i].perubahan_persentase, A[i].volume)
		fmt.Println()
	}
	fmt.Println("_______________________________________________________________________________________________________________________")
	fmt.Print("Ketik X untuk kembali > ")
	var kembali string
	fmt.Scan(&kembali)
}

func inputInt(words string) int {
	fmt.Print(words)
	var input int
	fmt.Scan(&input)
	return input
}

func inputString(words string) string {
	fmt.Print(words)
	var input string
	fmt.Scan(&input)
	return input
}

func transaksi_saham(A *daftarSaham, saldo *int) {
	// menu transaksi saham
	var pilih int = inputInt("Transaksi (1.Beli/2.Jual/3.Kembali) > ")
	switch pilih {
	case 1:
		// menu beli saham
		handleBeliSaham(A, saldo)
	case 2:
		// menu jual saham
		handleJualSaham(A, saldo)
	case 3:
		// kembali ke main menu
		return
	}
}

func handleBeliSaham(A *daftarSaham, saldo *int) {
	var beli_kode_saham string = inputString("Masukkan kode saham > ")
	var hasil int = sequential_search(*A, beli_kode_saham)
	displayHasilCariTransaksi(A, hasil, saldo)
	var beli_jumlah_saham int = inputInt("Masukkan jumlah saham yang ingin dibeli > ")

	// menamppilkan saldo pengguna
	var total_harga_sementara float64 = float64(beli_jumlah_saham) * A[hasil].harga
	fmt.Printf("Saldo anda sekarang adalah %d \n", *saldo)
	if total_harga_sementara > float64(*saldo) {
		fmt.Println("Saldo anda tidak cukup")
		transaksi_saham(A, saldo)
	} else if beli_jumlah_saham < 0 {
		fmt.Println("Jumlah saham tidak valid")
		transaksi_saham(A, saldo)
	} else if beli_jumlah_saham > A[hasil].volume {
		fmt.Println("Jumlah saham yang ingin dibeli melebihi volume saham")
		transaksi_saham(A, saldo)
	} else {
		var total_harga float64 = float64(beli_jumlah_saham) * A[hasil].harga

		fmt.Printf("Total harga yang harus dibayar adalah Rp%.2f \n", total_harga)
		fmt.Printf("Saldo anda sekarang adalah %d \n", *saldo)

		// melanjutkan pembayaran
		fmt.Print("Bayar ? (y/n) > ")
		var pilih_bayar string
		fmt.Scan(&pilih_bayar)
		if pilih_bayar == "y" || pilih_bayar == "Y" {
			// menambahkan ke portofolio
			ownedSaham[hasil] += beli_jumlah_saham

			// melakukan pembayaran
			*saldo -= int(total_harga)
			A[hasil].volume -= beli_jumlah_saham

			// menambahkan ke history transaksi
			tambahHistori(&histori, "Beli", A[hasil].kode, A[hasil].nama, beli_jumlah_saham, A[hasil].harga, total_harga)

			fmt.Println("Pembayaran berhasil")
			fmt.Printf("Sisa volume saham %s adalah %d \n", A[hasil].kode, A[hasil].volume)
			fmt.Printf("Sisa saldo anda adalah %d \n", *saldo)
			fmt.Print("Ketik X untuk kembali > ")
			var kembali string
			fmt.Scan(&kembali)
			return
		} else if pilih_bayar == "n" || pilih_bayar == "N" {
			fmt.Println("Pembayaran dibatalkan")
			transaksi_saham(A, saldo)
		} else {
			fmt.Println("Pilihan tidak valid")
			transaksi_saham(A, saldo)
		}
	}
}

func handleJualSaham(A *daftarSaham, saldo *int) {
	var jual_kode_saham string = inputString("Masukkan kode saham yang ingin dijual > ")
	var hasil int = sequential_search(*A, jual_kode_saham)
	displayHasilCariTransaksi(A, hasil, saldo)

	fmt.Print("Masukkan jumlah saham yang ingin dijual > ")
	var jual_jumlah_saham int
	fmt.Scan(&jual_jumlah_saham)
	// menampilkan saldo pengguna sekarang dengan mengambil dari variabel jumSaldo
	fmt.Printf("Saldo anda sekarang adalah %d \n", *saldo)
	if jual_jumlah_saham < 0 {
		fmt.Println("Jumlah saham tidak valid")
		transaksi_saham(A, saldo)
	} else if jual_jumlah_saham > ownedSaham[hasil] {
		fmt.Println("Jumlah saham yang ingin dijual melebihi volume saham")
		transaksi_saham(A, saldo)
	} else {
		// portofolio
		ownedSaham[hasil] -= jual_jumlah_saham

		// mendapatkan pendapatan dari penjualan saham
		var jual_pendapatan float64 = float64(jual_jumlah_saham) * A[hasil].harga
		*saldo += int(jual_pendapatan)
		A[hasil].volume += jual_jumlah_saham
		fmt.Printf("Volume saham %s sekarang adalah %d \n", A[hasil].kode, A[hasil].volume)
		var total_pendapatan float64 = float64(jual_jumlah_saham) * A[hasil].harga
		fmt.Printf("Anda mendapat Rp%.2f \n", total_pendapatan)
		fmt.Printf("Saldo anda sekarang adalah %d \n", *saldo)

		// menambahkan ke history transaksi
		tambahHistori(&histori, "Jual", A[hasil].kode, A[hasil].nama, jual_jumlah_saham, A[hasil].harga, total_pendapatan)
	}
}

func displayHasilCariTransaksi(A *daftarSaham, hasil int, saldo *int) {
	if hasil == -1 {
		fmt.Println("Kode saham tidak ditemukan")
		transaksi_saham(A, saldo)
	} else {
		fmt.Println("Berikut merupakan hasil pencarian : ")
		fmt.Println("_______________________________________________________________________________________________________________________")
		fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
		fmt.Println("_______________________________________________________________________________________________________________________")
		// Tampilan ke user
		fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", A[hasil].kode, A[hasil].nama, A[hasil].harga/1000, A[hasil].perubahan_persentase, A[hasil].volume)
		fmt.Println()
		fmt.Println("_______________________________________________________________________________________________________________________")
	}
}

func tambahHistori(histori *[nHistori]arrTransaksi, jenis string, kode string, nama string, jumlah int, harga float64, total float64) {
	if hitungHistori < nHistori {
		histori[hitungHistori].kode_saham_transaksi = kode
		histori[hitungHistori].nama_saham_transaksi = nama
		histori[hitungHistori].jumlah_saham_transaksi = jumlah
		histori[hitungHistori].harga_lembar_saham_transaksi = harga
		histori[hitungHistori].harga_total_transaksi = total
		histori[hitungHistori].jenis_transaksi = jenis
		hitungHistori++
	}
}

func portofolio() {
	// menu portofolio
	fmt.Println("Berikut merupakan portofolio anda : ")
	fmt.Println("_______________________________________________________________________________________________________________________")
	fmt.Printf("| %-6s | %-8s | %-10s | %-6s | %-10s |\n", "Kode", "Lembar", "Harga/Unit", "Perubahan %", "Total")
	fmt.Println("_______________________________________________________________________________________________________________________")

	// Tampilan ke user
	var i int
	for i = 0; i < nSaham; i++ {
		var banyak_kepunyaan int
		banyak_kepunyaan = ownedSaham[i]
		if banyak_kepunyaan > 0 {
			var total_kepunyaan float64 = float64(banyak_kepunyaan) * listSaham[i].harga
			fmt.Printf("| %-6s | %-8d | Rp%-10.3f | %-6.2f | Rp%-10.3f |\n", listSaham[i].kode, banyak_kepunyaan, listSaham[i].harga/1000, listSaham[i].perubahan_persentase, total_kepunyaan/1000)
		}
	}
	fmt.Print("\n Ketik X untuk kembali > ")
	var kembali string
	fmt.Scan(&kembali)
	return
}

func histori_transaksi(h *[nHistori]arrTransaksi, hitungHistori int) {
	// menu histori transaksi
	fmt.Println("Berikut merupakan histori transaksi anda : ")
	fmt.Printf("| %-6s | %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Jenis", "Kode", "Nama Perusahaan", "Jumlah", "Harga/Unit", "Total")

	var i int
	for i = 0; i < hitungHistori; i++ {
		fmt.Printf("%s\t %s\t %s\t %d\t %.2f\t %.2f\n",
			h[i].jenis_transaksi,
			h[i].kode_saham_transaksi,
			h[i].nama_saham_transaksi,
			h[i].jumlah_saham_transaksi,
			h[i].harga_lembar_saham_transaksi/1000,
			h[i].harga_total_transaksi/1000)
	}
	fmt.Println()
	fmt.Print("Ketik X untuk kembali > ")
	var kembali string
	fmt.Scan(&kembali)
	return

}

func bantuan(pilih_bantuan int) {
	// menu bantuan
	fmt.Println("Bantuan")
	fmt.Println("1. Tentang aplikasi")
	fmt.Println("2. Cara menggunakan aplikasi")
	fmt.Println("3. Kembali")
	fmt.Print("Masukkan pilihan menu > ")
	fmt.Scan(&pilih_bantuan)

	switch pilih_bantuan {
	case 1:
		var opsi_about int
		fmt.Println("Aplikasi ini adalah aplikasi simulasi pasar saham virtual yang dibuat dengan menggunakan bahasa Go.")
		fmt.Println("Aplikasi ini dibuat sebagai bentuk pelaksanaan tugas besar dari mata kuliah Algoritma dan Pemrograman 2")
		fmt.Println("Aplikasi ini memungkinkan pengguna untuk mensimulasikan perdagangan saham dengan menggunakan saldo virtual. Data utama yang digunakan adalah daftar saham, harga saham yang berubah, dan portofolio pengguna. Pengguna aplikasi adalah individu yang ingin belajar cara trading saham tanpa risiko nyata.")
		fmt.Println("Spesifikasi :")
		fmt.Println("a. Pengguna dapat menambahkan, mengubah, dan menghapus transaksi pembelian dan penjualan saham.")
		fmt.Println("b. Sistem menghitung perubahan nilai portofolio berdasarkan fluktuasi harga saham.")
		fmt.Println("c. Pengguna dapat mencari saham berdasarkan kode atau nama perusahaan menggunakan Sequential dan Binary Search.")
		fmt.Println("d. Pengguna dapat mengurutkan saham berdasarkan harga tertinggi atau volume transaksi menggunakan Selection dan Insertion Sort.")
		fmt.Println("Sistem menampilkan statistik keuntungan dan kerugian pengguna dalam simulasi trading.")
		fmt.Println("Pembuat : Kanaka Pradipta Arya Wismaya & Farrel Malik Pirade")
		fmt.Println()
		fmt.Print("Untuk kembali, pilih 0 > ")
		fmt.Scan(&opsi_about)
		if opsi_about == 0 {
			bantuan(pilih_bantuan)
		}

	case 2:
		var opsi_tutorial int
		fmt.Println("Cara menggunakan aplikasi ini adalah dengan memilih menu yang tersedia.")
		fmt.Println("Berikut merupakan penjelasan dari masing-masing menu utama :")
		fmt.Println("1. Saldo anda : Menampilkan saldo yang anda miliki, serta melakukan top up dan withdraw.")
		fmt.Println("2. Daftar saham : Menampilkan daftar saham yang tersedia dan yang dibantu dengan fitur pencarian dan pengurutan.")
		fmt.Println("3. Transaksi saham : Melakukan transaksi pembelian dan penjualan saham dengan fitur pencarian dan pengurutan.")
		fmt.Println("4. Portofolio : Menampilkan portofolio yang anda miliki.")
		fmt.Println("5. Histori transaksi : Menampilkan histori transaksi yang pernah anda lakukan.")
		fmt.Println("6. Artikel saham : Menampilkan artikel yang berisi informasi-informasi menarik tentang saham.")
		fmt.Println("7. Bantuan : Menampilkan informasi tentang aplikasi dan cara menggunakan aplikasi.")
		fmt.Println("8. Keluar : Keluar dari aplikasi.")
		fmt.Println()
		fmt.Print("Untuk kembali, ketik 0 > ")
		fmt.Scan(&opsi_tutorial)
		if opsi_tutorial == 0 {
			bantuan(pilih_bantuan)
		}

	case 3:
		// kembali ke main menu
		return
	}
}
