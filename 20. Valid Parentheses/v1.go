func isValid(s string) bool {
    stackPtr := -1
    stack := make([]rune, len(s))

    for _, r := range s {
        if isOpen(r) {
            stackPtr++
            stack[stackPtr] = r
        } else {
            if stackPtr < 0 {
                return false
            }

            if  !isPair(stack[stackPtr], r) {
                return false
            }

            stackPtr = stackPtr - 1
        }
    }

    if stackPtr != -1 {
        return false
    }

    return true
}

func isOpen(r rune) bool {
    return r == '(' || r == '[' || r == '{'
}

func isPair(open rune, close rune) bool {
    if open == '(' && close == ')' {
        return true
    }

    if open == '[' && close == ']' {
        return true
    }

    if open == '{' && close == '}' {
        return true
    }

    return false
}