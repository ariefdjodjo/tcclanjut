package main

import(
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"
	_"github.com/lib/pq"
	
)

type mahasiswa struct {
	Nim			int 	
	Nama 		string 	
	Jk			int 	
	Alamat 		string 	
	TempatLahir	string 	
	TglLahir 	string 	
	Agama 		int 	
}

type jurusan struct {
	Id_Jurusan 	int 	
	NamaJurusan	string	
}

func dbConn() (db *sql.DB)  {
	if len(os.Getenv("ROACH_HOST")) ==0 {
		os.Setenv("ROACH_HOST", "localhost")
	}

	dbUser := "root"
	host 	:= "localhost"
	dbName := "kampus"
	db, err := sql.Open("postgres", "postgresql://"+dbUser+"@"+host+":26257/"+dbName+"?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl, err = template.ParseGlob("view/*")

func index(w http.ResponseWriter, r *http.Request) {
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "golang CRUD",
	}

	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("bulma/css/"))))
}

func dataJurusan(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM jurusan ORDER BY id_jurusan ASC")
	if err != nil {
		panic(err.Error())
	}
	emp := jurusan{}
	res := []jurusan{}
	for selDB.Next() {
		var id_jurusan int
		var nama_jurusan string
		err = selDB.Scan(&id_jurusan, &nama_jurusan)
		if err != nil {
			panic(err.Error())
		}
		emp.Id_Jurusan = id_jurusan
		emp.NamaJurusan = nama_jurusan
		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "dataJurusan", res)
	defer db.Close()
}

func tambahJurusan(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"title": "golang CRUD",
	}

	err = tmpl.ExecuteTemplate(w, "tambahJurusan", data)
}

func tambah(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id_jurusan := r.FormValue("id_jurusan")
		nama_jurusan := r.FormValue("nama_jurusan")

		insForm, err := db.Prepare("INSERT INTO jurusan (id_jurusan, nama_jurusan) VALUES ($1,$2)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(id_jurusan, nama_jurusan)
		log.Println("INSERT: ID: " + id_jurusan + " | Jurusan: " + nama_jurusan)
	}
	defer db.Close()
	http.Redirect(w, r, "/dataJurusan", 301)
}

func editJurusan(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM jurusan WHERE id_jurusan=$1", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := jurusan{}
	for selDB.Next() {
		var id_jurusan int
		var nama_jurusan string
		err = selDB.Scan(&id_jurusan, &nama_jurusan)
		if err != nil {
			panic(err.Error())
		}
		emp.Id_Jurusan = id_jurusan
		emp.NamaJurusan = nama_jurusan
	}
	tmpl.ExecuteTemplate(w, "editJurusan", emp)
	defer db.Close()
}

func hapusJurusan(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM jurusan WHERE id_jurusan=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE BERHASIL")
	defer db.Close()
	http.Redirect(w, r, "/dataJurusan", 301)
}

func editJ(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id_jurusan := r.FormValue("id_jurusan")
		nama_jurusan := r.FormValue("nama_jurusan")

		insForm, err := db.Prepare("UPDATE jurusan SET nama_jurusan=$2 WHERE id_jurusan=$1")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(id_jurusan, nama_jurusan)
		log.Println("UPDATE: ID: " + id_jurusan + " | Jurusan: " + nama_jurusan)
	}
	defer db.Close()
	http.Redirect(w, r, "/dataJurusan", 301)
}

func dataMahasiswa(w http.ResponseWriter, r *http.Request) {
	err = tmpl.ExecuteTemplate(w, "dataMahasiswa", nil)
}

func tambahMahasiswa(w http.ResponseWriter, r *http.Request) {
	err = tmpl.ExecuteTemplate(w, "tambahMahasiswa", nil)
}

func tambahM(w http.ResponseWriter, r *http.Request) {
	err = tmpl.ExecuteTemplate(w, "tambahM", nil)
}


//membuat main
func main() {
	log.Println("Server started on: http://localhost:8010")
	http.HandleFunc("/", index)
	http.HandleFunc("/dataJurusan", dataJurusan)
	http.HandleFunc("/tambahJurusan", tambahJurusan)
	http.HandleFunc("/tambah", tambah)
	http.HandleFunc("/editJurusan", editJurusan)
	http.HandleFunc("/editJ", editJ)
	http.HandleFunc("/hapusJurusan", hapusJurusan)

	//handle untuk mahasiswa
	http.HandleFunc("/dataMahasiswa", dataMahasiswa)
	http.HandleFunc("/tambahMahasiswa", tambahMahasiswa)
	http.HandleFunc("/tambahM", tambahM)


	http.ListenAndServe(":8010", nil)
}