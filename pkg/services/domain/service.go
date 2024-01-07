package domain

import "github.com/opoccomaxao/wblitz-watcher/pkg/repo"

type Service struct {
	repo repo.Repository
}

func NewService(
	repo repo.Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}
