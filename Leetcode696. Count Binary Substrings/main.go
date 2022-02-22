package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'getSubstringCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func countBinarySubstrings(s string) int {
	return int(getSubstringCount(s))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getSubstringCount(s string) int32 {
	// Write your code here
	ret := 0
	pre := 0
	cur := 1

	for i, v := range s {
		if i == 0 {
			continue
		}
		if v == rune(s[i-1]) {
			cur++
			continue
		}

		//different
		ret += min(cur, pre)

		pre = cur
		cur = 1

	}

	ret += min(cur, pre)

	return int32(ret)
}

func getSubstringCount2(s string) int32 {
	// Write your code here
	ret := 0
	pre := 0
	cur := 1

	for i, v := range s {
		if i == 0 {
			continue
		}
		if v == rune(s[i-1]) {
			cur++
		} else {
			pre = cur
			cur = 1
		}

		//can count
		if cur <= pre {
			ret++
		}
	}

	return int32(ret)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := getSubstringCount(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
