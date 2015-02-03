package main
 
import (
    "github.com/emicklei/go-restful"
    "log"
    "net/http"
    "os"
    "./userservice"
)
 
func main() {

	httpPort := os.Getenv("HTTP_PLATFORM_PORT")

	if (len(httpPort) == 0) {
		httpPort = "8080"
	}

    restful.Add(userservice.New())
    log.Fatal(http.ListenAndServe(":" + httpPort, nil))
}