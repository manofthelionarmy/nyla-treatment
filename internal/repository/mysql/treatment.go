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
	timeRecorded := record.TimeTaken.Format(datetime)
	query := fmt.Sprintf(`insert into treatment_time(recorded_time_taken, medicine_id) values(?, (select id from medicine where name = '%s') )`, record.Name)
	_, err := d.conn.Exec(query, timeRecorded)
	if err != nil {
		return err
	}

	return nil
}

func (d *TreatmentDB) GetLatestTreatment() (*medicine.MedicineRecord, error) {
	query := `select m.name, tt.recorded_time_taken 
	from treatment_time tt inner join medicine m
	on tt.medicine_id = m.id
	order by tt.id desc limit 1;`
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
