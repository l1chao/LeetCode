package main

// 100011001
func shortestBeautifulSubstring(s string, k int) string {
	left := 0
	cnt := 0
	ans := s

	for i, v := range s {
		cnt += int(v & 1)
		if cnt < k {
			continue
		}

		for cnt > k || s[left] == '0' { // cnt>k要减少'1'，是为了保证窗口内只有k个'1'；而左端的所有'0'都要除掉，是为了保证整个窗口最小。
			if s[left] == '1' {
				cnt--
			}
			left++
		}

		t := s[left : i+1]
		if len(t) < len(ans) || len(t) == len(ans) && t < ans {
			ans = t
		}
	}

	if cnt < k { // 一旦有至少一个满足k个'1'的窗口出现，后续cnt>=k。所以如果cnt没有达到k，就表明没有满足的子串。
		return ""
	}

	return ans
}

func shortestBeautifulSubstring1(s string, k int) string {
	left := 0
	cnt := 0
	ans := s

	for right, v := range s {
		cnt += int(v & 1)
		if cnt < k {
			continue
		}

		for cnt > k { // 修正窗口内cnt个数，修正完成后cnt == k
			cnt -= int(s[left] & 1)
			left++
		}
		for s[left] == 0 {
			left++
		}

		// if cnt == k
		t := s[left : right+1]
		if len(t) < len(ans) || len(t) == len(ans) && t < ans {
			ans = t
		}
	}
	if cnt < k {
		return ""
	}
	return ans
}
