package main

//https://leetcode.com/problems/open-the-lock/

// You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.

// The lock initially starts at '0000', a string representing the state of the 4 wheels.

// You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.

// Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.

// Example 1:

// Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
// Output: 6
// Explanation:
// A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
// Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
// because the wheels of the lock become stuck after the display becomes the dead end "0102".
// Example 2:

// Input: deadends = ["8888"], target = "0009"
// Output: 1
// Explanation: We can turn the last wheel in reverse to move from "0000" -> "0009".
// Example 3:

// Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
// Output: -1
// Explanation: We cannot reach the target without getting stuck.

// Constraints:

// 1 <= deadends.length <= 500
// deadends[i].length == 4
// target.length == 4
// target will not be in the list deadends.
// target and deadends[i] consist of digits only.

//怎麼套到 BFS 的框架裡呢？首先明確一下起點 start 和終點 target 是什麼，怎麼判斷到達了終點？

//normal slice queue + replaceAtIndex2
// Runtime: 178 ms, faster than 44.90% of Go online submissions for Open the Lock.
// Memory Usage: 8.4 MB, less than 26.53% of Go online submissions for Open the Lock.

//normal slice queue + replaceAtIndex1
// Runtime: 128 ms, faster than 65.31% of Go online submissions for Open the Lock.
// Memory Usage: 8.3 MB, less than 34.69% of Go online submissions for Open the Lock.

//pre allow (addQueue2)+  replaceAtIndex1 [為什麼比較慢??]
// Runtime: 189 ms, faster than 36.74% of Go online submissions for Open the Lock.
// Memory Usage: 8.6 MB, less than 18.37% of Go online submissions for Open the Lock.

//pre allow (addQueue2)+  replaceAtIndex1 [拿掉一行重跑 有快一點但沒快很多]
// Runtime: 124 ms, faster than 69.39% of Go online submissions for Open the Lock.
// Memory Usage: 8.8 MB, less than 17.35% of Go online submissions for Open the Lock.

func replaceAtIndex1(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
}

// func replaceAtIndex2(str string, replacement rune, index int) string {
//     return str[:index] + string(replacement) + str[index+1:]
// }

func rollup(s string, idx int) string {
	c := s[idx]
	if c == '9' {
		c = '0'
	} else {
		c += 1
	}

	return replaceAtIndex1(s, rune(c), idx)
	// s1 := s[:idx]
	// s2 := s[idx+1:]
	// return s1 + string(c) + s2
}

func rolldown(s string, idx int) string {
	c := s[idx]
	if c == '0' {
		c = '9'
	} else {
		c -= 1
	}

	return replaceAtIndex1(s, rune(c), idx)
	// s1 := s[:idx]
	// s2 := s[idx+1:]
	// return s1 + string(c) + s2
}

func addQueue(s string, queue *[]string, visited map[string]string) {
	_, ok := visited[s]
	if ok {
		return
	}
	visited[s] = s
	//queue = append(queue, s) //if not pointer //this value of dead is never used (SA4006)
	*queue = append(*queue, s)
}

func addQueue2(s string, queue []string, visited map[string]string) []string {
	_, ok := visited[s]
	if ok {
		return queue
	}
	visited[s] = s

	n := len(queue)
	if n > 0 && n == cap(queue) {
		newCap := n * 2
		//not: newQ := make([]string, newCap, newCap)
		//not: newQ := make([]string, 0, newCap)
		// newQ 的元素數量如果「多於」被複製進去的元素時，會用 zero value 去補。例如，當 cloneScores 的長度是 4，但只複製 3 個元素進去時，最後位置多出來的元素會補 zero value。
		// newQ 的元素數量如果「少於」被複製進去的元素時，超過的元素不會被複製進去。例如，當 cloneScores 的長度是 1，但卻複製了 3 個元素進去時，只會有 1 個元素被複製進去。
		newQ := make([]string, n, newCap) //!!!!
		copy(newQ, queue)
		return append(newQ, s)
	}

	return append(queue, s)
}

func openLock(deadends []string, target string) int {
	visited := make(map[string]string)

	for _, v := range deadends {
		visited[v] = v
	}

	_, ok := visited["0000"]
	if ok {
		return -1
	}

	turn := 0
	queue := []string{"0000"}

	for {
		n := len(queue)
		if n <= 0 {
			break
		}

		for k := 0; k < n; k++ {
			cur := queue[0]
			if cur == target {
				return turn
			}

			queue = queue[1:]
			for i := range cur {
				//addQueue(rollup(cur, i), &queue, visited)
				//addQueue(rolldown(cur, i), &queue, visited)
				queue = addQueue2(rollup(cur, i), queue, visited)
				queue = addQueue2(rolldown(cur, i), queue, visited)

			}

		}
		turn++
	}

	return -1
}

func main() {

}
