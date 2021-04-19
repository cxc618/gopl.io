package greetings

import "testing"

func TestHello(t *testing.T) {
	name := "error"
	message, err := Hello(name)
	if  err == nil {
		t.Fatalf(`Hello(%v) = %v`, name, message)
	}
}

func TestHellos(t *testing.T) {
	names := []string{"cc"}
	messages, err := Hellos(names)
	if len(messages) != len(names) {
		t.Fatalf("%v %v", names, messages)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}