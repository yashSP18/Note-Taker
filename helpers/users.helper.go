package helpers

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yash-gkmit/NOTE-TAKER/constants"
)

func GetUserIDFromContext(ctx context.Context) (string, error) {

	decodedUser, ok := ctx.Value(constants.UserContextKey).(jwt.MapClaims)
	if !ok || decodedUser == nil {
		return "", errors.New("no user found in context")
	}

	userId, ok := decodedUser["userId"].(string)
	if !ok || userId == "" {
		return "", errors.New("invalid or missing userId in token")
	}

	return userId, nil
}
