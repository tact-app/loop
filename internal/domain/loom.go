package domain

type (
	FolderID    = string
	LoomID      = string
	OwnerID     = int
	SpaceID     = int
	WorkspaceID = int
)

type Workspace struct {
	ID      WorkspaceID
	Name    string
	Type    string
	Archive map[OwnerID]Folder
	Library map[OwnerID]Folder
	Spaces  map[SpaceID]Space
}

type Space struct {
	ID          SpaceID
	WorkspaceID WorkspaceID
	Name        string
	IsPrivate   bool
	IsPrimary   bool
	Folders     map[FolderID]Folder
}

type Folder struct {
	ID          FolderID
	OwnerID     OwnerID
	WorkspaceID WorkspaceID
	Name        string
	IsEditable  bool
	IsShared    bool
	Parent      *Folder
	Folders     map[FolderID]Folder
	Looms       map[LoomID]Loom
}

type Loom struct {
	ID   LoomID
	Name string
	URL  string
	Tags []Tag
}

type Tag string
