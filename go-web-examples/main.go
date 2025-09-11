package main

import (
	"fmt"
	"log"
	"time"

	"HuaBeiPingYuan/go-web-examples/db"
	"HuaBeiPingYuan/go-web-examples/models"
)

// test all the data access functions directly
func main() {
	// ✅ Confirm DB connection
	fmt.Println("Ready to query", db.DB)

	// ✅ Create table
	if err := models.CreateTables(db.DB); err != nil {
		log.Fatal("Failed to create tables:", err)
	}
	fmt.Println("Table ready")
	// ✅ Insert user
	id, err := models.InsertUser(db.DB)
	if err != nil {
		log.Fatal("Failed to insert user:", err)
	}
	fmt.Println("Inserted user with ID:", id)

	// ✅ Query single user (id=10 for now)
	username, err := models.QuerySingleUser(db.DB)
	if err != nil {
		log.Fatal("Failed to query single user:", err)
	}
	fmt.Println("Queried single user username:", username)

 	// ✅ Delete user (id=1 for now)
	rowsDeleted, err := models.DeleteUser(db.DB)
	if err != nil {
		log.Fatal("Failed to delete user:", err)
	}
	fmt.Println("Deleted rows:", rowsDeleted)


	// ✅ Query all users
	users, err := models.QueryAllUsers(db.DB)
	if err != nil {
		log.Fatal("Failed to query all users:", err)
	}

	for _, u := range users {
		fmt.Printf("User: ID=%d Username=%s Password=%s CreatedAt=%s\n",
			u.ID, u.Username, u.Password, u.CreatedAt.Format(time.RFC3339))
	}
}
