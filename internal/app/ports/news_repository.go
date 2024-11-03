package ports

import "news-back-go/internal/app/core"

type NewsRepository interface {
	Create(news *core.News) error
	GetById(id string) (*core.News, error)
	Update(news *core.News) error
	Delete(id string) error
	GetAll() ([]*core.News, error)
}
