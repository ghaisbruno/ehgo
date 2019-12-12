package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//Conecta no DB MySQL e confere as credenciais
func hasUser(user string, pass string) (bool, int) {
	db, err := sql.Open("mysql", "read:[^#r);oXPW1]uH&hAU@tcp(localhost)/jmeter") //Vulnerabilidade
	if err != nil {
		panic(err.Error()) //Throw error
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error()) //Throw error
	}

	var id int
	sqlSt := "SELECT id FROM cred WHERE username='" + user + "' and password=SHA1('" + pass + "')"
	row := db.QueryRow(sqlSt)

	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User/Pass incorrect")
			return false, id
		}
		panic(err)
	}

	defer db.Close()
	return true, id
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		response, id := hasUser(r.PostFormValue("username"), r.PostFormValue("password"))
		fmt.Println("response:", response)
		if response == true {
			fmt.Println("User id:", id)
			http.Redirect(w, r, "/txt", 301)
			// // make a post into /txt
			// request, err := json.Marshal(map[string]string{
			// 	"response": strconv.FormatBool(response),
			// 	"id":       string(id),
			// })
			// resp, err := http.Post("/txt", "application/json", bytes.NewBuffer(request))
			// if err != nil {
			// 	log.Fatalln(err)
			// }
			// fmt.Println(resp)
		} else {
			http.Redirect(w, r, "/", 301)
		}
	} else if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
	}

}
