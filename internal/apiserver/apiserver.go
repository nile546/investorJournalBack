package apiserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nile546/diplom/config"
	"github.com/nile546/diplom/internal/models"
	"github.com/nile546/diplom/internal/pgstore"
	"github.com/nile546/diplom/internal/store"
)

var (
	production bool
	tokenKey   string
)

//APIServer ...
type APIServer struct {
	config *config.Config
}

type server struct {
	router     *mux.Router
	repository store.Repository
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
	tokenKey = c.TokenKey

	addr := c.Address + ":" + c.Port

	db, err := newDB(c.ConnectionString)
	if err != nil {
		return err
	}

	r := pgstore.New(db)
	srv := newServer(r)

	fmt.Println("Started server at ", addr)

	return http.ListenAndServe(addr, srv)

}

func newServer(r store.Repository) *server {

	srv := &server{
		router:     mux.NewRouter(),
		repository: r,
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

func newDB(cs string) (*sql.DB, error) {
	db, err := sql.Open("postgres", cs)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
