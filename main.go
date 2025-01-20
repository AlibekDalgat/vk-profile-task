package main

import (
	"fmt"
	"math"
)

type coor struct {
	i, j int
}

func main() {
	var (
		n, m, stI, stJ, fnI, fnJ int
	)
	fmt.Scan(&n, &m)
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	fmt.Scan(&stI, &stJ, &fnI, &fnJ)
	sumPaths := make([][]int, n)
	parentCoors := make([][]coor, n)
	for i := 0; i < n; i++ {
		sumPaths[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sumPaths[i][j] = math.MaxInt
		}
		parentCoors[i] = make([]coor, m)
		for j := 0; j < m; j++ {
			parentCoors[i][j] = coor{-1, -1}
		}
	}
	que := make([]coor, 0)

	que = append(que, coor{stI, stJ})
	sumPaths[stI][stJ] = matrix[stI][stJ]
	for len(que) > 0 {
		curr := que[0]
		i, j := curr.i, curr.j
		if j < m-1 && matrix[i][j+1] != 0 {
			if sumPaths[i][j]+matrix[i][j+1] < sumPaths[i][j+1] {
				sumPaths[i][j+1] = sumPaths[i][j] + matrix[i][j+1]
				parentCoors[i][j+1] = coor{i, j}
				que = append(que, coor{i, j + 1})
			}
		}
		if i < n-1 && matrix[i+1][j] != 0 {
			if sumPaths[i][j]+matrix[i+1][j] < sumPaths[i+1][j] {
				sumPaths[i+1][j] = sumPaths[i][j] + matrix[i+1][j]
				parentCoors[i+1][j] = coor{i, j}
				que = append(que, coor{i + 1, j})
			}
		}
		if j > 0 && matrix[i][j-1] != 0 {
			if sumPaths[i][j]+matrix[i][j-1] < sumPaths[i][j-1] {
				sumPaths[i][j-1] = sumPaths[i][j] + matrix[i][j-1]
				parentCoors[i][j-1] = coor{i, j}
				que = append(que, coor{i, j - 1})
			}
		}
		if i > 0 && matrix[i-1][j] != 0 {
			if sumPaths[i][j]+matrix[i-1][j] < sumPaths[i-1][j] {
				sumPaths[i-1][j] = sumPaths[i][j] + matrix[i-1][j]
				parentCoors[i-1][j] = coor{i, j}
				que = append(que, coor{i - 1, j})
			}
		}
		que = que[1:]
	}

	res := make([]coor, 0)
	res = append(res, coor{fnI, fnJ})
	for res[len(res)-1].i != stI || res[len(res)-1].j != stJ {
		curr := res[len(res)-1]
		res = append(res, parentCoors[curr.i][curr.j])
	}
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Println(res[i].i, res[i].j)
	}
	fmt.Println(".")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if sumPaths[i][j] == math.MaxInt {
				fmt.Print(0, " ")
			} else {
				fmt.Print(sumPaths[i][j], " ")
			}
		}
		fmt.Println()
	}
}

/*
3 3
1 2 0
2 0 1
9 1 0
0 0 2 1

5 5
9 1 1 1 1
9 0 0 0 1
9 0 0 0 1
9 9 1 1 1
0 1 0 0 0
0 0 4 1

4 4
1 1 1 1
0 9 0 1
2 1 1 1
1 1 0 1
0 0 3 1

6 7
3 0 4 0 2 3 0
0 1 2 3 2 1 0
2 3 0 2 0 0 3
0 4 3 2 9 9 2
0 0 0 1 0 0 1
1 0 0 1 1 1 1
1 1 2 6
*/
