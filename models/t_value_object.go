package models

import "encoding/json"

type ValueObject[T any] struct {
	value    T
	notEmpty bool
}

func (t *ValueObject[T]) Get() (T, bool) {
	return t.value, t.notEmpty
}

func (t *ValueObject[T]) Set(value T) {
	t.value = value
	t.notEmpty = true
}

func (t *ValueObject[T]) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &t.value); err != nil {
		return err
	}

	t.notEmpty = true
	return nil
}
