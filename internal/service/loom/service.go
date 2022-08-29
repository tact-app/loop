package loom

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"go.octolab.org/tact/loop/internal/domain"
)

func New(client ServiceClient) *Service {
	return &Service{c: client}
}

type Service struct {
	c ServiceClient
}

func (s *Service) Workspaces(ctx context.Context) ([]domain.Workspace, error) {
	var (
		workspaces []domain.Workspace
		spaces     []domain.Space
	)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		workspaces, err = s.workspaces(ctx)
		if err != nil {
			return fmt.Errorf("fetch workspaces: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		spaces, err = s.spaces(ctx)
		if err != nil {
			return fmt.Errorf("fetch spaces: %w", err)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("fetch top-level structure: %w", err)
	}

	// TODO:performance O(n^2)
	for i := range workspaces {
		for j := range spaces {
			if workspaces[i].ID == spaces[j].WorkspaceID {
				workspaces[i].Spaces = append(workspaces[i].Spaces, spaces[j])
			}
		}
	}

	return workspaces, nil
}
