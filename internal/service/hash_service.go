package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crabocod/golang-test/internal/utils"

	"github.com/Crabocod/golang-test/internal/model"
)

type HashService interface {
	CreateHash(ctx context.Context, request model.HashRequest) (*model.HashResponse, error)
}

type CacheService interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value *model.HashResponse) error
}

type hashService struct {
	cacheService CacheService
}

func NewHashService(cacheService CacheService) HashService {
	return &hashService{
		cacheService: cacheService,
	}
}

func (s *hashService) CreateHash(ctx context.Context, request model.HashRequest) (*model.HashResponse, error) {
	var response *model.HashResponse

	str := fmt.Sprintf("%s:%s", request.Text, request.Algorithm)

	// Try to get from cache
	if cached, err := s.cacheService.Get(ctx, str); err == nil {
		if err := json.Unmarshal([]byte(cached), &response); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cached value: %w", err)
		}

		return response, nil
	}

	// Generate hash
	hashedText, err := utils.HashGenerate(request.Text, request.Algorithm)
	if err != nil {
		return nil, fmt.Errorf("failed to create hash: %w", err)
	}

	response = &model.HashResponse{
		Original:  request.Text,
		Hashed:    hashedText,
		Algorithm: request.Algorithm,
	}

	// Store in cache
	if err := s.cacheService.Set(ctx, str, response); err != nil {
		return nil, fmt.Errorf("failed to cache result: %w", err)
	}

	return response, nil
}
