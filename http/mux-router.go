package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	muxDispatcher = mux.NewRouter()
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}
func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")

}
func (*muxRouter) SERVE(port string) {
	fmt.Printf("MUX http server runnin at port:%v", port)
	http.ListenAndServe(port, muxDispatcher)
}
