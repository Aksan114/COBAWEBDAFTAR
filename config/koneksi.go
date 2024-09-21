package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var DB  *sql.DB

func GetConnection() {
	db, _ := sql.Open("mysql", "root:ojJmVodbLptqotLHyMyfCILelDATrjsL@tcp(mysql.railway.internal:3306)/pendaftaran?")

	if err := db.Ping(); err != nil {
		panic(err)
	}
	print("koneksi berhasil")
	DB = db
}