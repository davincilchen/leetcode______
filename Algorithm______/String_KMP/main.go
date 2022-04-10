package main

import "fmt"

//https://www.youtube.com/watch?v=t6xa2p6fFS8
//https://www.youtube.com/watch?v=3IFxpozBs2I
//https://www.youtube.com/watch?v=S3ckBIJG6fo
//https://www.evanlin.com/about-kmp/
//次長共同前後綴（Longest Proper Prefix Which Is Also Suffix
//LPS

// 『失配字符』（匹配失敗的字符）是 兩邊要比第一個不知道批不匹配的位置
// 是前一次失敗的點

// ababc

// a     0
// ab    0
// aba   1   前"a" = 後"a"
// abab  2   前"ab" = 後"ab"
// ababc 0

//  ababc  次長前綴表
//  00120  ->比對字串時 , 下一個一樣的位置

//  ababc  匹配表
// -10012  ->比對字串時 , 往右多一格, 下一個不知道一不一樣要比較的點

// 0 代表比index 0
// -1 代表連第一個都不一樣 ,多移一格

//*找多個匹配
//那就用目前的匹配表(最後一個字母)再往右移即可
//

//***建構 prefix表 LPPS
//一直更新i和len(len一樣就可以往後更新,不一樣就要縮短,怎麼縮)

//移過來比較不相等要填啥
//  ABA
//ABABCABAA
//0012?

//ABABCABAA
//00120123?

//ABAB
//
//ABABCABAA
//0012,0123
//比不同 不同要填啥

//ABAB
//0012  <-2 抓2 但現在抓同一位一定一樣會不同,所以抓前一個來比
//ABAB CABAA                         //即修改len長度(縮短) 然後讓下次比 //都不一樣就縮到無法再縮
//0012,0123                          //簡單說不一樣就要抓前一位, 就像第一次比也是抓前一位當len
//抓len當index ,修正len為前一位
//因為比過不用重比 ,可少抓一位

//2-2不一樣時 如果len可以退就退, 不能退則目前為長度就是0,往下一位檢測
//移過來相同
//ABABCABAA
//001201?

// *抓前一位len 1 當 index 1過來
//ABABCABAA
//001201?
//ABABCABAA
//001201?

//建構時可以用前一個長度來幫助運算
//如果目前比較又一樣 ,那目前長度就是, 那就是ori len ++

// Failure Function = F(i)
// 1. 看前一個索引值i-1的Failure Function: F(i-1)
// 2. 判斷在LPPaS的位置跟當前字元有無相同：if P[F(i-1)] == P[i]:
// 3. 如果不同的話繼續找目前LPPaS的LPPas重複第二步驟，直到F等於零

// n = len(s)
// m = len(p)
// failure_function = [0] * m
// # Build failure function recording longest proper suffix and prefix
// # dp
// for i in range(1,m): //針對每個i處理
// 	j = failure_function[i-1] //前一個 已知的長度, 當成index就是要比較的地方
// 	                          //1 index to 0 index

// 	while p[j] != p[i]: //先把退的前面最不一樣的?
// 		if j == 0:  //無法再退
// 			break
// 		j = failure_function[j-1] //

// 	if p[j] == p[i]:
// 		failure_function[i] = j + 1
// 	else:
// 		failure_function[i] = 0

func PrefixTableNormal1(text string) []int {
	newTxt := []rune(text)
	cnt := len(newTxt)
	if cnt < 1 {
		return []int{}
	}
	ret := make([]int, cnt)

	length := 0
	i := 1
	for i < cnt {
		if newTxt[length] == newTxt[i] {
			length++
			ret[i] = length
			i++
			continue
		}

		//not the same
		//update length
		if length > 0 {
			length = ret[length-1]
			continue
		}
		//無法處理了 長度是0
		//找不到可以比較的length了
		length = 0
		ret[i] = 0
		i++
	}

	return ret
}

//直接用PrefixTableNormal1移比較會對
//不然要修正很多地方
func PrefixTableShift1(text string) []int {
	newTxt := []rune(text)
	cnt := len(newTxt)
	if cnt < 1 {
		return []int{}
	}
	ret := make([]int, cnt)
	ret[0] = -1
	length := 0
	i := 2
	for i < cnt {
		checkIdx := i - 1                       //與normal板差在不是比i是比 i-1
		if newTxt[length] == newTxt[checkIdx] { //前一位的LPS
			length++
			ret[i] = length
			i++
			continue
		}
		//修正不match的length
		//if length > 0 && ret[length-1] >= 0 {
		if length > 0 {
			length = ret[length-1+1] //因為有多移了一位 要移回來

			if length < 0 {
				length = 0
			}
			continue
		}
		//無法處理了 長度是0
		//找不到可以比較的length了
		length = 0 //ret[length-1] 會有-1,有機會沒reset length = 0
		ret[i] = 0
		i++
	}

	return ret
}

// 28. Implement strStr()
// Easy

// 3839

// 3580

// Add to List

// Share
// Implement strStr().

// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

// Clarification:

// What should we return when needle is an empty string? This is a great question to ask during an interview.

// For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

// Example 1:

// Input: haystack = "hello", needle = "ll"
// Output: 2
// Example 2:

// Input: haystack = "aaaaa", needle = "bba"
// Output: -1

// Constraints:

// 1 <= haystack.length, needle.length <= 104
// haystack and needle consist of only lowercase English characters
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	text := []rune(haystack)
	pattern := []rune(needle)
	table := PrefixTableShift1(needle)

	fmt.Println(len(text), len(pattern))
	i := 0
	j := 0
	for i < len(text) {
		// if i == 115 {
		// 	fmt.Println(text[115:])
		// }
		for j < len(pattern) {
			// if i == 115 {
			// 	fmt.Println(string(text[115:]))
			// }
			if i >= len(text) {
				return -1
			}
			if text[i] != pattern[j] {
				// if i > 110 {
				// 	fmt.Println(string(text[i:]))
				// }
				// fmt.Printf("i = %d, new j = %d, cnt = %d\n", i, j, cnt)
				j = table[j]
				// cnt = j
				// fmt.Printf("i = %d, new j = %d, cnt = %d\n", i, j, cnt)

				if j < 0 {
					//cnt = 0
					i++
					j++
				}
				break
			}
			if j == len(pattern)-1 {
				return i - len(pattern) + 1
			}
			// if i > 113 {
			// 	cnt++
			// 	fmt.Printf("i = %d, new j = %d, cnt = %d\n", i, j, cnt)
			// }

			i++
			j++
		}
	}
	return -1
}
