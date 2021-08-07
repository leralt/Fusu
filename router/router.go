package router

import (
	"Fusu/ws"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)



func Start() {
	//fmt.Println("init...")
	http.HandleFunc("/static/",doExecute)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("username") == "long" && r.FormValue("password") == "123" {
			http.ServeFile(w,r,"views/room.html")
		}
	})
	http.HandleFunc("/room", ws.EchoMessage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Listen and serve error",err)
		return
	}
}

var realPath ="/Users/long/Projects/Fusu"
func doExecute( response http.ResponseWriter,request *http.Request) {
	requestUrl :=request.URL.String()
	fmt.Println(requestUrl[:])
	filePath := requestUrl[len("/static"):]
	fmt.Println("requestUrl =",filePath)
	file,err :=os.Open(realPath + requestUrl)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("file Close error",err)
		}
	}(file)
	if err != nil {
		log.Println("static resource:", err)
		response.WriteHeader(404)
	} else {
		bs,_ := ioutil.ReadAll(file)

		_, err := response.Write(bs)
		if err != nil {
			return
		}
	}
}
