package storage

import "github.com/walteranderson/twtr/types"

type Storage interface {
	Migrate() error
	GetAllPosts() []*types.Post
	GetPost(id string) *types.Post
}
