package models_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"git.github.com/pantafive/demo-go-generic/models"
)

func TestJSONUnmarshal(t *testing.T) {
	type Some struct {
		Value models.ValueObject[int] `json:"value"`
	}

	tests := []struct {
		name      string
		given     string // json string, e.g. `{"value": 42}`
		wantValue int
		wantOK    bool
		wantErr   bool
	}{
		{name: "no value", given: `{"something": "Alice"}`, wantValue: 0, wantOK: false, wantErr: false},
		{name: "zero value", given: `{"value": 0}`, wantOK: true, wantErr: false},
		{name: "usual value", given: `{"value": 1}`, wantValue: 1, wantOK: true, wantErr: false},
		{name: "broken value", given: `{"value": "hello"}`, wantValue: 0, wantOK: false, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var user Some

			err := json.Unmarshal([]byte(tt.given), &user)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			value, ok := user.Value.Get()
			assert.Equal(t, tt.wantValue, value)
			assert.Equal(t, tt.wantOK, ok)
		})
	}
}

func TestValueObject_Clean(t1 *testing.T) {
	var vo models.ValueObject[int]

	vo.Set(1)
	vo.Clean()

	value, ok := vo.Get()
	assert.Equal(t1, 0, value)
	assert.False(t1, ok)
}
