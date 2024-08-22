package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/jwt"
	"gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/models"
	"gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

const (
	Success string = "successfully registred"
	Fail    string = "registration failed"
)

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	appProvider AppProvider // Я использую для одного secret
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (statusMsg string, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
}

type AppProvider interface {
	Secret(ctx context.Context, id int) (models.Secret, error)
	GetPayload(ctx context.Context, payload *jwt.MyClaims) (models.User, error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider, // Я использую для одного secret
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		usrSaver:    userSaver,
		usrProvider: userProvider,
		log:         log,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
	}
}

// Login checks if user with given credentials exists in the system and returns access token.
//
// If user exists, but password is incorrect, returns error.
// If user doesn't exist, returns error.
func (a *Auth) Login(ctx context.Context, email string, password string, idSec int) (string, error) {
	const op = "Auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("username", email),
	)

	log.Info("attempting to login user")

	user, err := a.usrProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", slog.String("err", err.Error()))

			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", slog.String("err", err.Error()))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		a.log.Info("invalid credentials", slog.String("err", err.Error()))

		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	sec, err := a.appProvider.Secret(ctx, idSec)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user logged in successfully")

	token, err := jwt.NewToken(user, sec, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token", slog.String("err", err.Error()))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, password string) (string, error) {

	const op = "auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", slog.String("err", err.Error()))

		return Fail, fmt.Errorf("%s: %w", op, err)
	}

	msg, err := a.usrSaver.SaveUser(ctx, email, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			log.Warn("user already exists", slog.String("err", err.Error()))

			return Fail, fmt.Errorf("%s: %w", op, ErrUserExists)
		}

		log.Error("failed to save user", slog.String("err", err.Error()))

		return Fail, fmt.Errorf("%s: %w", op, err)
	}

	return msg, nil
}

func (a *Auth) Logout(ctx context.Context, token string) (string, error) {
	const op = "Auth.Logout"

	log := a.log.With(
		slog.String("op", op),
		slog.String("token", token),
	)

	log.Info("logging out ...")

	token = ""

	// TODO: Хз как, нужно подумать как испортить токен, как костыль мб переприсваивать новый невалидный токен.
	// Но я не могу ибо в proto файле я в респонсе указал чисто как bool, а не string
	//P.S поменял

	log.Info("successfully logged out")

	return token, nil
}

func (a *Auth) ValidateToken(ctx context.Context, token string, idSec int) (int, error) {
	const op = "Auth.ValidateToken"

	log := a.log.With(
		slog.String("op", op),
		slog.String("token", token),
	)

	log.Info("validating token ...")

	sec, err := a.appProvider.Secret(ctx, idSec)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	MyPayload, err := jwt.ValidateToken(token, sec)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	user, err := a.appProvider.GetPayload(ctx, MyPayload)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}


	return int(user.ID), nil
}
