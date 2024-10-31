package jwt_maker

import (
	"github.com/golang-jwt/jwt/v5"
	models "github.com/pedroxer/ordermanagmentsystem/internal/storage"
	"time"
)

type JwtMaker struct {
	secret string
	Dur    time.Duration
}

func NewJwtMaker(secret string, dur time.Duration) *JwtMaker {
	return &JwtMaker{
		secret: secret,
		Dur:    dur,
	}
}

// NewToken generates new jwt token for user
func (m *JwtMaker) NewToken(user models.User) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.UserID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(m.Dur).Unix()
	t, err := token.SignedString([]byte(m.secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (m *JwtMaker) VerifyToken(token string) (int64, error) {
	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(m.secret), nil
	})
	if err != nil {
		return -1, err
	}
	if claims, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid && claims["exp"].(float64) > float64(time.Now().Unix()) {
		return int64(claims["uid"].(float64)), nil
	}
	return -1, nil
}
