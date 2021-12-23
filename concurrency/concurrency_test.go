package concurrency_test

import (
	"github.com/vicgrygorchyk/example_hello/concurrency"
    "reflect"
    "testing"
)

var websites = []string{
	"http://google.com",
	"http://blog.gypsydave5.com",
	"waat://furhurterwe.geds",
}

func mockWebsiteChecker(url string) bool {
    if url == "waat://furhurterwe.geds" {
        return false
    }
    return true
}

func TestCheckWebsites(t *testing.T) {


    want := map[string]bool{
        "http://google.com":          true,
        "http://blog.gypsydave5.com": true,
        "waat://furhurterwe.geds":    false,
    }

    got := concurrency.CheckWebsites(mockWebsiteChecker, websites)

    if !reflect.DeepEqual(want, got) {
        t.Fatalf("Wanted %v, got %v", want, got)
    }
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrency.CheckWebsites(mockWebsiteChecker, websites)
	}
}