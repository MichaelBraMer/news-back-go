package services

import (
	"news-back-go/internal/app/core"
	"news-back-go/internal/app/ports"
)

type NewsService struct {
	repository ports.NewsRepository
}

func NewNewsService(repo ports.NewsRepository) *NewsService {
	return &NewsService{repository: repo}
}

func (s *NewsService) Create(news *core.News) error {
	return s.repository.Create(news)
}

func (s *NewsService) GetById(id string) (*core.News, error) {
	return s.repository.GetById(id)
}

func (s *NewsService) Update(news *core.News) error {
	return s.repository.Update(news)
}

func (s *NewsService) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *NewsService) GetAll() ([]*core.News, error) {
	return s.repository.GetAll()
}
