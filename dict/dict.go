package dict

import (
	"errors"
)

var (
	errNotFound = errors.New("not found")
	errSameKey  = errors.New("same key in dictionary")
)

type Dictionary map[string]interface{}

func (d Dictionary) Search(Value string) (interface{}, error) {
	value, exists := d[Value]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(Key, Value string) error {
	_, err := d.Search(Value)
	switch err {
	case errNotFound:
		d[Key] = Value
	case nil:
		return errSameKey
	}
	return nil
}
