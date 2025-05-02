package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
)

func TransferFunds(db *sql.DB, fromId string, toId string, amount int) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	ctx := context.Background()

	_, err = tx.ExecContext(ctx, "update customer set balance = balance - ? where id = ?", amount, fromId)
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.ExecContext(ctx, "update customer set balance = balance + ? where id = ?", amount, toId)
	if err != nil {
		log.Fatal(err)
	}

	return tx.Commit()
}

func TestDatabaseTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	err := TransferFunds(db, "siti", "budi", 800)
	if err != nil {
		t.Errorf("TransferFunds failed: %v", err)
		return
	}

	fmt.Println("Transfer successful")
}
