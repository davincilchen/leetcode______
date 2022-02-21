package main

//(()))(()))()())))  //最後三個會扣掉前面的
func minInsertions(s string) int {

	rr := 0
	ans := 0
	for _, v := range s {
		chk := string(v)
		if chk == "(" { //可能有需要結算
			if rr%2 == 1 {
				// need to insert one ')'
				// case '(()(' -> '(())('  // 沒辦法從外圍補的 )*n 被夾在 ((中間
				// case '()(' -> '())('
				// case ')(' -> '())('
				rr--  //結算挪移 才不會被reset
				ans++ //結算挪移 , == ll+rr
			}
			rr += 2
		} else {
			rr--
		}

		if rr < 0 { //多1個或2個r,確定左邊少l可以先處理
			ans++
			rr += 2
		}
	}

	return ans + rr

}

//r == -1 時可以先補ll了, 因為r多個一個要退 , 已經可以先補ll
func minInsertions2(s string) int {
	ans := 0
	rr := 0
	//用) 的前一個)判斷?
	for _, v := range s {
		chk := string(v)
		if chk == ")" { //
			rr--
			if rr < 0 { //-1先處理 不會有-2 , r == -1 時可以先補ll了
				ans++
				rr += 2
			}
		} else {
			if rr%2 == 1 {
				// need to insert one ')'
				// case '(()(' -> '(())('  // ??
				// case '()(' -> '())('
				rr--  //結算挪移
				ans++ //ll+已結案rr
			}
			rr += 2 // need two ')'s
		}

	}

	return ans + rr

}

func main() {

}

// func costToBalance(s string) int {

// 	ret := 0
// 	rr := 0
// 	ll := 0
// 	p := []rune(s)
// 	n := len(p)
// 	for i, v := range p {
// 		chk := string(v)
// 		if chk == "(" {
// 			//process ( in count
// 			if ll == 0 {
// 				ret += rr / 2
// 				if rr%2 == 1 {
// 					ret += 1
// 				}
// 				rr = 0
// 			} else {
// 				need := ll*2 - rr
// 				if need > 0 {
// 					ret += need
// 				} else {
// 					cnt := rr - (ll * 2)
// 					if cnt == 1 {
// 						ret += 2
// 					} else if cnt == 2 {
// 						ret += 1
// 					} else {
// 						ret += (cnt / 2) //補 cnt / 2的 ()
// 						if cnt%2 == 0 && cnt != 0 {
// 							ret += 2 //補()
// 						}
// 					}
// 					// (())))) ->() , (())))))
// 				}
// 				rr = 0
// 				ll = 0
// 			}

// 			if i == n-1 {
// 				ret += 2
// 				break
// 			} else if i == n-2 {
// 				if string(p[i+1]) == ")" {
// 					ret += 1
// 					break
// 				}
// 				ret += 4
// 				break
// 			}
// 			ll++

// 			continue
// 		}
// 		rr++
// 	}
// 	return ret
// }
