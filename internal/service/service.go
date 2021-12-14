package service

import (
	"context"
	"github.com/kanyuanzhi/web-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) *Service {
	return &Service{
		ctx: ctx,
		dao: dao.New(),
	}
}
