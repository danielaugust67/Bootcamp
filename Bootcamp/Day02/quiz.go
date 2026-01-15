package main

import (
	"fmt"
)

func main() {
	var pilihan int

	fmt.Println("1. Soal 6")
	fmt.Println("2. Soal 7")
	fmt.Println("3. Soal 8")
	fmt.Println("4. Soal 9")
	fmt.Print("Pilih soal (1-4): ")
	fmt.Scan(&pilihan)
	fmt.Println()

	if pilihan == 1 {
		var matrixsoal [5][5]int
		matrixsoal = soal6(matrixsoal)
		displayMatrixquiz(matrixsoal)
	} else if pilihan == 2 {
		var matrixsoal2 [5][5]int
		matrixsoal2 = soal7(matrixsoal2)
		displayMatrixquiz(matrixsoal2)
	} else if pilihan == 3 {
		var matrixsoal3 [7][7]int
		matrixsoal3 = soal8(matrixsoal3)
		displayMatrixquiz2(matrixsoal3)
	} else if pilihan == 4 {
		var matrixsoal4 [7][7]int
		matrixsoal4 = soal9(matrixsoal4)
		displayMatrixquiz2(matrixsoal4)
		fmt.Println()
		soal9sum(matrixsoal4)
	}
}

func soal6(matrix [5][5]int) [5][5]int {
	counter := 1
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == col {
				matrix[row][col] = counter
				counter++
			} else if row < col {
				matrix[row][col] = 10
			} else {
				matrix[row][col] = 20
			}
		}
	}
	return matrix
}

func soal7(matrix [5][5]int) [5][5]int {
	counter := len(matrix)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == col {
				matrix[row][col] = counter
				counter--
			} else if row < col {
				matrix[row][col] = 10
			} else {
				matrix[row][col] = 20
			}
		}
	}
	return matrix
}

func soal8(matrix [7][7]int) [7][7]int {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == 0 {
				matrix[row][col] = col
			} else if col == 0 {
				matrix[row][col] = row
			} else if col == len(matrix)-1 {
				matrix[row][col] = row + len(matrix) - 1
			} else if row == len(matrix)-1 {
				matrix[row][col] = col + len(matrix) - 1
			} else {
				matrix[row][col] = 0
			}
		}
	}
	return matrix
}

func soal9(matrix [7][7]int) [7][7]int {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == 0 {
				matrix[row][col] = col
			} else if row == len(matrix)-1 {
				matrix[row][col] = (len(matrix) - 1) + col
			} else {
				matrix[row][col] = row + col
			}
		}
	}
	return matrix
}

func soal9sum(matrix [7][7]int) {
	sumBaris := make([]int, len(matrix))
	sumKolom := make([]int, len(matrix[0]))
	totalSum := 0

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			sumBaris[row] += matrix[row][col]
			sumKolom[col] += matrix[row][col]
			totalSum += matrix[row][col]
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%5d", matrix[row][col])
		}
		fmt.Printf("%5d", sumBaris[row])
		fmt.Println()
	}

	for col := 0; col < len(matrix[0]); col++ {
		fmt.Printf("%5d", sumKolom[col])
	}
	fmt.Printf("%5d\n", totalSum)
}

func displayMatrixquiz(matrix [5][5]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%5d", matrix[row][col])
		}
		fmt.Println()
	}
}

func displayMatrixquiz2(matrix [7][7]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%5d", matrix[row][col])
		}
		fmt.Println()
	}
}
