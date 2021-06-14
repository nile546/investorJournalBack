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
	"github.com/nile546/diplom/internal/brokersgrab"
	"github.com/nile546/diplom/internal/brokersgrab/grabs"
	"github.com/nile546/diplom/internal/investinstruments"
	"github.com/nile546/diplom/internal/investinstruments/instruments"
	"github.com/nile546/diplom/internal/mailer"
	"github.com/nile546/diplom/internal/mailer/emailer"
	"github.com/nile546/diplom/internal/models"
	"github.com/nile546/diplom/internal/store"
	"github.com/nile546/diplom/internal/store/pgstore"
	"github.com/sirupsen/logrus"
)

const (
	ctxKeyRequestID ctxKey = iota
)

type ctxKey int8

var (
	production bool
	tokenKey   string
	addr       string
	protocol   string
	addrLand   string
	logLevel   string
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
	session     session
	logger      *logrus.Logger
	brokersGrab brokersgrab.Grab
}

type spaHandler struct {
	staticPath string
	indexPath  string
}
type session struct {
	User *models.User
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

	s.router.Use(requestIDMiddleware)
	s.router.Use(s.loggerMiddleware)
	s.router.Use(s.GetUserSession)

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

	// Open routes, use without session

	api.HandleFunc(updateSessionRoute, s.updateSession)
	api.HandleFunc(clearSessionRoute, s.clearSession)

	auth := api.PathPrefix(authRoute).Subrouter()
	auth.HandleFunc(signupRoute, s.signup).Methods(http.MethodPost)
	auth.HandleFunc(confirmSignupRoute, s.confirmSignup).Methods(http.MethodPost)
	auth.HandleFunc(signinRoute, s.signin).Methods(http.MethodPost)

	// Closed routes, with use session

	stockDeals := api.PathPrefix(stockDealsRoute).Subrouter()
	stockDeals.HandleFunc(getAllRoute, s.getAllStockDeals).Methods(http.MethodPost)

	spa := spaHandler{staticPath: "web", indexPath: "index.html"}

	s.router.PathPrefix("/").Handler(spa)
}

//Start ...
func Start(c *config.Config) error {

	production = c.Production
	tokenKey = c.TokenKey
	logLevel = c.LogLevel

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

	l := logrus.New()

	i := instruments.New(l)

	cI := &updateInstrumentsConfig{
		spbExchangeUrl: c.SpbexchangeAddress,
		mskStocksUrl:   c.MskexchangeAddress,
		bankiUrl:       c.BankiUrl,
		cryptoUrl:      c.CryptoUrl,
		cryptoKey:      c.CryptoKey,
		hours:          c.HoursUpdateInstruments,
		minutes:        c.MinutesUpdateInstruments,
		seconds:        c.SecondsUpdateInstruments,
	}

	b := grabs.New()

	//b.TinkoffGrab().GetTinkoffStockDeals("t.vOoTCbcoO8nacpr4lKaVx0Hv5OwOju93dctV0umgOmhp8OKX9bq-9EcyGRzoBBp_yIbs4iQgITafLD0r5tvyYg")

	srv := newServer(r, m, i, l, b)

	if err = srv.configureLogger(); err != nil {
		return err
	}

	err = srv.updateInstruments(srv.callUpdateHandlers, cI)
	if err != nil {
		srv.logger.Errorf("Error update instruments: %+v", err)
	}

	fmt.Println("Started server at ", addr)

	return http.ListenAndServe(addr, srv)

}

func newServer(r store.Repository, m mailer.Mailer, i investinstruments.Instruments, l *logrus.Logger, b brokersgrab.Grab) *server {

	srv := &server{
		router:      mux.NewRouter(),
		repository:  r,
		mailer:      m,
		logger:      l,
		instruments: i,
		brokersGrab: b,
	}
	srv.ConfugureRouter()

	if err := srv.configureLogger(); err != nil {
		srv.logger.Errorf("Error configure logger: %+v", err)
	}

	return srv

}

func (s *server) error(w http.ResponseWriter, errorMessage string) {
	res := models.Result{
		Status:       models.Error,
		ErrorMessage: errorMessage,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		s.logger.Errorf("error encode data %+v to json: %+v", res, err)
	}
}

func (s *server) respond(w http.ResponseWriter, payload interface{}) {
	res := models.Result{
		Status:       models.Ok,
		ErrorMessage: "",
		Payload:      payload,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		s.logger.Errorf("error encode data %+v to json: %+v", res, err)
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

func (s *server) configureLogger() error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	// Uncomment for get full path of called method in log message.
	// s.logger.SetReportCaller(true)

	return nil
}
