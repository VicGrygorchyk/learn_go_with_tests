package adder_test

import (
	"fmt"
	"github.com/vicgrygorchyk/example_hello/adder"
	"testing"

)

func TestAdd(t *testing.T) {
	t.Run("Check Add function", func (t *testing.T) {
		res, err := adder.Add(1, 2)
		if err != nil {
			t.Fatalf("Go error")
		}
		expected := 3

		if res != expected {
			t.Errorf("expected %d, got %d", res, expected)
		}
	})
}

// This is example of using Adder
func ExampleAdd() {
	sum, err := adder.Add(3, 5)
	if err != nil {
		fmt.Print(err)
	} 
	fmt.Println(sum)
	// Output: 8
}
