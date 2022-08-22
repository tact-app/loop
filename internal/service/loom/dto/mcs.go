package dto

type ClosedSpaces struct {
	Data struct {
		Result struct {
			Memberships *struct {
				Edges []struct {
					Node struct {
						ID    string `json:"id"`
						Space struct {
							ID        string `json:"id"`
							IsPrimary bool   `json:"is_primary"`
							Name      string `json:"name"`
							Privacy   any    `json:"privacy"`
						} `json:"space"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"memberships,omitempty"`
			Message string `json:"message,omitempty"`
		} `json:"result"`
	} `json:"data"`
}
