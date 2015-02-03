package main
 
import (
    "github.com/emicklei/go-restful"
    "log"
    "net/http"
    "./userservice"
)
 
func main() {
    restful.Add(userservice.New())
    log.Fatal(http.ListenAndServe(":8080", nil))
}