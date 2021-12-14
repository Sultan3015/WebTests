package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

//Method named save() that takes as its receiver p, a pointer to page.
//It takes no parameters, and returns a value of type error.
//It saves the Pages body to a text file. For simplicity. it uses the title as the file name
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) //0600 indicates filesshould be created with read-write permissions for the current user only
}

//Func loadPage constructs file name from the tit;e parameter and
//then reads the file ontents into a new variable body and returns pointer to a Page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) //io.Readfile func returns []byte and error
	if err != nil {
		return nil, err
	}
	//Callers of this function can now check if the second parameter after body is nil
	// then it successfully loaded a Page
	//But if not, it will be an error that can be handled by the caller
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)

	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	// t, _ := template.ParseFiles("edit.html") //reads content in the edit.html file
	// t.Execute(w, p)                          //writes the generated HTML to the http.responsewriter
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
