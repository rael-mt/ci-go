package main

import "testing"

func TestSoma(t *testing.T) {
	total := Mult(15, 15)

	if total != 225 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperando: %d", total, 30)
	}
}
