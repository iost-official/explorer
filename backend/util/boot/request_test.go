package boot

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckBase58(t *testing.T) {
	assert.Equal(t, -1, strings.IndexFunc("ABCDEFG", checkBase58))
	assert.NotEqual(t, -1, strings.IndexFunc("0", checkBase58))
	assert.NotEqual(t, -1, strings.IndexFunc("!@#", checkBase58))
}

func TestCheckId(t *testing.T) {
	assert.Equal(t, -1, strings.IndexFunc("qwertyjack", checkId))
	assert.Equal(t, -1, strings.IndexFunc("da_shen", checkId))
	assert.NotEqual(t, -1, strings.IndexFunc("Qwerty", checkId))
	assert.NotEqual(t, -1, strings.IndexFunc("!@#", checkId))
}
