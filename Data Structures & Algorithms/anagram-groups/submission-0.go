func groupAnagrams(strs []string) [][]string {

    serialized := make(map[[26]int][]string)

    for _, s := range strs {
        cs := freq(s)
        serialized[cs] = append(serialized[cs], s)
    }

    return convertToArr(serialized)

}

func freq(s string) [26]int {
    c := [26]int{}
    for _, ch := range s {
        c[ch - 'a']++
    }
    return c
}

func convertToArr(m map[[26]int][]string) [][]string {
    out := [][]string{}

    for _, v := range m {
        out = append( out, v)
    }

    return out
}
