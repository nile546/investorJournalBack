package apiserver

const (
	apiRoute = "/api"

	// auth
	authRoute           = "/auth"
	signupRoute         = "/signup"
	signinRoute         = "/signin"
	confirmSignupRoute  = "/confirm-signup"
	signoutRoute        = "/signout"
	refreshRoute        = "/refresh"
	getCurrentUserRoute = "/get-current-user"

	// landing
	landingRoute = "/landing"

	// users
	usersRoute = "/users"

	// stock deals
	stockDealsRoute     = "/stock-deals"
	getBrokerDealsRoute = "/get-brocker-deals"

	// stock instruments
	stockInstrumentsRoute = "/stock-instruments"
	getPopularInstrument  = ""

	// crypto
	cryptoDealsRoute = "/crypto"

	// deposit
	depositDealsRoute = "/deposit"

	// currencies
	currencyRoute = "/currency"

	// grab deal from brocker

	// common
	getAllRoute = "/getAll"
	getRoute    = "/get"
	createRoute = "/create"
	deleteRoute = "/delete"
	updateRoute = "/update"
)
