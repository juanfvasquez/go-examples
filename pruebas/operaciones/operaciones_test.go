package main

import "testing"

func TestSuma(t *testing.T) {

	assertEqual := func (resultado, esperado int, t *testing.T) {
		if resultado != esperado {
			t.Errorf("Resultado: %d, se esperaba: %d", resultado, esperado)
		}
	}

	assertNotEqual := func (resultado, esperado int, t *testing.T) {
		if resultado == esperado {
			t.Errorf("El resultado no debe ser %d", esperado)
		}
	}

	// assertNotError := func (err error, t *testing.T) {
	// 	if err != nil {
	// 		t.Errorf("NO se esperaba error")
	// 	}
	// }

	assertIsError := func (err error, t *testing.T) {
		if err == nil {
			t.Errorf("Se esperaba error")
		}
	}

	t.Run("Probar suma correcta", func(t *testing.T) {
		resultado, _ := sumar(3, 4)
		esperado := 7
		assertEqual(resultado, esperado, t)
	})

	t.Run("Probar que los resultados son diferentes", func(t *testing.T) {
		resultado, _ := sumar(3, 4)
		esperado := 8
		assertNotEqual(resultado, esperado, t)
	})

	t.Run("Probar que hay un error", func(t *testing.T) {
		_, err := sumar(-1, 1)
		assertIsError(err, t)
	})
}

func TestResta(t *testing.T) {
	t.Run("Probar que el resultado es correcto", func(t *testing.T) {
		resultado, _ := restar(10, 5)
		esperado := 5

		if resultado != esperado {
			t.Errorf("El resultado fue: %d, se esperaba: %d", resultado, esperado)
		}
	})

	t.Run("Probar que los resultados no coinciden", func(t *testing.T) {
		resultado, _ := restar(8, 3)
		esperado := 4

		if resultado == esperado {
			t.Errorf("Se esperaba un resultado diferente")
		}
	})

	t.Run("Probar que Y no es menor a X", func(t *testing.T) {
		_, err := restar(4, 8)
		if err == nil {
			t.Error("Se esperaba error por Y mayor a X")
		}
	})

	t.Run("Probar error cuando X o Y menores a cero", func(t *testing.T) {
		_, err := restar(1, -1)

		if err == nil {
			t.Error("Se esperaba error por números menores a cero")
		}
	})
}

func TestMultiplicar(t *testing.T) {
	t.Run("Probar resultado correcto", func(t *testing.T) {
		resultado, _ := multiplicar(2, 2)
		esperado := 4

		if resultado != esperado {
			t.Errorf("Resuldato: %d, se esperaba: %d", resultado, esperado)
		}
	})

	t.Run("Probar error cuando números son menores a cero", func(t *testing.T) {
		_, err := multiplicar(-1, -2)
		if err == nil {
			t.Errorf("Se esperaba error por números menores a cero")
		}
	})
}

func TestDividir(t *testing.T) {
	t.Run("La división por cero debe enviar error", func (t *testing.T) {
		_, err := dividir(10, 0)
		if err == nil {
			t.Errorf("Se esperaba error por división por cero")
		}
	})

	t.Run("Probar resultado correcto", func (t *testing.T) {
		resultado, _ := dividir(4, 2)
		esperado := 2
		if resultado != esperado {
			t.Errorf("Resultado %d, se esperaba %d", resultado, esperado)
		}
	})

	t.Run("Probar error con números negativos", func (t *testing.T) {
		_, err := dividir(-1, -1)
		if err == nil {
			t.Errorf("Se esperaba error por números negativos")
		}
	})

	t.Run("Probar divisiones no exactas", func (t *testing.T) {
		resultado, _ := dividir(10, 3)
		esperado := 3
		if resultado != esperado {
			t.Errorf("Resultado: %d, se esperaba: %d", resultado, esperado)
		}
	})
}