package main

import (
	"fmt"
	"reflect"
)

const PREFIX string = "testing:"
var arr [10]int = [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var slice = arr[:3]

// func main() {
// 	log.SetPrefix(PREFIX)
// 	log.SetFlags(0)

// 	for i := range arr {
// 		fmt.Println(i)
// 	}

// 	slice = append(slice, 4)
// 	fmt.Println("Len ", len(slice))
// 	fmt.Println("Cap ", cap(slice))

// 	slice2 := make([]int, 2, 2)
// 	slice2[0] = 10
// 	slice2[1] = 20
	

// 	for i := range slice {
// 		fmt.Println("slice is", i)
// 	}

// 	names := []string {"Viktor", "Ira", "Feya"}
// 	msgs, err := mascot.GetGreetings(names)
// 	fmt.Println(msgs)

// 	test := map[string]string {
// 		"test": "test",
// 		"test2": "test2",
// 	}

// 	fmt.Println("Test is ", test["test3"])

// 	val, ok := test["test3"]
// 	if ok == true {
// 		fmt.Println(val)
// 	} else {
// 		fmt.Println("no val for a key")
// 	}

// 	msg1, err1 := mascot.BestMascot("Viktor")
// 	if err1 != nil {
// 		log.Fatal(err1)
// 	}
// 	fmt.Println(msg1)

// 	msg, err := mascot.BestMascot("")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(msg)
// }
// type DefaultSleeper struct {}

// func (d *DefaultSleeper) Sleep() {
//     time.Sleep(1 * time.Second)
// }

// func main() {
// 	sleeper := DefaultSleeper{}
// 	counter.Countdown(os.Stdout, &sleeper)
// }


var test int = 5

func main() {
	val := reflect.ValueOf(test)

	fmt.Println(val)

	type StructX struct {
		fieldZero int
		fieldOne string
	}
	
	var example StructX = StructX{3, "test"}
	
	valOfStruct := reflect.ValueOf(example)
	fieldZero := valOfStruct.Field(0)
	fieldOne := valOfStruct.Field(1)
	fmt.Println(fieldZero)
	fmt.Println(fieldOne)
}