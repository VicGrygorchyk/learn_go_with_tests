package counter_test

import (
	"testing"
	"bytes"
	"github.com/vicgrygorchyk/example_hello/counter"
)

func TestCounter(t *testing.T) {
	buffer := bytes.Buffer{}
	spy := counter.SpySleeper{}

	counter.Countdown(&buffer, &spy)

    got := buffer.String()
    want := "3\n2\n1\nGo!"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}