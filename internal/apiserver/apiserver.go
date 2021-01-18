package apiserver

import (
	"net/http"
	"os"
	"path/filepath"

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

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = filepath.Join(s.staticPath, path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(s.staticPath, s.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(s.staticPath)).ServeHTTP(w, r)
}

func (s *server) ConfugureRouter() {

	//s.router.HandleFunc("/test", test)
	spa := spaHandler{staticPath: "web", indexPath: "index.html"}

	s.router.PathPrefix("/").Handler(spa)
}

//func test(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("111")
//}

//Start ...
func Start(c *config.Config) error {

	addr := c.Address + ":" + c.Port

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
