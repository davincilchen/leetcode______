package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_ReverseList(t *testing.T) {

// 	assert.Equal(t, "bca", v)
// }

func Test_getSubstringCount(t *testing.T) {

	r := getSubstringCount("001100011")
	assert.Equal(t, int32(6), r)

	r = getSubstringCount("010101010")
	assert.Equal(t, int32(8), r)

	r = getSubstringCount("000110")
	assert.Equal(t, int32(3), r)

}

func Test_getSubstringCount2(t *testing.T) {

	r := getSubstringCount2("001100011")
	assert.Equal(t, int32(6), r)

	r = getSubstringCount2("010101010")
	assert.Equal(t, int32(8), r)

	r = getSubstringCount2("000110")
	assert.Equal(t, int32(3), r)

}
