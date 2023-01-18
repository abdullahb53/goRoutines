package main

import (
	"fmt"
	"sync"
)

func GenerateMatrix(row, col int) *[][]uint16 {
	a := make([][]uint16, row)
	for i := range a {
		a[i] = make([]uint16, col)
	}
	return &a

}

func main() {

	row := 5
	col := 5

	wg := new(sync.WaitGroup)
	wg.Add(row * col * col)

	First_Matrix := GenerateMatrix(row, col)
	Second_Matrix := GenerateMatrix(row, col)
	Result_Matrix := GenerateMatrix(row, col)

	for i := 0; i < len(*First_Matrix); i++ {
		for j := 0; j < len(*First_Matrix); j++ {

			(*First_Matrix)[i][j] = uint16(1)
			(*Second_Matrix)[i][j] = uint16(2)
		}
	}

	for i := 0; i < len(*First_Matrix); i++ {

		for j := 0; j < len(*Second_Matrix); j++ {

			(*Result_Matrix)[i][j] = uint16(0)

			for k := 0; k < len(*Second_Matrix); k++ {

				go func(wg_in_routine *sync.WaitGroup, a *[][]uint16, b *[][]uint16, c *[][]uint16, i int, j int, k int) {
					(*c)[i][j] += (*a)[i][k] * (*b)[k][j]
					wg_in_routine.Done()
				}(wg, First_Matrix, Second_Matrix, Result_Matrix, i, j, k)

			}

		}

	}
	wg.Wait()

	fmt.Println(Result_Matrix)

}
