package storage

import (
	"database/sql"
	"errors"

	"github.com/walteranderson/twtr/types"
)

var (
	ErrNotExists = errors.New("row not exists")
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(db *sql.DB) *SQLiteStorage {
	return &SQLiteStorage{
		db: db,
	}
}

func (s *SQLiteStorage) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS posts(
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      body TEXT NOT NULL,
      view_count INTEGER NOT NULL
    );
  `

	_, err := s.db.Exec(query)
	return err
}

func (s *SQLiteStorage) GetAllPosts() ([]types.Post, error) {
	rows, err := s.db.Query("SELECT id, body, view_count FROM posts")
	if err != nil {
		return nil, err
	}

	var posts []types.Post = []types.Post{}
	for rows.Next() {
		var post types.Post
		if err := rows.Scan(&post.ID, &post.Body, &post.ViewCount); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *SQLiteStorage) GetPost(id string) (*types.Post, error) {
	row := s.db.QueryRow("SELECT id, body, view_count FROM posts")

	var post types.Post
	if err := row.Scan(&post.ID, &post.Body, &post.ViewCount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}

		return nil, err
	}

	return &post, nil
}
