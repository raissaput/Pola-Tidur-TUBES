package main

import "fmt"

const MaksData = 100

type CatatanTidur struct {
	Tanggal     string
	JamTidur    int
	MenitTidur  int
	JamBangun   int
	MenitBangun int
	DurasiMenit int
	Segar       string
	Ngantuk     string
	Terganggu   string
	Stres       string
	Kopi        string
}

func hitungDurasi(ct *CatatanTidur) {
	mulai := ct.JamTidur*60 + ct.MenitTidur
	selesai := ct.JamBangun*60 + ct.MenitBangun
	selisih := selesai - mulai
	if selisih < 0 {
		selisih += 24 * 60
	}
	ct.DurasiMenit = selisih
}

func rekap7Hari(data [MaksData]CatatanTidur, jumlah int) {
	fmt.Println("Rekapitulasi 7 hari terakhir:")
	awal := 0
	if jumlah > 7 {
		awal = jumlah - 7
	}
	for i := awal; i < jumlah; i++ {
		jam := data[i].DurasiMenit / 60
		menit := data[i].DurasiMenit % 60
		fmt.Printf("%s -> %d jam %d menit\n", data[i].Tanggal, jam, menit)
	}
}

func rataRataMingguan(data [MaksData]CatatanTidur, jumlah int) (int, int) {
	total := 0
	hari := jumlah
	if hari > 7 {
		hari = 7
	}
	for i := jumlah - hari; i < jumlah; i++ {
		total += data[i].DurasiMenit
	}
	rata := total / hari
	return rata / 60, rata % 60
}

func cariTanggal(data [MaksData]CatatanTidur, jumlah int, target string) int {
	bawah, atas := 0, jumlah-1
	for bawah <= atas {
		tengah := (bawah + atas) / 2
		if data[tengah].Tanggal == target {
			return tengah
		}
		if data[tengah].Tanggal < target {
			bawah = tengah + 1
		} else {
			atas = tengah - 1
		}
	}
	return -1
}

func urutkanTanggal(data *[MaksData]CatatanTidur, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		indeksMin := i
		for j := i + 1; j < jumlah; j++ {
			if data[j].Tanggal < data[indeksMin].Tanggal {
				indeksMin = j
			}
		}
		data[i], data[indeksMin] = data[indeksMin], data[i]
	}
}

func urutkanDurasi(data *[MaksData]CatatanTidur, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		indeksMin := i
		for j := i + 1; j < jumlah; j++ {
			if data[j].DurasiMenit < data[indeksMin].DurasiMenit {
				indeksMin = j
			}
		}
		data[i], data[indeksMin] = data[indeksMin], data[i]
	}
}

func tampilkanSemua(data [MaksData]CatatanTidur, jumlah int) {
	fmt.Println("Riwayat Tidur:")
	for i := 0; i < jumlah; i++ {
		jam := data[i].DurasiMenit / 60
		menit := data[i].DurasiMenit % 60
		fmt.Printf("%d. %s | Durasi: %d jam %d menit\n",
			i+1, data[i].Tanggal, jam, menit)
	}
}

func tampilkanKualitasTidur(data [MaksData]CatatanTidur, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data tidur yang dimasukkan.")
		return
	}

	totalJamTidur := 0
	totalDurasi := 0
	for i := 0; i < jumlah; i++ {
		totalJamTidur += data[i].JamTidur
		totalDurasi += data[i].DurasiMenit
	}

	rataJamTidur := totalJamTidur / jumlah
	rataDurasi := totalDurasi / jumlah
	jamDurasi := rataDurasi / 60

	kualitas := "Baik"
	if rataJamTidur >= 22 || jamDurasi > 8 {
		kualitas = "Buruk"
	}

	fmt.Println("\n======================== KUALITAS TIDUR =======================")
	for i := 0; i < jumlah; i++ {
		jam := data[i].DurasiMenit / 60
		menit := data[i].DurasiMenit % 60
		fmt.Printf("\n ✨ HARI KE-%d ✨\n", i+1)
		fmt.Printf("     📌 Tanggal : %s\n", data[i].Tanggal)
		fmt.Printf("     ⏰ Jam Tidur -> %02d.%02d\n", data[i].JamTidur, data[i].MenitTidur)
		fmt.Printf("     ⏳ Durasi Tidur -> %d jam %d menit\n", jam, menit)
		fmt.Printf("     🔍 Kualitas Tidur -> %s\n", kualitas)
		fmt.Println(" ")

		fmt.Println(" ✨🔎 SCREENING 🔍✨")
		fmt.Print("     Apakah Anda merasa segar saat bangun tidur (Ya/Tidak)? ")
		fmt.Scan(&data[i].Segar)
		fmt.Print("     Apakah Anda merasa ngantuk di siang hari meskipun sudah tidur cukup (Ya/Tidak)? ")
		fmt.Scan(&data[i].Ngantuk)
		fmt.Print("     Apakah Anda merasa tidur terganggu oleh mimpi buruk, suara, atau gangguan lainnya (Ya/Tidak)? ")
		fmt.Scan(&data[i].Terganggu)
		fmt.Print("     Apakah Anda merasa stres atau cemas sebelum tidur (Ya/Tidak)? ")
		fmt.Scan(&data[i].Stres)
		fmt.Print("     Seberapa sering Anda minum kopi (Sering/Kadang/Tidak)? ")
		fmt.Scan(&data[i].Kopi)
	}

	fmt.Println("\nTekan ENTER untuk kembali ke menu utama...")
	fmt.Scanln()
	fmt.Scanln()
}

func main() {
	var dataTidur [MaksData]CatatanTidur
	var jumlahData int
	fmt.Println("\n     APLIKASI PEMANTAUAN KESEHATAN DAN POLA TIDUR SEDERHANA")
	for {
		fmt.Println(" ")
		fmt.Println("  =============================================================")
		fmt.Println("                          DAFTAR PILIHAN")
		fmt.Println("  =============================================================")
		fmt.Println("   1. Jadwal Tidur")
		fmt.Println("   2. Tampilkan Data Jadwal Tidur")
		fmt.Println("   3. Kualitas Tidur")
		fmt.Println("   4. Riwayat Tidur")
		fmt.Println("  =============================================================")
		fmt.Print("   Silahkan pilih salah satu opsi yang anda butuhkan : ")

		var pilihan int
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			// Input jadwal tidur
			fmt.Println("\n  ======================= JADWAL TIDUR ========================")
			// fmt.Println("• Format penulisan yang digunakan adalah tahun-bulan-hari
			// 			   (yyyy-mm-dd), contoh: 2024-05-11)")
			// fmt.Println("• Format penulisan jajam dan menit: hh mm (contoh: Jam Tidur: 22 30)")
			// fmt.Println("====================================================================")

			fmt.Print("   Berapa hari data untuk jadwal tidur yang ingin Anda masukan? ")
			fmt.Scan(&jumlahData)

			for i := 0; i < jumlahData; i++ {
				fmt.Printf("\n  HARI KE-%d\n", i+1)
				fmt.Print("  Tanggal : ")
				fmt.Scan(&dataTidur[i].Tanggal)
				fmt.Print("  Jam Tidur : ")
				fmt.Scan(&dataTidur[i].JamTidur, &dataTidur[i].MenitTidur)
				fmt.Print("  Jam Bangun : ")
				fmt.Scan(&dataTidur[i].JamBangun, &dataTidur[i].MenitBangun)
				hitungDurasi(&dataTidur[i])
			}
		} else if pilihan == 2 {
			for i := 0; i < jumlahData; i++ {
				jam := dataTidur[i].DurasiMenit / 60
				menit := dataTidur[i].DurasiMenit % 60
				fmt.Println(" ")
				fmt.Printf("    ✨ HARI KE-%d ✨\n", i+1)
				fmt.Printf("      📌 Tanggal -> %s\n", dataTidur[i].Tanggal)
				fmt.Printf("      ⏰ Jam Tidur -> %02d.%02d\n", dataTidur[i].JamTidur, dataTidur[i].MenitTidur)
				fmt.Printf("      💡 Jam Bangun -> %02d.%02d\n", dataTidur[i].JamBangun, dataTidur[i].MenitBangun)
				fmt.Printf("      ⏳ Durasi Tidur -> %d jam %d menit\n", jam, menit)
				// fmt.Println("_")
			}

			fmt.Println("  ======================================================")
			fmt.Println("   Opsi Tindakan :")
			fmt.Println("    1. Ubah Data")
			fmt.Println("    2. Hapus Data")
			fmt.Println("    3. Kembali ke Daftar Pilihan")
			fmt.Println("  ======================================================")
			fmt.Print("    Silahkan pilih salah satu opsi : ")
			var poin int
			fmt.Scan(&poin)

			if poin == 1 {
				var ubah int
				fmt.Println(" ")
				fmt.Print("     Hari ke : ")
				fmt.Scan(&ubah)
				if ubah >= 1 && ubah <= jumlahData {
					idx := ubah - 1
					fmt.Print("     Tanggal : ")
					fmt.Scan(&dataTidur[idx].Tanggal)
					fmt.Print("     Jam Tidur : ")
					fmt.Scan(&dataTidur[idx].JamTidur, &dataTidur[idx].MenitTidur)
					fmt.Print("     Jam Bangun : ")
					fmt.Scan(&dataTidur[idx].JamBangun, &dataTidur[idx].MenitBangun)
					hitungDurasi(&dataTidur[idx])
				} else {
					fmt.Println("    Hari tidak valid.")
				}
			} else if poin == 2 {
				var hapus int
				fmt.Print("     Hari ke : ")
				fmt.Scan(&hapus)
				if hapus >= 1 && hapus <= jumlahData {
					idx := hapus - 1
					for i := idx; i < jumlahData-1; i++ {
						dataTidur[i] = dataTidur[i+1]
					}
					jumlahData--
				} else {
					fmt.Println("     Hari tidak valid.")
				}

			}
		} else if pilihan == 3 {
			tampilkanKualitasTidur(dataTidur, jumlahData)
		}
	}

}
