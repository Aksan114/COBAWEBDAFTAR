package controller

import (
	"fmt"
	"golang/entities"
	"golang/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	cat := models.See()

	data := map[string]any{
		"cat": cat,
	}

	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func Information(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/informasi-prosedur.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func Informationjdwl(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/informasi-jadwal.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func Informationpnggmn(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/informasi-prosedur.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func Haldepan(w http.ResponseWriter, r *http.Request) {
	cat := models.See()

	data := map[string]any{
		"cat": cat,
	}

	temp, err := template.ParseFiles("views/peserta.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func Buat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/pendaftaran.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}
		err = temp.Execute(w, nil)
		if err != nil {
			http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		var cat entities.User
		cat.NAMA = r.FormValue("name")
		semesterStr := r.FormValue("namo")
		semester, err := strconv.Atoi(semesterStr)
		if err != nil {
			http.Error(w, "Invalid semester value", http.StatusBadRequest)
			return
		}
		cat.SEMESTER = semester
		cat.ASAL_KAMPUS = r.FormValue("nama")

		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Unable to upload image", http.StatusBadRequest)
			return
		}
		defer file.Close()

		uploadDir := "./static/fotopendaftar"
		filePath := filepath.Join(uploadDir, fileHeader.Filename)

		out, err := os.Create(filePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to create the file for writing. Check your write access privilege: %v", err), http.StatusInternalServerError)
			return
		}
		defer out.Close()

		if _, err := io.Copy(out, file); err != nil {
			http.Error(w, fmt.Sprintf("Error saving file: %v", err), http.StatusInternalServerError)
			return
		}

		cat.Gambar = filePath

		if ok := models.Create(cat); !ok {
			temp, err := template.ParseFiles("views/pendaftaran.html")
			if err != nil {
				http.Error(w, "Failed to load template", http.StatusInternalServerError)
				return
			}
			err = temp.Execute(w, nil)
			if err != nil {
				http.Error(w, "Failed to execute template", http.StatusInternalServerError)
			}
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func Selesai(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := models.Delete(id); err != nil {
		http.Error(w, "Failed to delete entry", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ShowImage(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	imagePath, err := models.GetImagePath(id)
	if err != nil {
		http.Error(w, "Failed to retrieve image path", http.StatusInternalServerError)
		return
	}

	fullPath := filepath.Join("static/fotopendaftar", filepath.Base(imagePath))

	http.ServeFile(w, r, fullPath)
}



func Serverhandlestatic() {
	anu := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", anu))
}