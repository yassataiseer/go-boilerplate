package main

import(
	"fmt"
	"html/template"
	"net/http"
)


func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("Starting server")
	var tpl = template.Must(template.ParseFiles("templates/index.html"))
	tpl.Execute(w, nil)
}


func main(){
	//Main function simply handles routing
	http.HandleFunc("/", login) 
	http.ListenAndServe(":8000",nil)
}