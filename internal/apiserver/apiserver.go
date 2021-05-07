package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nile546/diplom/config"
	"github.com/nile546/diplom/internal/models"
)

var (
	production bool
)

//APIServer ...
type APIServer struct {
	config *config.Config
}

type server struct {
	router     *mux.Router
	repository *store.Repository
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

	if !production {
		cors := handlers.CORS(
			handlers.AllowedHeaders([]string{"Content-Type"}),
			handlers.AllowedOrigins([]string{"http://localhost:4200"}),
			handlers.AllowedMethods([]string{http.MethodPost, http.MethodOptions}),
			handlers.AllowCredentials(),
		)

		s.router.Use(cors)
	}

	api := s.router.PathPrefix(apiRoute).Subrouter()

	users := api.PathPrefix(usersRoute).Subrouter()
	users.HandleFunc(signupRoute, s.signup).Methods(http.MethodPost)

	spa := spaHandler{staticPath: "web", indexPath: "index.html"}

	s.router.PathPrefix("/").Handler(spa)
}

//Start ...
func Start(c *config.Config) error {

	production = c.Production

	addr := c.Address + ":" + c.Port

	srv := newServer()

	fmt.Println("Started server at ", addr)

	return http.ListenAndServe(addr, srv)

}

func newServer() *server {

	srv := &server{
		router:     mux.NewRouter(),
		repository: ,
	}
	srv.ConfugureRouter()
	return srv

}

func (s *server) error(w http.ResponseWriter, errorMessage string) {
	res := models.Result{
		Status:       models.Error,
		ErrorMessage: errorMessage,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		//TODO: Добавить сохрание ошибки в логгер.
	}
}
