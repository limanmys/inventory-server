package reporter

import (
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/client"
	"github.com/limanmys/inventory-server/internal/constants"
	"gorm.io/gorm"
)

func CreatePackageReport(job *entities.Job, db *gorm.DB, columns []string) {
	// Update job as in progress
	job.Update(entities.StatusInProgress, "Report in progress.")

	// Find data
	var data []map[string]interface{}
	if err := db.Find(&data).Error; err != nil {
		job.Update(entities.StatusError, "error when calculating data, err: "+err.Error())
	}

	// Build report body
	body := map[string]interface{}{
		"date":             time.Now().Format("01-02-2006 15:04:05"),
		"header":           "Inventory Server | Packages Report",
		"readable_columns": humanize(columns),
		"template_id":      "template.docx",
		"columns":          columns,
		"data":             data,
		"seperator":        ",",
	}

	// Check report engine url exists
	if os.Getenv("REPORT_ENGINE_URL") == "" {
		job.Update(entities.StatusError, "please set report engine url")
	}

	// Set API_URL & save path
	API_URL := os.Getenv("REPORT_ENGINE_URL") + "/" + string(job.FileType)
	PATH := constants.REPORT_SAVE_PATH + job.ID.String() + "." + string(job.FileType)

	// Send request to report engine
	_, err := client.NewRequest(body, PATH).Post(API_URL)
	if err != nil {
		job.Update(entities.StatusError, "error when creating report, err: "+err.Error())
	}

	job.UpdateAsDone(PATH)
}

func humanize(columns []string) []string {
	var humanized []string
	for _, column := range columns {
		humanized = append(humanized, capitalize(strings.ReplaceAll(column, "_", " ")))
	}

	return humanized
}

func capitalize(str string) string {
	var result string
	for idx, value := range strings.Split(str, " ") {
		runes := []rune(value)
		runes[0] = unicode.ToUpper(runes[0])
		if idx != len(strings.Split(str, " "))-1 {
			result += string(runes) + " "
		} else {
			result += string(runes)
		}
	}

	return result
}
