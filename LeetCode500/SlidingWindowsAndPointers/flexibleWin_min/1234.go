package main

func balancedString(s string) int {
	cnt := ['X']int{} // 也可以用哈希表，不过数组更快一些
	for _, c := range s {
		cnt[c]++
	}

	m := len(s) / 4
	if cnt['Q'] == m && cnt['W'] == m && cnt['E'] == m && cnt['R'] == m {
		return 0 // 已经符合要求啦
	}

	ans := len(s)
	left := 0
	for right, c := range s { // 枚举子串右端点
		cnt[c]--
		// 下面这么处理是错误的，原因是只有满足循环条件的时候才能够更新ans，下面的写法会导致即使没有满足循环条件也会更新ans。（"QQQQ"）
		// for cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m {
		// 	cnt[s[left]]++
		// 	left++
		// }
		// ans = min(ans, right-(left-1)+1)
		for cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m {
			ans = min(ans, right-(left-1)+1)
			cnt[s[left]]++
			left++
		}
	}
	return ans
}

func balancedString1(s string) int {
	left := 0
	cnt := ['X']int{}
	ans := len(s)

	m := len(s) / 4
	for _, v := range s {
		cnt[v]++
	}

	if cnt['Q'] == m && cnt['W'] == m && cnt['E'] == m && cnt['R'] == m {
		return 0
	}
	for right := range s {
		cnt[s[right]]--

		for cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m {
			ans = min(ans, right-left+1)
			cnt[s[left]]++
			left++
		}
	}

	return ans
}
