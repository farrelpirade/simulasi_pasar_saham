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

// variabel untuk handle nSaham di menu penjual
var countSaham int = nSaham

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
	initKodeNama()
	var jumSaldo int
	var pilihan int

	// selama pilihan tidak 8 (keluar), program akan terus berjalan
	for pilihan != 8 {
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
		case 7:
			// fitur penjual
			var pilih_menu_jual int
			penjual(pilih_menu_jual)
		case 8:
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
	fmt.Println("7. Menu penjual")
	fmt.Println("8. Keluar")
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
		if rand.Intn(2) == 0 {
			A[i].perubahan_persentase = -1 * (float64(rand.Intn(100)) + 1)
		} else {
			A[i].perubahan_persentase = float64(rand.Intn(100)) + 1
		}

		// Randomisasi volume saham dengan fluktuasi kurang lebih antara 100 hingga 3000000000
		A[i].volume = rand.Intn(3000000000) + 100
	}
}

func saldo(s *int) {
	// prosedur saldo untuk menampilkan saldo, top up, dan withdraw
	var pilihan int
	for pilihan != 3 {
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
	randomisasi(A)
	sel_sort_for_bin_search(A)
	displayDaftarSaham(A)
	var pilih int = 0
	for pilih != 3 {
		pilih = pilihMenuDaftarSaham()
		switch pilih {
		case 1:
			// menu searching saham
			menuCariSaham(A)
		case 2:
			// menu sorting saham
			menuSortSaham(A)
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
		fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %+15.2f | %-30d |", A[i].kode, A[i].nama, A[i].harga/1000, A[i].perubahan_persentase, A[i].volume)
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
	var pilih_search int = 0
	for pilih_search != 3 {
		fmt.Println("Pilih metode pencarian")
		fmt.Println("1. Pencarian dengan kode saham")
		fmt.Println("2. Pencarian dengan nama saham")
		fmt.Println("3. Kembali")
		fmt.Print("Masukkan pilihan menu > ")
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
		}
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
		fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %+15.2f | %-30d |", A[hasil].kode, A[hasil].nama, A[hasil].harga/1000, A[hasil].perubahan_persentase, A[hasil].volume)
		fmt.Println()
		fmt.Println("_______________________________________________________________________________________________________________________")
		var kembali string = ""
		for kembali != "X" && kembali != "x" {
			fmt.Print("Ketik X untuk kembali > ")
			fmt.Scan(&kembali)
		}

	}
}

func menuSortSaham(A *daftarSaham) {
	var pilih int = inputInt("Metode Pengurutan (1.Harga/2.Volume) > ")
	switch pilih {
	case 1:
		// pengurutan dengan harga saham dengan selection sort
		subSortHarga(A)
	case 2:
		// pengurutan dengan volume saham dengan insertion sort
		subSortVolume(A)
	}
}

func subSortHarga(A *daftarSaham) {
	var urutan int = inputInt("Urutan (1.Tertinggi/2.Terendah) > ")
	if urutan == 1 {
		sortHargaDesc(A)
	} else if urutan == 2 {
		sortHargaAsc(A)
	}
	displaySaham(A)
}

func subSortVolume(A *daftarSaham) {
	var urutan int = inputInt("Urutan (1.Tertinggi/2.Terendah) > ")
	if urutan == 1 {
		sortVolumeDesc(A)
	} else if urutan == 2 {
		sortVolumeAsc(A)
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
		fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %+15.2f | %-30d |", A[i].kode, A[i].nama, A[i].harga/1000, A[i].perubahan_persentase, A[i].volume)
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
	for {
		var pilih int = inputInt("Pilih menu transaksi (1.Beli/2.Jual/3.Kembali) > ")
		if pilih == 3 {
			break // repeat until akan berhenti jika pilih == 3
		}
		switch pilih {
		case 1:
			// menu beli saham
			randomisasi(A)
			handleBeliSaham(A, saldo)
		case 2:
			// menu jual saham
			randomisasi(A)
			handleJualSaham(A, saldo)
		}
	}
}

func handleBeliSaham(A *daftarSaham, saldo *int) {
	var kode string
	var idx, jumlah int
	var doneOuter bool = false // flag untuk keluar seluruh prosedur

	// loop utama akan terus tanya kode sampai doneOuter itu adalah true
	for !doneOuter {
		fmt.Print("Masukkan kode saham (X untuk kembali) > ")
		fmt.Scan(&kode)

		// Jika user pilih keluar, cukup set flag doneOuter true
		if kode == "X" || kode == "x" {
			doneOuter = true

		} else {
			// Cari dan tampilkan hasilnya
			idx = sequential_search(*A, kode)
			displayHasilCariTransaksi(A, idx)

			// Hanya jika kode valid (=! -1), maka akan masuk ke loop jumlah
			if idx != -1 {
				var doneInner bool = false // flag keluar loop jumlah

				// Loop jumlah & validasi
				for !doneInner {
					fmt.Print("Masukkan jumlah saham yang ingin dibeli > ")
					fmt.Scan(&jumlah)
					var total float64 = float64(jumlah) * A[idx].harga

					if jumlah <= 0 {
						fmt.Println("Jumlah saham tidak valid.")
					} else if total > float64(*saldo) {
						fmt.Println("Saldo anda tidak cukup.")
					} else if jumlah > A[idx].volume {
						fmt.Println("Jumlah melebihi volume tersedia.")
					} else {
						// jika valid maka lanjutkan ke proses transaksi
						fmt.Printf("Total harga: Rp%.2f | Saldo sekarang: %d\n", total, *saldo)
						fmt.Print("Bayar ? (y/n) > ")
						var c string
						fmt.Scan(&c)

						if c == "y" || c == "Y" {
							ownedSaham[idx] += jumlah
							*saldo -= int(total)
							A[idx].volume -= jumlah
							tambahHistori(&histori, "Beli", A[idx].kode, A[idx].nama, jumlah, A[idx].harga, total)

							fmt.Println("Pembayaran berhasil.")
							fmt.Printf("Sisa saldo: %d | Sisa volume %s: %d\n", *saldo, A[idx].kode, A[idx].volume)

							// Tandai agar keluar dari kedua loop
							doneInner = true
							doneOuter = true

							// tunggu X untuk benar-benar kembali ke transaksi_saham
							var back string
							for back != "X" && back != "x" {
								fmt.Print("Ketik X untuk kembali > ")
								fmt.Scan(&back)
							}
						} else {
							fmt.Println("Pembayaran dibatalkan. Masukkan ulang jumlah.")
						}
					}
				}
			}
		}
	}
}

func handleJualSaham(A *daftarSaham, saldo *int) {
	var kode string
	var idx, jumlah int
	var doneOuter bool = false

	for !doneOuter {
		fmt.Print("Masukkan kode saham (X untuk kembali) > ")
		fmt.Scan(&kode)

		if kode == "X" || kode == "x" {
			doneOuter = true

		} else {
			idx = sequential_search(*A, kode)
			displayHasilCariTransaksi(A, idx)

			if idx != -1 {
				var doneInner bool = false

				for !doneInner {
					fmt.Print("Masukkan jumlah saham yang ingin dijual > ")
					fmt.Scan(&jumlah)

					if jumlah <= 0 {
						fmt.Println("Jumlah tidak valid.")
					} else if jumlah > ownedSaham[idx] {
						fmt.Println("Volume yang dijual melebihi milik Anda.")
					} else {
						// Proses jual
						var pendapatan float64 = float64(jumlah) * A[idx].harga
						ownedSaham[idx] -= jumlah
						*saldo += int(pendapatan)
						A[idx].volume += jumlah
						tambahHistori(&histori, "Jual", A[idx].kode, A[idx].nama, jumlah, A[idx].harga, pendapatan)

						fmt.Printf("Anda mendapat Rp%.2f | Saldo sekarang: %d\n", pendapatan, *saldo)
						fmt.Printf("Volume %s kini: %d\n", A[idx].kode, A[idx].volume)

						doneInner = true
						doneOuter = true

						// Tunggu X agar benar-benar keluar
						var back string
						for back != "X" && back != "x" {
							fmt.Print("Ketik X untuk kembali > ")
							fmt.Scan(&back)
						}
					}
				}
			}
		}
	}
}

func displayHasilCariTransaksi(A *daftarSaham, hasil int) {
	if hasil == -1 {
		fmt.Println("Kode saham tidak ditemukan")
	} else {
		fmt.Println("Berikut merupakan hasil pencarian : ")
		fmt.Println("_______________________________________________________________________________________________________________________")
		fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
		fmt.Println("_______________________________________________________________________________________________________________________")
		// Tampilan ke user
		fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %+15.2f | %-30d |", A[hasil].kode, A[hasil].nama, A[hasil].harga/1000, A[hasil].perubahan_persentase, A[hasil].volume)
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
	fmt.Println("________________________________________________________________")
	fmt.Printf("| %-6s | %-8s | %-10s | %-10s | %-10s |\n", "Kode", "Lembar", "Harga/Unit", "Perubahan %", "Total")
	fmt.Println("________________________________________________________________")

	// Tampilan ke user
	var i int
	for i = 0; i < nSaham; i++ {
		var banyak_kepunyaan int
		banyak_kepunyaan = ownedSaham[i]
		if banyak_kepunyaan > 0 {
			var total_kepunyaan float64 = float64(banyak_kepunyaan) * listSaham[i].harga
			fmt.Printf("| %-6s | %-8d | Rp%-10.3f | %+10.2f | Rp%-10.3f |\n", listSaham[i].kode, banyak_kepunyaan, listSaham[i].harga/1000, listSaham[i].perubahan_persentase, total_kepunyaan/1000)
		}
	}
	fmt.Print("\n Ketik X untuk kembali > ")
	var kembali string
	fmt.Scan(&kembali)

}

func histori_transaksi(h *[nHistori]arrTransaksi, hitungHistori int) {
	// menu histori transaksi
	fmt.Println("Berikut merupakan histori transaksi anda : ")
	fmt.Printf("%s\t %s\t %s\t\t\t %s\t %s\t %s\n", "Jenis", "Kode", "Nama Perusahaan", "Jumlah", "Harga/Unit", "Total")

	var i int
	for i = 0; i < hitungHistori; i++ {
		fmt.Printf("%s\t %s\t %s\t\t %d\t %.2f\t\t %.2f\n",
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

}

func bantuan(pilih_bantuan int) {
	// menu bantuan
	for pilih_bantuan != 3 {
		fmt.Println("Bantuan")
		fmt.Println("1. Tentang aplikasi")
		fmt.Println("2. Cara menggunakan aplikasi")
		fmt.Println("3. Kembali")
		fmt.Print("Masukkan pilihan menu > ")
		fmt.Scan(&pilih_bantuan)

		switch pilih_bantuan {
		case 1:
			var opsi_about int
			fmt.Println("+-------------------------------------------------------+")
			fmt.Println("| Aplikasi ini adalah aplikasi simulasi pasar saham     |")
			fmt.Println("| virtual yang dibuat dengan menggunakan bahasa Go.     |")
			fmt.Println("| Aplikasi ini dibuat sebagai tugas besar Algoritma     |")
			fmt.Println("| dan Pemrograman 2.                                    |")
			fmt.Println("|                                                       |")
			fmt.Println("| Spesifikasi:                                          |")
			fmt.Println("| a. Tambah, ubah, hapus transaksi beli/jual saham.     |")
			fmt.Println("| b. Hitung nilai portofolio berdasarkan fluktuasi.     |")
			fmt.Println("| c. Cari saham dengan Sequential & Binary Search.      |")
			fmt.Println("| d. Urut saham berdasarkan harga/volume transaksi.     |")
			fmt.Println("| menggunakan Selection & Insertion Sort                |")
			fmt.Println("| Pembuat: Kanaka Pradipta Arya Wismaya & Farrel Malik  |")
			fmt.Println("| Pirade                                                |")
			fmt.Println("+-------------------------------------------------------+")
			fmt.Println()
			fmt.Print("Untuk kembali, pilih 0 > ")
			fmt.Scan(&opsi_about)

		case 2:
			var opsi_tutorial int
			fmt.Println("+-------------------------------------------------------+")
			fmt.Println("| Cara menggunakan aplikasi:                            |")
			fmt.Println("| 1. Saldo: Cek saldo, top up, withdraw.                |")
			fmt.Println("| 2. Daftar saham: Lihat, cari, urut daftar saham.      |")
			fmt.Println("| 3. Transaksi saham: Beli/jual dengan pencarian & urut.|")
			fmt.Println("| 4. Portofolio: Lihat portofolio Anda.                 |")
			fmt.Println("| 5. Histori: Lihat riwayat transaksi.                  |")
			fmt.Println("| 6. Bantuan: Menu bantuan.                             |")
			fmt.Println("| 7. Keluar: Tutup aplikasi.                            |")
			fmt.Println("+-------------------------------------------------------+")
			fmt.Println()
			fmt.Print("Untuk kembali, ketik 0 > ")
			fmt.Scan(&opsi_tutorial)
		case 3:
		}
	}
}

func penjual(pilih_menu_jual int) {
	for pilih_menu_jual != 3 {
		pilih_menu_jual = inputInt("Pilih menu penjual (1.Hapus Saham/2.Tambah Saham/3.Kembali) > ")
		switch pilih_menu_jual {
		case 1:
			// menu hapus saham
			var kode string
			fmt.Print("masukkan kode saham yang ingin dihapus > ")
			fmt.Scan(&kode)
			var idx int = sequential_search(listSaham, kode)
			if idx == -1 || idx >= countSaham {
				fmt.Println("Kode saham tidak ditemukan.")
			} else {
				var i int = idx
				for i < countSaham-1 {
					listSaham[i] = listSaham[i+1]
					ownedSaham[i] = ownedSaham[i+1]
					i++
				}
				countSaham--
			}
			fmt.Println("Saham berhasil dihapus.")
		case 2:
			// menu tambah saham
			if countSaham < nSaham {
				var kode, nama string
				var harga, perubahan_persentase float64
				var volume int

				fmt.Print("Masukkan kode saham > ")
				fmt.Scan(&kode)
				fmt.Print("Masukkan nama perusahaan (gunakan underscore untuk spasi) > ")
				fmt.Scan(&nama)
				fmt.Print("Masukkan harga saham > ")
				fmt.Scan(&harga)
				fmt.Print("Masukkan perubahan % saham (gunakan tanda +/-) > ")
				fmt.Scan(&perubahan_persentase)
				fmt.Print("Masukkan volume saham > ")
				fmt.Scan(&volume)

				// menyimpan data saham baru ke dalam listSaham
				listSaham[countSaham].kode = kode
				listSaham[countSaham].nama = nama
				listSaham[countSaham].harga = harga
				listSaham[countSaham].perubahan_persentase = perubahan_persentase
				listSaham[countSaham].volume = volume

				ownedSaham[countSaham] = 0 // jumlah saham yang dimiliki awalnya 0
				countSaham++               // menambah jumlah saham yang ada
				fmt.Println("Saham berhasil ditambahkan.")
			} else {
				fmt.Println("Market sudah penuh.")
			}
		case 3:
		}
	}
}
