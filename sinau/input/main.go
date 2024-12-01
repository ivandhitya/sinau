package main

import (
	"fmt"
	"os"
)

func hitungPersegiPanjang(p int, l int) int {
	return p * l
}
func hitungSegitiga(a float32, t float32) float32 {
	return 0.5 * a * t
}

func main() {

	for {
		var menu string
		fmt.Println("=======Pilih Menu========")
		fmt.Println("1. Hitung Persegi Panjang")
		fmt.Println("2. Hitung Segitiga")
		fmt.Println("3. Keluar")
		fmt.Print("Ketik angka lalu press enter:")
		fmt.Scan(&menu)
		switch menu {
		case "1":
			var p, l int
			fmt.Print("Masukan Panjang: ")
			fmt.Scan(&p)
			fmt.Print("Masukan Lebar: ")
			fmt.Scan(&l)
			fmt.Printf("Luas persegi panjang dengan panjang %d dan lebar %d adalah %d\n", p, l, hitungPersegiPanjang(p, l))
		case "2":
			var a, t float32
			fmt.Print("Masukan Alas: ")
			fmt.Scan(&a)
			fmt.Print("Masukan Tinggi: ")
			fmt.Scan(&t)
			fmt.Printf("Luas segitiga dengan alas %f dan tinggi %f adalah %f\n", a, t, hitungSegitiga(a, t))
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Menu tidak tersedia")
		}
	}

}
