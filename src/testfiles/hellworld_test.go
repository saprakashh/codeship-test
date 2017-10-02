package src

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	var st string = "Hello World"
	assert.Equal(t, st, "Hello World", "The two string should be the same")
}
