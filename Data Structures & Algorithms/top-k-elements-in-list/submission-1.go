func topKFrequent(nums []int, k int) []int {

    freq := map[int]int{}

    for _, n := range nums {
        freq[n]++
    }

    pairs := make([]pair, len(freq))

    for k,v := range freq {
        pairs = append(pairs, pair{k,v})
    }

    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].v > pairs[j].v 
    })

    out := make([]int, k)

    for i := range k {
        out[i] = pairs[i].k
    }

    return out   

}

type pair struct {
    k, v int
}
