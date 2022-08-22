package dto

type OpenSpaces struct {
	Data struct {
		Result struct {
			Message string `json:"message,omitempty"`
			Spaces  *struct {
				Edges []struct {
					Node struct {
						ID        string `json:"id"`
						IsPrimary bool   `json:"is_primary"`
						Name      string `json:"name"`
						Privacy   string `json:"privacy"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"spaces,omitempty"`
		} `json:"result"`
	} `json:"data"`
}
