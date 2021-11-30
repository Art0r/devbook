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
