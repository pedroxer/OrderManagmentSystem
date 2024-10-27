package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	models "github.com/pedroxer/ordermanagmentsystem/internal/storage"
	"time"
)

// NewToken generates new jwt token for user
func NewToken(user models.User, secretPhrase string, duration time.Duration) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.UserID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	t, err := token.SignedString([]byte(secretPhrase))
	if err != nil {
		return "", err
	}
	return t, nil
}
