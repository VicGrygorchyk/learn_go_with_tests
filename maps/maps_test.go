package maps_test

import (
	"github.com/vicgrygorchyk/example_hello/maps"
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("Test with word in the dict", func(t *testing.T) {
		dictionary := maps.Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"
	
		assertStrings(t, got, want)
	})

	t.Run("Test with word absent from the dict", func(t *testing.T) {
		dictionary := maps.Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("test2")
	
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}