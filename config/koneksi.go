package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var DB  *sql.DB

func GetConnection() {
	db, _ := sql.Open("mysql", "mysql://root:rZGuemzqgBwXhPZotNrWkSMdUsfOCDRo@junction.proxy.rlwy.net:46866/pendaftaran")

	if err := db.Ping(); err != nil {
		panic(err)
	}
	print("koneksi berhasil")
	DB = db
}