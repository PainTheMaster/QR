package main

import (
	"fmt"

	"./qr"
)

var A [][]float64
var X, Y []float64

func main() {

	/*	A = make([][]float64, 3)
		for i := 0; i <= 2; i++ {
			A[i] = make([]float64, 3)
		}
	*/
	A = [][]float64{{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
	}

	H, diag := qr.HouseHolder(A, 0)

	C := multi(H, A)

	fmt.Println("diag:", diag)
	fmt.Println("H:")
	for i := range H {
		for j := range H[i] {
			fmt.Print(H[i][j], ", ")
		}
		fmt.Println()
	}

	fmt.Println("HA")
	for i := range C {
		for j := range C[i] {
			fmt.Print(C[i][j], ", ")
		}
		fmt.Println()
	}

}

func multi(A [][]float64, B [][]float64) (C [][]float64) {

	size := len(A)

	C = make([][]float64, size)
	for i := 0; i <= size-1; i++ {
		C[i] = make([]float64, size)
	}

	for i := 0; i <= size-1; i++ {
		for j := 0; j <= size-1; j++ {
			for k := 0; k <= size-1; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return
}
