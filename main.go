package main

import(
	"fmt"
	"log"
  "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hallo Wald\n")
}

func main(){
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started and listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))

}