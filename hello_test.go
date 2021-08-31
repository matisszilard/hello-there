package main

import "testing"

func TestSatHello(t *testing.T) {
	if sayHello() != "Hello there!" {
		t.Error("invalid greetings value")
	}
}
