package main

import (
	"math"
	"testing"
)

func TestPenjumlahan(t *testing.T) {
	result := penjumlahan(5, 7)
	if result != 12 {
		t.Errorf("Expected result to be 12, but got %f", result)
	}
}

func TestPengurangan(t *testing.T) {
	result := Pengurangan(10, 3)
	if result != 7 {
		t.Errorf("Expected result to be 7, but got %f", result)
	}
}

func TestPerkalian(t *testing.T) {
	result := Perkalian(4, 3)
	if result != 12 {
		t.Errorf("Expected result to be 12, but got %f", result)
	}
}

func TestPembagian(t *testing.T) {
	result, err := Pembagian(10, 2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected result to be 5, but got %f", result)
	}

	_, err = Pembagian(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero, but got none")
	}
}

func TestAkarKuadrat(t *testing.T) {
	result := akar_kuadrat(25)
	if result != 5 {
		t.Errorf("Expected result to be 5, but got %f", result)
	}
}

func TestPangkatDua(t *testing.T) {
	result := pangkat_dua(3)
	if result != 9 {
		t.Errorf("Expected result to be 9, but got %f", result)
	}
}

func TestPangkatTiga(t *testing.T) {
	result := pangkat_tiga(2)
	if result != 8 {
		t.Errorf("Expected result to be 8, but got %f", result)
	}
}

func TestSin(t *testing.T) {
	result := Sin(math.Pi / 2)
	if result != 1 {
		t.Errorf("Expected result to be 1, but got %f", result)
	}
}

func TestCos(t *testing.T) {
	result := Cos(math.Pi)
	if result != -1 {
		t.Errorf("Expected result to be -1, but got %f", result)
	}
}

func TestTan(t *testing.T) {
	result := Tan(math.Pi / 4)
	if result != 1 {
		t.Errorf("Expected result to be 1, but got %f", result)
	}
}

func TestSinD(t *testing.T) {
	result := SinD(90)
	if result != 1 {
		t.Errorf("Expected result to be 1, but got %f", result)
	}
}

func TestCosD(t *testing.T) {
	result := CosD(180)
	if result != -1 {
		t.Errorf("Expected result to be -1, but got %f", result)
	}
}

func TestTanD(t *testing.T) {
	result := TanD(45)
	if result != 1 {
		t.Errorf("Expected result to be 1, but got %f", result)
	}
}
