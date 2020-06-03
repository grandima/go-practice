package spiralMatrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	cases := []struct {
		in  [][]int
		out []int
	}{
		{
			[][]int{},
			[]int{},
		},
		{
			[][]int{{5}},
			[]int{5},
		},
		{
			[][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
	}

	for _, data := range cases {
		if output := SpiralOrder(data.in); !reflect.DeepEqual(output, data.out) {
			t.Errorf("Expected %v to be %v", output, data.out)
		}

	}
}
