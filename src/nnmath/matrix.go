package nnmath

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	Rows int
	Cols int
	Data []float64
}

type Vector = Matrix

func NewMatrix(rows, cols int) Matrix {
	return Matrix{
		Rows: rows,
		Cols: cols,
		Data: make([]float64, rows*cols),
	}
}

func NewMatrixData(rows, cols int, data []float64) Matrix {
	if len(data) != rows*cols {
		panic(fmt.Sprintf("data length %d does not match matrix size %d,%d", len(data), rows, cols))
	}

	return Matrix{
		Rows: rows,
		Cols: cols,
		Data: data,
	}
}

func NewMatrixRandom(rows, cols int) Matrix {
	M := NewMatrix(rows, cols)

	for i := range rows*cols {
		M.Data[i] = rand.NormFloat64()
	}

	return M
}

func NewVector(size int) Vector {
	return NewMatrix(size, 1)
}

func NewVectorData(size int, data []float64) Vector {
	if len(data) != size {
		panic(fmt.Sprintf("data length %d does not match vector size %d", len(data), size))
	}

	return NewMatrixData(size, 1, data)
}

func NewVectorRandom(size int) Vector {
	V := NewVector(size)

	for i := range size {
		V.Data[i] = rand.NormFloat64()
	}

	return V
}

func (M Matrix) Get(row, col int) float64 {
	if row < 0 || row >= M.Rows || col < 0 || col >= M.Cols {
		panic(fmt.Sprintf("index out of range [%d][%d] with length %d,%d", row, col, M.Rows, M.Cols))
	}

	return M.Data[row*M.Cols+col]
}

func (M Matrix) Set(row, col int, value float64) {
	if row < 0 || row >= M.Rows || col < 0 || col >= M.Cols {
		panic(fmt.Sprintf("index out of range [%d][%d] with length %d,%d", row, col, M.Rows, M.Cols))
	}

	M.Data[row*M.Cols+col] = value
}

func Add(A, B Matrix) Matrix {
	if A.Rows != B.Rows || A.Cols != B.Cols {
		panic("matrix dimensions do not match")
	}

	C := NewMatrix(A.Rows, A.Cols)

	for i := range A.Rows {
		for j := range A.Cols {
			C.Set(i, j, A.Get(i, j)+B.Get(i, j))
		}
	}

	return C
}

func Mul(A, B Matrix) Matrix {
	if A.Cols != B.Rows {
		panic("matrix dimensions do not match")
	}

	C := NewMatrix(A.Rows, B.Cols)

	for i := range A.Rows {
		for j := range B.Cols {
			var sum float64
			for k := range A.Cols {
				sum += A.Get(i, k) * B.Get(k, j)
			}
			C.Set(i, j, sum)
		}
	}

	return C
}

func Apply(A Matrix, fn func(float64) float64) Matrix {
	R := NewMatrix(A.Rows, A.Cols)

	for i := range A.Rows {
		for j := range A.Cols {
			R.Set(i, j, fn(A.Get(i, j)))
		}
	}

	return R
}