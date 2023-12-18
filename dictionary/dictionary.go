package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	var err error

	definition, ok := d[word]
	if !ok {
		err = ErrNotFound
	}

	return definition, err
}

func (d Dictionary) Add(word, definition string) error {
	var err error

	_, status := d.Search(word)

	switch status {
	case ErrNotFound:
		d[word] = definition
	case nil:
		err = ErrWordExists
	default:
		err = status
	}

	return err
}

