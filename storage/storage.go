package storage

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v7"

)

type Storage struct {
	redis *redis.Client
}

func NewStorage(redis *redis.Client) *Storage {
	return &Storage{
		redis,
	}
}

type Page struct {
	ID          string `json:id`
	TITLE       string `json:title`
	DESCRIPTION string `json:DESCRIPTION`
}

func (s *Storage) GetPageByID(id string) (p *Page, err error) {

	b, err := s.redis.Get(id).Bytes()
	if err != nil {
		err = fmt.Errorf("failed to get page %s: %v", id, err)
		return

	}

	err = json.Unmarshal(b, &p)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal %s: %v", id, err)
	}

	return
}
