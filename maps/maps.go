package maps

import (
	"errors"
)

type Dictionary map[string]string

var ErrWordNoExists = errors.New("could not find the word you were looking for")
var ErrWordAlreadyExists = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(text string) (string, error) {
	item, ok := d[text]

	if !ok {
		return "", ErrWordNoExists
	}

	return item, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, exist := d[key]

	if exist {
		return ErrWordAlreadyExists
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, exist := d[key]

	if !exist {
		return ErrWordNoExists
	}

	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, exist := d[key]

	if !exist {
		return ErrWordNoExists
	}

	delete(d, key)
	return nil
}
