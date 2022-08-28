package dto

type PrivateSpaces struct {
	Data struct {
		Result struct {
			Memberships *struct {
				Edges []struct {
					Node struct {
						ID    string `json:"id"`
						Space struct {
							ID          string `json:"id"`
							IsPrimary   bool   `json:"is_primary"`
							Name        string `json:"name"`
							Privacy     any    `json:"privacy"`
							TotalVideos int    `json:"totalVideos"`
							WorkspaceID int    `json:"workspace_id"`
						} `json:"space"`
					} `json:"node"`
				} `json:"edges"`
				PageInfo struct {
					EndCursor   string `json:"endCursor"`
					HasNextPage bool   `json:"hasNextPage"`
				} `json:"pageInfo"`
			} `json:"memberships,omitempty"`
			Message string `json:"message,omitempty"`
		} `json:"result"`
	} `json:"data"`
}
