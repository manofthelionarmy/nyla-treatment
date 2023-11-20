package mysql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMedicine(t *testing.T) {
	db, err := NewMedicineDB()
	require.NoError(t, err)
	medicineList, err := db.List()
	require.NoError(t, err)
	require.NotEmpty(t, medicineList)
	require.NotZero(t, len(medicineList))
}
