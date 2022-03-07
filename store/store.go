package store

import (
	"fmt"
	"sync"

	"github.com/api-paradise/model"
)

type Store interface {
	GetCategories() []model.Category
	GetCategory(id int64) (*model.Category, error)
	
	GetJokes() []model.Joke
	GetJoke(id int64) (*model.Joke, error)
}

var (
	initCtx       sync.Once
	storeInstance *store
)

type store struct {
	jokes      []model.Joke
	categories []model.Category
}

func GetInstance() Store {
	initCtx.Do(func() {
		storeInstance = &store{}
		storeInstance.init()
	})
	return storeInstance
}

func (s *store) init() {
	s.jokes = readJokes()
	s.categories = readCategories()
}

func (s *store) GetJokes() []model.Joke {
	return s.jokes
}

func (s *store) GetJoke(id int64) (*model.Joke, error) {
	for _, v := range s.jokes {
		if v.Id == id {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("failed to get joke with id %d", id)
}

func (s *store) GetCategories() []model.Category {
	return s.categories
}

func (s *store) GetCategory(id int64) (*model.Category, error) {
	for _, v := range s.categories {
		if v.Id == id {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("failed to get category with id %d", id)
}
