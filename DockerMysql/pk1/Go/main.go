package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"html/template"

	_ "github.com/go-sql-driver/mysql"

)

var tpl *template.Template 

func Debugger(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")

	tpl.ExecuteTemplate(w, "index.html", "")
	return

}

func PinfoTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Info page")

	db, err := sql.Open("mysql", "test:toor@(db:3306)/sqldock")
	Debugger(err)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("USE sqldock")
	Debugger(err)
	log.Print("Connected to DB!!!")



	tpl.ExecuteTemplate(w, "Pinfo.html", "")
	
	return
}


func AppRoutes() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	http.HandleFunc("/", Home)
	http.HandleFunc("/Info", PinfoTest)

	log.Fatal(http.ListenAndServe(":8088", nil))

}

func main() {
	
	tpl, _ = template.ParseGlob("./static/templates/*html")


	log.Print("Listening.....")
	

	AppRoutes()

}