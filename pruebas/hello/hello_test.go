package main

import "testing"

func TestSaludar(t *testing.T) {
	t.Run("Probar cuando el nombre está vacío", func(t *testing.T) {
		saludo := saludar("")
		esperado := "Hola mundo!"
		if saludo != esperado {
			t.Errorf("Obtuvimos %s pero esperábamos %s", saludo, esperado)
		}
	})
	
	t.Run("Probar cuando el nombre no esté vacío", func(t *testing.T) {
		saludo := saludar("Juan")
		esperado := "Hola Juan"
		if saludo != esperado {
			t.Errorf("Obtuvimos %s pero esperábamos %s", saludo, esperado)
		}
	})
}

// func TestSaludarNombre(t *testing.T) {
// 	saludo := saludar("Juan")
// 	esperado := "Hola Pedro"
// 	if saludo != esperado {
// 		t.Errorf("Obtuvimos %s pero esperábamos %s", saludo, esperado)
// 	}
// }

// func TestSaludarNombreVacio(t *testing.T) {
// 	saludo := saludar("")
// 	esperado := "Hola mundo"
// 	if saludo != esperado {
// 		t.Errorf("Obtuvimos %s pero esperábamos %s", saludo, esperado)
// 	}
// }