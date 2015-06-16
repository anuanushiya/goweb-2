package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
    "log"
    "time"
    "crypto/md5"
    "io"
    "strconv"
)

func sayhelloName(w http.ResponseWriter, r *http.Request){
    r.ParseForm()  // parse, default will not do it. 
    fmt.Println(r.Form)
    fmt.Println("Path: ", r.URL.Path)
    fmt.Println("scheme: ", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key: ", k)
	fmt.Println("val: ", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello chris")  // to client
}

func login(w http.ResponseWriter, r *http.Request){
    fmt.Println("Method: ", r.Method)
    if r.Method == "GET" {
	curtime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(curtime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("login.gtpl")
	t.Execute(w, token)
    } else {
	r.ParseForm()
        // do some ligin logically
	token := r.Form.Get("token")
	if token != "" {
	    // check token
	} else {
	    // report error of no token
	}
	fmt.Println("username lenght: ", len(r.Form.Get("username")))
	fmt.Println("username; ", template.HTMLEscapeString(r.Form.Get("username")))
	fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
	template.HTMLEscape(w, []byte(r.Form.Get("username")))
    }
}

func main(){
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
