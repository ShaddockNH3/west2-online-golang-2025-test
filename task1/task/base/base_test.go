package base

import (
	"testing"
)

func TestSolveLuoguP1001(t *testing.T) {
	a, b := 1, 1
	result := solve_luogu_p1001(a, b)
	t.Logf("输入: a=%d, b=%d, 输出: %d", a, b, result)
}

func TestSolveLuoguP1046(t *testing.T) {
	apple := [10]int{100, 200, 150, 140, 129, 134, 167, 198, 200, 111}
	height := 110
	result := solve_luogu_p1046(apple, height)
	t.Logf("可摘苹果数: %d", result)
}

func TestSolveLuoguP5737(t *testing.T) {
	x, y := 2000, 2023
	cnt, years := solve_luogu_p5737(x, y)
	yearsStr := formatYears(years)
	t.Logf("闰年数量: %d", cnt)
	t.Logf("闰年列表: %s", yearsStr)
}
