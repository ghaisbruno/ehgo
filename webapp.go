package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page : Creates a simple WebPage
type Page struct {
	Title string
	Body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "POST" {

	// }
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

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/teste", test)
	http.HandleFunc("/txt", handler)
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":80", nil))
}
