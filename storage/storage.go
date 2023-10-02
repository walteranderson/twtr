package storage

import "github.com/walteranderson/twtr/types"

type Storage interface {
	Migrate() error
	GetAllPosts() ([]types.Post, error)
	GetPost(id string) (*types.Post, error)
	CreatePost(post types.Post) (*types.Post, error)
}
