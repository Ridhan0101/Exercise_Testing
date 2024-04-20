package main_test

import (
	. "EXERCISE_Testing/Ginkgo_Testing" // Sesuaikan dengan path package calculatorscience.go

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	Context("Testing basic arithmetic operations", func() {
		It("should correctly add two numbers", func() {
			Expect(Penjumlahan(2, 3)).To(Equal(5.0))
		})

		It("should correctly subtract two numbers", func() {
			Expect(Pengurangan(5, 3)).To(Equal(2.0))
		})

		It("should correctly multiply two numbers", func() {
			Expect(Perkalian(2, 3)).To(Equal(6.0))
		})

		It("should correctly divide two numbers", func() {
			result, err := Pembagian(10, 2)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(5.0))
		})

		It("should return an error when dividing by zero", func() {
			_, err := Pembagian(10, 0)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Testing scientific functions", func() {
		It("should correctly calculate square root", func() {
			Expect(Akar_Kuadrat(16)).To(Equal(4.0))
		})

		It("should correctly calculate square", func() {
			Expect(Pangkat_Dua(4)).To(Equal(16.0))
		})

		It("should correctly calculate cube", func() {
			Expect(Pangkat_Tiga(2)).To(Equal(8.0))
		})

		It("should correctly calculate sine", func() {
			Expect(Sin(0)).To(Equal(0.0))
		})

		It("should correctly calculate cosine", func() {
			Expect(Cos(0)).To(Equal(1.0))
		})

		It("should correctly calculate tangent", func() {
			Expect(Tan(0)).To(Equal(0.0))
		})

		It("should correctly calculate sine in degrees", func() {
			Expect(SinD(90)).To(Equal(1.0))
		})

		It("should correctly calculate cosine in degrees", func() {
			Expect(CosD(0)).To(Equal(1.0))
		})

		It("should correctly calculate tangent in degrees", func() {
			Expect(TanD(45)).To(Equal(1.0))
		})
	})
})
