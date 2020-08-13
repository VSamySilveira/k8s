package main

import "testing"

func TestGreeting(t *testing.T) {
	resultado := greeting("Olá!")
	if resultado != "<b>Olá!</b>" {
	  t.Errorf("Greeting esperada: <b>Olá!</b>, obtida: %s", resultado)
	}
}