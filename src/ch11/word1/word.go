package word

func IsPalindrome(s string) bool {
	s_len := len(s)
	for i := range s {
		if s[i] != s[s_len-1-i] {
			return false
		}
	}
	return true
}
