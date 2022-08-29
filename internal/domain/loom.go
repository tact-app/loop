package domain

type Workspace struct {
	ID     int
	Name   string
	Type   string
	Spaces []Space
}

type Space struct {
	ID          int
	Name        string
	IsPrivate   bool
	IsPrimary   bool
	WorkspaceID int

	Root Folder
}

type Folder struct {
	Folders []Folder
	Looms   []Loom
}

type Loom struct {
	Name string
	URL  string
	Tags []Tag
}

type Tag string
