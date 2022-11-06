package main

import "testing"

func TestSample(t *testing.T) {
	greet := MakeHello("Azusa")
	expect := "Hello World! Azusa"
	if greet != expect {
		t.Fatal("expect [" + expect + "] but, [" + greet + "].")
	}
}
