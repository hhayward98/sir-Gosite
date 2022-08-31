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


func Options(){

	fmt.Println("Enter 1 to create GoFrame-L ")
	fmt.Println("Enter 2 for GoApp-skeleton ")
	fmt.Println("Enter 3 for ....")
	fmt.Println("Enter 4 for ....")


}


func GoFRAME_L(Num int){
	fmt.Println("Creating GOFRAME....")

	InitAppHead()

	InitDocker()

	InitTemplates(Num)

	InitCSS()

	fmt.Println("GoFRAME Complete!!!")

}


func InitAppHead() {
	fmt.Println("Creating webapp File....")

	Tem := "package main\n\nimport (\n\t'fmt'\n\t'log'\n\t'net/http'\n\t'html/template'\n)\n\nfunc Home(w http.ResponseWriter, r *http.Request) {\n\n\tfmt.Println('Home')\n\n\ttmpl := template.Must(template.ParseFiles('static/templates/index.html'))\n\ttmpl.Execute(w, 'Home')\n\treturn\n\n}\n\n\n//here\n\n\nfunc main() {\n\n\thttp.HandleFunc('/', Home)\n\n\tlog.Print('Listening....')\n\tlog.Fatal(http.ListenAndServe(':8080', nil))\n\n}\n\n"
	New := strings.ReplaceAll(Tem, "'", `"`,)
	f, err := os.Create("./app.go")
	_, err2 := f.WriteString(New)
	Debugger(err)
	Debugger(err2)

	defer f.Close()


	fmt.Println("Done")

}

func InitDocker(){
	fmt.Println("Creating Docker File....")

	Tem := "FROM golang:1.18\n\nRUN mkdir /GoWeb\n\nADD . /GoWeb\n\nWORKDIR /GoWeb\n\nCOPY go.* ./\n\nRUN go mod download && go mod verify\n\nRUN go build -o app .\n\nEXPOSE 8080\n\nCMD ['/GoWeb/app']"
	New := strings.ReplaceAll(Tem, "'", `"`,)
	f, err := os.Create("./Dockerfile")
	_, err2 := f.WriteString(New)
	Debugger(err)
	Debugger(err2)

	defer f.Close()
	fmt.Println("Done")
}

func InitCSS() {
	fmt.Println("Creating CSS....")
	if err := os.Mkdir("./static/css/", os.ModePerm); err != nil {
		log.Fatal(err)
	}


	f, err := os.Create("./static/css/main.css")
	_, err2 := f.WriteString("html {\n\t\n\t\n}\n\nhead {\n\t\n\t\n}\n\nbody {\n\t\n\t\n}\n\nfooter {\n\t\n\t\n}\n")
	Debugger(err)
	Debugger(err2)

	defer f.Close()

}



func InitTemplates(Num int) {
	fmt.Println("Creating Templates....")
	if err := os.MkdirAll("./static/templates/", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if Num == 1 {
		WriteHTML("index.html")
	}else if Num > 1 {
		WriteHTML("index.html")

		for i := 0; i < Num; i++ {
			strBuf := "Page"
			s1 := strconv.Itoa(i)
			strBuf += s1
			strBuf += ".html"
			fmt.Println(strBuf)
			WriteHTML(strBuf)
		}
	}

	AppendRoutsMain(Num)

	fmt.Println("Done")
}


func WriteHTML(Fname string) {

	Tem := "<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset='utf-8'>\n\t<meta name='viewport' content='width=device-width, initial-scale=1'>\n\t<meta http-equiv='X-UA-Compatible' content='ie=edge' />\n\t<title>Home</title>\n</head>\n<body>\n\t<h3>{{.}}</h3>\n\t<p>Information</p>\n</body>\n</html>"
	New := strings.ReplaceAll(Tem, "'", `"`,)

	f, err := os.Create("./static/templates/"+ Fname)
	_, err2 := f.WriteString(New)
	Debugger(err)
	Debugger(err2)

	defer f.Close()

}


func AppendRoutsMain(Num int) {

	var strbuffer string
	f, err := os.Open("app.go")
	Debugger(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		Line := scanner.Text()

		// adds functions for routs
		// allow users to choose methods for template serving
			// parsing Glob of html vs parsing individual files  

		if strings.Contains(Line, "//here") {

			for j := 0; j < Num; j++ {
				StJ := strconv.Itoa(j)
				Rfunc := "func Page"+StJ+"(w http.ResponseWriter, r *http.Request) {\n\n\ttmpl := template.Must(template.ParseFiles('static/templates/Page"+StJ+".html'))\n\ttmpl.Execute(w, 'Page"+StJ+"')\n\treturn\n\n}\n\n"
				Rinject := strings.ReplaceAll(Rfunc, "'", `"`,)
				//append to string Buffer code for injection
				strbuffer += Rinject 

			}

		}

		strbuffer += Line +"\n"

		if strings.Contains(Line, "main()"){
			for j := 0; j < Num; j++ {
				StJ := strconv.Itoa(j)
				strinj := "\thttp.HandleFunc('/Page"+StJ+"', Page"+StJ+")\n"
				Rinject := strings.ReplaceAll(strinj, "'", `"`,)
				//append to string Buffer code for injection
				strbuffer += Rinject 

			}
		}

	}

	h, err := os.Create("./app.go")
	_, err2 := h.WriteString(strbuffer)
	Debugger(err)
	Debugger(err2)

	defer h.Close()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}


func main() {
	fmt.Println("Starting up Hephaestus....")

	
	fmt.Println("Welcome!!!!")
	Options()

	var CHOICE int
	var TempNum int
	fmt.Scanln(&CHOICE)
	fmt.Println(CHOICE)

	if CHOICE == 1 {
		fmt.Println("Enter the number html Templates you want")
		fmt.Scanln(&TempNum)
		GoFRAME_L(TempNum)
	}else if CHOICE == 2 {
		fmt.Println("Enter the number html Templates you want")
		fmt.Scanln(&TempNum)
		


	}else {
		fmt.Println("Sorry, this program is still undergoing development!")
	}

}