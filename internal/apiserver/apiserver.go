package apiserver

import "github.com/nile546/diplom/config"

//APIServer ...
type APIServer struct {
	config *config.Config
}

//NewAPIServer ...
func NewAPIServer(c *config.Config) (*APIServer, error) {
	var a *APIServer
	a.config = c

	return a, nil
}
