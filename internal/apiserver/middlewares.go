package apiserver

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nile546/diplom/internal/models"
	"github.com/sirupsen/logrus"
)

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		r = r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id))

		next.ServeHTTP(w, r)
	})
}

// Logger
func (s *server) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})

		logger.Infof("started %s %s", r.Method, r.RequestURI)
		start := time.Now()

		resp := &response{w, http.StatusOK}

		next.ServeHTTP(resp.writer, r)

		s.logger.Infof(
			"completed with %d %s in %v",
			resp.code,
			http.StatusText(resp.code),
			time.Since(start),
		)
	})
}

func (s *server) GetUserSession(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			ck, err := r.Cookie("at")
			if err != nil {
				s.respond(w, err.Error())
				return
			}

			at := &models.Token{}

			if err = at.GetClaims(ck.Value, tokenKey); err != nil {
				s.respond(w, err.Error())
				return
			}

			s.session.User.ID = at.UserID

			next.ServeHTTP(w, r)
		},
	)
}
