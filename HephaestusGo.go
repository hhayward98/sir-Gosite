package main

import (
	// "bufio"
	"fmt"
	"os"
	"log"
	"strconv"

)

func Debugger(err error) {
	if err != nil{
		log.Fatal(err)
	}
}


func Options(){

	fmt.Println("Enter 1 to create GoFrame ")
	fmt.Println("Enter 2 for .... ")
	fmt.Println("Enter 3 for ....")
	fmt.Println("Enter 4 for ....")


}


func GoFRAME(Num int){
	fmt.Println("Creating GOFRAME....")

	InitAppHead()

	InitDocker()

	InitTemplates(Num)

	InitCSS()

	fmt.Println("GoFRAME Complete!!!")

}


func InitAppHead() {
	fmt.Println("Creating webapp File....")

	f, err := os.Create("./app.go")
	_, err2 := f.WriteString("package main\n\nimport ()\n\nfunc Home(w http.ResponseWriter, r *http.Response) {\n\tfmt.Println(`Home`)\n}\n\nfunc main() {\n\tfmt.Println(`hello`)\n}\n\n")
	Debugger(err)
	Debugger(err2)

	defer f.Close()
	fmt.Println("Done")

}

func InitDocker(){

	fmt.Println("Creating Docker File....")
	f, err := os.Create("./Dockerfile")
	_, err2 := f.WriteString("FROM golang:1.18\n\nRUN mkdir /GoWeb\n\nADD . /GoWeb\n\nWORKDIR /GoWeb\n\nCOPY go.* ./\n\nRUN go mod download && go mod verify\n\nRUN go build -o app .\n\nEXPOSE 8080\n\nCMD ['/GoWeb/app']")
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
			strBuf := "page"
			s1 := strconv.Itoa(i)
			strBuf += s1
			strBuf += ".html"
			fmt.Println(strBuf)
			WriteHTML(strBuf)
		}
	}

	fmt.Println("Done")
}


func WriteHTML(Fname string) {

	f, err := os.Create("./static/templates/"+ Fname)
	_, err2 := f.WriteString("<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset='utf-8'>\n\t<meta name='viewport' content='width=device-width, initial-scale=1'>\n\t<meta http-equiv='X-UA-Compatible' content='ie=edge' />\n\t<title>Home</title>\n</head>\n<body>\n\t<h3>Title</h3>\n\t<p>Information</p>\n</body>\n</html>")
	Debugger(err)
	Debugger(err2)

	defer f.Close()

}


func AppendToMain() {
	// append Temp(each line) to the string buffer
	// can make conditions to search for content
	// if Temp contains main() append to StrBuff and stepinto condition
		// after appending current line from file, we can inject code to run in main()
		// include appropiate tabs('\t') and new lines('\n')

	var strbuffer string

	f, err := os.Open("app.go")
	Debugger(err)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		Line := scanner.Text()

		if strings.Contains(Temp, "main()"){
			fmt.Println(Temp)
		}

	}

	defer f.Close()

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
		GoFRAME(TempNum)
	} else {
		fmt.Println("Sorry, this program is still undergoing development!")
	}

}