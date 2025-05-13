package nnmath

import (
	"errors"
	"fmt"
)

type Matrix struct {
	Rows int
	Cols int
	Data []float64
}

type Vector = Matrix

func NewMatrix(rows, cols int) (*Matrix, error) {

	if rows <= 0 || cols <= 0 {
		return nil, errors.New("nnmath: rows and columns must be positive integers")
	}

	return &Matrix{
		Rows: rows,
		Cols: cols,
		Data: make([]float64, rows*cols),
	}, nil
}

func NewMatrixInit(rows, cols int, matrix []float64) (*Matrix, error) {

	if rows > 0 || cols > 0 {
		return nil, errors.New("nnmath: rows and columns must be positive integers")
	}

	size := rows * cols

	if size != len(matrix) {
		return nil, errors.New("nnmath: the input matrix dimensions do not match the input size")
	}

	return &Matrix{
		Rows: rows,
		Cols: cols,
		Data: make([]float64, rows*cols),
	}, nil
}

func NewIdMatrix(dim int) (*Matrix, error) {

	if dim <= 0 {
		return nil, errors.New("nnmath: the dimension must be a positive integer")
	}

	data := make([]float64, dim*dim)
	for i := range dim {
		data[i*dim+i] = 1
	}

	return &Matrix{
		Rows: dim,
		Cols: dim,
		Data: data,
	}, nil
}

func (M *Matrix) Get(row, col int) float64 {

	if row < 0 || M.Rows <= row || col < 0 || M.Cols <= col {
		panic(fmt.Sprintf("nnmath: index out of bounds [%d][%d] with dimensions %d, %d", row, col, M.Rows, M.Cols))
	}

	return M.Data[row*M.Cols+col]
}

func (M *Matrix) Set(row, col int, value float64) {

	if row < 0 || M.Rows <= row || col < 0 || M.Cols <= col {
		panic(fmt.Sprintf("nnmath: index out of bounds [%d][%d] with dimensions %d, %d", row, col, M.Rows, M.Cols))
	}

	M.Data[row*M.Cols+col] = value
}

func Add(A, B *Matrix) *Matrix {

	if A.Rows != B.Rows || A.Cols != B.Cols {
		panic("nnmath: dimensions do not match for addition")
	}

	R, err := NewMatrix(A.Rows, A.Cols)
	if err != nil {
		panic(err)
	}

	for i := range A.Rows {
		for j := range A.Cols {
			R.Set(i, j, A.Get(i, j)+B.Get(i, j))
		}
	}

	return R
}

func Sub(A, B *Matrix) *Matrix {

	if A.Rows != B.Rows || A.Cols != B.Cols {
		panic("nnmath: dimensions do not match for subtration")
	}

	R, err := NewMatrix(A.Rows, A.Cols)
	if err != nil {
		panic(err)
	}

	for i := range A.Rows {
		for j := range A.Cols {
			R.Set(i, j, A.Get(i, j)-B.Get(i, j))
		}
	}

	return R
}

func SMul(A *Matrix, s float64) *Matrix {

	R, err := NewMatrix(A.Rows, A.Cols)
	if err != nil {
		panic(err)
	}

	for i := range A.Rows {
		for j := range A.Cols {
			R.Set(i, j, s*A.Get(i, j))
		}
	}

	return R
}

func Mul(A, B *Matrix) *Matrix {

	if A.Cols != B.Rows {
		panic("nnmath: dimensions do not match for multiplication")
	}

	R, err := NewMatrix(A.Rows, B.Cols)
	if err != nil {
		panic(err)
	}

	for i := range A.Rows {
		for j := range A.Cols {
			var sum float64
			for k := range A.Cols {
				sum += A.Get(i, k) * B.Get(k, j)
			}
			R.Set(i, j, sum)
		}
	}

	return R
}
