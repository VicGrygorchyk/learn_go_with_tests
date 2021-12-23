package reflect_test

import (
	"fmt"
	reflect_std "reflect"
	"github.com/vicgrygorchyk/example_hello/reflect"
	"testing"
)

func TestWalt(t *testing.T) {
	expected := "Vik"
	var got []string

	x := struct {
		Name string
	}{expected}

	reflect.Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1{
		t.Errorf("Expected %d, got %d", 1, len(got))
	}

	res := got[0]
	if res != expected {
		t.Errorf("Expected %q, got %q", expected, res)
	}
}

func TestWalkParametrize(t *testing.T) {

	testCases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			"One string field",
			struct {
				Name string
			}{"Vik"},
			[]string{"Vik"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			struct {
				Name string
				Profile struct {
					Age  int
					City string
				}
			}{"Chris", struct {
				Age  int
				City string
			}{33, "London"}},
			[]string{"Chris", "London"},
		},
	}

	fmt.Println(len(testCases))
	fmt.Println(cap(testCases))

	for _, test := range testCases {
		t.Run(test.Name, func (t *testing.T){
			// got := make([]string, 1) requires length of slice
			var got []string // empty so far

			reflect.Walk(test.Input, func(inp string) {
				got = append(got, inp)
			})

			if !reflect_std.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Got %v, expected %v", got, test.ExpectedCalls)
			}
		})
	}

}