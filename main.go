package main

import (
	"fmt"
	"os"
)

type coor struct {
	i, j int
}

func main() {
	var (
		n, m, stI, stJ, fnI, fnJ int
	)
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input: ", err)
		os.Exit(1)
	}
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_, err = fmt.Scan(&matrix[i][j])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Invalid input: ", err)
				os.Exit(1)
			}
		}
	}
	_, err = fmt.Scan(&stI, &stJ, &fnI, &fnJ)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input: ", err)
		os.Exit(1)
	}

	sumPaths := make([][]int, n)
	parentCoors := make([][]coor, n)
	for i := 0; i < n; i++ {
		sumPaths[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if matrix[i][j] != 0 {
				sumPaths[i][j] = -1
			}
		}
		parentCoors[i] = make([]coor, m)
		for j := 0; j < m; j++ {
			parentCoors[i][j] = coor{-1, -1}
		}
	}
	sumPaths[stI][stJ] = matrix[stI][stJ]

	que := make([]coor, 0)
	que = append(que, coor{stI, stJ})
	for len(que) > 0 {
		curr := que[0]
		i, j := curr.i, curr.j
		directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, direct := range directions {
			nextI, nextJ := i+direct[0], j+direct[1]
			if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m && matrix[nextI][nextJ] != 0 {
				if sumPaths[nextI][nextJ] == -1 || sumPaths[i][j]+matrix[nextI][nextJ] < sumPaths[nextI][nextJ] {
					sumPaths[nextI][nextJ] = sumPaths[i][j] + matrix[nextI][nextJ]
					parentCoors[nextI][nextJ] = coor{i, j}
					que = append(que, coor{nextI, nextJ})
				}
			}
		}
		que = que[1:]
	}

	if sumPaths[fnI][fnJ] == -1 {
		fmt.Fprintln(os.Stderr, "No path found for finish")
		os.Exit(1)
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
			if sumPaths[i][j] == -1 {
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

3 3
1 1 1
1 1 0
1 0 2
0 0 2 2

*/
