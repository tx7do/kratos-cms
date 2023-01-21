package auth

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	authn "github.com/tx7do/kratos-authn/middleware"
	"github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"
)

var action = engine.Action("ANY")

func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			authnClaims, ok := authn.FromContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			sub := engine.Subject(authnClaims.Subject)
			path := engine.Resource(tr.Operation())
			authzClaims := engine.AuthClaims{
				Subject:  &sub,
				Action:   &action,
				Resource: &path,
			}

			ctx = authz.NewContext(ctx, &authzClaims)

			return handler(ctx, req)
		}
	}
}
