package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
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
	TokenTTL time.Duration
	Secret   string
}

type UserSaver interface {
	SaveUser(ctx context.Context, params models.SaveUserParams) (int32, error)
}

type UserProvider interface {
	GetUser(ctx context.Context, email pgtype.Text) (models.User, error)
}

var (
	ErrInvalidCreds = errors.New("invalid credentials")
)

// New returns a new Auth instance
func New(log *log.Logger, saver UserSaver, provider UserProvider, tokenTTL time.Duration) *Auth {
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

	user, err := a.provider.GetUser(ctx, pgtype.Text{email, true})
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

	token, err := jwt.NewToken(user, a.Secret, a.TokenTTL)
	if err != nil {
		log.Error("failed to create token: ", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) RegisterUser(ctx context.Context, login, email, password string) (int64, error) {
	const op = "auth.RegisterUser"
	log := a.log.WithField("op", op)

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash: ", err)
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	arg := models.SaveUserParams{
		Username:     login,
		Email:        pgtype.Text{email, true},
		PasswordHash: string(passHash),
	}
	id, err := a.saver.SaveUser(ctx, arg)

	if err != nil {
		if errors.Is(err, models.ErrUserAlreadyExists) {
			log.Warn("user already exists", err)
			return -1, fmt.Errorf("%s: %w", op, models.ErrUserAlreadyExists)
		}

		log.Error("failed to save user: ", err)
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return int64(id), nil
}
