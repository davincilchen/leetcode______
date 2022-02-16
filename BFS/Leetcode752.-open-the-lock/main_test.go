package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_openLock(t *testing.T) {

	dead := []string{}
	v := 0

	dead = []string{"0201", "0101", "0102", "1212", "2002"}
	v = openLock(dead, "0202")
	assert.Equal(t, 6, v)

}
