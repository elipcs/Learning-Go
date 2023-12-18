package dictionary

type Dictionary map[string]string

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(word, definition string) error {
	var err error

	_, status := d.Search(word)

	switch status {
	case ErrNotFound:
		err = ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		err = status
	}

	return err
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
