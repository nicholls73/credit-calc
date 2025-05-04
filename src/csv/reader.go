package csv

import (
	"encoding/csv"
	"os"
)

type CSVReader struct {
	reader *csv.Reader
	file   *os.File
}

func CreateCSVReader(filePath string) (*CSVReader, func(), error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, ErrFailedToOpenFile(err)
	}

	return &CSVReader{
		reader: csv.NewReader(file),
		file:   file,
	}, func() { file.Close() }, nil
}

func (r *CSVReader) ReadRow() ([]string, error) {
	row, err := r.reader.Read()
	if err != nil {
		return nil, err
	}

	return row, nil
}
