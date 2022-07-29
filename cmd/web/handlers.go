package main
 
import (
    "net/http"
    "html/template"
    "log"
)
 
func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    
    tm, err := template.ParseFiles(".ui/html/home.page.tmpl")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Server error", 500)
        return
    }

    err = tm.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Server error", 500)
    }

    w.Write([]byte("Привет из Snippetbox"))
}