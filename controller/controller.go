package controller

import (
	"net/http"
	"golang/models"
	"strconv"
	"golang/entities"
	"text/template"
	"time"
)


func Haldepan(w http.ResponseWriter, r *http.Request) {
	cat := models.See()

	data := map[string]any{
		"cat": cat,
	}

	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}



func Buat(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/created.html")


		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var cat entities.User

		cat.Nama_buku = r.FormValue("name")
		cat.Waktu_pengambilan = time.Now()
		cat.Nama_peminjam = r.FormValue("nama")

		if ok := models.Create(cat); !ok {
			temp, err := template.ParseFiles("views/created.html")

			if err != nil {
				panic(err)
			}

			temp.Execute(w, nil)
		}
		http.Redirect(w, r ,"/", http.StatusSeeOther)
	}
}


func Update(w http.ResponseWriter, r *http.Request)  {
	if  r.Method == "GET" {
		temp, err :=template.ParseFiles(("views/edit.html"))

		if err != nil {
			panic(err)
		}	

		idkeys := r.URL.Query().Get("id")

		id , err := strconv.Atoi(idkeys)

		if err != nil {
			panic(err)
		}

		cat := models.Detail(id)
		data := map[string]any{
			"zz" : cat,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var category entities.User

		idname := r.FormValue("id")
		res, err := strconv.Atoi(idname)

		if err != nil {
			panic(err)
		}
		category.Nama_buku = r.FormValue("name")
		category.Nama_peminjam = r.FormValue("nama")

		if err := models.Update(res, category); !err{
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return

		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}


func Selesai(w http.ResponseWriter, r *http.Request)  {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	
	if err != nil {
		panic(err)
	}

	if err := models.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}


func Serverhandlestatic() {
    anu := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", anu))
}