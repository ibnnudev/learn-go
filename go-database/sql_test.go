package godatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlMutation := "insert into customer (id, name) values ('nada', 'nada salsabila hakim')"

	_, err := db.ExecContext(ctx, sqlMutation)
	if err != nil {
		t.Errorf("Error executing SQL: %v", err)
	}
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "select * from customer"
	rows, err := db.QueryContext(ctx, sqlQuery)

	if err != nil {
		t.Errorf("Error executing SQL: %v", err)
		return
	}

	defer rows.Close()

	var id string
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			t.Errorf("Error scanning row: %v", err)
		}

		t.Logf("Retrieved customer: id=%s, name=%s", id, name)

	}

	if err := rows.Err(); err != nil {
		t.Errorf("Error with rows: %v", err)
	}
}

func TestGetLastId(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlScript := "insert into comments (email, comment) values (?, ?)"
	result, err := db.ExecContext(ctx, sqlScript, "test@example.com", "This is a test comment")
	if err != nil {
		t.Fatalf("Failed to execute query: %v", err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("success insert id:", insertId)
}
