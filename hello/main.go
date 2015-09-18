package main

import (
        "fmt"
        "net/http"
        "os"
        "io/ioutil"
)

const (
        HostVar = "HOST"
        PortVar = "PORT"
)

func main() {
        http.HandleFunc("/", index)
        http.HandleFunc("/env", env)
        var port string
        if port = os.Getenv(PortVar); port == "" {
                port = "5000"
        }
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                panic(err)
        }

}

func index(res http.ResponseWriter, req *http.Request) {
        contents,err:= ioutil.ReadFile("./static/index.html" )
        if err !=nil{
        fmt.Fprintf( res,"404" )
        return
		}
        fmt.Fprintf( res,"%s\n",contents )
}


func env(res http.ResponseWriter, req *http.Request) {
        fmt.Fprint(res, "System Environment:\n\n")
        env := os.Environ()
        for _, e := range env {
                fmt.Fprintln(res, e)
        }
}
