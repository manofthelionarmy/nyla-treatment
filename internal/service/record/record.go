package record

import (
	"nylatreatment/internal/model/medicine"
	"time"
)

// Service is our record service
type Service interface {
	Record(name string, timeRecorded time.Time) error
}

// NewService initializes our record service
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

type service struct {
	repo Repository
}

// Repository states our api for retrieving data
type Repository interface {
	Record(record medicine.MedicineRecord) error
}

// Record records the time of the medicine treatment
func (s *service) Record(name string, timeRecorded time.Time) error {
	medicineRecord := medicine.MedicineRecord{
		Name:      name,
		TimeTaken: timeRecorded,
	}
	return s.repo.Record(medicineRecord)
}
