package csv

import (
	"credit-calc/src/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func openFile(t *testing.T, filename string) *os.File {
	t.Helper()

	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		file.Close()
	})

	return file
}

func TestOpenCSVFile_FileExists(t *testing.T) {
	t.Parallel()

	content := []byte("one, two, three")
	filename := createTestFile(t, content)

	file, err := OpenCSVFile(filename)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestOpenCSVFile_FileDoesNotExist(t *testing.T) {
	t.Parallel()

	filename := "non-existent-file.csv"

	file, err := OpenCSVFile(filename)
	assert.ErrorIs(t, err, ErrFileNotFound)
	assert.NotNil(t, file)
}

func TestReadCSVRow_ValidRows(t *testing.T) {
	t.Parallel()

	content := []byte(`20/03/2025,500.00,VENDOR ONE
										21/03/2025,1000.00,VENDOR TWO
										21/03/2025,-100.00,VENDOR THREE`)

	filename := createTestFile(t, content)

	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}

	row, err := ReadCSVRow(file, 0)
	assert.NoError(t, err)
	assert.Equal(t, []string{"20/03/2025", "500.00", "VENDOR"}, row)
}

func TestReadCSVRow_InvalidRow(t *testing.T) {
	t.Parallel()

	content := []byte("20/03/2025,500.00")

	filename := createTestFile(t, content)

	file := openFile(t, filename)

	row, err := ReadCSVRow(file, 0)
	assert.ErrorIs(t, err, ErrInvalidRow)
	assert.Nil(t, row)
}

func TestReadCSVRow_Amount(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		content []byte
	}{
		{
			name:    "invalid amount",
			content: []byte("20/03/2025,INVALID,VENDOR"),
		},
		{
			name:    "invalid amount format",
			content: []byte("20/03/2025,5.00.00,VENDOR"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filename := createTestFile(t, tc.content)
			file := openFile(t, filename)

			row, err := ReadCSVRow(file, 0)
			assert.ErrorIs(t, err, ErrInvalidAmount)
			assert.Nil(t, row)
		})
	}
}

func TestReadCSVRow_Date(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		content []byte
	}{
		{
			name:    "invalid date",
			content: []byte("INVALID,500.00,VENDOR"),
		},
		{
			name:    "invalid date format",
			content: []byte("2025/03/02,500.00,VENDOR"),
		},
		{
			name:    "incomplete date",
			content: []byte("02/03,500.00,VENDOR"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filename := createTestFile(t, tc.content)
			file := openFile(t, filename)

			row, err := ReadCSVRow(file, 0)
			assert.ErrorIs(t, err, ErrInvalidDate)
			assert.Nil(t, row)
		})
	}
}

func TestReadCSVFile_ValidFile(t *testing.T) {
	t.Parallel()

	csvContent := []byte(`20/03/2025,500.00,VENDER ONE
										21/03/2025,1000.00,VENDOR TWO
										21/03/2025,-160.00,VENDOR THREE`)

	expected := []models.Transaction{
		{
			Date:        "20/03/2025",
			Amount:      500.00,
			Description: "VENDER ONE",
		},
		{
			Date:        "21/03/2025",
			Amount:      1000.00,
			Description: "VENDOR TWO",
		},
		{
			Date:        "21/03/2025",
			Amount:      -160.00,
			Description: "VENDOR THREE",
		},
	}

	filename := createTestFile(t, csvContent)

	content, err := ReadCSVFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, expected, content)
}

func TestReadCSVFile_MissingFile(t *testing.T) {
	t.Parallel()

	filename := "non-existent-file.csv"

	content, err := ReadCSVFile(filename)
	assert.ErrorIs(t, err, ErrFileNotFound)
	assert.Nil(t, content)
}

func TestReadCSVFile_EmptyFile(t *testing.T) {
	t.Parallel()

	filename := createTestFile(t, []byte{})

	content, err := ReadCSVFile(filename)
	assert.ErrorIs(t, err, ErrFileEmpty)
	assert.Nil(t, content)
}
