package main

import (
	c "credit-calc/config"
	"credit-calc/csv"
	"credit-calc/summary"
	"credit-calc/transactions"
	"io"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	config, err := c.LoadConfig("../config.yaml")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	reader, closeFile, err := csv.CreateCSVReader("../ANZ.csv")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open csv file")
	}
	defer closeFile()

	transactionsList := make([]*transactions.Transaction, 0)

	for row, err := reader.ReadRow(); err != io.EOF; row, err = reader.ReadRow() {
		if err != nil {
			log.Fatal().Err(err).Msg("failed to read row")
		}

		transaction, err := transactions.FromCSVRow(row)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to parse transaction")
		}

		transactionsList = append(transactionsList, transaction)
	}

	sum, err := summary.GenerateSummary(transactionsList, config, time.Now())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate summary")
	}
	sum.Display()
}
