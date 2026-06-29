import "slices"

func evalRPN(tokens []string) int {
	// if num, push stack
	// if op, pop last 2 and push the result
	// in the end pop and return

	ops := []string{"+","-","*","/"}
	stack := &Stack{}

	for _, tkn := range tokens {
		if slices.Contains(ops, tkn){
			v2 := stack.Pop()
			v1 := stack.Pop()
			ans := apply(tkn, v1, v2)
			stack.Push(ans)
		} else {
			stack.Push(tkn)
		}

	}

	val,_ := strconv.Atoi(stack.Pop())	
	return val
}


func apply(op, v1, v2 string) string{

	nv1,_ := strconv.Atoi(v1)
	nv2,_ := strconv.Atoi(v2)
	ans := 0
	switch op {
		case "+":
			ans = nv1 + nv2
		case "-":
			ans = nv1 - nv2
		case "*":
			ans = nv1 * nv2
		case "/":
			ans = nv1 / nv2
	}
	return strconv.Itoa(ans)	
}


type Stack struct{
	n []string	
}


func(s *Stack) Push(v string){
	s.n = append(s.n, v)
}

func (s *Stack) Pop() string{
	v := s.n[len(s.n)-1]
	s.n = s.n[:len(s.n)-1]
	return v
}
