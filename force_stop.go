package palworld

func (c *Client) ForceStop() error {
	apiPath := "/v1/api/stop"

	if err := c.post(apiPath, nil); err != nil {
		return err
	}

	return nil
}
