/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	m := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	n := len(s)
	num := m[string(s[n-1])]

	for i := n - 2; i >= 0; i-- {
		// handle case IX, IV etc.
		if m[string(s[i])] < m[string(s[i+1])] {
			num -= m[string(s[i])]
		} else {
			num += m[string(s[i])]
		}
	}

	return num
}

// @lc code=end

