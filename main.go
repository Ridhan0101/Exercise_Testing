package main

import (
	"errors"
	"fmt"
	"math"
)

// Fungsi untuk melakukan penjumlahan
func penjumlahan(a, b float64) float64 {
	return a + b
}

// Fungsi untuk melakukan pengurangan
func Pengurangan(a, b float64) float64 {
	return a - b
}

// Fungsi untuk melakukan perkalian
func Perkalian(a, b float64) float64 {
	return a * b
}

// Fungsi untuk melakukan pembagian
func Pembagian(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak dapat melakukan pembagian dengan nol")
	}
	return a / b, nil
}

// Fungsi untuk menghitung akar kuadrat
func akar_kuadrat(x float64) float64 {
	return math.Sqrt(x)
}

// Fungsi untuk menghitung pangkat dua
func pangkat_dua(x float64) float64 {
	return x * x
}

// Fungsi untuk menghitung pangkat tiga
func pangkat_tiga(x float64) float64 {
	return x * x * x
}

// Fungsi untuk menghitung sinus
func Sin(x float64) float64 {
	return math.Sin(x)
}

// Fungsi untuk menghitung kosinus
func Cos(x float64) float64 {
	return math.Cos(x)
}

// Fungsi untuk menghitung tangen
func Tan(x float64) float64 {
	return math.Tan(x)
}

// Fungsi untuk menghitung sinus dari sudut (dalam derajat)
func SinD(angle float64) float64 {
	return math.Sin(angle * math.Pi / 180)
}

// Fungsi untuk menghitung kosinus dari sudut (dalam derajat)
func CosD(angle float64) float64 {
	return math.Cos(angle * math.Pi / 180)
}

// Fungsi untuk menghitung tangen dari sudut (dalam derajat)
func TanD(angle float64) float64 {
	return math.Tan(angle * math.Pi / 180)
}

func main() {
	var choice int
	var input1, input2 float64

	// Menu
	fmt.Println("Kalkulator Ilmiah")

	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")
	fmt.Println("5. Akar Kuadrat")
	fmt.Println("6. Pangkat Dua")
	fmt.Println("7. Pangkat Tiga")
	fmt.Println("8. Sinus")
	fmt.Println("9. Kosinus")
	fmt.Println("10. Tangen")
	fmt.Println("11. Sinus° (Dalam Derajat)")
	fmt.Println("12. Kosinus° (Dalam Derajat)")
	fmt.Println("13. Tangen°( Dalam Derajat)")
	fmt.Print("Pilih operasi yang diinginkan: ")
	fmt.Scan(&choice)

	// Meminta input sesuai dengan operasi yang dipilih
	switch choice {
	case 1, 2, 3, 4:
		fmt.Print("Masukkan angka pertama: ")
		fmt.Scan(&input1)
		fmt.Print("Masukkan angka kedua: ")
		fmt.Scan(&input2)
	default:
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&input1)
	}

	switch choice {
	case 1:
		fmt.Println("Hasil Penjumlahan:", penjumlahan(input1, input2))
	case 2:
		fmt.Println("Hasil pengurangan:", Pengurangan(input1, input2))
	case 3:
		fmt.Println("Hasil perkalian:", Perkalian(input1, input2))
	case 4:
		result, err := Pembagian(input1, input2)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Hasil pembagian:", result)
		}
	case 5:
		fmt.Println("Hasil akar kuadrat:", akar_kuadrat(input1))
	case 6:
		fmt.Println("Hasil pangkat dua:", pangkat_dua(input1))
	case 7:
		fmt.Println("Hasil pangkat tiga:", pangkat_tiga(input1))
	case 8:
		fmt.Println("Hasil sinus:", Sin(input1))
	case 9:
		fmt.Println("Hasil kosinus:", Cos(input1))
	case 10:
		fmt.Println("Hasil tangen:", Tan(input1))
	case 11:
		fmt.Println("Hasil sinus°(dalam derajat):", SinD(input1))
	case 12:
		fmt.Println("Hasil kosinus°(dalam derajat):", CosD(input1))
	case 13:
		fmt.Println("Hasil tangen°(dalam derajat):", TanD(input1))
	default:
		fmt.Println("Operasi tidak valid")
	}
}
