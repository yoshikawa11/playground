package main

func generateParenthesis(n int) []string {
	var res []string

	curr := make([]byte, 0, 2*n)
	var backtrack func(open, close int)
	backtrack = func(open, close int) {

		if open == n && close == n {
			res = append(res, string(curr))
			return
		}

		if open < n {
			curr = append(curr, '(')
			backtrack(open+1, close)
			curr = curr[:len(curr)-1]
		}

		if close < open {
			curr = append(curr, ')')
			backtrack(open, close+1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, 0)
	return res
}
