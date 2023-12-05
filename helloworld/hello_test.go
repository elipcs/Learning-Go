package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Eli", "")
		want := "Hello, Eli"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("Eli", "Spanish")
		want := "Hola, Eli"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in portuguese", func(t *testing.T) {
		got := Hello("Eli", "Portuguese")
		want := "Hello, Eli"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Oi, World' when an empty string is supplied and language is portuguese", func(t *testing.T){
		got := Hello("", "Portuguese")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})	
	t.Run("say 'Hola, World' when an empty string is supplied and language is spanish", func(t *testing.T){
		got := Hello("", "Spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
    t.Helper()	
	if got != want {
			t.Errorf("got %q want %q", got, want)
		}
}
