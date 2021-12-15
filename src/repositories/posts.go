package repositories

import (
	"database/sql"
	"devbook-api/src/models"
)

type posts struct {
	db *sql.DB
}

func NewRepositoryFromPosts(db *sql.DB) *posts {
	return &posts{db}
}

func (repository posts) Create(post models.Post) (uint64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(post.Title, post.Content, post.AuthorId)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (repository posts) SearchById(pid uint64) (models.Post, error) {
	rows, err := repository.db.Query("SELECT p.*, u.nick FROM posts p INNER JOIN users u ON u.id = p.author_id WHERE p.id = ?", pid)
	if err != nil {
		return models.Post{}, nil
	}
	defer rows.Close()

	var post models.Post

	if rows.Next() {
		if err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}
	return post, nil
}

func (repository posts) Search(uid uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT DISTINCT  p.*, u.nick FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		INNER JOIN followers f ON p.author_id = f.user_id
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 DESC;`, uid, uid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (repository posts) Update(pid uint64, post models.Post) error {
	stmt, err := repository.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(post.Title, post.Content, pid); err != nil {
		return err
	}

	return nil
}

func (repository posts) Delete(pid uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(pid); err != nil {
		return err
	}

	return nil
}

func (repository posts) SearchByUser(uid uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
	SELECT DISTINCT p.*, u.nick FROM posts p 
	INNER JOIN users u ON p.author_id = u.id 
	WHERE u.id = ?;
	`, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func (repository posts) Like(pid uint64) error {
	stmt, err := repository.db.Prepare(`
		UPDATE posts 
		SET likes = likes + 1
		WHERE id = ?`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(pid); err != nil {
		return err
	}
	return nil
}

func (repository posts) Unlike(pid uint64) error {
	stmt, err := repository.db.Prepare(`
		UPDATE posts SET likes =
		CASE 
			WHEN likes > 0 THEN likes - 1
		END
		WHERE id = ?`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(pid); err != nil {
		return err
	}
	return nil
}
