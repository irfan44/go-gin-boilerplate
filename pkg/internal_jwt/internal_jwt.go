package internal_jwt

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/irfan44/go-gin-boilerplate/pkg/errs"
)

const (
	bearer                   = "Bearer"
	invalidTokenErrorMessage = "Invalid token."
)

type (
	InternalJwt interface {
		GenerateToken(jwtClaim jwt.MapClaims, secretKey string) string
		ValidateBearerToken(bearerToken string, secretKey string) (jwt.MapClaims, errs.MessageErr)
	}

	internalJwt struct {
	}
)

func (ij *internalJwt) signToken(claims jwt.MapClaims, secretKey string) string {
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func (ij *internalJwt) GenerateToken(jwtClaim jwt.MapClaims, secretKey string) string {
	return ij.signToken(jwtClaim, secretKey)
}

func (ij *internalJwt) parseToken(stringToken string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(invalidTokenErrorMessage)
		}

		return []byte(secretKey), nil
	})

	if err != nil {

		var vErr *jwt.ValidationError

		if errors.As(err, &vErr) {
			if vErr.Errors == jwt.ValidationErrorExpired {
				return nil, errors.New("expired token")
			}
		}

		return nil, errors.New(invalidTokenErrorMessage)
	}

	return token, nil
}

func (ij *internalJwt) ValidateBearerToken(
	bearerToken string,
	secretKey string,
) (jwt.MapClaims, errs.MessageErr) {

	if bearer := strings.HasPrefix(bearerToken, bearer); !bearer {
		return nil, errs.NewUnauthenticatedError(invalidTokenErrorMessage)
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return nil, errs.NewUnauthenticatedError(invalidTokenErrorMessage)
	}

	token, err := ij.parseToken(splitToken[1], secretKey)

	if err != nil {
		return nil, errs.NewUnauthenticatedError(invalidTokenErrorMessage)
	}

	var mapClaims jwt.MapClaims

	if v, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {

		return nil, errs.NewUnauthenticatedError(invalidTokenErrorMessage)
	} else {
		mapClaims = v

	}

	return mapClaims, nil
}

func NewInternalJwt() InternalJwt {
	return &internalJwt{}
}
