package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type Account struct {
	ID       int
	Name     string
	Email    string
	Password string
	Address  string
	Job      string
	Reason   string
}

var accounts = []Account{
	{ID: 0, Name: "Admin", Email: "admin@gmail.com", Password: "", Address: "Aceh", Job: "Admin IT", Reason: "Alasan Admin"},
	{ID: 1, Name: "Ahmad", Email: "ahmad@gmail.com", Password: "", Address: "Aceh", Job: "Agen rahasia", Reason: "Alasan Ahmad"},
	{ID: 2, Name: "Beni", Email: "beni@gmail.com", Password: "", Address: "Bandung", Job: "Backend", Reason: "Alasan Beni"},
	{ID: 3, Name: "Chanif", Email: "chanif@gmail.com", Password: "", Address: "Colomadu", Job: "Chef", Reason: "Alasan Chanif"},
	{ID: 4, Name: "Didik", Email: "didik@gmail.com", Password: "", Address: "Demak", Job: "Dokter", Reason: "Alasan Didik"},
	{ID: 5, Name: "Eko", Email: "eko@gmail.com", Password: "", Address: "Empat Lawang", Job: "Embedded", Reason: "Alasan Eko"},
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/user-validation", userValidation)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var filePath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, accounts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func emailVerif(s string, a []Account) (bool, Account) {
	valid := false
	acc := Account{}

	for i := 0; i < len(a); i++ {
		if s == a[i].Email {
			valid = true
			acc = a[i]
			break
		}
	}

	return valid, acc
}

func userValidation(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		valid, acc := emailVerif(email, accounts)

		var filepath string
		if valid {
			filepath = path.Join("views", "user.html")
		} else {
			filepath = path.Join("views", "no-user.html")
		}

		tmpl, err := template.ParseFiles(filepath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, acc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
