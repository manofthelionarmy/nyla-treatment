package list

import "nylatreatment/internal/model/medicine"

// Service is our record service
type Service interface {
	List() (medicine.MedicineList, error)
	GetLatestTreatment() (*medicine.MedicineRecord, error)
	GetMedicineLastTreatment(name string) (*medicine.MedicineRecord, error)
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
	List() (medicine.MedicineList, error)
	GetLatestTreatment() (*medicine.MedicineRecord, error)
	GetMedicineLastTreatment(name string) (*medicine.MedicineRecord, error)
}

func (s *service) List() (medicine.MedicineList, error) {
	return s.repo.List()
}

func (s *service) GetLatestTreatment() (*medicine.MedicineRecord, error) {
	return s.repo.GetLatestTreatment()
}

func (s *service) GetMedicineLastTreatment(name string) (*medicine.MedicineRecord, error) {
	return s.repo.GetMedicineLastTreatment(name)
}
