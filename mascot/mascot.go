package mascot

import (
	"time"
	"fmt"
	"errors"
	"math/rand"
)

// Best mascot function
func BestMascot(name string) (string, error) {

	if name == "" {
		return "no name", errors.New("empty name")
	}
	return fmt.Sprintf(getRandGreetings(), name), nil
}

func GetGreetings(names []string) (map[string]string, error) {
	msgs := make(map[string]string)

	for _, name := range names {
		msg, err := BestMascot(name)
		if err != nil {
			return nil, err
		}
		msgs[name] = msg
	}

	return msgs, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandGreetings() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Aloha, %v!",
	}

	return formats[rand.Intn(len(formats))]
}
