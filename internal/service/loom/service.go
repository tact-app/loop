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
		workspaces       []domain.Workspace
		spaces           []domain.Space
		archive, library *domain.Folder
	)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		workspaces, err = s.workspaces(ctx)
		if err != nil {
			err = fmt.Errorf("fetch workspaces: %w", err)
		}
		return err
	})
	g.Go(func() error {
		var err error
		spaces, err = s.spaces(ctx)
		if err != nil {
			err = fmt.Errorf("fetch spaces: %w", err)
		}
		return err
	})
	g.Go(func() error {
		var err error
		scope := Vars{
			"first":          99,
			"parentFolderId": nil,
			"source":         "ACTIVE",
			"sortOrder":      "ASC",
			"sortType":       "NAME",
		}
		library, err = s.folders(ctx, scope, true)
		if err != nil {
			err = fmt.Errorf("fetch personal library: %w", err)
		}
		return err
	})
	g.Go(func() error {
		var err error
		scope := Vars{
			"first":          99,
			"parentFolderId": nil,
			"source":         "ARCHIVED",
			"sortOrder":      "ASC",
			"sortType":       "NAME",
		}
		archive, err = s.folders(ctx, scope, true)
		if err != nil {
			err = fmt.Errorf("fetch archive: %w", err)
		}
		return err
	})
	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("fetch top-level structure: %w", err)
	}

	// TODO:performance O(n^2)
	for i := range workspaces {
		for j := range spaces {
			if workspaces[i].ID == spaces[j].WorkspaceID {
				workspaces[i].Spaces[spaces[j].ID] = spaces[j]
			}
		}
	}
	_, _ = archive, library

	return workspaces, nil
}
