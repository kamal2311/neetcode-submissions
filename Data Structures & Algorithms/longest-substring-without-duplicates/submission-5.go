func lengthOfLongestSubstring(s string) int {

	seen := map[byte]int{}
	var l, r, longest int

	for r < len(s) {
		if i, ok := seen[s[r]]; ok {
			l = max(l, i + 1)			
		} 
		seen[s[r]] = r			
		len := r - l + 1
		if len > longest {
			longest = len
		}
		r++
	}

	return longest

}
