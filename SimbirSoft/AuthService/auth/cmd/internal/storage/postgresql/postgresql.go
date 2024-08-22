package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
	"gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/jwt"
	"gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/models"
)

const (
	Success string = "successfully registred"
	Fail    string = "registration failed"

	saveCommand   string = "INSERT INTO users(email, pass_hash) VALUES($1, $2)"
	selectCommand string = "SELECT id, email, pass_hash FROM users WHERE email = $1"
	secretCommand string = "SELECT id, secret FROM secrets WHERE id = $1"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgresql.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (string, error) {
	const op = "storage.postgresql.SaveUser"

	stmt, err := s.db.Prepare(saveCommand)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return Fail, fmt.Errorf("%s: %s", op, "user already exists")
		}

		return Fail, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, email, passHash)

	if err != nil {
		return Fail, fmt.Errorf("%s, %s: %w", res, op, err)
	}

	return Success, nil // Стрингу возвращать нехорошо
}

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
	const op = "storage.postgresql.User"

	stmt, err := s.db.Prepare(selectCommand)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, email)

	var user models.User
	err = row.Scan(&user.ID, &user.Email, &user.PassHash)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// Secret returns Secret.
func (s *Storage) Secret(ctx context.Context, id int) (models.Secret, error) {
	const op = "storage.sqlite.Secret"

	stmt, err := s.db.Prepare(secretCommand)
	if err != nil {
		return models.Secret{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, id)

	var sec models.Secret
	err = row.Scan(&sec.ID, &sec.Secret)
	if err != nil {
		return models.Secret{}, fmt.Errorf("%s: %w", op, err)
	}

	return sec, nil
}

func (s *Storage) GetPayload(ctx context.Context, payload *jwt.MyClaims) (models.User, error) {
	const op = "storage.sqlite.GetPayload"

	stmt, err := s.db.Prepare("SELECT id FROM users WHERE email = $1")
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, payload.Email)

	var user models.User
	err = row.Scan(&user.ID)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
