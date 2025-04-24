package main

func main() {
	var siswa = [5]string{"ibnu", "budi", "andi", "siti", "ani"}
	for _, v := range siswa {
		println(v)
	}

	// multidimensional array
	const (
		rows = 2
		cols = 3
	)
	var matrix [rows][cols]int = [rows][cols]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			print(matrix[i][j])
			print(" ")
		}
		println()
	}
	println("Matrix printed successfully.")
}
