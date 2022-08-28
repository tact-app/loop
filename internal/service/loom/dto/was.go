package dto

type ArchivedSpaces struct {
	Data struct {
		Result struct {
			Message string `json:"message,omitempty"`
			Spaces  *struct {
				Nodes []struct {
					ID          string  `json:"id"`
					IsPrimary   bool    `json:"is_primary"`
					Name        string  `json:"name"`
					Privacy     *string `json:"privacy"`
					TotalVideos int     `json:"totalVideos"`
					WorkspaceID int     `json:"workspace_id"`
				} `json:"nodes"`
				PageInfo struct {
					EndCursor   string `json:"endCursor"`
					HasNextPage bool   `json:"hasNextPage"`
				} `json:"pageInfo"`
			} `json:"spaces,omitempty"`
		} `json:"result"`
	} `json:"data"`
}
