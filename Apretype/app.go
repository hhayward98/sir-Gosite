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

type InputForm struct {
	name string
	color string
	Favthing favthings
}

type favthings struct {
	T1 string
	T2 string
	T3 string
}



func Debugger(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running Home Page\n")
	
	tmpl := template.Must(template.ParseFiles("static/templates/index.html"))
	tmpl.Execute(w, "null")
	return

}


func Page3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running Page3\n")
	tmpl := template.Must(template.ParseFiles("static/templates/page3.html")
	
	obj1 := favthings{r.FormValue("F1"), r.FormValue("F2"), r.FormValue("F3")}
	fmt.Println(obj1)
	obj2 := InputForm{r.FormValue("Uname"), r.FormValue("Color"), obj1}
	fmt.Println(obj2)

	tmpl.Execute(w, obj2)
	log.Print("Running web-page")


}

func Page4(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running Page4\n")

	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	Debugger(err)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE aesir")
	Debugger(err)

	log.Print("Connected to DB")

	tpl.ExecuteTemplate(w, "Page4.html", "Connected To Database")

}



func Page2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running Page2\n")

	tmpl := template.Must(template.ParseFiles("static/templates/Page2.html"))
	tmpl.Execute(w, "null")
	return

}


func AppRoutes() {

	//testing for docker container
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/Page2", Page2)
	http.HandleFunc("/Page3", Page3)
	http.HandleFunc("/Page4", Page4)
	log.Fatal(http.ListenAndServe(":8080", nil))

}


func main() {
	tpl, _ = template.ParseGlob("static/templates/*html")

	log.Print("Listening......")
	AppRoutes()

}