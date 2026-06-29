type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var ans strings.Builder
	for _, str := range strs {
		l := strconv.Itoa(len(str))
		ans.WriteString(l)
		ans.WriteByte('#')
		ans.WriteString(str)
	}
	return ans.String()
}

func (s *Solution) Decode(encoded string) []string {
		res := []string{}
		
		start := 0
		
		for i := 0; i < len(encoded); {
			r := encoded[i]
			if r == '#'{
				len,_ := strconv.Atoi(encoded[start:i])
				res = append(res, encoded[i+1:i+1+len])
				start = i + 1 + len
				i = start
			}
			i++

		}

		return res
}
