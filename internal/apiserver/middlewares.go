package apiserver

import (
	"context"
	"net/http"
	"strings"
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

func (s *server) sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			openRoutes := []string{
				authRoute,
				updateSessionRoute,
				clearSessionRoute,
			}

			for _, rt := range openRoutes {
				if strings.Contains(r.URL.Path, rt) {
					next.ServeHTTP(w, r)
					return
				}
			}

			c, err := r.Cookie("at")
			if err != nil {
				s.logger.Error(err)
				s.unauthorized(w)
				return
			}

			at := &models.Token{}

			if err = at.GetClaims(c.Value, tokenKey); err != nil {
				s.logger.Error(err)
				s.unauthorized(w)
				return
			}

			s.session.userId = at.UserID
			next.ServeHTTP(w, r)
		},
	)
}
