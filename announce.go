package palworld

type AnnounceRequestBody struct {
	Message string `json:"message"`
}

func (c *Client) Announce(message string) error {
	apiPath := "/v1/api/announce"

	body := &AnnounceRequestBody{
		Message: message,
	}

	if err := c.post(apiPath, body); err != nil {
		return err
	}

	return nil
}
