package loom

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"

	"go.octolab.org/tact/loop/internal/domain"
	"go.octolab.org/tact/loop/internal/service/loom/dto"
)

func (s *Service) spaces(ctx context.Context) ([]domain.Space, error) {
	var mu sync.Mutex
	index := make(map[string]domain.Space)
	limit := map[string]interface{}{"first": 1000}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var r dto.OpenSpaces
		if err := s.c.Do(ctx, OpenSpaces, limit, &r); err != nil {
			return fmt.Errorf("fetch open spaces: %w", err)
		}
		if err := r.Data.Result.Message; err != "" {
			return fmt.Errorf("fetch open spaces: %s", err)
		}

		mu.Lock()
		for _, edge := range r.Data.Result.Spaces.Edges {
			index[edge.Node.ID] = domain.Space{
				ID:        edge.Node.ID,
				Name:      edge.Node.Name,
				IsPrivate: edge.Node.Privacy != "workspace",
				IsPrimary: edge.Node.IsPrimary,
			}
		}
		mu.Unlock()
		return nil
	})
	g.Go(func() error {
		var r dto.ClosedSpaces
		if err := s.c.Do(ctx, ClosedSpaces, limit, &r); err != nil {
			return fmt.Errorf("fetch closed spaces: %w", err)
		}
		if err := r.Data.Result.Message; err != "" {
			return fmt.Errorf("fetch closed spaces: %s", err)
		}

		mu.Lock()
		for _, edge := range r.Data.Result.Memberships.Edges {
			index[edge.Node.Space.ID] = domain.Space{
				ID:        edge.Node.Space.ID,
				Name:      edge.Node.Space.Name,
				IsPrivate: edge.Node.Space.Privacy == nil,
				IsPrimary: edge.Node.Space.IsPrimary,
			}
		}
		mu.Unlock()
		return nil
	})
	g.Go(func() error {
		var r dto.ArchivedSpaces
		if err := s.c.Do(ctx, ArchivedSpaces, limit, &r); err != nil {
			return fmt.Errorf("fetch closed spaces: %w", err)
		}
		if err := r.Data.Result.Message; err != "" {
			return fmt.Errorf("fetch closed spaces: %s", err)
		}

		mu.Lock()
		for _, node := range r.Data.Result.Spaces.Nodes {
			index[node.ID] = domain.Space{
				ID:        node.ID,
				Name:      node.Name,
				IsPrivate: node.Privacy == nil,
				IsPrimary: node.IsPrimary,
			}
		}
		mu.Unlock()
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("fetch spaces: %w", err)
	}

	spaces := make([]domain.Space, 0, len(index))
	for _, space := range index {
		spaces = append(spaces, space)
	}
	return spaces, nil
}
