package main

import (
	"credit-calc/csv"
	"io"

	"github.com/rs/zerolog/log"
)

func main() {
	reader, closeFile, err := csv.CreateCSVReader("../ANZ.csv")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open csv file")
	}
	defer closeFile()

	transactions := make([]*csv.Transaction, 0)

	for row, err := reader.ReadRow(); err != io.EOF; row, err = reader.ReadRow() {
		if err != nil {
			log.Fatal().Err(err).Msg("failed to read row")
		}

		transactions = append(transactions, row)
		log.Info().Msgf("transaction: %+v", row)
	}
}
