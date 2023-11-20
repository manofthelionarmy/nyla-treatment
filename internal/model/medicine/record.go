package medicine

import "time"

type MedicineRecord struct {
	Name      string
	TimeTaken time.Time
}

type NextMedicineTime struct {
	Name     string
	NextTime time.Time
}
