package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {

	//  2*(1+2*1)                        //(1+ ((1)*2))*2
	assert.Equal(t, 6, scoreOfParentheses("(()(()))"))

	//  (1+1+1*2)*2                     //  (1+ 1+((1)*2))*2
	assert.Equal(t, 8, scoreOfParentheses("(()()(()))"))

	assert.Equal(t, 1, scoreOfParentheses("()"))

	assert.Equal(t, 2, scoreOfParentheses("(())"))

	assert.Equal(t, 2, scoreOfParentheses("()()"))

	assert.Equal(t, 4, scoreOfParentheses("(())()()"))
}
