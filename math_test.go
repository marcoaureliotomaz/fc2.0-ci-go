package main

import "testing"

func TestSoma( t *testing.T) {

	total := soma(15,15)

	if total != 30{
		t.Errorf("Resultado da função somaestá  inválido: Resultado %d. Esperado %d", total, 30)
	}
}