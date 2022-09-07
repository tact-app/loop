package dto

import (
	"time"
)

type Folders struct {
	Data struct {
		GetPublishedFolders struct {
			Folders *struct {
				Edges []struct {
					Cursor string `json:"cursor"`
					Node   struct {
						ContentLastUpdated   time.Time `json:"contentLastUpdated"`
						CurrentUserCanEdit   bool      `json:"currentUserCanEdit"`
						HasDuplicateInSpaces bool      `json:"hasDuplicateInSpaces"`
						HasSubFolders        bool      `json:"hasSubFolders"`
						ID                   string    `json:"id"`
						IsArchived           bool      `json:"isArchived"`
						IsTopLevelFolder     bool      `json:"isTopLevelFolder"`
						Name                 string    `json:"name"`
						Organization         struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"organization"`
						OrganizationID int `json:"organization_id"`
						Owner          struct {
							Avatars []struct {
								Name  string `json:"name"`
								Thumb string `json:"thumb"`
							} `json:"avatars"`
							Email     string `json:"email"`
							FirstName string `json:"first_name"`
							LastName  string `json:"last_name"`
						} `json:"owner"`
						OwnerID      int `json:"owner_id"`
						ParentFolder struct {
							ID    string `json:"id"`
							Name  string `json:"name"`
							Owner struct {
								Email     string `json:"email"`
								FirstName string `json:"first_name"`
								LastName  string `json:"last_name"`
							} `json:"owner"`
							OwnerID   int     `json:"owner_id"`
							SpecialID *string `json:"special_id"`
						} `json:"parent_folder"`
						Shared bool `json:"shared"`
						Space  *struct {
							ID        string `json:"id"`
							IsPrimary bool   `json:"is_primary"`
							Name      string `json:"name"`
						} `json:"space"`
						SpecialID         *string `json:"special_id"`
						TotalNestedVideos int     `json:"totalNestedVideos"`
						Visibility        string  `json:"visibility"`
					} `json:"node"`
				} `json:"edges"`
				PageInfo struct {
					EndCursor   string `json:"endCursor"`
					HasNextPage bool   `json:"hasNextPage"`
				} `json:"pageInfo"`
				TotalCount int `json:"totalCount"`
			} `json:"folders"`
		} `json:"getPublishedFolders"`
	} `json:"data"`
	Errors []struct {
		Extensions struct {
			Code string `json:"code"`
		} `json:"extensions"`
		Locations []struct {
			Column int `json:"column"`
			Line   int `json:"line"`
		} `json:"locations"`
		Message string   `json:"message"`
		Path    []string `json:"path"`
	} `json:"errors,omitempty"`
}
