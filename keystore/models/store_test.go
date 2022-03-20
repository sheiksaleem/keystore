package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tempStore KeyValueData

func TestStore(t *testing.T) {
	tempStore.store = map[string]string{"NameA": "Employee A", "NameB": "Employee B", "User-1": "User A", "User-2": "User-B"}
	assert.Equal(t, "Employee A", tempStore.GetKey("NameA"))
	assert.Equal(t, true, tempStore.SetKey("NameC", "Employee C"))
	assert.Equal(t, "Employee C", tempStore.GetKey("NameC"))
	assert.Equal(t, len([]string{"NameA", "NameB", "NameC"}), len(tempStore.SearchPrefix("Name")))
	assert.Equal(t, []string{"NameA"}, tempStore.SearchSuffix("A"))
}
