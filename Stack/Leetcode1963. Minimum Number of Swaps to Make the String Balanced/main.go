package main

func costToBalance(s string) int {
	ll := 0
	rr := 0
	p := []rune(s)
	//n := len(p)
	for _, v := range p {
		chk := string(v)
		if chk == "(" {
		}
	}

	return ll + rr

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
