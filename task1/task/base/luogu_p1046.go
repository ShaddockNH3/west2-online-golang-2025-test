package base

func solve_luogu_p1046(apple [10]int, height int) int {
	cnt := 0
	height += 30
	for _, appleHeight := range apple {
		if height >= appleHeight {
			cnt += 1
		}
	}
	return cnt
}
