func removeInvalidParentheses(s string) []string {
    // count misplaced left and right parentheses
    left, right := 0, 0
    for _, ch := range s {
        if ch == '(' {
            left++
        } else if ch == ')' {
            if left == 0 {
                right++
            }
            if left > 0 {
                left--
            }
        }
    }
    exprSet := make(map[string]int)
    dfs(s, "", 0, 0, 0, left, right, exprSet)
    res := make([]string, 0)
    for key, _ := range exprSet {
        res = append(res, key)
    }
    return res
}

func dfs(s, expr string, pos, left, right, leftRem, rightRem int, exprSet map[string]int) {
    if pos == len(s) {
        if leftRem == 0 && rightRem == 0 {
            exprSet[expr] = 1
        }
        return
    }
    ch := s[pos]
    // the discard case (pruning condition)
    if ch == '(' && leftRem > 0 {
        dfs(s, expr, pos+1, left, right, leftRem-1, rightRem, exprSet)
    } else if ch == ')' && rightRem > 0 {
        dfs(s, expr, pos+1, left, right, leftRem, rightRem-1, exprSet)
    }
    // simply recurse one step further
    if ch == '(' { // consider an opening bracket
        dfs(s, expr+string(ch), pos+1, left+1, right, leftRem, rightRem, exprSet)
    } else if ch == ')' && left > right { // consider an closing bracket
        dfs(s, expr+string(ch), pos+1, left, right+1, leftRem, rightRem, exprSet)
    } else if ch != ')' { // considering letters other than parentheses
        dfs(s, expr+string(ch), pos+1, left, right, leftRem, rightRem, exprSet)
    }
}
