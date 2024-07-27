package main

import ("errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got  /hello request\n")
	io.WriteString(w, "Hello , HTTP\n"
}
func main(){
	http.handleFunc("/" , getRoot)	
	http.handleFunc("/hell" , getHello)
err := http.ListenAndServer(":3333",nil)

}

