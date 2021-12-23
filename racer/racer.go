package racer

import (
	"net/http"
)

func Racer(FirstUrl string, SecondUrl string) string {
	select {
	case <- ping(FirstUrl):
		return FirstUrl
	case <- ping(SecondUrl):
		return SecondUrl
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
