package main

// func PrefixTable(str string) []int {
// 	ss := []rune(str)
// 	cnt := len(ss)
// 	ret := make([]int, cnt)
// 	if cnt > 0 {
// 		ret[0] = -1
// 	}
// 	i := 2
// 	length := 0
// 	for i < cnt {
// 		checkIdx := i - 1
// 		if ss[checkIdx] == ss[length] {
// 			length++
// 			ret[i] = length
// 			i++
// 			continue
// 		}
// 		//not match
// 		if length > 0 {
// 			length = ret[length-1]
// 			if length < 0 {
// 				length = 0
// 			}
// 			continue
// 		}

// 		//length == 0 && chenk failed
// 		length = 0
// 		ret[i] = 0
// 		i++

// 	}

// 	return ret

// }

func PrefixTable(str string) []int {
	ss := []rune(str)
	cnt := len(ss)
	ret := make([]int, cnt)

	i := 1
	length := 0
	for i < cnt {
		checkIdx := i
		if ss[checkIdx] == ss[length] {
			length++
			ret[i] = length
			i++
			continue
		}
		//not match
		if length > 0 {
			length = ret[length-1]
			if length < 0 {
				length = 0
			}
			continue
		}

		//length == 0 && chenk failed
		length = 0
		ret[i] = 0
		i++

	}

	return ret

}

//加數版
func repeatedSubstringPattern(s string) bool {

	cnt := len(s)
	if cnt <= 0 {
		return false
	}
	table := PrefixTable(s)
	lps := table[cnt-1]

	if cnt%(cnt-lps) == 0 {
		return true
	}
	return false
}

// func repeatedSubstringPattern(s string) bool {

// 	cnt := len(s)
// 	for l := 1; l < cnt; l++ {
// 		if cnt%l != 0 {
// 			continue
// 		}
// 		cur := s[0:l]
// 		i := 0
// 		same := true
// 		for i <= cnt-l {
// 			check := s[i : i+l]
// 			if check != cur {
// 				same = false
// 				break
// 			}
// 			i += l
// 		}
// 		if same {
// 			return true
// 		}
// 	}

// 	return false
// }
