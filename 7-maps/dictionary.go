package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(searchkey string) (string, error) {
	definition, ok := d[searchkey]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}
