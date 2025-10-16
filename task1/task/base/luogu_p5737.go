package base

import (
	"fmt"
)

func isrun(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

func solve_luogu_p5737(x, y int) (int, []int) {
	cnt := 0
	runnian := make([]int, 0)

	for i := x; i <= y; i++ {
		if isrun(i) {
			runnian = append(runnian, i)
			cnt += 1
		}
	}

	return cnt, runnian
}

func formatYears(years []int) string {
	result := ""
	for i, year := range years {
		result += fmt.Sprintf("%d", year)
		if i != len(years)-1 {
			result += " "
		}
	}
	return result
}
