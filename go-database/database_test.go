package godatabase

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
}
