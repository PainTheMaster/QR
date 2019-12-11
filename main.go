package main

import (
	"fmt"

	"./qr"
)

//A is a matrix
var A [][]float64
var y []float64

func main() {

	A = [][]float64{{1.0, 2.0, 3.0},
		{4.0, 5.0, 7.0},
		{10, 8.0, 15},
	}

	y = []float64{14.0, 35.0, 71.0}

	qr.QR(A, y)

	x := qr.Solve(A, y)

	fmt.Println("A:")
	for i := range A {
		for j := range A[i] {
			fmt.Printf("%9.5f, ", A[i][j])
		}
		fmt.Println("")
	}

	fmt.Println("y:")
	for i := range y {
		fmt.Println(y[i])
	}

	fmt.Println("x:")
	for i := range x {
		fmt.Println(x[i])
	}
}

/*
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
*/
