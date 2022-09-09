package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"

)

func Debugger(err error) {
	if err != nil{
		log.Fatal(err)
	}
}








func Opt2Int() {
	fmt.Println("Starting up Opt2....")

	//Initalize appHead with approuts as function
	InitAppHead2()

	// Ask user the amount of HTML pages they want
	// when Init HTML pages, the routs will go under AppRoutes rather than main

	// Ask user if they want to add css styling
	// if yes ask for backround color, Text color, and secondary color

	// Ask user if they need any user input forms in the HTML
	// Ask user which page and form name
	// ask user for form size
	// inside of HTML page write div with formname as class and write form within div


	// Ask user if they want to add a database?
	// Ask which sql database to use
	// Ask user for name of database
	// Ask if user wants to use docker
	// If user wants to use Docker than create docker-compose.yml file 
	// else just append Database code under a page method that the user picks
	// Run Database Initalizer 
	// create initDB.go with code to add table and populate it so the web app will run 


	// Ask use if they want to add comments 
}



func InitAppHead2() {
	fmt.Println("Creating webapp File....")
	Tem := "package main\n\nimport (\n\t'fmt'\n\t'log'\n\t'net/http'\n\t'html/template'\n)\n\n\nvar tpl *template.Template\n\nfunc Home(w http.ResponseWriter, r *http.Request) {\n\n\tfmt.Println('Home')\n\n\ttpl.ExecuteTemplate(w, 'index.html', '')\n\n\treturn\n\n}\n\n\n//here\n\n\nfunc AppRouts() {\n\n\thttp.Handle('/static/', http.StripPrefix('/static/', http.FileServer(http.Dir('static'))))\n\thttp.HandleFunc('/', Home)\n\n\tlog.Fatal(http.ListenAndServe(':8080', nil))\n\n}\n\n\nfunc main() {\n\n\ttpl, _ = template.ParseGlob('./static/templates/*html')\n\n\tlog.Print('Listening....')\n\tAppRouts()\n\n}\n\n"
	New := strings.ReplaceAll(Tem, "'", `"`,)
	f, err := os.Create("./main.go")
	_, err2 := f.WriteString(New)
	Debugger(err)
	Debugger(err2)

	defer f.Close()


	fmt.Println("Done")

}


func AppendStyling(BK_color string, Text_color string, Second_color string) {
	fmt.Println("Adding Styling")
	// add styling to css file based on user input
}

// ask if user needs a database
func AppendDataBase(DBname string) {
	fmt.Println("Implementing Database....")
	// must create and setup database before trying to use in our program
	InitDatabase(DBname)
	// add code that connects to a database as part of a page function 
	// let the user choose what route/page will be modified

}

// check if the Database already exists before trying to create a newone
func InitDatabase(DBname string) {
	fmt.Println("Setting Up Database....")
	// create a database to use for web app 
	//full interface might be required

	fmt.Println("Database Creation complete!")

}

// ask if user needs any input forms
func AppendForms(FormName string, PageName string) {
	fmt.Println("Adding Forms....")
	// let the user specify form parameters (size, var_names, var_types, etc...) 
	// user should also choose what Page gets the form (letting users choose the location on the page for a bonus)
	// might need a string buffer to write working code 

}


func InitUnitTest() {
	fmt.Println("Creating Test....")
	//add a unit test file that will create test for everything created and test to see if the creation was successful
}