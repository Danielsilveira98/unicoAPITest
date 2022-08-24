package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type loggerRequest interface {
	Infof(context.Context, string, map[string]interface{})
}

type LogRequestMiddleware struct {
	logger loggerRequest
}

func NewLogRequestMiddleware(logger loggerRequest) *LogRequestMiddleware {
	return &LogRequestMiddleware{logger}
}

func (m *LogRequestMiddleware) Middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			path := r.URL.Path
			method := r.Method
			msg := fmt.Sprintf("Receiving request %s on path %s", method, path)

			m.logger.Infof(ctx, msg, map[string]interface{}{
				"path":   path,
				"method": method,
			})

			next.ServeHTTP(w, r)
		})
	}
}
