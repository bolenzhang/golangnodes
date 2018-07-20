package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,`
	<meta name="go-import" content="bolen.com/test git https://github.com/bolenzhang/test">
`)
}

func main() {
	http.HandleFunc("/test",handler)
	http.ListenAndServe(":80",nil)
}
