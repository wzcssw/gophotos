package dao

import "photos/models"

// Db ddd
// var Db *sql.DB

// GetAllHospital GetAllHospital
func GetAllHospital(limit string) []models.Hospital {
	hospitals := make([]models.Hospital, 0)
	rows, _ := Db.Query("SELECT id,name,address FROM hospitals LIMIT ?", limit)
	for rows.Next() {
		hospital := &models.Hospital{}
		rows.Scan(&hospital.ID, &hospital.Name, &hospital.Address)
		hospitals = append(hospitals, *hospital)
	}
	return hospitals
}

// GetHospitalByID GetHospitalByID
func GetHospitalByID(id string) models.Hospital {
	hospital := models.Hospital{}
	rows, _ := Db.Query("SELECT id,name,address FROM hospitals WHERE id = ?", id)
	for rows.Next() {
		rows.Scan(&hospital.ID, &hospital.Name, &hospital.Address)
	}
	return hospital
}

// EditHospital EditHospital
func EditHospital(hospital models.Hospital) bool {
	stmt, _ := Db.Prepare("UPDATE hospitals SET name=?,address=? WHERE id=?")
	res, _ := stmt.Exec(hospital.Name, hospital.Address, hospital.ID)
	affect, _ := res.RowsAffected()
	return (affect > 0)
}

// AddHospital AddHospital
func AddHospital(hospital models.Hospital) int64 {
	stmt, _ := Db.Prepare("INSERT INTO hospitals SET name=?,address=? ")
	res, _ := stmt.Exec(hospital.Name, hospital.Address)
	id, _ := res.LastInsertId()
	return id
}

// DelHospital DelHospital
func DelHospital(id string) bool {
	stmt, _ := Db.Prepare("DELETE FROM hospitals WHERE id=?")
	res, _ := stmt.Exec(id)
	affect, _ := res.RowsAffected()
	return (affect > 0)
}
