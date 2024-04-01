package main

import (
    "fmt"
    "log"
    "net/http"
)
func main() {

	// about page handler 
	http.HandleFunc("/about",aboutPage)

	// login page handler 
	http.HandleFunc("/login",loginPage)

	http.HandleFunc("/",homePage)

    http.HandleFunc("/auth",authentication)

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
    })

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func aboutPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "about page")
}

func loginPage(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./static/login.html")
}
func homePage(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./static/index.html")
}
func authentication(w http.ResponseWriter, r *http.Request){
    // admin = "chimi"
    // password = "meopen"
    fmt.Println("working")
}