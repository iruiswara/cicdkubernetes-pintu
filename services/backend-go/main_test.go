package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	result := helloWorld()
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
