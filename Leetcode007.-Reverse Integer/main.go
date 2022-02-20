package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
// Memory Usage: 2.3 MB, less than 38.32% of Go online submissions for Reverse Integer.
func reverse(x int) int {

	if x > 0 {
		return reverseP(x)
	}
	return -reverseP(-x)
}

//const INT_MAX = int(^uint(0) >> 1)
const INT_MAX32 = int32(^uint32(0) >> 1)

func reverseP(x int) int {

	v := x / 10
	if v == 0 {
		return x
	}
	t := []int{x % 10}
	for v > 0 {
		t = append(t, v%10)
		v /= 10
	}

	ret := 0
	n := len(t)
	for i := 0; i < n; i++ {
		if i == 0 {
			ret = t[i]
			continue
		}
		ret = ret*10 + t[i]
	}

	if ret > int(INT_MAX32) {
		return 0
	}

	return ret
}

func main() {

}

// Runtime: 8 ms, faster than 10.47% of C++ online submissions for Reverse Integer.
// Memory Usage: 5.8 MB, less than 76.65% of C++ online submissions for Reverse Integer.
// class Solution {
// 	public:
// 		int reverse(int x) {
// 			int rev = 0;
// 			while (x != 0) {
// 				int pop = x % 10;
// 				x /= 10;
// 				if (rev > INT_MAX/10 || (rev == INT_MAX / 10 && pop > 7)) return 0;
// 				if (rev < INT_MIN/10 || (rev == INT_MIN / 10 && pop < -8)) return 0;
// 				rev = rev * 10 + pop;
// 			}
// 			return rev;
// 		}
// 	};
