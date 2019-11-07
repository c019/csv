package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadAll - Retorna todo o conteudo do arquivo
// Retorno: [][]string, error
func ReadAll(file, delimiter, comment string) ([][]string, error) {
	// Open the file
	csvfile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("Couldn't open the csv file: %v", err)
	}

	// Parse the file
	reader := csv.NewReader(csvfile)
	reader.Comma = []rune(delimiter)[0]
	reader.Comment = []rune(comment)[0]

	return reader.ReadAll()
}
