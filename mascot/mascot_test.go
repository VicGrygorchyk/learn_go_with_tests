package mascot_test

import (
	"github.com/vicgrygorchyk/example_hello/mascot"
	"testing"
)

func TestMascot(t *testing.T) {
	res, err := mascot.BestMascot("test")
	if (res != "Hello, test") && (err == nil) {
		t.Errorf("Expected to be hello, got %s", res)
	}

	t.Run("Tetsing with empty argument", func(t *testing.T) {
		res, err := mascot.BestMascot("")
		if (res != "Hello, test") && (err == nil) {
			t.Errorf("Expected to be hello, got %s", res)
		} else {
			t.Errorf("Got error: %v", err)
		}
	})
}
