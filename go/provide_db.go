package swagger

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ProvideDBConnection() (*sql.DB, error) {
	return sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bank_mock")
}