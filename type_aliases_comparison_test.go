package unexpected

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeAliases_Comparison(t *testing.T) {
	type TestString string
	var s TestString = "test"

	assert.True(t, "test" == s)   // equals
	assert.NotEqual(t, "test", s) // not equals
	// the reason is the different types

	// it happens when we want to change field type,
	// but after that the tests are broken
	var v struct {
		OldType string     `json:"old_type"`
		NewType TestString `json:"new_type"`
	}
	j := `{"old_type": "test", "new_type": "test"}`
	err := json.Unmarshal([]byte(j), &v)
	require.NoError(t, err)

	assert.Equal(t, "test", v.OldType)    // old test
	assert.NotEqual(t, "test", v.NewType) // new test
}
