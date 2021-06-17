package apiserver

import (
	"net/http"
)

func (s *server) getCurrenciesRatio(w http.ResponseWriter, r *http.Request) {

	c, err := s.currencyGrab.GrabCbr().GrabUsdEur()
	if err != nil {
		s.logger.Errorf("Error grab currency with error:%+v", err)
		s.error(w, err.Error())
	}

	s.respond(w, c)

}
