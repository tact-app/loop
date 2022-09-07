package domain

type (
	FolderID    = string
	LoomID      = string
	OwnerID     = int
	SpaceID     = int
	WorkspaceID = int
)

type Workspace struct {
	ID      WorkspaceID        `json:"id" yaml:"id"`
	Name    string             `json:"name" yaml:"name"`
	Type    string             `json:"type" yaml:"type"`
	Archive map[OwnerID]Folder `json:"archive" yaml:"archive"`
	Library map[OwnerID]Folder `json:"library" yaml:"library"`
	Spaces  map[SpaceID]Space  `json:"spaces" yaml:"spaces"`
}

type Space struct {
	ID          SpaceID             `json:"id" yaml:"id"`
	WorkspaceID WorkspaceID         `json:"workspace_id" yaml:"workspace_id"`
	Name        string              `json:"name" yaml:"name"`
	IsArchived  bool                `json:"is_archived" yaml:"is_archived"`
	IsPrimary   bool                `json:"is_primary" yaml:"is_primary"`
	IsPrivate   bool                `json:"is_private" yaml:"is_private"`
	Folders     map[FolderID]Folder `json:"folders,omitempty" yaml:"folders,omitempty"`
}

type Folder struct {
	ID          FolderID            `json:"id" yaml:"id"`
	OwnerID     OwnerID             `json:"owner_id" yaml:"owner_id"`
	WorkspaceID WorkspaceID         `json:"workspace_id" yaml:"workspace_id"`
	Name        string              `json:"name" yaml:"name"`
	IsEditable  bool                `json:"is_editable" yaml:"is_editable"`
	IsShared    bool                `json:"is_shared" yaml:"is_shared"`
	ParentID    *FolderID           `json:"parent_id" yaml:"parent_id"`
	Folders     map[FolderID]Folder `json:"folders,omitempty" yaml:"folders,omitempty"`
	Looms       map[LoomID]Loom     `json:"looms,omitempty" yaml:"looms,omitempty"`
}

type Loom struct {
	ID   LoomID `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url" yaml:"url"`
	Tags []Tag  `json:"tags,omitempty" yaml:"tags,omitempty"`
}

type Tag string
