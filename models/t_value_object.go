package models

import "encoding/json"

type ValueObject[T comparable] struct {
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

//func (t *ValueObject[T]) Equals(other ValueObject[T]) bool {
//	if t.value != other.value || t.notEmpty != other.notEmpty {
//		return false
//	}
//	return true
//}

func (t *ValueObject[T]) Clean() {
	var noop T
	t.value = noop
	t.notEmpty = false
}

func (t *ValueObject[T]) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &t.value); err != nil {
		return err
	}

	t.notEmpty = true
	return nil
}
