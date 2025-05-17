package model

type Sample[T comparable] struct {
	Label  T
	Values []float64
}
