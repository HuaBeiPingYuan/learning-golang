package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
}

func CreateTables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	)	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(db *sql.DB) (int64, error) {

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	query := `INSERT INTO users (username, password, created_at)
	VALUES (?,?,?) 
	ON DUPLICATE KEY UPDATE 
	   password = VALUES(password),
	   created_at = VALUES(created_at)`

	result, err := db.Exec(query, username, password, createdAt)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func QuerySingleUser(db *sql.DB) (string, error) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := "SELECT id, username, password, created_at FROM users WHERE id= ?"

	if err := db.QueryRow(query, 4).Scan(&id, &username, &password, &createdAt); err != nil {
		return "", err
	}
	return username, nil

}

func QueryAllUsers(db *sql.DB) ([]User, error) {

	query := "SELECT id, username, password, created_at FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User

	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil

}

func DeleteUser(db *sql.DB) (int64, error) {
	query := `DELETE FROM users WHERE id = ?`
	result, err := db.Exec(query, 1)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}
	return rowsDeleted, nil

}
