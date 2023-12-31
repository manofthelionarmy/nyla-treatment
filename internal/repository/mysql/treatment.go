package mysql

import (
	"database/sql"
	"fmt"
	"nylatreatment/internal/model/medicine"
	"time"
)

const datetime string = "2006-01-02 15:04:05"

type TreatmentDB struct {
	conn *sql.DB
}

func NewTreatmentDB() (*TreatmentDB, error) {
	conn, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/nyla_treatment")
	if err != nil {
		return nil, err
	}
	db := TreatmentDB{
		conn: conn,
	}
	return &db, nil
}

func (d *TreatmentDB) Record(record medicine.MedicineRecord) error {
	intervalQuery := "select id, time_period_hr from medicine where name = ?"
	row := d.conn.QueryRow(intervalQuery, record.Name)
	if row.Err() != nil {
		return row.Err()
	}
	var medicineID, interval int
	err := row.Scan(&medicineID, &interval)
	if err != nil {
		return err
	}

	// the date format is datetime in mysql
	timeRecorded := record.TimeTaken.Format(datetime)

	// at the application level, add 12 hours
	nextTime := record.TimeTaken.Add(time.Duration(interval) * time.Hour).Format(datetime)

	stmt := fmt.Sprintf(`insert into treatment_time(recorded_time_taken, medicine_id, next_treatment_time) values(?, ?, ?)`)
	_, err = d.conn.Exec(stmt, timeRecorded, medicineID, nextTime)
	if err != nil {
		return err
	}

	return nil
}

func (d *TreatmentDB) GetAllMedicineLatestTreatment() ([]medicine.MedicineRecord, error) {
	query := `select m.name, tt.recorded_time_taken from medicine m inner join treatment_time tt on tt.medicine_id = m.id where tt.recorded_time_taken = ( select max(treatment_time.recorded_time_taken) from medicine inner join treatment_time on treatment_time.medicine_id = medicine.id where medicine.name = m.name) group by m.name, tt.recorded_time_taken;`
	rows, err := d.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]medicine.MedicineRecord, 0)
	for rows.Next() {
		item := medicine.MedicineRecord{}
		var timeTaken string
		rows.Scan(&item.Name, &timeTaken)
		parsedTimeTaken, err := time.Parse(datetime, timeTaken)
		if err != nil {
			return nil, err
		}
		item.TimeTaken = parsedTimeTaken
		result = append(result, item)
	}
	return result, nil
}

func (d *TreatmentDB) GetMedicineLastTreatment(name string) (*medicine.MedicineRecord, error) {
	fmtStr := `select m.name, tt.recorded_time_taken
	from treatment_time tt inner join medicine m
	on tt.medicine_id = m.id
	where m.name = '%s'
	order by tt.id desc limit 1;`
	query := fmt.Sprintf(fmtStr, name)
	row := d.conn.QueryRow(query)
	if row.Err() != nil {
		return nil, row.Err()
	}
	item := medicine.MedicineRecord{}
	var timeTaken string
	row.Scan(&item.Name, &timeTaken)
	parsedTimeTaken, err := time.Parse(datetime, timeTaken)
	if err != nil {
		return nil, err
	}
	item.TimeTaken = parsedTimeTaken
	return &item, nil
}

func (d *TreatmentDB) GetMedicineNextTreatment(name string) (*medicine.MedicineRecord, error) {
	// why don't I add a new column that performs this calculation on each entry?
	// NOTE: this adds the interval to the date column on every row, I want to add it on one row
	fmtStr := `select m.name as name, tt.next_treatment_time as next_treatment_time
			from treatment_time tt
			inner join medicine m on m.id = tt.medicine_id
			where m.name = '%s' order by tt.recorded_time_taken desc limit 1`
	query := fmt.Sprintf(fmtStr, name)
	row := d.conn.QueryRow(query)
	if row.Err() != nil {
		return nil, row.Err()
	}
	item := medicine.MedicineRecord{}
	var timeTaken string
	row.Scan(&item.Name, &timeTaken)
	parsedTimeTaken, err := time.Parse(datetime, timeTaken)
	if err != nil {
		return nil, err
	}
	item.TimeTaken = parsedTimeTaken
	return &item, nil
}

func (db *TreatmentDB) GetAllMedicinesNextTreatment() ([]medicine.MedicineRecord, error) {
	query := `select m.name, next_treatment_time from medicine m inner join treatment_time tt on tt.medicine_id = m.id where tt.recorded_time_taken = ( select max(treatment_time.recorded_time_taken)
from medicine inner join treatment_time on treatment_time.medicine_id = medicine.id where medicine.name = m.name);`
	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records := make([]medicine.MedicineRecord, 0)
	for rows.Next() {
		mr := medicine.MedicineRecord{}
		var nextTime string
		rows.Scan(&mr.Name, &nextTime)
		// TODO: fix this when it's null, actually, we should set a rule in the db where an entry shouldn't be null
		parsedTimeTaken, err := time.Parse(datetime, nextTime)
		if err != nil {
			return nil, err
		}
		mr.TimeTaken = parsedTimeTaken
		records = append(records, mr)
	}
	return records, nil
}
