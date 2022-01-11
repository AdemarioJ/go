package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Ademario")
	esperado := "Olá, Ademario"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
