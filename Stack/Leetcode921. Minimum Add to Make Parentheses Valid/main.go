package main

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Add to Make Parentheses Valid.
//Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Minimum Add to Make Parentheses Valid.
//Next challenges:

func minAddToMakeValid(s string) int {
	ll := 0
	rr := 0
	//p := []rune(s)
	//n := len(p)
	for _, v := range s {
		chk := string(v)
		if chk == "(" { //右邊需要)
			rr += 1
		} else { //右邊可以少一個)
			rr -= 1
		}
		//要保證rr >= 0 ?
		if rr == -1 { //多一個r,確定左邊少l可以先處理
			ll++
			rr++
		}
	}

	return ll + rr

}

// //失敗
//因為左邊讀到右邊 , (要在左邊 ,讀到右邊的-1 無法補左邊缺少的
// func minAddToMakeValid2(s string) int {
// 	ll := 0
// 	rr := 0
// 	//p := []rune(s)
// 	//n := len(p)
// 	for _, v := range s {
// 		chk := string(v)
// 		if chk == "(" { //可以少一個(
// 			ll--
// 		} else { //需要一個(
// 			ll++
// 		}
// 		//要保證ll >= 0 ?
// 		if ll == -1 {
// 			ll++
// 			rr++
// 		}
// 	}

// 	return ll + rr

// }

func main() {

}
