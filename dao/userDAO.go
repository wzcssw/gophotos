package dao

import "photos/models"

// GetAllUser GetAllUser
func GetAllUser(limit string) []models.User {
	users := make([]models.User, 0)
	rows, _ := Db.Query("SELECT id,name,phone,realname,age,password FROM users LIMIT ?", limit)
	for rows.Next() {
		user := &models.User{}
		rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Realname, &user.Age, &user.Password)
		users = append(users, *user)
	}
	return users
}

// GetUser GetUser
func GetUser(name string, password string) (models.User, bool) {
	user := models.User{}
	rows, _ := Db.Query("SELECT id,name,phone,realname,age,password FROM users WHERE name = ? AND PASSWORD = ?", name, password)
	hasResults := false
	for rows.Next() {
		hasResults = true
		rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Realname, &user.Age, &user.Password)
	}
	return user, hasResults
}

// // EditHospital EditHospital
// func EditHospital(hospital models.Hospital) bool {
// 	stmt, _ := Db.Prepare("UPDATE hospitals SET name=?,address=? WHERE id=?")
// 	res, _ := stmt.Exec(hospital.Name, hospital.Address, hospital.ID)
// 	affect, _ := res.RowsAffected()
// 	return (affect > 0)
// }

// // AddHospital AddHospital
// func AddHospital(hospital models.Hospital) int64 {
// 	stmt, _ := Db.Prepare("INSERT INTO hospitals SET name=?,address=? ")
// 	res, _ := stmt.Exec(hospital.Name, hospital.Address)
// 	id, _ := res.LastInsertId()
// 	return id
// }

// // DelHospital DelHospital
// func DelHospital(id string) bool {
// 	stmt, _ := Db.Prepare("DELETE FROM hospitals WHERE id=?")
// 	res, _ := stmt.Exec(id)
// 	affect, _ := res.RowsAffected()
// 	return (affect > 0)
// }
