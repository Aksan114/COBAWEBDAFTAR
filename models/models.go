package models

import (
	"golang/config"
	"golang/entities"
)

func See() []entities.User {
	query, err := config.DB.Query("SELECT * FROM  tugas2")

	if err != nil {
		panic(err)
	}

	defer query.Close()

	var category []entities.User

	for query.Next() {
		var cat entities.User
		if err := query.Scan(&cat.ID, &cat.Nama_buku, &cat.Waktu_pengambilan, &cat.Nama_peminjam,
			); err != nil {
			panic(err)
		}

		category = append(category, cat)
	}

	return category
}

func Create(category entities.User) bool {
	res, err := config.DB.Exec("INSERT INTO tugas2(Nama_buku, Waktu_pengambilan, NAma_peminjam) VALUES (?,?,?)",
		category.Nama_buku,
		category.Waktu_pengambilan,
		category.Nama_peminjam,
	)

	if err != nil {
		panic(err)
	}

	lastinsertid, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastinsertid > 0
}
func Update(id int, category entities.User) bool  {
	res, err := config.DB.Exec("UPDATE tugas2 set Nama_buku=?, Nama_peminjam =? WHERE ID = ?",
	category.Nama_buku,
	category.Nama_peminjam,
	id,
	)

	if err != nil {
		panic(err)
	}

	result, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Detail(Id int)entities.User{
	row := config.DB.QueryRow("SELECT ID, Nama_buku, Nama_peminjam FROM tugas2 WHERE ID = ?", Id)

	var cat entities.User
	if err := row.Scan(&cat.ID, &cat.Nama_buku ,&cat.Nama_peminjam); err != nil {
		panic(err)
	}
	return cat
}

func Delete(id int) error  {
	_, err := config.DB.Exec("DELETE FROM tugas2 WHERE ID = ? " ,id)

	return err
}