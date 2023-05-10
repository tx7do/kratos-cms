package auth

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

	var strUserName string

	strUserId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, "", errors.New("extract subject failed")
	}
	userId := uint32(strUserId)

	return userId, strUserName, nil
}
