package template

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

const (
	TryTimes   = 10
	ServerPort = "80"
	FilePath   = "../static/"
	FileName   = "index.html"
)

type User struct {
	Name string
	Age  int
	Sex  string
}


func showUser(w http.ResponseWriter, r *http.Request) {
	t := template.New(FileName)
	t, _ = t.ParseFiles(FilePath + FileName)

	user := User{Name: "zw7", Age: 11, Sex: "man"}
	err := t.Execute( w, user)
	if err!=nil {
		_, _ = w.Write([]byte("server error: " +err.Error()))
	}
}

func SetupHttpServer() {

	for n:=TryTimes; n>0; n-- {
		fmt.Println("Try", n, ":", "Try to Set Up Server, Port:", ServerPort)

		// 成功会阻塞
		err := http.ListenAndServe( ":"+ServerPort, nil)
		if err!=nil && n!=1 {
			time.Sleep(3*time.Second)
		}
	}

	panic("SetupHttpServer: fail to set up http server")
}

func Example() {

	http.HandleFunc("/", showUser)
	SetupHttpServer()
}