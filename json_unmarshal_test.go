package unexpected

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONUnmarshal(t *testing.T) {
	var testName = "test name"
	type st struct {
		MyName string `json:"MyName"`
	}
	j := fmt.Sprintf(`{"MyName": "%s", "myName": "wrong"}`, testName)

	var v st
	err := json.Unmarshal([]byte(j), &v)
	require.NoError(t, err)

	assert.NotEqual(t, testName, v.MyName)
	// how to fix - add one extra field in struct:
	// NoName string `json:"myName"`
}
