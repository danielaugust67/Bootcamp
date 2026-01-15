package main

import (
	"fmt"
)

func main() {
	var tipe int
	fmt.Println("Pilih tipe matrix:")
	fmt.Println("1. Integer")
	fmt.Println("2. String")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scanln(&tipe)

	if tipe == 1 {
		var pilihan int
		fmt.Println("Pilih pola matrix integer:")
		fmt.Println("1. Box Start")
		fmt.Println("2. Hollow Box")
		fmt.Println("3. Triangle 1")
		fmt.Println("4. Triangle 2")
		fmt.Println("5. Triangle 3")
		fmt.Println("6. Triangle 4")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&pilihan)
		var matrix [5][5]int

		if pilihan == 1 {
			matrix = boxstart(matrix)
		} else if pilihan == 2 {
			matrix = hollowbox(matrix)
		} else if pilihan == 3 {
			matrix = triangle1(matrix)
		} else if pilihan == 4 {
			matrix = triangle2(matrix)
		} else if pilihan == 5 {
			matrix = triangle3(matrix)
		} else if pilihan == 6 {
			matrix = triangle4(matrix)
		} else {
			fmt.Println("Pilihan tidak valid")
			return
		}
		displayMatrix(matrix)

	} else if tipe == 2 {
		var pilihan int
		fmt.Println("Pilih pola matrix string:")
		fmt.Println("1. Box Start")
		fmt.Println("2. Hollow Box")
		fmt.Println("3. Triangle 1")
		fmt.Println("4. Triangle 2")
		fmt.Println("5. Triangle 3")
		fmt.Println("6. Triangle 4")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&pilihan)
		var matrix [5][5]string

		if pilihan == 1 {
			matrix = boxstartStr(matrix)
		} else if pilihan == 2 {
			matrix = hollowboxStr(matrix)
		} else if pilihan == 3 {
			matrix = triangle1str(matrix)
		} else if pilihan == 4 {
			matrix = triangle2str(matrix)
		} else if pilihan == 5 {
			matrix = triangle3str(matrix)
		} else if pilihan == 6 {
			matrix = triangle4str(matrix)
		} else {
			fmt.Println("Pilihan tidak valid")
			return
		}
		displayMatrixStr(matrix)

	} else {
		fmt.Println("Not Found")
	}
}

// Box Start
func boxstart(matrix [5][5]int) [5][5]int {
	counter := 1
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			matrix[row][col] = counter
			counter++
		}
	}
	return matrix
}

func boxstartStr(matrix [5][5]string) [5][5]string {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			matrix[row][col] = "*"
		}
	}
	return matrix
}

//Hollow Box

func hollowbox(matrix [5][5]int) [5][5]int {
	counter := 1
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == 0 || row == len(matrix)-1 || col == 0 || col == len(matrix[row])-1 {
				matrix[row][col] = counter
				counter++
			}
		}
	}
	return matrix
}

func hollowboxStr(matrix [5][5]string) [5][5]string {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == 0 || row == len(matrix)-1 || col == 0 || col == len(matrix[row])-1 {
				matrix[row][col] = "*"
			}
		}

	}
	return matrix
}

// triangle1
func triangle1(matrix [5][5]int) [5][5]int {
	counter := 1
	for row := 0; row < len(matrix); row++ {
		for col := 0; col <= row; col++ {
			matrix[row][col] = counter
			counter++
		}
	}
	return matrix
}

func triangle1str(matrix [5][5]string) [5][5]string {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col <= row; col++ {
			matrix[row][col] = "*"
		}
	}
	return matrix
}

// trinagle2
func triangle2(matrix [5][5]int) [5][5]int {
	counter := 1
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix)-row; col++ {
			matrix[row][col] = counter
			counter++
		}
	}
	return matrix
}

func triangle2str(matrix [5][5]string) [5][5]string {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix)-row; col++ {
			matrix[row][col] = "*"
		}
	}
	return matrix
}

// triangle3
func triangle3(matrix [5][5]int) [5][5]int {
	counter := 1
	for row := 0; row < len(matrix[row]); row++ {
		for col := len(matrix[row]) - 1; col >= len(matrix[row])-row-1; col-- {
			matrix[row][col] = counter
			counter++
		}
	}
	return matrix
}

func triangle3str(matrix [5][5]string) [5][5]string {
	for row := 0; row < len(matrix[row]); row++ {
		for col := len(matrix[row]) - 1; col >= len(matrix[row])-row-1; col-- {
			matrix[row][col] = "*"
		}
	}
	return matrix
}

// triangle4
func triangle4(matrix [5][5]int) [5][5]int {
	counter := 1
	for row := 0; row < len(matrix); row++ {
		for col := len(matrix[row]) - 1; col >= row; col-- {
			matrix[row][col] = counter
			counter++
		}
	}
	return matrix
}

func triangle4str(matrix [5][5]string) [5][5]string {
	for row := 0; row < len(matrix); row++ {
		for col := len(matrix[row]) - 1; col >= row; col-- {
			matrix[row][col] = "*"
		}
	}
	return matrix
}

func displayMatrix(matrix [5][5]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%5d", matrix[row][col])
		}
		fmt.Println()
	}
}

func displayMatrixStr(matrix [5][5]string) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%5s", matrix[row][col])
		}
		fmt.Println()
	}
}
