package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"

)

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

		//connect to database
	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	Debugger(err)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE aesir")
	Debugger(err)

	log.Print("Connected to DB")
	
	Sdata := favthings{
		Thing1: r.FormValue("F1")
		Thing2: r.FormValue("F2")
		Thing3: r.FormValue("F3")

	}
	_ = Sdata

	log.Print(Sdata)

	data := InputForm {
		Name: r.FormValue("Uname")
		Color: r.FormValue("Color")
		Fthings: Sdata
	}

	_ = data

	log.Print(data)
	
	tmpl := template.Must(template.ParseFiles("static/templates/Page3.html"))
	tmpl.Execute(w, "null")
	log.Print("Running web-page")
	return

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
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs)))
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/Page2", Page2)
	http.HandleFunc("/Page3", Page3)
	log.Fatal(http.ListenAndServe(":8080", nil))

}


func main() {

	log.Print("Listening......")
	AppRoutes()

}