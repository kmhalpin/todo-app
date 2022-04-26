package sqlite

import (
	"context"
	"database/sql"
	"errors"

	errorCommon "github.com/kmhalpin/todoapp/common/error"
	uModel "github.com/kmhalpin/todoapp/internal/model/user"
)

func (r sqliteUserRepository) GetUser(ctx context.Context) (users []uModel.User, err error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT
			id,
			permission,
			username,
			created_at,
			updated_at
		FROM users
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users = []uModel.User{}
	for rows.Next() {
		var user uModel.User
		if err := rows.Scan(
			&user.ID,
			&user.Permission,
			&user.Username,
			&user.TimeCreated,
			&user.TimeUpdated,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r sqliteUserRepository) GetUserByID(ctx context.Context, id string) (user uModel.User, err error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT
			id,
			permission,
			username,
			created_at,
			updated_at
		FROM users
		WHERE id = ?
	`, id)
	err = row.Scan(
		&user.ID,
		&user.Permission,
		&user.Username,
		&user.TimeCreated,
		&user.TimeUpdated,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return user, errorCommon.NewNotFoundError("user not found")
	}
	return user, err
}

func (r sqliteUserRepository) GetPasswordByUsername(ctx context.Context, username string) (password string, err error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT
			password
		FROM users
		WHERE username = ?
	`, username)
	err = row.Scan(&password)
	if errors.Is(err, sql.ErrNoRows) {
		return password, errorCommon.NewNotFoundError("user not found")
	}
	return password, err
}

func (r sqliteUserRepository) GetIDByUsername(ctx context.Context, username string) (id string, err error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT
			id
		FROM users
		WHERE username = ?
	`, username)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}

func (r sqliteUserRepository) GetUsernameIsAvailable(ctx context.Context, username string) (err error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT
			username
		FROM users
		WHERE username = ?
	`, username)
	var u string
	err = row.Scan(&u)
	if u == username {
		return errorCommon.NewInvariantError("username not available")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}
