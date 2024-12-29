package service

import (
	repository "a21hc3NpZ25tZW50/repository/fileRepository"
	"encoding/csv"
	"fmt"
	"strings"
)

type FileService struct {
	Repo *repository.FileRepository
}

func (s *FileService) ProcessFile(fileContent string) (map[string][]string, error) {
	// fileByte, err := s.Repo.ReadFile(fileContent)
	// if err != nil {
	// 	return nil, err
	// }

	csvReader := csv.NewReader(strings.NewReader(string(fileContent)))
	csvData, err := csvReader.ReadAll()
	if len(csvData) == 0 {
		return nil, fmt.Errorf("kOSONG bANG")
	}

	if err != nil {
		return nil, err
	}

	// prepare the return value
	resultTable := make(map[string][]string)

	// create the header of csv
	headers := csvData[0]

	for _, header := range headers {
		resultTable[header] = make([]string, 0)
	}

	// create data below header
	for _, dataRows := range csvData[1:] {
		// loop row in rows
		for indexRow, dataRow := range dataRows {
			// add data in column n-header row indexRow
			resultTable[headers[indexRow]] = append(resultTable[headers[indexRow]], dataRow)
		}
	}

	return resultTable, nil
}
