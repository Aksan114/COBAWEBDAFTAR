package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var DB  *sql.DB

func GetConnection() {
	db, _ := sql.Open("mysql", "root:@tcp/tugas2?parseTime=true")

	if err := db.Ping(); err != nil {
		// print("gagal koneksi boss")
		// return
		panic(err)
	}
	print("koneksi berhasil")
	DB = db
}