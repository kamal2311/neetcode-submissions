func evalRPN(tokens []string) int {
	stack := NewStack(len(tokens))

	for _, tkn := range tokens {
		switch tkn {
		case "+", "-", "*", "/":
			v2 := stack.Pop()
			v1 := stack.Pop()
			stack.Push(apply(tkn, v1, v2))
		default:
			n, _ := strconv.Atoi(tkn)
			stack.Push(n)
		}
	}

	return stack.Pop()
}

func apply(op string, v1, v2 int) int {
	switch op {
	case "+":
		return v1 + v2
	case "-":
		return v1 - v2
	case "*":
		return v1 * v2
	case "/":
		return v1 / v2
	}
	return 0
}

type Stack struct {
	n   []int
	top int
}

func NewStack(cap int) *Stack {
	return &Stack{n: make([]int, cap)}
}

func (s *Stack) Push(v int) {
	s.n[s.top] = v
	s.top++
}

func (s *Stack) Pop() int {
	s.top--
	return s.n[s.top]
}
