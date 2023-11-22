package calendar

import (
	"nylatreatment/internal/model/medicine"
)

type Adapter interface {
	AddToCalendar(mr medicine.MedicineRecord) error
}

type adapter struct {
	svc Service
}

func NewAdapter(svc Service) Adapter {
	return &adapter{
		svc: svc,
	}
}

type Service interface {
	AddToCalendar(mr medicine.MedicineRecord) error
}

func (s *adapter) AddToCalendar(mr medicine.MedicineRecord) error {
	return s.svc.AddToCalendar(mr)
}
