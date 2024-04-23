package palworld

func (c *Client) Save() error {
	apiPath := "/v1/api/save"

	if err := c.post(apiPath, nil); err != nil {
		return err
	}

	return nil
}
