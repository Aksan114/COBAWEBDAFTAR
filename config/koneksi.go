package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var DB  *sql.DB

func GetConnection() {
	db, err := sql.Open("mysql", "root:qCNrNlUMyCCJFxufCseOPuMGXDCyRwGd@junction.proxy.rlwy.net:50004/railway")

	if err != nil {
		panic(err)
	}
	print("koneksi berhasil")
	DB = db
}	