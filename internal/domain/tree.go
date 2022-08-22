package domain

type Tree struct {
	Spaces []Space
}

type Space struct {
	ID        string
	Name      string
	IsPrivate bool
	IsPrimary bool

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
