package palworld

type Information struct {
	Version     string `json:"version"`
	ServerName  string `json:"servername"`
	Description string `json:"description"`
}

func (c *Client) Information() (*Information, error) {
	apiPath := "/v1/api/info"

	var v Information
	if err := c.get(apiPath, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
