package storage

import (
	"database/sql"

	"github.com/walteranderson/twtr/types"
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

func (s *SQLiteStorage) GetAllPosts() []*types.Post {
	var posts []*types.Post

	for i := 0; i < 5; i++ {
		posts = append(posts, &types.Post{})
	}

	return posts
}

func (s *SQLiteStorage) GetPost(id string) *types.Post {
	return &types.Post{}
}
