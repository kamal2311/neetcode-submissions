type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    
    var encoded strings.Builder

    for _, s := range strs {
        fmt.Fprintf(&encoded, "%d#%s", len(s), s)
    }

    return encoded.String()

}


func (s *Solution) Decode(encoded string) []string {
    
    out := []string{}
    i := 0
    for i < len(encoded) {
        j := i
        for encoded[j] != '#'{
            j++
        }
        len, _ := strconv.Atoi(encoded[i:j])
        out = append(out, encoded[j+1:j+1+len])
        i = j + 1 + len
    }

    return out

}
