package parser

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/P-franczak/swift-api/database"
	"github.com/P-franczak/swift-api/models"
)

func ParseSwiftCSV(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    for _, record := range records {
        isHQ := strings.HasSuffix(record[0], "XXX")
        swiftCode := models.SwiftCode{
            SwiftCode:     record[0],
            BankName:      record[1],
            Address:       record[2],
            CountryISO2:   strings.ToUpper(record[3]),
            CountryName:   strings.ToUpper(record[4]),
            IsHeadquarter: isHQ,
        }
        database.DB.Create(swiftCode)
    }
    return nil
}
