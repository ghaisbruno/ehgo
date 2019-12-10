package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a simple page")}
	p1.save()
	p2, _ := loadPage("TestPage.txt")
	fmt.Fprintf(w, "Hi there, %s!", string(p2.Body))
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testando 1,2,3....")
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	fmt.Println(p.Title)
	return ioutil.WriteFile(filename, p.Body, 0600) //(filename, content_to_save, permission_lvl)
}

func loadPage(title string) (*Page, error) {
	filename := title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func index(w http.ResponseWriter, r *http.Request) {
	// p, _ := loadPage("index.html")
	// fmt.Fprintf(w, string(p.Body))
	// ou
	http.ServeFile(w, r, "index.html")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := r.PostFormValue("username")
		hasher := sha1.New()
		pss := []byte(r.PostFormValue("password"))
		hasher.Write(pss)
		pass := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		if user == "a@a.com" && pass == "hvfkN_qlp_zhXR3cuerq6jd2Z7g=" {
			fmt.Printf("\nPasssou, user tá safe\n")
			http.Redirect(w, r, "/txt", 302)
		} else {
			http.Redirect(w, r, "/", 301)
		}
	} else if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
	}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/txt", handler)
	http.HandleFunc("/teste", test)
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
