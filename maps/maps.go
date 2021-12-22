package maps

import (
	"fmt"
	"errors"
)


type Dictionary map[string]string

var (
    ErrNotFound   = errors.New("could not find the word you were looking for")
    ErrWordExists = errors.New("cannot add word because it already exists")
)


func (dict Dictionary) Search(find string) (string, error) {
	res, ok := dict[find]
	if !ok {
		return "", errors.New(fmt.Sprintf("Not found %v in the %v", dict, find))
	}
	return res, nil
}

func (dict Dictionary) Add(newKey string, newWord string) (string, error) {
	_, ok := dict[newKey]

	if ok {
		return "", ErrWordExists
	}

	dict[newKey] = newWord
	return "", nil
}