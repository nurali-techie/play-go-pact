package provider

import (
	"fmt"
	"net/http"
)

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	msg := fmt.Sprintf("Hello %s!", name)
	res.Write([]byte(msg))
}
