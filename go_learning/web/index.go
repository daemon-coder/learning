package web

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, request.Method)
	fmt.Fprintln(writer, request.URL.Path)
	fmt.Fprintln(writer, request.Proto)
	fmt.Fprintln(writer, request.Header)
	fmt.Fprintln(writer, request.Body)
}

func WebDemo() {
	ServeWithDefaultServeMux()
	// ServeWithCustomServeMux()
	// HttpsDemo()
	// HttpRouterDemo()
}

func ServeWithDefaultServeMux() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/request/url", Request)
	http.HandleFunc("/response/welcome", ResponseContent)
	http.HandleFunc("/response/redirect", ResponseHeader)
	http.HandleFunc("/cookie", Cookie)
	http.ListenAndServe(":9000", nil)
}

func ServeWithCustomServeMux() {
	mux := http.NewServeMux()
	// files := http.FileServer(http.Dir("/public"))
	// mux.Handle("/static", http.StripPrefix("/static", files))
	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr: ":9001",
		Handler: mux,
	}
	server.ListenAndServe()
}

func HttpsDemo() {
	http.ListenAndServeTLS(":9002", "cert.pem", "key.pem", nil)
}

func article(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Params:%+v", p)
}

func HttpRouterDemo() {
	mux := httprouter.New()
	mux.GET("/article/:name/:id", article)
	server := http.Server{
		Addr: ":9002",
		Handler: mux,
	}
	server.ListenAndServe()
}