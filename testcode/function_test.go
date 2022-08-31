package testcode

import (
	f "golang-tuckers/function/libs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	rst := f.Add(3, 5)
	if rst != 8 {
		t.Errorf("8 should be returned %d", rst)
	}
	assert := assert.New(t)
	assert.Equal(11, f.Add(5, 6), "should be 11")
}

func TestMultiReturn(t *testing.T) {
	assert := assert.New(t)
	max, message := f.MultiReturn(1, 3)
	assert.NotEqualf(3, max, "not equal")
	assert.NotNilf(message, "not nil")
}

func TestArray(t *testing.T) {
	assert := assert.New(t)
	array := []int{1, 2}
	assert.Len(array, 2, "array length is not 2")
}
