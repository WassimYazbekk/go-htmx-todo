package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct{
  Content string
  Done bool
  }

func mainHandler(w http.ResponseWriter, _ *http.Request){
  tmpl := template.Must(template.ParseFiles("views/index.html"))
  todos := map[string][]Todo{
    "todos" : {
      {Content: "Learn Go", Done: true},
      {Content: "Learn HTMX", Done: true},
      {Content: "Find a job", Done: false},
    },
  }
    tmpl.Execute(w, todos)
  }

func addTodoHandler(w http.ResponseWriter, r *http.Request){
  tmpl := template.Must(template.ParseFiles("views/index.html"))
  content := r.PostFormValue("todo")
  tmpl.ExecuteTemplate(w,"todo-item" ,Todo{Content: content, Done: false})
}

func deleteTodoHandler(w http.ResponseWriter, _ *http.Request){
  tmpl, _ := template.New("t").Parse("")
  tmpl.Execute(w,nil)
}

func doneTodoHandler(w http.ResponseWriter, _ *http.Request){
  tmpl := template.Must(template.ParseFiles("views/index.html"))
  tmpl.ExecuteTemplate(w,"undo-btn" ,nil)
}

func undoTodoHandler(w http.ResponseWriter, _ *http.Request){
  tmpl := template.Must(template.ParseFiles("views/index.html"))
  tmpl.ExecuteTemplate(w,"done-btn" ,nil)
}

func main(){
  http.HandleFunc("/", mainHandler)
  http.HandleFunc("/add-todo/", addTodoHandler)
  http.HandleFunc("/delete-todo", deleteTodoHandler)
  http.HandleFunc("/done", doneTodoHandler)
  http.HandleFunc("/undo", undoTodoHandler)

  fmt.Println("Server running at localhost:6969")
  log.Fatal(http.ListenAndServe(":6969", nil))
}
