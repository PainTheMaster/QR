package qr

import (
	"math"
)

//QR performs QR decomposition
func QR(A [][]float64, y []float64) {

	size := len(A)

	for pivotCol := 0; pivotCol <= size-2; pivotCol++ {
		v, normVecSq, diag := Householder(A, pivotCol)
		A[pivotCol][pivotCol] = diag

		for i := pivotCol + 1; i <= size-1; i++ {
			A[i][pivotCol] = 0.0
		}

		for j := pivotCol + 1; j <= size-1; j++ {
			scale := 0.0

			for i := pivotCol; i <= size-1; i++ {
				iLocal := i - pivotCol
				scale += A[i][j] * v[iLocal]
			}

			scale = scale * 2.0 / normVecSq

			for i := pivotCol; i <= size-1; i++ {
				iLocal := i - pivotCol
				A[i][j] -= v[iLocal] * scale
			}
		}

		scale := 0.0

		for i := pivotCol; i <= size-1; i++ {
			iLocal := i - pivotCol
			scale += y[i] * v[iLocal]
		}

		scale = scale * 2.0 / normVecSq

		for i := pivotCol; i <= size-1; i++ {
			iLocal := i - pivotCol
			y[i] -= v[iLocal] * scale
		}
	}
}

//Householder is a function that retruns Householder matrix
func Householder(A [][]float64, col int) (v []float64, normVecSq, diag float64) {

	size := len(A)

	v = make([]float64, size-col)

	normCol := 0.0

	for i := col; i <= size-1; i++ {
		normCol += A[i][col] * A[i][col]
	}

	normCol = math.Sqrt(normCol)

	sign := -1.0

	if A[col][col] < 0 {
		sign = 1.0
	}

	diag = normCol * sign

	normVecSq = 0.0

	v[0] = A[col][col] - normCol*sign
	normVecSq += v[0] * v[0]

	for i := 1; i+col <= size-1; i++ {
		v[i] = A[col+i][col]
		normVecSq += v[i] * v[i]
	}

	return
}

//Solve solves QR decomposed equasion
func Solve(A [][]float64, y []float64) (x []float64) {

	size := len(A)

	x = make([]float64, size)

	x[size-1] = y[size-1] / A[size-1][size-1]

	for i := size - 2; i >= 0; i-- {
		ans := y[i]
		for k := size - 1; k >= i+1; k-- {
			ans -= A[i][k] * x[k]
		}
		x[i] = ans / A[i][i]
	}
	return
}
