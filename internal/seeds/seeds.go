package seeds

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/constants"
	"github.com/limanmys/inventory-server/internal/database"
)

// Init, seeds alternative packages
func Init() {
	// Open alternatives.csv
	file, err := os.Open(constants.ALTERNATIVES_CSV_PATH)
	if err != nil {
		log.Println("error when reading package alternatives, " + err.Error())
		return
	}

	// Close file after function end
	defer file.Close()

	// Read all rows from csv
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println("error when reading records, " + err.Error())
		return
	}

	// Create records
	for idx, record := range records {
		// If row is wrong or title row
		if len(record) != 3 || idx == 0 {
			continue
		}
		// Create record on database
		if err := database.Connection().Where("name = ?", record[0]).FirstOrCreate(&entities.AlternativePackage{
			Name:        record[0],
			URL:         record[1],
			PackageName: record[2],
		}).Error; err != nil {
			log.Println("error when creating alternative package, " + err.Error())
		}
	}
}
