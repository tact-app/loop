package loom

import (
	"context"
	"fmt"
	"strconv"

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
			return fmt.Errorf("fetch spaces: %w", err)
		}
		g, ctx := errgroup.WithContext(ctx)
		for i := range spaces {
			idx := i
			g.Go(func() error {
				scope := Vars{
					"first":       99,
					"sortType":    "NAME",
					"sortOrder":   "ASC",
					"source":      "SPACE",
					"sourceValue": strconv.Itoa(spaces[idx].ID),
				}
				folders, err := s.folders(ctx, scope, true)
				if err != nil {
					return fmt.Errorf("fetch folders of space %d: %w", spaces[idx].ID, err)
				}

				spaces[idx].Folders = make(map[domain.FolderID]domain.Folder, len(folders))
				for _, folder := range folders {
					spaces[idx].Folders[folder.ID] = folder
				}
				return nil
			})
		}
		return g.Wait()
	})
	g.Go(func() error {
		var err error
		scope := Vars{
			"first":     99,
			"sortOrder": "ASC",
			"sortType":  "NAME",
			"source":    "ACTIVE",
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
			"first":     99,
			"sortType":  "NAME",
			"sortOrder": "ASC",
			"source":    "ARCHIVED",
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

	return workspaces, g.Wait()
}
