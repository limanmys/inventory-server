package entities

import "github.com/limanmys/inventory-server/internal/database"

type JobType string

type ReportType string

type FileType string

var (
	JobTypeReport JobType = "report"

	FileTypePDF FileType = "pdf"
	FileTypeCSV FileType = "csv"

	ReportTypeAssetPackage ReportType = "asset_packages"
	ReportTypeAsset        ReportType = "asset"
	ReportTypePackage      ReportType = "package"
)

type Job struct {
	Base
	Status     Status     `json:"status"`
	JobType    JobType    `json:"job_type"`
	FileType   FileType   `json:"file_type"`
	ReportType ReportType `json:"report_type"`

	Message string `json:"message"`
	Path    string `json:"path"`
}

func (j *Job) Update(status Status, message string) {
	j.Status = status
	j.Message = message

	database.Connection().Model(j).Save(j)
}

func (j *Job) UpdateAsDone(path string) {
	j.Status = StatusDone
	j.Message = "Report created successfully."
	j.Path = path

	database.Connection().Model(j).Save(j)
}
