package models_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
			user := Some{}

			err := json.Unmarshal([]byte(tt.given), &user)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			value, ok := user.Value.Get()
			assert.Equal(t, tt.wantValue, value)
			assert.Equal(t, tt.wantOK, ok)
		})
	}
}
