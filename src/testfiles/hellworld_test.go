package src

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIntCompare(t *testing.T) {
	var num1 int = 45
	var num2 int = 45
	assert.Equal(t, num1, num2, "The numbers are different")
}

func TestHelloWorld(t *testing.T) {
	var st string = "Hello World"
	assert.Equal(t, st, "Hello World", "The two string should be the same")
}
