package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ivandhitya/sinau/database/util"
)

type student struct {
	id        int    `json:"id"`
	name      string `json:"name"`
	birthDate string `json:"birth_date"`
}

func main() {
	db, err := util.ConnectDB("ivan", "12345", "postgres")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// menggunakan parameterized query
	student := student{
		name:      "Marjiman",
		birthDate: "10-21-2000",
	}
	if err := createStudent(db, student); err != nil {
		log.Fatal("Failed to insert user:", err)
	}
	fmt.Println("User added successfully!")

	students, err := getAllStudents(db)
	if err != nil {
		log.Fatal("Failed to get users:", err)
	}

	// Tampilkan data user
	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %s\n", student.id, student.name, student.birthDate)
	}
}

func createStudent(db *sql.DB, student student) error {
	query := `INSERT INTO student (name, birth_date) VALUES ($1, $2)`
	_, err := db.Exec(query, student.name, student.birthDate)
	return err
}

func getAllStudents(db *sql.DB) ([]student, error) {
	query := "select id, name, birth_date from student"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []student
	for rows.Next() {
		var student student
		if err := rows.Scan(&student.id, &student.name, &student.birthDate); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	// Cek error dari iterasi
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
