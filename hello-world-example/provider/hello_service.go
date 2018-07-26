package provider

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start")

	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	msg := fmt.Sprintf("Hello %s!", name)
	res.Write([]byte(msg))
}
