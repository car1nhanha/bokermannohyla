package service

import (
	"context"
	"thoropa/internal/model"
	"thoropa/internal/repository"
)

type LinkService struct {
	repo *repository.LinkRepository
}

func NewLinkService(r *repository.LinkRepository) *LinkService {
	return &LinkService{repo: r}
}

func (s *LinkService) Create(ctx context.Context, link *model.Link) error {
	// 👉 aqui entra regra de negócio depois
	return s.repo.Create(ctx, link)
}

func (s *LinkService) FindByID(ctx context.Context, id string) (*model.Link, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *LinkService) IncrementAccesses(ctx context.Context, link *model.Link) error {
	link.Accesses++
	return s.repo.Create(ctx, link)
}

func (s *LinkService) FindByIP(ctx context.Context, ip string) ([]*model.Link, error) {
	return s.repo.FindByIP(ctx, ip)
}
