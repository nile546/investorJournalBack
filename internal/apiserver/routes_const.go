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

	// crypto
	cryptoDealsRoute = "/crypto"

	// crypto instruments
	cryptoInstrumentsRoute = "/crypto-instruments"

	// deposit
	depositDealsRoute = "/deposit"

	// crypto instruments
	bankInstrumentsRoute = "/bank-instruments"

	// currencies
	currencyRoute = "/currency"

	// common
	getAllRoute          = "/getAll"
	getRoute             = "/get"
	createRoute          = "/create"
	deleteRoute          = "/delete"
	updateRoute          = "/update"
	getPopularInstrument = "/popular"
)
