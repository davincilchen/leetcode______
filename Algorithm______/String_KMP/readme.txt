*
https://www.youtube.com/watch?v=dgPabAsTFa8
https://www.youtube.com/watch?v=3IFxpozBs2I


https://www.itread01.com/content/1505656572.html
https://zhuanlan.zhihu.com/p/36190375

https://medium.com/nlp-tsupei/kmp%E7%AE%97%E6%B3%95%E8%A9%B3%E8%A7%A3-1b1050a45850



最近在看算法，覺得kmp算法的一些學習心得可以記錄一下。

我本身是在看《算法》的，裏面介紹kmp算法時，實在看的一臉懵逼，就看了別人的心得，這裏推薦2篇博文：

1.阮一峰大大的：http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html

2.實現1中求next的博文：http://www.cnblogs.com/c-cloud/p/3224788.html

根據1文我實現了kmp，但留空next的求法，畢竟1文也只是大概說了next的意義，對2文我在自己的理解上也自己實現了一波，覺得更加容易懂些吧。

個人添加的理解（也都在註釋裏了）：

　　1.next的本質是最長共同前綴後綴，然後在匹配失敗時，指向模板字符串的指針直接重置為next[k-1]（即重置到最長共同前後綴的下一個位置），而長字符串的指針勿需動，即相當於指針新位置前的那一坨已經匹配成功了，接著從新指針那裏匹配下去即可

　　2.（參考下面代碼）計算next時，如果p[i]=p[k]則共同前後綴+1，不等則在已經匹配成功的部分p[0~k-1]中找最大共同前後綴next[k-1]，然後繼續比較p[i]和p[新k]，如果k<=0即得不到已經存在的next的幫助了，這時k=0，即木有共同前後綴了。

// next的本質是最長共同前綴後綴，然後在匹配失敗時，指向模板字符串的指針直接重置為next[k-1]，而長字符串的指針勿需動，即相當於指針新位置前的那一坨已經匹配成功了，接著從新指針那裏匹配下去即可
    void CalculateNext(string pattern, ref int[] next)
    {
        next[0] = 0;
        for (int i = 1, k = 0; i < pattern.Length; ++i)
        {
            while(pattern[i] != pattern[k])
            {
                if (k > 0)
                    k = next[k - 1];//這個有點難理解，因為她是匹配失敗後(p[i]!=p[k])，
                                    //直接使用匹配成功部分的共同前後綴接著進行匹配p[i]和p[新k]，此時因為新k=next[k-1],所以p[x~i]=p[0-新k]
                else
                {
                    k = -1;//得不到前面next的幫助了，需要置k=0，並且需要跳出while
                    break;
                }
            }

            k++;
            next[i] = k;
        }
    }

    void KMP(string inStr, string pattern, int[] next)
    {
        for (int i = 0, k = 0; i < inStr.Length; i++)
        {
            if (pattern[k] == inStr[i])
            {
                k++;
                if (k == pattern.Length)
                {
                    break;
                }
            }
            else
            {
                if (k != 0)//k=0特殊處理，直接移動i即可，k依然是0
                {
                    k = next[k-1];//一開始我是用1文的說法（移動位數 = 已匹配的字符數 - 對應的部分匹配值）：pattern右移d=k-next[k-1]，即k需回退x才能維持和i對應(分別對著兩字符串的下一個比較字符)，x=k-d=next[k-1]
　　　　　　　　　　　　　　　　　　　　//即，k=next[k-1]，所以理解為：k應該重置到最長共同前後綴的下一個位置，這樣比1文的說法容易理解些（因為next裏是長度,k是下標，所以剛好next[k-1]就是最長共同前後綴的下一個位置的下標）
                }
            }
        }
    }