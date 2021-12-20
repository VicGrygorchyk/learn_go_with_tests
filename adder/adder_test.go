package adder_test

import (
	"reflect"
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

func TestAddAll(t *testing.T) {
	t.Run("Check AddAll with two param", func (t *testing.T) {
		res := adder.AddAll([]int{1, 2}, []int{5, 6})
		expected := []int {3, 11}
		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Expected %d, got %d", res, expected)
		}
	})

	t.Run("Check AddAll with three param", func (t *testing.T) {
		res := adder.AddAll([]int{1, 2}, []int{5, 6}, []int{2, 2, 2})
		expected := []int {3, 11, 6}
		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Expected %d, got %d", res, expected)
		}
	})

	t.Run("Check AddAll with big slices", func (t *testing.T) {
		res := adder.AddAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, []int{5, 6}, []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2})
		expected := []int {66, 11, 20}
		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Expected %d, got %d", res, expected)
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
