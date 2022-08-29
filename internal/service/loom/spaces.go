package loom

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"

	"go.octolab.org/tact/loop/internal/domain"
	"go.octolab.org/tact/loop/internal/pkg/unsafe"
	"go.octolab.org/tact/loop/internal/service/loom/dto"
)

func (s *Service) workspaces(ctx context.Context) ([]domain.Workspace, error) {
	var r dto.Workspaces
	if err := s.c.Do(ctx, Workspaces, nil, &r); err != nil {
		return nil, fmt.Errorf("fetch user workspaces: %w", err)
	}

	workspaces := make([]domain.Workspace, 0, len(r.Data.UserWorkspaceMemberships))
	for _, workspace := range r.Data.UserWorkspaceMemberships {
		org := workspace.Organization
		workspaces = append(workspaces, domain.Workspace{
			ID:     unsafe.ReturnInt(strconv.Atoi(org.ID)),
			Name:   org.Name,
			Type:   org.Type,
			Spaces: make(map[int]domain.Space, org.Counts.Spaces.TotalActiveSpaces),
		})
	}
	return workspaces, nil
}

func (s *Service) spaces(ctx context.Context) ([]domain.Space, error) {
	var mu sync.Mutex
	index := make(map[string]domain.Space)
	limit := Vars{"first": 1000}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var r dto.PublicSpaces
		if err := s.c.Do(ctx, PublicSpaces, limit, &r); err != nil {
			return fmt.Errorf("fetch open spaces: %w", err)
		}
		if err := r.Data.Result.Message; err != "" {
			return fmt.Errorf("fetch open spaces: %s", err)
		}

		mu.Lock()
		for _, edge := range r.Data.Result.Spaces.Edges {
			space := edge.Node
			index[edge.Node.ID] = domain.Space{
				ID:          unsafe.ReturnInt(strconv.Atoi(space.ID)),
				Name:        space.Name,
				IsPrivate:   space.Privacy != "workspace",
				IsPrimary:   space.IsPrimary,
				WorkspaceID: space.WorkspaceID,
			}
		}
		mu.Unlock()
		return nil
	})
	g.Go(func() error {
		var r dto.PrivateSpaces
		if err := s.c.Do(ctx, PrivateSpaces, limit, &r); err != nil {
			return fmt.Errorf("fetch closed spaces: %w", err)
		}
		if err := r.Data.Result.Message; err != "" {
			return fmt.Errorf("fetch closed spaces: %s", err)
		}

		mu.Lock()
		for _, edge := range r.Data.Result.Memberships.Edges {
			space := edge.Node.Space
			index[edge.Node.Space.ID] = domain.Space{
				ID:          unsafe.ReturnInt(strconv.Atoi(space.ID)),
				Name:        space.Name,
				IsPrivate:   space.Privacy == nil,
				IsPrimary:   space.IsPrimary,
				WorkspaceID: space.WorkspaceID,
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
			space := node
			index[node.ID] = domain.Space{
				ID:          unsafe.ReturnInt(strconv.Atoi(space.ID)),
				Name:        space.Name,
				IsPrivate:   space.Privacy == nil,
				IsPrimary:   space.IsPrimary,
				WorkspaceID: space.WorkspaceID,
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

func (s *Service) folders(ctx context.Context, scope Vars, nested bool) (*domain.Folder, error) {
	var r dto.Folders
	if err := s.c.Do(ctx, Folders, scope, &r); err != nil {
		return nil, err
	}
	return &domain.Folder{ID: "fake", Name: "Stub"}, nil
	//{
	//	"first": 99,
	//	"parentFolderId": null,
	//	"source": "SPACE",
	//	"sourceValue": "{{spaceId}}",
	//	"sortOrder": "ASC",
	//	"sortType": "NAME"
	//}
}
