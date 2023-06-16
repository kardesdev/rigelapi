package extra

import "github.com/vmkevv/rigelapi/app/models"

type ExtraRepository interface {
	GetYearsData() ([]models.YearData, error)
	GetGrades() ([]models.Grade, error)
	GetSubjects() ([]models.Subject, error)
	SaveAppErrors([]models.AppError) error
}
