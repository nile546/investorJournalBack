package apiserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nile546/diplom/config"
)

//APIServer ...
type APIServer struct {
	config *config.Config
}

type server struct {
	router *mux.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) ConfugureRouter() {

	s.router.HandleFunc("/test", test)

}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("111")
}

//Start ...
func Start(c *config.Config) error {

	addr := c.Address + ":" + c.Port
	fmt.Printf("Server start at %s\n", addr)

	srv := newServer()

	return http.ListenAndServe(addr, srv)

}

func newServer() *server {

	srv := &server{
		router: mux.NewRouter(),
	}
	srv.ConfugureRouter()
	return srv

}
