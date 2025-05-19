package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// aplikasi simulasi pasar saham virtual

type saham struct {
	kode                 string
	nama                 string
	harga                float64
	perubahan_persentase float64
	volume               int
}

var daftarSaham = []saham{
    {kode: "ASII", nama: "ASTRA_INTERNATIONAL"},
    {kode: "BBRI", nama: "BANK_RAKYAT_INDONESIA"},
    {kode: "BMRI", nama: "BANK_MANDIRI_(PERSERO)_TBK"},
    {kode: "TLKM", nama: "TELKOM_INDONESIA_(PERSERO)_TBK"},
    {kode: "UNTR", nama: "UNITED_TRACTORS"},
    {kode: "BBCA", nama: "BANK_CENTRAL_ASIA"},
    {kode: "AMRT", nama: "SUMBER_ALFARIA_TRIJAYA_TBK"},
    {kode: "INDF", nama: "PT_INDOFOOD_SUKSES_MAKMUR_TBK"},
    {kode: "GGRM", nama: "GUDANG_GARAM_TBK"},
    {kode: "BBNI", nama: "PT_BANK_NEGARA_INDONESIA_(PERSERO)"},
    {kode: "AADI", nama: "ADARO_ANDALAN_INDONESIA_TBK"},
    {kode: "SMAR", nama: "SINAR_MAS_AGRO_RESOURCES_TECHNOLOGY"},
    {kode: "ICBP", nama: "INDOFOOD_CBP_SUKSES_MAKMUR_TBK_PT"},
    {kode: "ANTM", nama: "ANEKA_TAMBANG"},
    {kode: "CPIN", nama: "CHAROEN_POKPHAND_INDONESIA"},
    {kode: "ERAA", nama: "ERAJAYA_TBK"},
    {kode: "PGAS", nama: "PERUSAHAAN_GAS_NEGARA_TBK"},
    {kode: "ISAT", nama: "INDOSAT_(PT_INDOSAT_TBK)"},
    {kode: "JPFA", nama: "JAPFA_COMFEED_INDONESIA"},
    {kode: "BYAN", nama: "BAYAN_RESOURCES_TBK"},
    {kode: "GIAA", nama: "GARUDA_INDONESIA_(PERSERO)_TBK"},
    {kode: "INKP", nama: "INDAH_KIAT_PULP_&_PAPER"},
    {kode: "HMSP", nama: "HANJAYA_MANDALA_SAMPOERNA"},
    {kode: "DSSA", nama: "DIAN_SWASTATIKA_SENTOSA_TBK"},
    {kode: "GEMS", nama: "GOLDEN_ENERGY_MINES_TBK"},
    {kode: "PTBA", nama: "BUKIT_ASAM_TBK"},
    {kode: "AMMN", nama: "AMMAN_MINERAL_INTERNASIONAL_TBK"},
    {kode: "INDY", nama: "INDIKA_ENERGY_TBK"},
    {kode: "AKRA", nama: "AKR_CORPORINDO_TBK"},
    {kode: "MEDC", nama: "MEDCO_ENERGI_INTERNATIONAL"},
    {kode: "MAPI", nama: "MITRA_ADIPERKASA_TBK"},
    {kode: "BRPT", nama: "BARITO_PACIFIC_TBK"},
    {kode: "ITMG", nama: "INDO_TAMBANGRAYA_MEGAH_TBK"},
    {kode: "SMGR", nama: "SEMEN_INDONESIA_(PERSERO)_TBK"},
    {kode: "MYOR", nama: "MAYORA_INDAH"},
    {kode: "MDKA", nama: "MERDEKA_COPPER_GOLD_TBK_PT"},
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

// variabel global
var jumSaldo int

func main() {
	// program main merupakan interface awal aplikasi saat dibuka
	rand.Seed(time.Now().UnixNano())
	randomisasi()

	for {
		var pilihan int
		fmt.Print(asciiArt + "\n")
		fmt.Println("Pilih menu")
		fmt.Println("1. Saldo anda")
		fmt.Println("2. Daftar saham")
		fmt.Println("3. Transaksi saham")
		fmt.Println("4. Portofolio")
		fmt.Println("5. Histori transaksi")
		fmt.Println("6. Bantuan")
		fmt.Println("7. Keluar")
		fmt.Print("Masukkan pilihan menu > ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			// menu saldo
			saldo()
		case 2:
			// menu daftar saham
			daftar_saham()
		case 3:
			// menu transaksi saham
			transaksi_saham()
		case 4:
			// menu portofolio
		case 5:
			// menu histori transaksi
		case 6:
			// menu bantuan
			var pilih_bantuan int
			bantuan(pilih_bantuan)
		case 7:
			// menu keluar
			os.Exit(0)

		default:
			fmt.Println("Pilihan anda tidak valid, silahkan coba lagi")
			return
		}
	}

}

func randomisasi() {
	var i int
	// fungsi randomisasi untuk mengacak harga, perubahan persentase, dan volume saham
	for i = 0; i < len(daftarSaham); i++ {
		// Randomisasi harga saham dengan range sampai 21000
		daftarSaham[i].harga = float64(rand.Intn(21000)) + 1
		if daftarSaham[i].harga < 1 {
			daftarSaham[i].harga = 1 // harga saham tidak boleh kurang dari 1
		}

		//Randomisasi perubahan persentase saham
		daftarSaham[i].perubahan_persentase = float64(rand.Intn(100)) + 1
		if daftarSaham[i].perubahan_persentase < 1 {
			daftarSaham[i].perubahan_persentase = 1 // perubahan persentase tidak boleh kurang dari 1
		}

		// Randomisasi volume saham dengan fluktuasi kurang lebih antara 100 hingga 3000000000
		daftarSaham[i].volume = rand.Intn(3000000000) + 100
	}
}

func saldo() {
	// prosedur saldo untuk menampilkan saldo, top up, dan withdraw
	var pilihan int
	fmt.Printf("Saldo anda adalah %d \n", jumSaldo)
	fmt.Println("1. Top up")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Kembali")
	fmt.Print("Masukkan pilihan menu > ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		// melakukan top up
		var topup int
		fmt.Print("Masukkan jumlah top up > ")
		fmt.Scan(&topup)
		if topup < 0 {
			fmt.Println("Jumlah top up tidak valid")
			saldo()
		}
		jumSaldo += topup
		fmt.Printf("Saldo anda sekarang adalah %d \n", jumSaldo)
		saldo()
	case 2:
		// melakukan withdraw
		var withdraw int
		fmt.Print("Masukkan jumlah withdraw > ")
		fmt.Scan(&withdraw)
		if withdraw > jumSaldo {
			fmt.Println("Saldo anda tidak cukup")
			saldo()
		}
		jumSaldo -= withdraw
		fmt.Printf("Saldo anda sekarang adalah %d \n", jumSaldo)
		saldo()
	case 3:
		// kembali ke main menu
		return
	}
}

func sequential_search(daftarSaham []saham, kode_saham string) int {
	var iterasi_kode_seq int
	var found bool = false
	for !found && iterasi_kode_seq < len(daftarSaham) {
		found = daftarSaham[iterasi_kode_seq].kode == kode_saham
		iterasi_kode_seq++
	}
	if found {
		return iterasi_kode_seq - 1
	} else {
		return -1
	}
}

func sel_sort_for_bin_search(daftarSaham []saham) {
	// pengurutan daftar saham berdasarkan nama perusahaan menggunakan selection sort
	var iterasi_nama_bin1, iterasi_nama_bin2 int
	var temp saham
	var min int
	for iterasi_nama_bin1 = 0; iterasi_nama_bin1 < len(daftarSaham); iterasi_nama_bin1++ {
		min = iterasi_nama_bin1
		for iterasi_nama_bin2 = iterasi_nama_bin1 + 1; iterasi_nama_bin2 < len(daftarSaham); iterasi_nama_bin2++ {
			if daftarSaham[iterasi_nama_bin2].nama < daftarSaham[min].nama {
				min = iterasi_nama_bin2
			}
		}
		temp = daftarSaham[iterasi_nama_bin1]
		daftarSaham[iterasi_nama_bin1] = daftarSaham[min]
		daftarSaham[min] = temp
	}
}

func binary_search(daftarSaham []saham, nama_saham string) int {
	// menggunakan binary search untuk mencari nama perusahaan
	var left, right, mid int
	left = 0
	right = len(daftarSaham) - 1
	for left <= right {
		mid = (left + right) / 2
		if nama_saham == daftarSaham[mid].nama {
			return mid
		} else if nama_saham > daftarSaham[mid].nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func daftar_saham() {
	sel_sort_for_bin_search(daftarSaham)

	// beberapa daftar saham yang bisa diakses
	for {
		var i int
		fmt.Println("Daftar saham : ")
		fmt.Println("_______________________________________________________________________________________________________________________")
		fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
		fmt.Println("_______________________________________________________________________________________________________________________")

		for i = 0; i < len(daftarSaham); i++ {
			// Tampilan ke user
			fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", daftarSaham[i].kode, daftarSaham[i].nama, daftarSaham[i].harga/1000, daftarSaham[i].perubahan_persentase, daftarSaham[i].volume)
			fmt.Println()
		}
		fmt.Println("_______________________________________________________________________________________________________________________")
		fmt.Println()
		fmt.Println("Pilih menu")
		fmt.Println("1. Cari saham")
		fmt.Println("2. Urutkan saham")
		fmt.Println("3. Kembali")
		fmt.Print("Masukkan pilihan menu > ")

		var pilih_daftar_menu_saham int
		fmt.Scan(&pilih_daftar_menu_saham)

		switch pilih_daftar_menu_saham {
		case 1:
			// menu searching saham
			var pilih_search int
			fmt.Println("Pilih metode pencarian")
			fmt.Println("1. Pencarian dengan kode saham")
			fmt.Println("2. Pencarian dengan nama saham")
			fmt.Println("3. Kembali")
			fmt.Print("Masukkan pilihan menu > ")
			fmt.Scan(&pilih_search)

			switch pilih_search {
			case 1:
				// pencarian dengan kode saham dengan sequential search
				fmt.Print("Masukkan kode saham > ")
				var kode_saham string
				fmt.Scan(&kode_saham)
				var hasil int = sequential_search(daftarSaham, kode_saham)
				if hasil == -1 {
					fmt.Println("Kode saham tidak ditemukan")
				} else {
					fmt.Println("Berikut merupakan hasil pencarian : ")
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
					fmt.Println("_______________________________________________________________________________________________________________________")
					// Tampilan ke user
					fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", daftarSaham[hasil].kode, daftarSaham[hasil].nama, daftarSaham[hasil].harga/1000, daftarSaham[hasil].perubahan_persentase, daftarSaham[hasil].volume)
					fmt.Println()
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Print("Ketik apa saja untuk kembali > ")
					var kembali int
					fmt.Scan(&kembali)
					return
				}
			case 2:
				// pencarian dengan nama perusahaan saham dengan binary search
				fmt.Println("Gunakan underscore (_) untuk spasi")
				fmt.Print("Masukkan nama perusahaan > ")
				var nama_saham_raw string
				fmt.Scan(&nama_saham_raw)

				// Memanggil fungsi binary search untuk mencari nama perusahaan
				var hasil int = binary_search(daftarSaham, nama_saham_raw)
				if hasil == -1 {
					fmt.Println("Nama perusahaan tidak ditemukan")
				} else {
					fmt.Println("Berikut merupakan hasil pencarian : ")
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
					fmt.Println("_______________________________________________________________________________________________________________________")
					// Tampilan ke user
					fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", daftarSaham[hasil].kode, daftarSaham[hasil].nama, daftarSaham[hasil].harga/1000, daftarSaham[hasil].perubahan_persentase, daftarSaham[hasil].volume)
					fmt.Println()
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Print("Ketik apa saja untuk kembali > ")
					var kembali int
					fmt.Scan(&kembali)
					return
				}
			case 3:
				// kembali ke menu daftar saham
				daftar_saham()
			}
		case 2:
			// menu sorting saham
			var pilih_sort int
			fmt.Println("Pilih metode pengurutan")
			fmt.Println("1. Pengurutan dengan harga saham")
			fmt.Println("2. Pengurutan dengan volume saham")
			fmt.Println("3. Kembali")
			fmt.Print("Masukkan pilihan menu > ")
			fmt.Scan(&pilih_sort)

			switch pilih_sort {
			case 1:
				// pengurutan dengan harga saham dengan selection sort
				fmt.Println("Mau diurutkan dari harga tertinggi atau terendah?")
				fmt.Println("1. Tertinggi")
				fmt.Println("2. Terendah")
				fmt.Print("Masukkan pilihan menu > ")
				var pilih_harga int
				fmt.Scan(&pilih_harga)

				if pilih_harga == 1 {
					// pengurutan dari harga tertinggi
					var temp saham
					var n int = len(daftarSaham)
					var iterasi1, iterasi2, max int
					for iterasi1 = 0; iterasi1 < n-1; iterasi1++ {
						max = iterasi1
						for iterasi2 = iterasi1 + 1; iterasi2 < n; iterasi2++ {
							if daftarSaham[iterasi2].harga > daftarSaham[max].harga {
								max = iterasi2
							}
						}
						temp = daftarSaham[iterasi1]
						daftarSaham[iterasi2] = daftarSaham[max]
						daftarSaham[max] = temp
					}
					// menampilkan hasil pengurutan
					fmt.Println("Berikut merupakan hasil pengurutan : ")
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
					fmt.Println("_______________________________________________________________________________________________________________________")
					for i = 0; i < len(daftarSaham); i++ {
						// Tampilan ke user
						fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", daftarSaham[i].kode, daftarSaham[i].nama, daftarSaham[i].harga/1000, daftarSaham[i].perubahan_persentase, daftarSaham[i].volume)
						fmt.Println()
					}
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Print("Ketik apa saja untuk kembali > ")
					var kembali int
					fmt.Scan(&kembali)
					return

				} else if pilih_harga == 2 {
					// pengurutan dari harga terendah
					var temp saham
					var n int = len(daftarSaham)
					var iterasi3, iterasi4, min int
					for iterasi3 = 0; iterasi3 < n-1; iterasi3++ {
						min = iterasi3
						for iterasi4 = iterasi3 + 1; iterasi4 < n; iterasi4++ {
							if daftarSaham[iterasi4].harga < daftarSaham[min].harga {
								min = iterasi4
							}
						}
						temp = daftarSaham[iterasi3]
						daftarSaham[iterasi3] = daftarSaham[min]
						daftarSaham[min] = temp
					}
					// menampilkan hasil pengurutan
					fmt.Println("Berikut merupakan hasil pengurutan : ")
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
					fmt.Println("_______________________________________________________________________________________________________________________")
					for i = 0; i < len(daftarSaham); i++ {
						// Tampilan ke user
						fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", daftarSaham[i].kode, daftarSaham[i].nama, daftarSaham[i].harga/1000, daftarSaham[i].perubahan_persentase, daftarSaham[i].volume)
						fmt.Println()
					}
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Print("Ketik apa saja untuk kembali > ")
					var kembali int
					fmt.Scan(&kembali)
					return

				} else {
					fmt.Println("Pilihan tidak valid, silahkan coba lagi")
					return
				}
			case 2:
				// pengurutan dengan volume saham dengan insertion sort
				fmt.Println("Mau diurutkan dari volume tertinggi atau terendah?")
				fmt.Println("1. Tertinggi")
				fmt.Println("2. Terendah")
				fmt.Print("Masukkan pilihan menu > ")
				var pilih_volume int
				fmt.Scan(&pilih_volume)

				if pilih_volume == 1 {
					// pengurutan dari harga tertinggi
					var n int = len(daftarSaham)
					var iterasi5, iterasi6 int
					for iterasi5 = 1; iterasi5 < n; iterasi5++ {
						var temp saham = daftarSaham[iterasi5]
						iterasi6 = iterasi5 - 1
						for iterasi6 >= 0 && daftarSaham[iterasi6].volume < temp.volume {
							daftarSaham[iterasi6+1] = daftarSaham[iterasi6]
							iterasi6--
						}
						daftarSaham[iterasi6+1] = temp
					}

					// menampilkan hasil pengurutan
					fmt.Println("Berikut merupakan hasil pengurutan : ")
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
					fmt.Println("_______________________________________________________________________________________________________________________")
					for i = 0; i < len(daftarSaham); i++ {
						// Tampilan ke user
						fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", daftarSaham[i].kode, daftarSaham[i].nama, daftarSaham[i].harga/1000, daftarSaham[i].perubahan_persentase, daftarSaham[i].volume)
						fmt.Println()
					}
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Print("Ketik apa saja untuk kembali > ")
					var kembali int
					fmt.Scan(&kembali)
					return

				} else if pilih_volume == 2 {
					// pengurutan dari harga terendah
					var n int = len(daftarSaham)
					var iterasi7, iterasi8 int
					for iterasi7 = 1; iterasi7 < n; iterasi7++ {
						var temp saham = daftarSaham[iterasi7]
						iterasi8 = iterasi7 - 1
						for iterasi8 >= 0 && daftarSaham[iterasi8].volume > temp.volume {
							daftarSaham[iterasi8+1] = daftarSaham[iterasi8]
							iterasi8--
						}
						daftarSaham[iterasi8+1] = temp
					}

					// menampilkan hasil pengurutan
					fmt.Println("Berikut merupakan hasil pengurutan : ")
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
					fmt.Println("_______________________________________________________________________________________________________________________")
					for i = 0; i < len(daftarSaham); i++ {
						// Tampilan ke user
						fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", daftarSaham[i].kode, daftarSaham[i].nama, daftarSaham[i].harga/1000, daftarSaham[i].perubahan_persentase, daftarSaham[i].volume)
						fmt.Println()
					}
					fmt.Println("_______________________________________________________________________________________________________________________")
					fmt.Print("Ketik apa saja untuk kembali > ")
					var kembali int
					fmt.Scan(&kembali)
					return

				} else {
					fmt.Println("Pilihan tidak valid, silahkan coba lagi")
					return
				}
			case 3:
				// kembali ke menu daftar saham
				daftar_saham()
			}
		case 3:
			// kembali ke main menu
			return
		}
	}
}

func transaksi_saham() {
	// menu transaksi saham
	fmt.Println("Transaksi saham")
	fmt.Println("1. Beli saham")
	fmt.Println("2. Jual saham")
	fmt.Println("3. Kembali")
	fmt.Print("Masukkan pilihan menu > ")
	var pilih_transaksi int
	fmt.Scan(&pilih_transaksi)
	switch pilih_transaksi {
	case 1:
		// menu beli saham
		fmt.Print("Masukkan kode saham > ")
		var beli_kode_saham string
		fmt.Scan(&beli_kode_saham)
		var hasil int = sequential_search(daftarSaham, beli_kode_saham)
		if hasil == -1 {
			fmt.Println("Kode saham tidak ditemukan")
			transaksi_saham()
		} else {
			fmt.Println("Berikut merupakan hasil pencarian : ")
			fmt.Println("_______________________________________________________________________________________________________________________")
			fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
			fmt.Println("_______________________________________________________________________________________________________________________")
			// Tampilan ke user
			fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", daftarSaham[hasil].kode, daftarSaham[hasil].nama, daftarSaham[hasil].harga/1000, daftarSaham[hasil].perubahan_persentase, daftarSaham[hasil].volume)
			fmt.Println()
			fmt.Println("_______________________________________________________________________________________________________________________")
		}

		fmt.Print("Masukkan jumlah saham yang ingin dibeli > ")
		var beli_jumlah_saham int
		fmt.Scan(&beli_jumlah_saham)

		// menampilkan saldo pengguna sekarang dengan mengambil dari variabel jumSaldo
		var total_harga_sementara float64 = float64(beli_jumlah_saham) * daftarSaham[hasil].harga
		fmt.Printf("Saldo anda sekarang adalah %d \n", jumSaldo)
		if total_harga_sementara > float64(jumSaldo) {
			fmt.Println("Saldo anda tidak cukup")
			transaksi_saham()
		} else if beli_jumlah_saham < 0 {
			fmt.Println("Jumlah saham tidak valid")
			transaksi_saham()
		} else if beli_jumlah_saham > daftarSaham[hasil].volume {
			fmt.Println("Jumlah saham yang ingin dibeli melebihi volume saham")
			transaksi_saham()
		} else {
			var total_harga float64 = float64(beli_jumlah_saham) * daftarSaham[hasil].harga
			jumSaldo -= int(total_harga)
			daftarSaham[hasil].volume -= beli_jumlah_saham

			fmt.Printf("Total harga yang harus dibayar adalah Rp%.2f \n", total_harga)
			fmt.Printf("Saldo anda sekarang adalah %d \n", jumSaldo)

			// melanjutkan pembayaran
			fmt.Print("Bayar ? (y/n) > ")
			var pilih_bayar string
			fmt.Scan(&pilih_bayar)
			if pilih_bayar == "y" || pilih_bayar == "Y" {
				fmt.Println("Pembayaran berhasil")
				daftarSaham[hasil].volume -= beli_jumlah_saham
				fmt.Printf("Sisa volume saham %s adalah %d \n", daftarSaham[hasil].kode, daftarSaham[hasil].volume)
				fmt.Printf("Sisa saldo anda adalah %d \n", jumSaldo)
				fmt.Println("Ketik apa saja untuk kembali > ")
				var kembali int
				fmt.Scan(&kembali)
				return
			} else if pilih_bayar == "n" || pilih_bayar == "N" {
				fmt.Println("Pembayaran dibatalkan")
				transaksi_saham()
			} else {
				fmt.Println("Pilihan tidak valid")
				transaksi_saham()
			}
		}

	case 2:
		// menu jual saham
		fmt.Print("Masukkan kode saham > ")
		var jual_kode_saham string
		fmt.Scan(&jual_kode_saham)
		var hasil int = sequential_search(daftarSaham, jual_kode_saham)
		if hasil == -1 {
			fmt.Println("Kode saham tidak ditemukan")
			transaksi_saham()
		} else {
			fmt.Println("Berikut merupakan hasil pencarian : ")
			fmt.Println("_______________________________________________________________________________________________________________________")
			fmt.Printf("| %-6s | %-40s | %-10s | %-15s | %-30s |\n", "Kode", "Nama Perusahaan", "Harga", "Perubahan %", "Volume")
			fmt.Println("_______________________________________________________________________________________________________________________")
			// Tampilan ke user
			fmt.Printf("| %-6s | %-40s | Rp%-10.3f | %-15.2f | %-30d |", daftarSaham[hasil].kode, daftarSaham[hasil].nama, daftarSaham[hasil].harga/1000, daftarSaham[hasil].perubahan_persentase, daftarSaham[hasil].volume)
			fmt.Println()
			fmt.Println("_______________________________________________________________________________________________________________________")
		}

		fmt.Print("Masukkan jumlah saham yang ingin dijual > ")
		var jual_jumlah_saham int
		fmt.Scan(&jual_jumlah_saham)
		// menampilkan saldo pengguna sekarang dengan mengambil dari variabel jumSaldo
		fmt.Printf("Saldo anda sekarang adalah %d \n", jumSaldo)
		if jual_jumlah_saham < 0 {
			fmt.Println("Jumlah saham tidak valid")
			transaksi_saham()
		} else if jual_jumlah_saham > daftarSaham[hasil].volume {
			fmt.Println("Jumlah saham yang ingin dijual melebihi volume saham")
			transaksi_saham()
		} else {
			var jual_pendapatan float64 = float64(jual_jumlah_saham) * daftarSaham[hasil].harga
			jumSaldo += int(jual_pendapatan)
			daftarSaham[hasil].volume += jual_jumlah_saham
			fmt.Printf("Volume saham %s sekarang adalah %d \n", daftarSaham[hasil].kode, daftarSaham[hasil].volume)
			var total_pendapatan float64 = float64(jual_jumlah_saham) * daftarSaham[hasil].harga
			fmt.Printf("Anda mendapat Rp%.2f \n", total_pendapatan)
			fmt.Printf("Saldo anda sekarang adalah %d \n", jumSaldo)
		}
	case 3:
		// kembali ke main menu
		return
	}
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
