package dataset

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/alan-b-lima/handwritten_digits_example/src/model"
)

type Digit = model.Sample[byte]
type Dataset []Digit

func LoadDataset(path string) (Dataset, error) {

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

func _ProcessCsv(csv_input [][]string) (Dataset, error) {

	if len(csv_input[0]) != 28*28+1 {
		return nil, errors.ErrUnsupported
	}

	var dataset Dataset

	for i := 1; i < len(csv_input); i++ {
		row := csv_input[i]

		label, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, errors.New("data: bad csv")
		}

		pixels := make([]float64, 28*28)

		for j := 1; j < len(csv_input[j]); j++ {
			pixel, err := strconv.Atoi(row[j])
			if err != nil {
				return nil, errors.New("data: bad csv")
			}

			pixels[j-1] = float64(pixel) / 255.0
		}

		dataset = append(dataset, Digit{
			Label:  byte(label),
			Values: pixels,
		})
	}

	return dataset, nil
}
