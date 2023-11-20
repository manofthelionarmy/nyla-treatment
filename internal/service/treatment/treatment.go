package treatment

import (
	"nylatreatment/internal/model/medicine"
	"time"
)

// Service is our record service
type Service interface {
	Record(name string, timeRecorded time.Time) error
	GetLatestTreatment() (*medicine.MedicineRecord, error)
	GetMedicineLastTreatment(name string) (*medicine.MedicineRecord, error)
	GetMedicineNextTreatment(name string) (*medicine.MedicineRecord, error)
	GetAllMedicinesNextTreatment() ([]medicine.MedicineRecord, error)
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
	GetLatestTreatment() (*medicine.MedicineRecord, error)
	GetMedicineLastTreatment(name string) (*medicine.MedicineRecord, error)
	GetMedicineNextTreatment(name string) (*medicine.MedicineRecord, error)
	GetAllMedicinesNextTreatment() ([]medicine.MedicineRecord, error)
}

// Record records the time of the medicine treatment
func (s *service) Record(name string, timeRecorded time.Time) error {
	medicineRecord := medicine.MedicineRecord{
		Name:      name,
		TimeTaken: timeRecorded,
	}
	return s.repo.Record(medicineRecord)
}

func (s *service) GetLatestTreatment() (*medicine.MedicineRecord, error) {
	return s.repo.GetLatestTreatment()
}

func (s *service) GetMedicineLastTreatment(name string) (*medicine.MedicineRecord, error) {
	return s.repo.GetMedicineLastTreatment(name)
}

func (s *service) GetMedicineNextTreatment(name string) (*medicine.MedicineRecord, error) {
	return s.repo.GetMedicineNextTreatment(name)
}

func (s *service) GetAllMedicinesNextTreatment() ([]medicine.MedicineRecord, error) {
	return s.repo.GetAllMedicinesNextTreatment()
}
