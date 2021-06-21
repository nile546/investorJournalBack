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
	cryptoDealsRoute = "/crypto-deals"

	// crypto instruments
	cryptoInstrumentsRoute = "/crypto-instruments"

	// deposit
	depositDealsRoute = "/deposit-deals"

	// crypto instruments
	bankInstrumentsRoute = "/bank-instruments"

	// currencies
	currencyRatesRoute = "/currency-rates"

	// strategies
	strategiesRoute = "/strategies"

	// patterns
	patternsRoute = "/patterns"

	// common
	getAllRoute          = "/getAll"
	getRoute             = "/get"
	createRoute          = "/create"
	deleteRoute          = "/delete"
	updateRoute          = "/update"
	getPopularInstrument = "/popular"
)
