package model

import "github.com/alan-b-lima/handwritten_digits_example/src/nnmath"

type Sample struct {
	Label  nnmath.Vector
	Values nnmath.Vector
}

type Dataset = []Sample