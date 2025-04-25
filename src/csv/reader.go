package csv

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Transaction struct {
	Date   time.Time
	Amount float64
	Vendor string
}

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

func (r *CSVReader) ReadRow() (*Transaction, error) {
	row, err := r.reader.Read()
	if err != nil {
		return nil, err
	}

	date, err := time.Parse("02/01/2006", row[0])
	if err != nil {
		return nil, ErrFailedToParseDate(err)
	}

	amount, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		return nil, ErrFailedToParseAmount(err)
	}

	return &Transaction{
		Date:   date,
		Amount: amount,
		Vendor: row[2],
	}, nil
}
