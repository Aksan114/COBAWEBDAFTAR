package models

import (
	"golang/config"
	"golang/entities"
	"log"
)

func See() []entities.User {
	query, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	var users []entities.User

	for query.Next() {
		var user entities.User
		err := query.Scan(&user.ID, &user.NAMA, &user.SEMESTER, &user.ASAL_KAMPUS, &user.Gambar)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users
}

func Create(user entities.User) bool {
	res, err := config.DB.Exec("INSERT INTO users (NAMA, SEMESTER, ASAL_KAMPUS, Gambar) VALUES (?, ?, ?, ?)",
		user.NAMA,
		user.SEMESTER,
		user.ASAL_KAMPUS,
		user.Gambar,
	)

	if err != nil {
		log.Fatal(err)
		return false
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return false
	}

	return lastInsertID > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE ID = ?", id)
	return err
}

func GetImagePath(id int) (string, error) {
	row := config.DB.QueryRow("SELECT Gambar FROM users WHERE ID = ?", id)

	var imagePath string
	err := row.Scan(&imagePath)
	if err != nil {
		return "", err
	}

	return imagePath, nil
}