package maps

import "errors"

type Dictionary map[string]string

var (
	ErrorNotFoundWord = errors.New("could not find the word you were looking for")
	ErrorWordExists   = errors.New("cannot add word because it already exists")
)

func (dictionary Dictionary) Search(word string) (string, error) {
	searchValue, ok := dictionary[word]
	if !ok {
		return "", ErrorNotFoundWord
	}
	return searchValue, nil
}

func (dictionary Dictionary) Add(word, value string) error {
	_, err := dictionary.Search(word)
	if err == nil {
		return ErrorWordExists
	}
	dictionary[word] = value
	return nil
}

func (dictionary Dictionary) Update(word, value string) error {
	_, err := dictionary.Search(word)
	if err == nil {
		dictionary[word] = value
		return nil
	}
	return err
}

func (dictionary Dictionary) Delete(word string) error {
	_, err := dictionary.Search(word)
	if err == nil {
		delete(dictionary, word)
		return nil
	}
	return err
}

func Search(dictionary Dictionary, word string) string {
	return dictionary[word]
}
