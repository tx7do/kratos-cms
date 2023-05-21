package authn

import (
	"context"
	"errors"
	"strconv"

	"github.com/tx7do/kratos-authn/engine"
)

func ParseFromContext(ctx context.Context) (uint32, string, error) {
	claims, ok := engine.AuthClaimsFromContext(ctx)
	if !ok {
		return 0, "", errors.New("no jwt token in context")
	}

	userId, err := strconv.ParseUint(claims.Subject, 10, 32)
	if err != nil {
		return 0, "", errors.New("extract subject failed")
	}

	var strUserName string

	return uint32(userId), strUserName, nil
}
