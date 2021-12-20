package iter_test

import (
	"testing"
	"github.com/vicgrygorchyk/example_hello/iter"
)

func TestRepeat(t *testing.T) {
	repeated := iter.Repeat("a", 5)
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("Expected %q, got %q", expected, repeated)
	}
}

var result string

func BenchmarkRepeat(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res += iter.Repeat("a", 6)
	}
	result = res
}