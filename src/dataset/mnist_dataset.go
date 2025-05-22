package dataset

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/alan-b-lima/handwritten_digits_example/src/model"
	"github.com/alan-b-lima/handwritten_digits_example/src/nnmath"
)

func LoadDataset(path string) (model.Dataset, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	csv_reader := csv.NewReader(file)
	csv_reader.FieldsPerRecord = 0
	csv_reader.Comma = ','

	result, err := csv_reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return _ProcessCsv(result)
}

func _ProcessCsv(csv_input [][]string) (model.Dataset, error) {

	if len(csv_input[0]) != 28*28+1 {
		return nil, errors.ErrUnsupported
	}

	var dataset model.Dataset

	for i := 1; i < len(csv_input); i++ {
		row := csv_input[i]

		label, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, errors.New("data: bad csv")
		}

		pixels := make([]float64, 28*28)

		for j := 1; j < len(row); j++ {
			pixel, err := strconv.Atoi(row[j])
			if err != nil {
				return nil, errors.New("data: bad csv")
			}

			pixels[j-1] = float64(pixel) / 255.0
		}

		dataset = append(dataset, model.Sample{
			Label:  LABELS[label],
			Values: nnmath.NewVectorData(28*28, pixels),
		})
	}

	return dataset, nil
}

var LABELS = []nnmath.Vector{
	{Rows: 10, Cols: 1, Data: []float64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	{Rows: 10, Cols: 1, Data: []float64{0, 1, 0, 0, 0, 0, 0, 0, 0, 0}},
	{Rows: 10, Cols: 1, Data: []float64{0, 0, 1, 0, 0, 0, 0, 0, 0, 0}},
	{Rows: 10, Cols: 1, Data: []float64{0, 0, 0, 1, 0, 0, 0, 0, 0, 0}},
	{Rows: 10, Cols: 1, Data: []float64{0, 0, 0, 0, 1, 0, 0, 0, 0, 0}},
	{Rows: 10, Cols: 1, Data: []float64{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}},
	{Rows: 10, Cols: 1, Data: []float64{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}},
	{Rows: 10, Cols: 1, Data: []float64{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}},
	{Rows: 10, Cols: 1, Data: []float64{0, 0, 0, 0, 0, 0, 0, 0, 1, 0}},
	{Rows: 10, Cols: 1, Data: []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
}