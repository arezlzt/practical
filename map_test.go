package practical

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapShift(t *testing.T) {
	expect := "mapShift"
	actual := MapShift(map[string]string{"shift": "mapShift"}, "shift")
	assert.Equal(t, expect, actual)
}
func TestMapKeyExists(t *testing.T) {
	expect := true
	actual := MapKeyExists(map[string]string{"shift": "mapShift"}, "shift")
	assert.Equal(t, expect, actual)
}
func TestMapKeys(t *testing.T) {
	expect := []string{
		"shift",
		"keyExists",
		"keys",
	}
	actual := MapKeys(map[string]string{"shift": "mapShift", "keyExists": "mapKeyExists", "keys": "mapKeys"})
	assert.Equal(t, expect, actual)
}
func TestMapValues(t *testing.T) {
	expect := []string{
		"mapShift",
		"mapKeyExists",
		"mapKeys",
	}
	actual := MapValues(map[string]string{"shift": "mapShift", "keyExists": "mapKeyExists", "keys": "mapKeys"})
	assert.Equal(t, expect, actual)
}
func TestMapFlip(t *testing.T) {
	expect := map[string]string{
		"mapShift": "shift", "mapKeyExists": "keyExists", "mapKeys": "keys",
	}
	actual := MapFlip(map[string]string{"shift": "mapShift", "keyExists": "mapKeyExists", "keys": "mapKeys"})
	assert.Equal(t, expect, actual)
}
