package main 

import (
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {

	const MSG_NO_VALIDO = "No es un número válido"
	const FIZZ = "fizz"
	const BUZZ = "buzz"
	const FIZZBUZZ = "fizzbuzz"

	t.Run("Probar cuando enviamos 0", func (t *testing.T) {
		salida := fizzbuzz(0)
		assert.Equal(t, MSG_NO_VALIDO, salida)
		// if salida != MSG_NO_VALIDO {
		// 	t.Errorf("Resultado: `%s`; se esperara `%s`", salida, MSG_NO_VALIDO)
		// }
	})

	t.Run("Probar cuando enviamos un número negativo", func (t *testing.T) {
		salida := fizzbuzz(-2)
		assert.Equal(t, MSG_NO_VALIDO, salida)
		// if salida != MSG_NO_VALIDO {
		// 	t.Errorf("Resultado: `%s`; se esperara `%s`", salida, MSG_NO_VALIDO)
		// }
	})

	t.Run("Probar cuando enviamos un múltiplo de 3", func (t *testing.T) {
		salida := fizzbuzz(3)
		assert.Equal(t, FIZZ, salida)
		// if salida != FIZZ {
		// 	t.Errorf("Resultado: `%s`; se esperara `%s`", salida, FIZZ)
		// }
	})

	t.Run("Probar cuando enviamos un múltiplo de 5", func (t *testing.T) {
		salida := fizzbuzz(5)
		assert.Equal(t, BUZZ, salida)
		// if salida != BUZZ {
		// 	t.Errorf("Resultado: `%s`; se esperara `%s`", salida, BUZZ)
		// }
	})

	t.Run("Probar cuando enviamos un múltiplo de 3 y de 5", func (t *testing.T) {
		salida := fizzbuzz(15)
		assert.Equal(t, FIZZBUZZ, salida)
		// if salida != FIZZBUZZ {
		// 	t.Errorf("Resultado: `%s`; se esperara `%s`", salida, FIZZBUZZ)
		// }
	})

	t.Run("Probar cuando enviamos un número que no es múltiplo de 3 o de 5", func (t *testing.T) {
		salida := fizzbuzz(22)
		esperado := strconv.Itoa(22)
		assert.Equal(t, esperado, salida)
		// if salida != esperado {
		// 	t.Errorf("Resultado: `%s`; se esperara `%s`", salida, esperado)
		// }
	})
}