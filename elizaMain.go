package main

import(
	"net/http"
	"html/template"
)

type myMsg struct{
	ER string
	UI string
}

func requestHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")

	
}

func eliza(w http.ResponseWriter, r *http.Request){
	elizaResponse := "Hi, how are you today?"

	userInput := r.FormValue("userInput")

	t,_ := template.ParseFiles("chatbot.html")

	t.Execute(w, &myMsg{ER:elizaResponse, UI:userInput})
}



func main(){
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/chatbot",eliza)

	http.ListenAndServe(":8080", nil)
}