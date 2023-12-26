package jobs

import (
	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/database"
	"gorm.io/gorm/clause"
)

// Create new job
func NewJob(file_type entities.FileType) (*entities.Job, error) {
	// Set job object
	job := entities.Job{
		ReportType: entities.ReportTypePackage,
		Status:     entities.StatusPending,
		JobType:    entities.JobTypeReport,
		FileType:   file_type,
		Message:    "Report in progress.",
		Path:       "",
	}

	// Create job on database
	if err := database.Connection().Clauses(clause.Returning{}).Create(&job).Error; err != nil {
		return nil, err
	}

	return &job, nil
}
