package loom

import (
	"context"
	"fmt"

	"go.octolab.org/tact/loop/internal/domain"
)

func New(client ServiceClient) *Service {
	return &Service{c: client}
}

type Service struct {
	c ServiceClient
}

func (s *Service) Tree(ctx context.Context, workspaceID int) (*domain.Tree, error) {
	var (
		tree domain.Tree
		err  error
	)

	tree.Spaces, err = s.spaces(ctx, workspaceID)
	if err != nil {
		return nil, fmt.Errorf("fetch tree: %w", err)
	}

	return &tree, nil
}
