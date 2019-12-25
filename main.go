package main

import(
	"log"
	"net/http"
	"webserice/http/controller"
)

func main(){
	http.HandleFunc("/", controller.Index)

	err := http.ListenAndServe(":8000",nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}else{
		log.Printf("listening at port 8000")
	}

}
