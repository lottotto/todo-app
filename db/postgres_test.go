package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	expect := "hoge"
	actual := getEnv("EXAMPLE", "hoge")
	assert.Equal(t, expect, actual)
}
