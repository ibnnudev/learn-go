package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

type Customer struct {
	ID        string
	Name      string
	Email     sql.NullString
	Balance   sql.NullInt64
	Rating    sql.NullFloat64
	CreatedAt time.Time
	BirthDate sql.NullTime
	Married   sql.NullBool
}

func TestBatchInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	customersToInsert := []Customer{
		{
			ID:    "budi",
			Name:  "Budi Santoso",
			Email: sql.NullString{String: "budi.santoso@example.com", Valid: true},
		},
		{
			ID:      "siti",
			Name:    "Siti Aminah",
			Email:   sql.NullString{String: "siti.aminah@example.com", Valid: true},
			Balance: sql.NullInt64{Int64: 150000, Valid: true},
			Rating:  sql.NullFloat64{Float64: 4.8, Valid: true},
		},
		{
			ID:   "andi",
			Name: "Andi Wijaya",
		},
	}

	sqlInsert := "INSERT INTO customer (id, name, email, balance, rating, created_at, birth_date, married) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.PrepareContext(ctx, sqlInsert)
	if err != nil {
		t.Fatalf("Failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	for _, customer := range customersToInsert {
		_, err := stmt.ExecContext(ctx,
			customer.ID,
			customer.Name,
			customer.Email,
			customer.Balance,
			customer.Rating,
			time.Now(),
			customer.BirthDate,
			customer.Married,
		)
		if err != nil {
			t.Fatalf("Failed to execute statement: %v", err)
		}
	}
	t.Logf("Inserted %d customers successfully", len(customersToInsert))
}

func TestSelectAll(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlQuery := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer"
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		t.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.Balance,
			&customer.Rating,
			&customer.CreatedAt,
			&customer.BirthDate,
			&customer.Married,
		)
		if err != nil {
			t.Fatalf("Failed to scan row: %v", err)
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	// Print the results
	fmt.Println("Customers:")
	fmt.Println("------------------------------------------------")
	fmt.Println("ID\tName\tEmail\tBalance\tRating\tCreated At\tBirth Date\tMarried")
	fmt.Println("------------------------------------------------")

	for _, customer := range customers {
		fmt.Printf("%s\t%s\t%s\t%d\t%.2f\t%s\t%s\t%t\n",
			customer.ID,
			customer.Name,
			customer.Email.String,
			customer.Balance.Int64,
			customer.Rating.Float64,
			customer.CreatedAt.Format("2006-01-02 15:04:05"),
			customer.BirthDate.Time.Format("2006-01-02"),
			customer.Married.Bool,
		)
	}

	t.Logf("Retrieved %d customers successfully", len(customers))
}

func TestSQLInjection(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "user2'; #"
	username := "user2"
	password := "pass2"

	// prevent SQL injection
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		t.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login successfully with username:", username)
	} else {
		fmt.Println("Login failed")
	}

	t.Log("SQL Injection test completed")
}
