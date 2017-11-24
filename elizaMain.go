package main

import(
	"net/http"
	"html/template"
)

type myMsg struct{
	ER string
	UR string
}

func requestHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")

	
}

func eliza(w http.ResponseWriter, r *http.Request){
	elizaResponse := "Hi, how are you today?"

	userResponse := r.FormValue("userResponse")

	t,_ := template.ParseFiles("chatbot.html")

	t.Execute(w, &myMsg{ER:elizaResponse, UR:userResponse})
}



func main(){
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/chatbot",eliza)

	http.ListenAndServe(":8080", nil)
}