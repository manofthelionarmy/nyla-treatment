package mysql

import (
	"database/sql"
	"nylatreatment/internal/model/medicine"
)

type MedicineDB struct {
	conn *sql.DB
}

func NewMedicineDB() (*MedicineDB, error) {
	conn, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/nyla_treatment")
	if err != nil {
		return nil, err
	}
	db := MedicineDB{
		conn: conn,
	}
	return &db, nil
}

func (d *MedicineDB) List() (medicine.MedicineList, error) {
	query := `select name, type, time_period_hr from medicine;`
	rows, err := d.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	medicineList := make([]medicine.MedicineItem, 0)
	for rows.Next() {
		item := medicine.MedicineItem{}
		if err := rows.Scan(&item.Name, &item.Type, &item.TimePeriod); err != nil {
			return nil, err
		}
		medicineList = append(medicineList, item)
	}
	return medicineList, nil
}
