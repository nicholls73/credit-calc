package main

import (
	"credit-calc/configuration"
	"credit-calc/csv"
	"credit-calc/transactions"
	"io"

	"github.com/rs/zerolog/log"
)

func main() {
	config, err := configuration.LoadConfig("../config.yaml")
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

		log.Info().Msgf("%s is eligible: %t", transaction.Vendor, transaction.IsEligible(config))
		transactionsList = append(transactionsList, transaction)
	}
}
