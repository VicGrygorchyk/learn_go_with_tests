package dep_inj_test

import (
	"github.com/vicgrygorchyk/example_hello/dep_inj"
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	dep_inj.Greet(&buffer, "Vik")

	got := buffer.String()
	want := "Hello, Vik"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}