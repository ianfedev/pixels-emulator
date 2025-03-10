package encode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEnable enables bitmask and test if correctly enabled.
func TestEnable(t *testing.T) {
	mask := 0

	mask = Enable(mask, Thumbnail)
	assert.True(t, Has(mask, Thumbnail), "Thumbnail should be enabled")

	mask = Enable(mask, Guild)
	assert.True(t, Has(mask, Guild), "Guild should be enabled")
}

// TestHas test if bitmasks are correctly enabled.
func TestHas(t *testing.T) {
	mask := Thumbnail | Owner

	assert.True(t, Has(mask, Thumbnail), "Thumbnail should be set")
	assert.True(t, Has(mask, Owner), "Owner should be set")
	assert.False(t, Has(mask, Pets), "Pets should not be set")
}
