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

func tampilkanKualitasTidur(data *[MaksData]CatatanTidur, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data tidur yang dimasukkan.")
		return
	}

	fmt.Println("\n======================== KUALITAS TIDUR =======================")
	var i int
	for i := 0; i < jumlah; i++ {
		jam := data[i].DurasiMenit / 60
		menit := data[i].DurasiMenit % 60

		var kualitas string
		if data[i].JamTidur <= 21 && jam == 8 {
			kualitas = "Sangat Baik"
		} else if data[i].JamTidur > 21 || jam == 8 {
			kualitas = "Baik"
		} else {
			kualitas = "Buruk"
		}

		fmt.Printf("\n ✨ HARI KE-%d ✨\n", i+1)
		fmt.Printf("     📌 Tanggal : %s\n", data[i].Tanggal)
		fmt.Printf("     ⏰ Jam Tidur -> %02d.%02d\n", data[i].JamTidur, data[i].MenitTidur)
		fmt.Printf("     ⏳ Durasi Tidur -> %d jam %d menit\n", jam, menit)
		fmt.Printf("     🔍 Kualitas Tidur -> %s\n", kualitas)
	}

	fmt.Println("\n ✨🔎 SCREENING 🔍✨")
	fmt.Print("     Apakah Anda merasa segar saat bangun tidur (Ya/Tidak)? ")
	fmt.Scan(&data[i].Segar)
	fmt.Print("     Apakah Anda merasa ngantuk setiap menjalani aktivitas (Ya/Tidak)? ")
	fmt.Scan(&data[i].Ngantuk)
	fmt.Print("     Apakah Anda merasa tidur terganggu oleh mimpi buruk, suara, atau gangguan lainnya (Ya/Tidak)? ")
	fmt.Scan(&data[i].Terganggu)
	fmt.Print("     Seberapa sering Anda minum kopi (Sering/Kadang/Tidak)? ")
	fmt.Scan(&data[i].Kopi)
}

func tampilkanRiwayatTidur(data [MaksData]CatatanTidur, jumlah int, menuPilihan int) {
	var i int
	awal := 0
	if jumlah > 7 {
		awal = jumlah - 7
	}
	totalDurasi := 0
	if menuPilihan != 5 {
		fmt.Println("\n📝 Rekapitulasi 7 Hari Terakhir:")
	}
	for i := awal; i < jumlah; i++ {
		jam := data[i].DurasiMenit / 60
		menit := data[i].DurasiMenit % 60
		fmt.Printf(" %s -> %d jam %d menit\n", data[i].Tanggal, jam, menit)
		totalDurasi += data[i].DurasiMenit
	}
	jumlahHari := jumlah
	if jumlahHari > 7 {
		jumlahHari = 7
	}
	if menuPilihan != 5 {
		rata := totalDurasi / jumlahHari
		fmt.Printf("\n📊 Rata-rata durasi tidur per minggu: %d jam %d menit\n", rata/60, rata%60)
		fmt.Println(" ")
		if data[i].Segar == "Tidak" {
			fmt.Println(" Penyebab :")
			fmt.Println("  1. Durasi tidur yang tidak cukup dan umumnya butuh 6-8 jam tidur per-malam")
			fmt.Println("  2. Tidur terlalu larut dan tidak teratur")
			fmt.Println("  3. Stres atau Overthinking")
			fmt.Println(" Saran : Tidur dan bangun pada jam yang sama setiap hari")
		}
		if data[i].Ngantuk == "Ya" {
			fmt.Println(" ")
			fmt.Println(" Penyebab :")
			fmt.Println("  1. Mengonsumsi makanan dengan kadar gula yang tinggi")
			fmt.Println("  2. Kurang paparan sinar matahari pagi")
			fmt.Println(" Saran : Aktif bergerak di siang hari, mengonsumsi makanan yang seimbang,")
			fmt.Println("         dapatkan paparan cahaya alami di pagi hari, dan power nap (tidur")
			fmt.Println("         siang singkat) 10-20 menit.")
		}
		if data[i].Terganggu == "Ya" {
			fmt.Println(" ")
			fmt.Println(" Penyebab :")
			fmt.Println("  1. Kondisi kamar yang tidak kondusif")
			fmt.Println("  2. Pengaruh obat-obatan tertentu")
			fmt.Println("  3. Stres dan kecemasan")
			fmt.Println(" Saran : Hindari menatap layar sebelum tidur, serta hindari konsumsi")
			fmt.Println("         kafein dan alkohol di malam hari.")
		}
		if data[i].Kopi == "Sering" {
			fmt.Println(" ")
			fmt.Println("Saran : Batasi konsumsi kopi maksimal 1-2 cangkir kopi per-hari,")
			fmt.Println("        karena dapat menyebabkan ketergantungan, sulit tidur, dan")
			fmt.Println("        gangguan pencernaan")
		} else if data[i].Kopi == "Kadang" {
			fmt.Println(" ")
			fmt.Println("  Tetap jaga konsumsi kopi agar tidak menjadi kebiasaan harian yang berlebihan.")
		} else {
			fmt.Println(" ")
		}
	}
	fmt.Println(" ")
}

func Insertionsort(A *[MaksData]CatatanTidur, jumlah int) {
	fmt.Println(" ")
	fmt.Println(" Terbesar ke Terkecil")
	for pass := 1; pass < jumlah; pass++ {
		temp := A[pass]
		i := pass

		for i > 0 && temp.DurasiMenit > A[i-1].DurasiMenit {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
	tampilkanRiwayatTidur(*A, jumlah, 5)

}

func DataJadwalTidur(A *[MaksData]CatatanTidur, jumlah *int) {
	for i := 0; i < *jumlah; i++ {
		jam := A[i].DurasiMenit / 60
		menit := A[i].DurasiMenit % 60
		fmt.Println(" ")
		fmt.Printf("    ✨ HARI KE-%d ✨\n", i+1)
		fmt.Printf("      📌 Tanggal -> %s\n", A[i].Tanggal)
		fmt.Printf("      ⏰ Jam Tidur -> %02d.%02d\n", A[i].JamTidur, A[i].MenitTidur)
		fmt.Printf("      💡 Jam Bangun -> %02d.%02d\n", A[i].JamBangun, A[i].MenitBangun)
		fmt.Printf("      ⏳ Durasi Tidur -> %d jam %d menit\n", jam, menit)
		fmt.Println(" ")
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
		if ubah >= 1 && ubah <= *jumlah {
			idx := ubah - 1
			fmt.Print("     Tanggal : ")
			fmt.Scan(&A[idx].Tanggal)
			fmt.Print("     Jam Tidur : ")
			fmt.Scan(&A[idx].JamTidur, &A[idx].MenitTidur)
			fmt.Print("     Jam Bangun : ")
			fmt.Scan(&A[idx].JamBangun, &A[idx].MenitBangun)
			hitungDurasi(&A[idx])
		} else {
			fmt.Println("    Hari tidak valid.")
		}
	} else if poin == 2 {
		var hapus int
		fmt.Print("     Hari ke : ")
		fmt.Scan(&hapus)
		if hapus >= 1 && hapus <= *jumlah {
			idx := hapus - 1
			for i := idx; i < *jumlah-1; i++ {
				A[i] = A[i+1]
			}
			*jumlah--
		} else {
			fmt.Println("     Hari tidak valid.")
		}

	}
}

func cariTanggal(A [MaksData]CatatanTidur, jumlah int, target string) {
	for i := 0; i < jumlah; i++ {
		if A[i].Tanggal == target {
			jam := A[i].DurasiMenit / 60
			menit := A[i].DurasiMenit % 60
			fmt.Printf("      ⏰ Jam Tidur -> %02d.%02d\n", A[i].JamTidur, A[i].MenitTidur)
			fmt.Printf("      💡 Jam Bangun -> %02d.%02d\n", A[i].JamBangun, A[i].MenitBangun)
			fmt.Printf("      ⏳ Durasi Tidur -> %d jam %d menit\n", jam, menit)
		}
	}

}

func main() {
	var dataTidur [MaksData]CatatanTidur
	var jumlahData int
	var pilihan int = 0
	fmt.Println("\n     APLIKASI PEMANTAUAN KESEHATAN DAN POLA TIDUR SEDERHANA")
	for pilihan != 7 {
		fmt.Println(" ")
		fmt.Println("  =============================================================")
		fmt.Println("                          DAFTAR PILIHAN")
		fmt.Println("  =============================================================")
		fmt.Println("   1. Jadwal Tidur")
		fmt.Println("   2. Lihat Jadwal Tidur")
		fmt.Println("   3. Kualitas Tidur")
		fmt.Println("   4. Riwayat Tidur")
		fmt.Println("   5. Urutkan Data")
		fmt.Println("   6. Cari Data")
		fmt.Println("   7. Keluar")
		fmt.Println("  =============================================================")
		fmt.Print("   Silakan pilih salah satu opsi yang anda butuhkan : ")

		fmt.Scan(&pilihan)

		if pilihan == 1 {

			fmt.Println("\n  ======================= JADWAL TIDUR ========================")
			fmt.Print("   Berapa hari data untuk jadwal tidur yang ingin Anda masukan? ")
			fmt.Scan(&jumlahData)

			for i := 0; i < jumlahData; i++ {
				fmt.Printf("\n   HARI KE-%d\n", i+1)
				fmt.Print("   Tanggal : ")
				fmt.Scan(&dataTidur[i].Tanggal)
				fmt.Print("   Jam Tidur : ")
				fmt.Scan(&dataTidur[i].JamTidur, &dataTidur[i].MenitTidur)
				fmt.Print("   Jam Bangun : ")
				fmt.Scan(&dataTidur[i].JamBangun, &dataTidur[i].MenitBangun)
				hitungDurasi(&dataTidur[i])
			}
		} else if pilihan == 2 {
			DataJadwalTidur(&dataTidur, &jumlahData)
		} else if pilihan == 3 {
			tampilkanKualitasTidur(&dataTidur, jumlahData)
		} else if pilihan == 4 {
			fmt.Println("\n======================== RIWAYAT TIDUR ========================")
			tampilkanRiwayatTidur(dataTidur, jumlahData, 0)
		} else if pilihan == 5 {
			Insertionsort(&dataTidur, jumlahData)
		} else if pilihan == 6 {
			fmt.Print("\n  🔍 Masukkan tanggal yang ingin dicari (yyyy-mm-dd): ")
			var target string
			fmt.Scan(&target)
			cariTanggal(dataTidur, jumlahData, target)
		} else {
			fmt.Println(" ")
		}
	}

}
