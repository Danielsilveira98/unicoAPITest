package middleware

import (
	"context"
	"net/http"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/gorilla/mux"
)

type idGen func() string

type TraceIDMiddleware struct {
	idGen idGen
}

func NewTraceIDMiddleware(idGen idGen) *TraceIDMiddleware {
	return &TraceIDMiddleware{idGen}
}

func (m *TraceIDMiddleware) Middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			traceIDHKey := http.CanonicalHeaderKey(string(domain.TraceIDCtxKey))
			var traceId string
			if r.Header[traceIDHKey] == nil {
				traceId = m.idGen()
			} else {
				traceId = r.Header[traceIDHKey][0]
			}

			ctx = context.WithValue(ctx, domain.TraceIDCtxKey, traceId)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
