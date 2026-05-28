func isAnagram(s string, t string) bool {
    
    if len(s) != len(t) {
        return false
    }

    sCharCounts := make(map[byte]int)
    tCharCounts := make(map[byte]int)
    
    for i := 0; i < len(s); i++ {
        sCharCounts[s[i]]++
        tCharCounts[t[i]]++
    }

    if len(sCharCounts) != len(tCharCounts) {
        return false
    }

    for k,v := range sCharCounts {
        if tCharCounts[k] != v {
            return false
        }
    }

    return true
}
