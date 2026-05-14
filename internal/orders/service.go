package orders

import (
	"context"

	repo "github.com/Tnozone/ecom/internal/adapters/postgresql/sqlc"
)

type svc struct {
	repo *repo.Queries
}

func NewService(repo repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {

}
