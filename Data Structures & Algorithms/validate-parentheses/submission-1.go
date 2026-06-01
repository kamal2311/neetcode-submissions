func isValid(s string) bool {
    
    stack := []byte{}

    for i :=0; i < len(s); i++ {
        ch := s[i]
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, ch)
        } else{
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if ch == ')' && top != '(' {return false}
            if ch == '}' && top != '{' {return false}
            if ch == ']' && top != '[' {return false}
        }
    }

    return len(stack) == 0

}
