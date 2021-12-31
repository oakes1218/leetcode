package main

func main() {
	//Array  Rotate Image
	var input = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(input)
}

/*
   1, 2, 3         7, 4, 1
   4, 5, 6    =>   8, 5, 2
   7, 8, 9         9, 6, 3
*/

func rotate(matrix [][]int) {
	l := len(matrix)
	newMatrix := make([][]int, l)
	for i := 0; i < l; i++ {
		newMatrix[i] = make([]int, l)
	}

	for k, v := range matrix {
		vl := l - 1 - k
		for k1, v1 := range v {
			newMatrix[k1][vl] = v1
		}
	}

	copy(matrix, newMatrix)
}
