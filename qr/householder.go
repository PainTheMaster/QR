package qr

import (
	"fmt"
	"math"
)

func QR(A [][]float64, y []float64) {
	size := len(A)

	for pivotCol := 0; pivotCol <= size-1; pivotCol++ {

		vTemp := make([]float64, size-pivotCol)

		H, diag := Householder(A, pivotCol)
		j := pivotCol
		i := pivotCol
		A[i][j] = diag
		for i = j + 1; i <= size-1; i++ {
			A[i][j] = 0.0
		}

		for j = pivotCol + 1; j <= size-1; j++ {
			for i = pivotCol; i <= size-1; i++ {
				iLocal := i - pivotCol
				vTemp[iLocal] = 0.0
				for k := pivotCol; k <= size-1; k++ {
					kLocal := k - pivotCol
					vTemp[iLocal] += H[iLocal][kLocal] * A[k][j]
				}
			}

			for k := pivotCol; k <= size-1; k++ {
				kLocal := k - pivotCol
				A[k][j] = vTemp[kLocal]
			}
		}

		for j = pivotCol; j <= size-1; j++ {
			jLocal := j - pivotCol

			vTemp[jLocal] = 0.0
			for k := pivotCol; k <= size-1; k++ {
				kLocal := k - pivotCol
				vTemp[jLocal] += H[jLocal][kLocal] * y[k]
			}
		}
		for j = pivotCol; j <= size-1; j++ {
			jLocal := j - pivotCol
			y[j] = vTemp[jLocal]
		}

	}
}

//HouseHolder is a function that retruns Householder matrix
func Householder(A [][]float64, col int) (H [][]float64, diag float64) {
	size := len(A)
	fmt.Println("size:", size)

	v := make([]float64, size-col)

	H = make([][]float64, size-col)
	for i := 0; i <= size-col-1; i++ {
		H[i] = make([]float64, size-col)
	}

	normCol := 0.0
	for i := col; i <= size-1; i++ {
		normCol += A[i][col] * A[i][col]
	}
	normCol = math.Sqrt(normCol)

	sign := 1.0
	if A[col][col] < 0 {
		sign = -1.0
	}

	diag = normCol * sign

	normVec := 0.0
	v[0] = A[col][col] + normCol*sign
	normVec += v[0] * v[0]
	for i := 1; i+col <= size-1; i++ {
		v[i] = A[i][col]
		normVec += v[i] * v[i]
	}
	normVec = math.Sqrt(normVec)
	for i := 0; i+col <= size-1; i++ {
		v[i] /= normVec
	}

	for i := 0; i <= size-col-1; i++ {
		for j := 0; j < i; j++ {
			H[i][j] = -2.0 * v[i] * v[j]
		}

		H[i][i] = 1.0 - 2.0*v[i]*v[i]

		for j := i + 1; j <= size-col-1; j++ {
			H[i][j] = -2.0 * v[i] * v[j]
		}
	}

	return
}
