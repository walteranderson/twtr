package storage

import (
	"database/sql"
	"errors"
	"time"

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
      view_count INTEGER NOT NULL,
      posted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
  `

	_, err := s.db.Exec(query)
	return err
}

func (s *SQLiteStorage) GetAllPosts() ([]types.Post, error) {
	rows, err := s.db.Query("SELECT id, body, view_count, posted_at FROM posts")
	if err != nil {
		return nil, err
	}

	var posts []types.Post = []types.Post{}
	for rows.Next() {
		var post types.Post
		if err := rows.Scan(&post.ID, &post.Body, &post.ViewCount, &post.PostedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *SQLiteStorage) GetPost(id string) (*types.Post, error) {
	row := s.db.QueryRow("SELECT id, body, view_count, posted_at FROM posts WHERE id = ?", id)

	var post types.Post
	if err := row.Scan(&post.ID, &post.Body, &post.ViewCount, &post.PostedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}

		return nil, err
	}

	return &post, nil
}

func (s *SQLiteStorage) CreatePost(post types.Post) (*types.Post, error) {
	res, err := s.db.Exec("INSERT INTO posts (body, view_count) values (?,?)", post.Body, 0)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	post.ID = id
	post.ViewCount = 0
	post.PostedAt = time.Now()

	return &post, nil
}

func (s *SQLiteStorage) UpdatePost(post types.Post) (*types.Post, error) {
  _, err := s.db.Exec("UPDATE posts SET body = ? WHERE id = ?", post.Body, post.ID)
  if err != nil {
    return nil, err
  }

  return &post, nil
}
