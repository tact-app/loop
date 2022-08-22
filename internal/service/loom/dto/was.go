package dto

type ArchivedSpaces struct {
	Data struct {
		Result struct {
			Message string `json:"message,omitempty"`
			Spaces  *struct {
				Nodes []struct {
					ID        string  `json:"id"`
					IsPrimary bool    `json:"is_primary"`
					Name      string  `json:"name"`
					Privacy   *string `json:"privacy"`
				} `json:"nodes"`
			} `json:"spaces,omitempty"`
		} `json:"result"`
	} `json:"data"`
}
