package spiralMatrix

import "fmt"

func SpiralOrder(matrix [][]int) []int {
	l := len(matrix)
	if l == 0 {
		return make([]int, 0)
	}
	if l == 1 {
		return matrix[0]
	}

	res := make([]int, 0)
	r1, r2 := 0, len(matrix)-1
	c1, c2 := 0, len(matrix[0])-1

	for r1 <= r2 && c1 <= c2 {
		for c := c1; c <= c2; c++ {
			res = append(res, matrix[r1][c])
		}
		for r := r1 + 1; r <= r2; r++ {
			res = append(res, matrix[r][c2])
		}

		if r1 < r2 && c1 < c2 {
			for c := c2 - 1; c > c1; c-- {
				res = append(res, matrix[r2][c])
			}
			for r := r2; r > r1; r-- {
				res = append(res, matrix[r][c1])
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return res
}

func Spiral() {
	m := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(SpiralOrder(m))

}
