package repositories

import (
	"database/sql"
	"devbook-api/src/models"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewRepositoryFromUsers(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	stmt, err := repository.db.Prepare(
		"INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInserted), nil
}

func (repository users) SearchByNickOrName(nickname string) ([]models.User, error) {
	nickname = fmt.Sprintf("%%%s%%", nickname) //  %nickname%

	rows, err := repository.db.Query(
		"SELECT id, name, email, nick, createdAt FROM users WHERE name LIKE ? OR nick LIKE ?",
		nickname, nickname,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.User

	for rows.Next() {
		var this_user models.User

		if err = rows.Scan(
			&this_user.ID,
			&this_user.Name,
			&this_user.Email,
			&this_user.Nick,
			&this_user.CreatedAt,
		); err != nil {
			return nil, err
		}

		result = append(result, this_user)
	}

	return result, nil
}

func (repository users) SearchById(id uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, name, email, nick, createdAt FROM users WHERE id =?",
		id,
	)

	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var this_user models.User

	if rows.Next() {
		if err = rows.Scan(
			&this_user.ID,
			&this_user.Name,
			&this_user.Email,
			&this_user.Nick,
			&this_user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return this_user, nil
}

func (repository users) Update(id uint64, user models.User) error {
	stmt, err := repository.db.Prepare(
		"UPDATE users SET name = ?, email = ?, nick = ? WHERE id = ?",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(user.Name, user.Email, user.Nick, id); err != nil {
		return err
	}

	return nil
}

func (repository users) Delete(id uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM users WHERE id = ?")

	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repotitory users) SearchByEmail(email string) (models.User, error) {
	row, err := repotitory.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, nil
		}
	}

	return user, nil
}

func (repository users) Follow(uid, fid uint64) error {
	stmt, err := repository.db.Prepare("INSERT INTO followers (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(uid, fid); err != nil {
		return err
	}
	return nil
}

func (repository users) Unfollow(uid, fid uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if err != nil {
		return nil
	}
	defer stmt.Close()

	if _, err := stmt.Exec(uid, fid); err != nil {
		return err
	}

	return nil
}

func (repository users) SearchFollowers(uid uint64) ([]models.User, error) {
	rows, err := repository.db.Query(
		"SELECT u.id, u.name, u.nick, u.email, u.createdAt FROM users u INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ?",
		uid,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) SearchFollowing(uid uint64) ([]models.User, error) {
	rows, err := repository.db.Query(
		"SELECT u.id, u.name, u.nick, u.email, u.createdAt FROM users u INNER JOIN followers f ON u.id = f.user_id WHERE f.follower_id = ?",
		uid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) SearchPassword(uid uint64) (string, error) {
	row, err := repository.db.Query("SELECT password FROM users WHERE id = ?", uid)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.Password); err != nil {
			return "", nil
		}
	}
	return user.Password, nil
}

func (repository users) UpdatePassword(uid uint64, password string) error {
	stmt, err := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(password, uid); err != nil {
		return err
	}

	return nil
}
