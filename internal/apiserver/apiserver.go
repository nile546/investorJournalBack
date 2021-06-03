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
	"github.com/nile546/diplom/internal/investinstruments"
	"github.com/nile546/diplom/internal/investinstruments/instruments"
	"github.com/nile546/diplom/internal/mailer"
	"github.com/nile546/diplom/internal/mailer/emailer"
	"github.com/nile546/diplom/internal/models"
	"github.com/nile546/diplom/internal/store"
	"github.com/nile546/diplom/internal/store/pgstore"
)

var (
	production bool
	tokenKey   string
	addr       string
	protocol   string
	addrLand   string
)

//APIServer ...
type APIServer struct {
	config *config.Config
}

type server struct {
	router      *mux.Router
	repository  store.Repository
	mailer      mailer.Mailer
	instruments investinstruments.Instruments
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
	users.HandleFunc(confirmSignupRoute, s.confirmSignup).Methods(http.MethodPost)

	users.HandleFunc(signinRoute, s.signin).Methods(http.MethodPost)

	spa := spaHandler{staticPath: "web", indexPath: "index.html"}

	s.router.PathPrefix("/").Handler(spa)
}

//Start ...
func Start(c *config.Config) error {

	//inst := instruments.New()

	//inst.Stocks().GrabAll(c.SpbexchangeAddress, c.MskexchangeAddress)

	production = c.Production
	tokenKey = c.TokenKey

	protocol = c.Protocol
	addr = c.Host + ":" + c.Port
	addrLand = c.LandingAddress

	db, err := newDB(c.ConnectionString)
	if err != nil {
		return err
	}

	r := pgstore.New(db)

	mConf := &emailer.Config{
		Login:  c.MailerLogin,
		Pass:   c.MailerPass,
		Host:   c.MailerHost,
		Port:   c.MailerPort,
		Sender: c.MailerSender,
	}

	m := emailer.New(mConf)

	i := instruments.New()

	// Сделать конфиг + добавить токен крипты

	srv := newServer(r, m, i)
	fmt.Println("Started server at ", addr)

	return http.ListenAndServe(addr, srv)

}

func newServer(r store.Repository, m mailer.Mailer, i investinstruments.Instruments) *server {

	srv := &server{
		router:      mux.NewRouter(),
		repository:  r,
		mailer:      m,
		instruments: i,
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

func (s *server) respond(w http.ResponseWriter, payload interface{}) {
	res := models.Result{
		Status:       models.Ok,
		ErrorMessage: "",
		Payload:      payload,
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
