package main

import (
	"fmt"
)

func main() {

	var pilihan int
	fmt.Print("Masukkan Soal: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		soal01()
	case 2:
		soal02()
	case 3:
		soal03a()
	case 4:
		soal03b()
	case 5:
		soal04()
	case 6:
		soal05()
	case 7:
		soal06()
	case 8:
		soal08()
	case 9:
		soal09()
	case 10:
		soal10()
	case 11:
		soal11()
	default:
		fmt.Println("pilihan tidak tersedia")
	}

}

func soal01() {
	var n int
	fmt.Print("masukkan value n: ")
	fmt.Scan(&n)

	for i := 1; i < n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}

func soal02() {
	var n int
	fmt.Scan(&n)

	for n > 0 {
		digit := n % 10
		fmt.Print(digit)
		if n >= 10 {
			fmt.Print(" ")
		}
		n /= 10
	}
}

func soal03a() {
	var n int
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		for s := 0; s < n-i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func soal03b() {
	var n int
	fmt.Scan(&n)

	for i := n; i > 1; i-- {
		for s := 0; s < n-i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func soal04() {
	var n int
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Print(j, " ")
		}
		for j := 2; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

}

func soal05() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {

			if j%2 == 1 {
				fmt.Print(i)
			} else {
				fmt.Print(n - i + 1)
			}
			if j < n {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func soal06() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {

			if (i+j)%2 == 0 {
				fmt.Print("-")
			} else {
				fmt.Print(j)
			}

			if j < n {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func soal08() {
	text := "Kasur ini rusak"

	kiri := 0
	kanan := len(text) - 1
	palindrom := true

	for kiri < kanan {

		if text[kiri] == ' ' {
			kiri++
			continue
		}
		if text[kanan] == ' ' {
			kanan--
			continue
		}

		left := text[kiri]
		right := text[kanan]

		if left >= 'A' && left <= 'Z' {
			left += 32
		}
		if right >= 'A' && right <= 'Z' {
			right += 32
		}

		if left != right {
			palindrom = false
			break
		}

		kiri++
		kanan--
	}

	fmt.Println(palindrom)
}

func soal09() {
	text := "ABCD"
	result := ""

	for i := len(text) - 1; i >= 0; i-- {
		result += string(text[i])
	}

	fmt.Println(result)
}

func soal10() {
	text := "(())"
	count := 0
	valid := true

	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '(':
			count++
		case ')':
			count--
		}

		if count < 0 {
			valid = false
			break
		}
	}

	if count != 0 {
		valid = false
	}

	fmt.Println(valid)
}

func soal11() {
	var n int
	fmt.Scan(&n)

	original := n
	reversed := 0

	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}

	fmt.Println(original == reversed)
}
