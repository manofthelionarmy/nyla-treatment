package calendar

type Adapter interface {
	AddToCalendar() error
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
	AddToCalendar() error
}

func (s *adapter) AddToCalendar() error {
	return s.svc.AddToCalendar()
}
