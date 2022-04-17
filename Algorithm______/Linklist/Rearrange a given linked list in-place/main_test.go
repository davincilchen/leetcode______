package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrefixTableNormal1(t *testing.T) {

	//no compare
	assert.Equal(t, []int{0, 0, 1, 2, 0, 1, 2, 3, 1}, PrefixTableNormal1("ABABCABAA"))
	assert.Equal(t, []int{0, 0, 1}, PrefixTableNormal1("ABA"))
	assert.Equal(t, []int{0, 0}, PrefixTableNormal1("AB"))
	assert.Equal(t, []int{0}, PrefixTableNormal1("A"))

	assert.Equal(t, []int{0, 0, 0, 0, 0, 0, 0, 1, 2, 0, 0, 0, 0, 0, 0, 1, 2, 3, 0, 0, 0, 0, 0, 0}, PrefixTableNormal1("PARTICIPATE IN PARACHUTE"))

	assert.Equal(t, []int{0, 1, 0, 1, 2, 3}, PrefixTableNormal1("bbabba"))
	assert.Equal(t, []int{0, 1, 0, 1, 2, 2, 0}, PrefixTableNormal1("aabaaac"))
	assert.Equal(t, []int{0, 0, 1, 2, 0, 1, 1, 2, 0}, PrefixTableNormal1("ababcaabc"))
}

func Test_PrefixTableShift1(t *testing.T) {

	//no compare
	assert.Equal(t, []int{-1, 0, 0, 1, 2, 0, 1, 2, 3}, PrefixTableShift1("ABABCABAA"))
	assert.Equal(t, []int{-1, 0, 0}, PrefixTableShift1("ABA"))
	assert.Equal(t, []int{-1, 0}, PrefixTableShift1("AB"))
	assert.Equal(t, []int{-1}, PrefixTableShift1("A"))

	assert.Equal(t, []int{-1, 0, 0, 0, 0, 0, 0, 0, 1, 2, 0, 0, 0, 0, 0, 0, 1, 2, 3, 0, 0, 0, 0, 0}, PrefixTableShift1("PARTICIPATE IN PARACHUTE"))

	assert.Equal(t, []int{-1, 0, 1, 0, 1, 2}, PrefixTableShift1("bbabba"))
	assert.Equal(t, []int{-1, 0, 1, 0, 1, 2, 2}, PrefixTableShift1("aabaaac"))

	assert.Equal(t, []int{-1, 0, 0, 1, 2, 0, 1, 1, 2}, PrefixTableShift1("ababcaabc"))
}

func Test_strStr(t *testing.T) {
	des := "bbababaaaababbaabbbabbbaaabbbaaababbabaabbaaaaabbaaabbbbaaabaabbaababbbaabaaababbaaabbbbbbaabbbbbaaabbababaaaaabaabbbababbaababaabbaa"
	pattern := "bbabba"
	//"aaaaabaabbb_ababbaababaabbaa" //133
	assert.Equal(t, -1, strStr(des, pattern))

	des = "hello"
	pattern = "ll"
	assert.Equal(t, 2, strStr(des, pattern))

	des = "aabaaabaaac"
	pattern = "aabaaac"
	assert.Equal(t, 4, strStr(des, pattern))

	des = "ababcaababcaabc"
	pattern = "ababcaabc"
	assert.Equal(t, 6, strStr(des, pattern))

}
