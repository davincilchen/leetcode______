package main

// func scoreOfParentheses(s string) int { //被包住總共只有兩分

// 	ret := 0
// 	stack := make([]string, 0)
// 	p := []rune(s)
// 	n := len(p)
// 	for i := 0; i < n; i++ {
// 		cur := string(p[i])
// 		if cur == "(" {
// 			next := string(p[i+1]) //s must balance
// 			if next == ")" {
// 				idx := i
// 				i++
// 				//drop "()"
// 				if idx+2 >= len(p) { //final close "")""
// 					ret += 1
// 					continue
// 				} else if idx-1 < 0 { // first start
// 					ret += 1
// 					continue

// 				} else if string(p[idx-1]) == ")" && string(p[idx+2]) == "(" {
// 					ret += 1
// 					continue
// 				}

// 				continue
// 			}
// 			stack = append(stack, cur)
// 			continue
// 		}

// 		if len(stack) == 0 {
// 			break
// 		}

// 		ret += 2
// 		for j := i + 1; j < n; j++ { //drop
// 			if string(p[j]) != ")" {
// 				break
// 			}
// 			i = j
// 			if len(stack) == 0 {
// 				break
// 			}
// 			stack = stack[0 : len(stack)-1]
// 		}
// 	}

// 	return ret
// }

// 其实也是将当前状态压入栈中，等递归退出后再从栈中取出之前的状态。
// 这里的实现思路是，遍历字符串S，当遇到左括号时，将当前的分数压入栈中
// ，并把当前得分清0，若遇到的是右括号，
// 说明此时已经形成了一个完整的合法的括号字符串了，而且除去外层的括号，
// 内层的得分已经算出来了，就是当前的结果 res，此时就要乘以2，
// 并且要跟1比较，取二者中的较大值，
// 这样操作的原因已经在上面解法的讲解中解释过了。然后还要加上栈顶的值，
// 因为栈顶的值是之前合法括号子串的值，跟当前的是并列关系，
// 所以是相加的操作，最后不要忘了要将栈顶元素移除即可，参见代码如下：
func scoreOfParentheses(s string) int {

	ret := 0
	stack := make([]int, 0)
	for _, v := range s {
		//看到(表示新的開始要把舊的分數存起來,下次
		//(( 還沒有值的新的開始. ret*2 , ret用0先暫存
		//)( 有值的新的開始.   ret + ?
		if string(v) == "(" { //store and reset ret
			stack = append(stack, ret)
			ret = 0
			continue
		}

		//可能是1分或*2
		tmp := 1
		if ret*2 > tmp { //可能是第一個) -> 1 , 或第n個) -> *2
			tmp = ret * 2
		}
		ret = stack[len(stack)-1] + tmp
		stack = stack[0 : len(stack)-1]
	}
	return ret

}

func main() {

}
