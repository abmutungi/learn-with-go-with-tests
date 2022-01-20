//Writing a test is just like writing a function, with a few rules
//It needs to be in a file with a name like xxx_test.go

package main

//In order to use the *testing.T type, you need:
import "testing"

//The test function must start with the word Test
//The test function takes one argument only t *testing.T

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}