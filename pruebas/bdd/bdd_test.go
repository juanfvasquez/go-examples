package main
 
import (
	. "github.com/onsi/ginkgo" // BDD
	. "github.com/onsi/gomega" // Asserts
)

var _ = Describe("Consignar un monto a una cuenta", func () {
	var iCuenta Cuenta

	BeforeSuite(func() {
		iCuenta = Cuenta{Monto: 10000}
	})

	Context("Si el valor a consignar es 2000", func () {
		BeforeEach(func () {
			iCuenta.Monto = 10000
		})

		It("El monto de la cuenta debe se de 12000", func () {
			iCuenta.Consignar(2000)
			Ω(iCuenta.Monto).Should(Equal(12000))
			// Expect(iCuenta.Monto).To(Equal(12000))
		})
	})

	Context("Si el valor a consignar es -2000", func () {
		BeforeEach(func () {
			iCuenta.Monto = 10000
		})
		It("El monto de la cuenta debe se de 10000", func () {
			_, err := iCuenta.Consignar(-2000)
			Ω(iCuenta.Monto).Should(Equal(10000))
			Expect(err.Error()).To(Equal("El valor no debe ser negativo"))
		})
	})
})