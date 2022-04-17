package main

// 我们可以对空间复杂度进行进一步的优化，
// 并不需要使用栈去保留所有的中间情况，
// 可以只用一个变量 cnt 来记录当前在第几层括号之中，
// 因为本题的括号累加值是有规律的，"()" 是1，
// 因为最中间的括号在0层括号内，2^0 = 1。"(())" 是2，
// 因为最中间的括号在1层括号内，2^1 = 2。"((()))"
// 是4，因为最中间的括号在2层括号内，2^2 = 4。
// 因此类推，其实只需要统计出最中间那个括号外变有几个括号，
// 就可以直接算出整个多重包含的括号字符串的值，参见代码如下：
// 解法三：

func scoreOfParentheses(s string) int {

	ret := 0
	cnt := 0
	p := []rune(s)

	for i, v := range p {
		if string(v) == "(" {
			cnt++
			//fmt.Printf("v = %s, i = %d,++ cnt= %d, ret = %d \n", string(v), i, cnt, ret)
		} else {
			cnt--
			//fmt.Printf("v = %s, i = %d,-- cnt= %d, ret = %d \n", string(v), i, cnt, ret)
		}

		if string(v) == ")" && string(p[i-1]) == "(" {
			//fmt.Printf("v = %s, i = %d, cnt= %d, old ret = %d, new ret = %d \n", string(v), i, cnt, ret, (ret + (1 << cnt)))
			ret += 1 << cnt

		}

	}
	return ret

}

func main() {

}
