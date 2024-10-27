package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/pedroxer/ordermanagmentsystem/internal/lib/jwt"
	models "github.com/pedroxer/ordermanagmentsystem/internal/storage"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Auth struct {
	log      *log.Logger
	provider UserProvider
	saver    UserSaver
	TokenTTL int
	Secret   string
}

type UserSaver interface {
	SaveUser(ctx context.Context, email, password string) (int64, error)
}

type UserProvider interface {
	GetUser(ctx context.Context, email string) (models.User, error)
}

var (
	ErrInvalidCreds = errors.New("invalid credentials")
)

// New returns a new Auth instance
func New(log *log.Logger, saver UserSaver, provider UserProvider, tokenTTL int) *Auth {
	return &Auth{
		log:      log,
		provider: provider,
		saver:    saver,
		TokenTTL: tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email, password string) (string, error) {
	const op = "auth.Login"
	log := a.log.WithField("op", op)

	user, err := a.provider.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			log.Warn("user not found", err)
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCreds)
		}

		log.Error("failed to get user: ", err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		log.Error("failed to compare password hash: ", err)
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCreds)
	}

	token, err := jwt.NewToken(user, a.Secret, time.Duration(a.TokenTTL)*time.Second)
	if err != nil {
		log.Error("failed to create token: ", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) RegisterUser(ctx context.Context, email, password string) (int64, error) {
	const op = "auth.RegisterUser"
	log := a.log.WithField("op", op)

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash: ", err)
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	id, err := a.saver.SaveUser(ctx, email, string(passHash))

	if err != nil {
		if errors.Is(err, models.ErrUserAlreadyExists) {
			log.Warn("user already exists", err)
			return -1, fmt.Errorf("%s: %w", op, models.ErrUserAlreadyExists)
		}

		log.Error("failed to save user: ", err)
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
