package mysql

import (
	"fmt"
	"nylatreatment/internal/model/medicine"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TODO: fix tests
func TestEnterRecord(t *testing.T) {
	db, err := NewTreatmentDB()
	require.NoError(t, err)
	mr := medicine.MedicineRecord{
		Name:      "Prednisolone",
		TimeTaken: time.Now(),
	}
	// TODO: make db be able to do transactions
	err = db.Record(mr)
	require.NoError(t, err)
}

func TestLatestRecord(t *testing.T) {
	db, err := NewTreatmentDB()
	require.NoError(t, err)
	item, err := db.GetAllMedicineLatestTreatment()
	require.NoError(t, err)
	require.NotNil(t, item)
	fmt.Println(item)
}

func TestMedicineLatestTreatment(t *testing.T) {
	db, err := NewTreatmentDB()
	require.NoError(t, err)
	item, err := db.GetMedicineLastTreatment("Prednisolone")
	require.NoError(t, err)
	require.NotNil(t, item)
	fmt.Println(item)
}

func TestGetMedicineLastTreatment(t *testing.T) {
	db, err := NewTreatmentDB()
	require.NoError(t, err)
	item, err := db.GetMedicineNextTreatment("Prednisolone")
	require.NoError(t, err)
	require.NotNil(t, item)
	fmt.Println(item)
}
