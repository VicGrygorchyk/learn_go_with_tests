package counter

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls ++
}

func Countdown(writer io.Writer, spy Sleeper) {
	count := 3
	for i := count; i > 0; i-- {
		spy.Sleep()
        fmt.Fprintln(writer, i)
    }
	for i := count; i > 0; i-- {
		spy.Sleep()
        fmt.Fprintln(writer, i)
    }
	spy.Sleep()
    fmt.Fprint(writer, "Go!")
}
