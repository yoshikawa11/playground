package main

func isValidSudoku(board [][]byte) bool {
	const size = 9

	check := func(nums []byte) bool {
		seen := make(map[byte]struct{})
		for _, ch := range nums {
			if ch != '.' {
				if _, exists := seen[ch]; exists {
					return false
				}
				seen[ch] = struct{}{}
			}
		}
		return true
	}

	for i := 0; i < size; i++ {
		if !check(board[i]) {
			return false
		}
	}

	for i := 0; i < size; i++ {
		col := make([]byte, size)
		for j := 0; j < size; j++ {
			col[j] = board[j][i]
		}
		if !check(col) {
			return false
		}
	}

	for boxRow := 0; boxRow < size; boxRow += 3 {
		for boxCol := 0; boxCol < size; boxCol += 3 {
			box := make([]byte, 0, 9)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					box = append(box, board[boxRow+i][boxCol+j])
				}
			}
			if !check(box) {
				return false
			}
		}
	}

	return true
}
