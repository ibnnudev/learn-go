package main

func luasPersegiPanjang(p, l int) int {
	return p * l
}

func isOdorEven(n int) bool {
	return n%2 == 0
}

func main() {
	var result int
	result = luasPersegiPanjang(5, 10)
	println("Luas Persegi Panjang:", result)

	var result2 bool
	result2 = isOdorEven(4)
	println("Apakah 4 adalah bilangan genap?", result2)
	println("Apakah 5 adalah bilangan genap?", isOdorEven(5))

	var nilaiAwal int = 10
	var nilaiAkhir int = 30

	for i := nilaiAwal; i < nilaiAkhir; i++ {
		if i%2 == 0 {
			println(i, "adalah bilangan genap")
		} else {
			println(i, "adalah bilangan ganjil")
		}
	}
	println("Proses selesai.")
}
