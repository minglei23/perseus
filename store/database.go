package store

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open(
		"mysql", MysqlUser+":"+MysqlPassword+"@tcp("+MysqlAddress+":"+MysqlPort+")/"+MysqlDB,
	)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
}

// LOGIN API

func GetUserIdByEmailAndPassword(email, password string) (id int64, err error) {
	query := "SELECT id FROM user_info WHERE email = ? AND password = ?"
	err = db.QueryRow(query, email, password).Scan(&id)
	if err == sql.ErrNoRows {
		// email or password is not correct
		return -1, nil
	}
	return id, err
}

// REGISTER API

func EmailExist(email string) (bool, error) {
	var id int
	query := "SELECT id FROM user_info WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, err
}

func InsertUser(password, email string) (id int64, err error) {
	query := "INSERT INTO user_info (password, email, activated, vip) values (?, ?, FALSE, FALSE)"
	result, err := db.Exec(query, password, email)
	if err != nil {
		return id, err
	}
	return result.LastInsertId()
}

// RESET API

func GetUserIdByEmail(email string) (id int64, err error) {
	query := "SELECT id FROM user_info WHERE email = ?"
	err = db.QueryRow(query, email).Scan(&id)
	if err == sql.ErrNoRows {
		// email is not correct
		return -1, nil
	}
	return id, err
}

func UpdateUserPassword(id, password string) error {
	query := "UPDATE user_info SET password = ? WHERE id = ?"
	_, err := db.Exec(query, password, id)
	return err
}

// VERIFY API

func ActivateUser(id string) error {
	query := "UPDATE user_info SET activated = TRUE WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}
