package synclearn_test

import (
	"github.com/vicgrygorchyk/example_hello/synclearn"
	"testing"
	"sync"
)

func TestCounter(t *testing.T) {
    t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
        counter := synclearn.Counter{}
        counter.Inc()
        counter.Inc()
        counter.Inc()

		assertCounter(&counter, 3, t)
    })

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := synclearn.Counter{}
	
		var wg sync.WaitGroup
		wg.Add(wantedCount)
	
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
	
		assertCounter(&counter, wantedCount, t)
	})
}

func assertCounter(c *synclearn.Counter, expected int, t testing.TB) {
	t.Helper()
	if c.Value() != expected {
		t.Errorf("got %d, want %d", c.Value(), expected)
	}
}