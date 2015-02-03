package main

import (
  "os"
  "fmt"
  "net/http"
  "time"
  "encoding/json"
)

func main() {
  
  httpPort := os.Getenv("HTTP_PLATFORM_PORT")

  if len(httpPort) == 0  {
    httpPort = "8080"
  }

  http.HandleFunc("/hello", viewHandlerHelloWorld1)
  http.HandleFunc("/hello2", viewHandlerHelloWorld2)
  http.ListenAndServe(":" + httpPort, nil)
}

func viewHandlerHelloWorld1(w http.ResponseWriter, r *http.Request) {  
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-type", "application/json")  

  jsonMsg, err := getResponse("Title", "Hello World!") 
  if err != nil {
    http.Error(w, "Oops", http.StatusInternalServerError)
  }
  fmt.Fprintf(w, jsonMsg)
}

func viewHandlerHelloWorld2(w http.ResponseWriter, r *http.Request) {  
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-type", "application/json")  

  jsonMsg, err := getResponse("Title 2", "Hello World 2!") 
  if err != nil {
    http.Error(w, "Oops", http.StatusInternalServerError)
  }
  fmt.Fprintf(w, jsonMsg)
}

func getResponse(t string, m string) (string, error){
  unixtime := int32(time.Now().Unix())
  msg := Message{t, m, unixtime}
  jbMsg, err := json.Marshal(msg)

  if err != nil {    
    return "", err
  }

  jsonMsg := string(jbMsg[:]) // converting byte array to string
  return jsonMsg, nil
}

type Message struct {
  Title string
  Body string
  Time int32
}