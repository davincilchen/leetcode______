package main

//https://www.youtube.com/watch?v=t6xa2p6fFS8
//https://www.youtube.com/watch?v=3IFxpozBs2I
//https://www.youtube.com/watch?v=S3ckBIJG6fo
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


Failure Function = F(i)
1. 看前一個索引值i-1的Failure Function: F(i-1)
2. 判斷在LPPaS的位置跟當前字元有無相同：if P[F(i-1)] == P[i]:
3. 如果不同的話繼續找目前LPPaS的LPPas重複第二步驟，直到F等於零

n = len(s)
m = len(p)
failure_function = [0] * m
# Build failure function recording longest proper suffix and prefix
# dp
for i in range(1,m): //針對每個i處理
	j = failure_function[i-1] //前一個 已知的長度, 當成index就是要比較的地方 
	                          //1 index to 0 index

	while p[j] != p[i]: //先把退的前面最不一樣的?
		if j == 0:  //無法再退
			break
		j = failure_function[j-1] //

	if p[j] == p[i]:
		failure_function[i] = j + 1
	else:
		failure_function[i] = 0