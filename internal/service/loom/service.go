package loom

import (
	"context"
	"fmt"

	"go.octolab.org/pointer"
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
		archive, library []domain.Folder
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
			err = fmt.Errorf("fetch archived: %w", err)
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
		for j := range archive {
			if workspaces[i].ID == archive[j].WorkspaceID {
				root, has := workspaces[i].Archive[archive[j].OwnerID]
				if !has {
					root = domain.Folder{
						ID:          pointer.ValueOfString(archive[j].ParentID),
						OwnerID:     archive[j].OwnerID,
						WorkspaceID: archive[j].WorkspaceID,
						Name:        "Archived",
						Folders:     make(map[domain.FolderID]domain.Folder),
					}
				}
				root.Folders[archive[j].ID] = archive[j]
				workspaces[i].Archive[archive[j].OwnerID] = root
			}
		}
		for j := range library {
			if workspaces[i].ID == library[j].WorkspaceID {
				root, has := workspaces[i].Library[library[j].OwnerID]
				if !has {
					root = domain.Folder{
						ID:          pointer.ValueOfString(library[j].ParentID),
						OwnerID:     library[j].OwnerID,
						WorkspaceID: library[j].WorkspaceID,
						Name:        "Personal Library",
						Folders:     make(map[domain.FolderID]domain.Folder),
					}
				}
				root.Folders[library[j].ID] = library[j]
				workspaces[i].Library[library[j].OwnerID] = root
			}
		}
	}

	//{
	//	"first": 99,
	//	"parentFolderId": null,
	//	"source": "SPACE",
	//	"sourceValue": "{{spaceId}}",
	//	"sortOrder": "ASC",
	//	"sortType": "NAME"
	//}

	return workspaces, nil
}
