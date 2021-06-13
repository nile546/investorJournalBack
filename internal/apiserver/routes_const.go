package apiserver

const (
	apiRoute = "/api"

	// auth
	authRoute          = "/auth"
	signupRoute        = "/signup"
	signinRoute        = "/signin"
	confirmSignupRoute = "/confirm-signup"

	// landing
	landingRoute = "/landing"

	// users
	usersRoute = "/users"

	// stocks
	stockDealsRoute = "/stock-deals"

	// common
	getAllRoute = "/getAll"
	get         = "/get"
	create      = "/create"
	delete      = "/delete"
	update      = "/update"

	// session routes
	updateSessionRoute = "/refresh"
	clearSessionRoute  = "/delete"

	//Any routes
	privateRoute = "/private"
)
