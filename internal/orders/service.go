package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/Tnozone/ecom/internal/adapters/postgresql/sqlc"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProdcutNoStock = error.New("product not enough stock")
)

type svc struct {
	repo *repo.Queries
	db *pgx.Conn
}

func NewService(repo repo.Queries) Service {
	return &svc{
		repo: repo,
		db: db,
	}
}

func (s *svc) ListProducts(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("customer ID is required")
	}
	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("at least one item is required")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	deter tx.Rollback(ctx)
	qtx := s.repo.WithTX(tx)

	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}

	for _, item := range tempOrder.items {
		product, err := qtx.repo.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProdcutNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProdcutNoStock
		}

		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID: order.ID,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
			PriceCents: item.PriceInCents,
		})
		if err != nil {
			return repo.Order{}, err
		}
	}

	tx.Commit(ctx)

	return order, nil
}
