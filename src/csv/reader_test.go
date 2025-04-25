package csv_test

import (
	"credit-calc/csv"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const testFileContent = `20/03/2025,500.00,VENDOR ONE
21/03/2025,1000.00,VENDOR TWO
24/03/2025,-100.00,VENDOR THREE`

var testTransactions = []*csv.Transaction{
	{
		Date:   time.Date(2025, 3, 20, 0, 0, 0, 0, time.UTC),
		Amount: 500.00,
		Vendor: "VENDOR ONE",
	},
	{
		Date:   time.Date(2025, 3, 21, 0, 0, 0, 0, time.UTC),
		Amount: 1000.00,
		Vendor: "VENDOR TWO",
	},
	{
		Date:   time.Date(2025, 3, 24, 0, 0, 0, 0, time.UTC),
		Amount: -100.00,
		Vendor: "VENDOR THREE",
	},
}

func createTestFile(t *testing.T, content []byte) string {
	t.Helper()

	testFile, err := os.CreateTemp("", "test-*.csv")
	if err != nil {
		t.Fatal(err)
	}

	defer testFile.Close()

	if _, err := testFile.Write(content); err != nil {
		t.Fatal(err)
	}

	filename := testFile.Name()
	t.Cleanup(func() {
		os.Remove(filename)
	})

	return filename
}

func TestCreateCSVReader_ValidFile(t *testing.T) {
	t.Parallel()

	filename := createTestFile(t, []byte(testFileContent))

	reader, closeFile, err := csv.CreateCSVReader(filename)
	t.Cleanup(closeFile)

	require.NoError(t, err)
	require.NotNil(t, reader)
	require.NotNil(t, closeFile)
}

func TestCreateCSVReader_InvalidFile(t *testing.T) {
	t.Parallel()

	filename := "invalid-file.csv"

	reader, closeFile, err := csv.CreateCSVReader(filename)

	require.ErrorContains(t, err, csv.ErrFailedToOpenFileMsg)
	require.Nil(t, reader)
	require.Nil(t, closeFile)
}

func TestReadRow_ValidRows(t *testing.T) {
	t.Parallel()

	filename := createTestFile(t, []byte(testFileContent))

	reader, closeFile, _ := csv.CreateCSVReader(filename)
	t.Cleanup(closeFile)

	for _, expected := range testTransactions {
		transaction, err := reader.ReadRow()
		require.NoError(t, err)
		require.Equal(t, expected.Date, transaction.Date)
		require.Equal(t, expected.Amount, transaction.Amount)
		require.Equal(t, expected.Vendor, transaction.Vendor)
	}
}

func TestReadRow_InvalidDateParsing(t *testing.T) {
	t.Parallel()

	filename := createTestFile(t, []byte("dog,500.00,VENDOR ONE"))

	reader, closeFile, _ := csv.CreateCSVReader(filename)
	t.Cleanup(closeFile)

	transaction, err := reader.ReadRow()

	require.ErrorContains(t, err, csv.ErrFailedToParseDateMsg)
	require.Nil(t, transaction)
}

func TestReadRow_InvalidAmountParsing(t *testing.T) {
	t.Parallel()

	filename := createTestFile(t, []byte("20/03/2025,dog,VENDOR ONE"))

	reader, closeFile, _ := csv.CreateCSVReader(filename)
	t.Cleanup(closeFile)

	transaction, err := reader.ReadRow()

	require.ErrorContains(t, err, csv.ErrFailedToParseAmountMsg)
	require.Nil(t, transaction)
}
