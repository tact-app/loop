package domain

type Workspace struct {
	ID      int
	Name    string
	Type    string
	Archive map[int]Folder // Owner ID => Archived
	Library map[int]Folder // Owner ID => Personal Library
	Spaces  map[int]Space  // Space ID => Space
}

type Space struct {
	ID          int
	WorkspaceID int
	Name        string
	IsPrivate   bool
	IsPrimary   bool
	Folders     map[string]Folder // Folder ID => Folder
}

type Folder struct {
	ID          string
	OwnerID     int
	WorkspaceID int
	Name        string
	IsEditable  bool
	IsShared    bool
	Parent      *Folder
	Folders     map[string]Folder // Folder ID => Folder
	Looms       map[string]Loom   // Loom ID => Loom
}

type Loom struct {
	ID   string
	Name string
	URL  string
	Tags []Tag
}

type Tag string
