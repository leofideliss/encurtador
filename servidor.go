package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"


    "github.com/leofideliss/encurtador/url"
)

var (
    porta int
    urlBase string
)

func init(){
    porta = 8888
    urlBase = fmt.Sprintf("http://localhost:%d",porta)
}

func main(){
    http.HandleFunc("/api/encurtar",Encurtador)
    http.HandleFunc("/r/",Ridirecionar)

    log.Fatal(http.ListenAndServe(
        fmt.Sprintf(":%d",porta),nil))
}


func Encurtador (w http.ResponseWriter , r *http.Request){
    if r.Method != "POST" {
        responderCom(w,http.StatusMethodNotAllowed , Headers{
            "Allow" : "Post",
        })
        return
    }
}

func responderCom( w http.ResponseWriter, status int , headers Headers){
    for k , v := range headers{
        w.Header().Set(k,v)
    }
    w.WriteHeader(status)
}

func extrairUrl(r *http.Request) string {
    url := make([]byte,r.ContentLength,r.ContentLength) // criar um array de bytes
    r.Body.Read(url) // lendo o body para dentro de url - uma cópia para dentro de url
    return string(url) // casting e retorna
}

func Redirecionador(w http.ResponseWriter , r *http.Request){
    caminho:= strings.Split(r.URL.Path,"/")
    id := caminho[len(caminho)-1]

    if url:= url.Buscar(id); url != nil{
        http.Redirect(w,r,url.Destino,http.StatusMovedPermanently)
    }else {
        http.NotFound(w,r)
    }
}
