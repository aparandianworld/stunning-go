package mathops

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, Add(2.0, 3.0), 5.0)
}

func TestSub(t *testing.T) {
	assert.Equal(t, Sub(5.0, 3.0), 2.0)
}

func TestMul(t *testing.T) {
	assert.Equal(t, Mul(2.0, 3.0), 6.0)
}

func TestDiv(t *testing.T) {
	result, err := Div(6.0, 3.0)

	require.NoError(t, err)
	assert.Equal(t, 2.0, result)
}

func TestDivZero(t *testing.T) {
	_, err := Div(6.0, 0.0)

	require.Error(t, err)
}
