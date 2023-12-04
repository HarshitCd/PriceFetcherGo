package main

import (
	"context"
	"fmt"
)

type metricsService struct {
	next PriceFetcher
}

func NewMetricsService(next PriceFetcher) PriceFetcher {
	return &metricsService{
		next: next,
	}
}

func (s *metricsService) FetchPrice(ctx context.Context, ticker string) (float64, error) {

	price, err := s.next.FetchPrice(ctx, ticker)
	fmt.Println("Add the metrics sending logic here")

	return price, err

}
