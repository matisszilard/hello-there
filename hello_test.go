package main

import "testing"

func TestSatHello(t *testing.T) {
	if sayHello() != "General Kenobi" {
		t.Error("invalid greetings value")
	}
}
