package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var DB  *sql.DB

func GetConnection() {
	db, err := sql.Open("mysql", "mysql://root:qCNrNlUMyCCJFxufCseOPuMGXDCyRwGd@mysql.railway.internal:3306/railway")

	if err != nil {
		panic(err)
	}
	print("koneksi berhasil")
	DB = db
}	