package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
