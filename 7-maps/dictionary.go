package maps

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(searchkey string) (string, error) {
	definition, ok := d[searchkey]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key string, value string) {
	d[key] = value
}
